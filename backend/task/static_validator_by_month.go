package task

import (
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/explorer/backend/orm/document"
	"gopkg.in/mgo.v2/txn"
	"gopkg.in/mgo.v2/bson"
	"time"
	"github.com/irisnet/explorer/backend/types"
	"fmt"
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/service"
	"github.com/irisnet/explorer/backend/lcd"
)

type StaticValidatorByMonthTask struct {
	mStaticModel              document.ExStaticValidatorMonth
	staticModel               document.ExStaticValidator
	txModel                   document.CommonTx
	AddressCoin               map[string]document.Coin
	AddressPeriodCommission   map[string]document.Coin
	AddressTerminalCommission map[string]document.Coin
	AddressBeginCommission    map[string]document.Coin
}

func (task *StaticValidatorByMonthTask) Name() string {
	return "static_validator_month"
}
func (task *StaticValidatorByMonthTask) Start() {
	taskName := task.Name()
	timeInterval := 3600 * 24 * 30

	if err := tcService.runTask(taskName, timeInterval, task.DoTask); err != nil {
		logger.Error(err.Error())
	}
}

func (task *StaticValidatorByMonthTask) SetAddressCoinMapData(rewards, pcommission, bcommission, tcommission map[string]document.Coin) {
	task.AddressCoin = rewards
	task.AddressPeriodCommission = pcommission
	task.AddressBeginCommission = bcommission
	task.AddressTerminalCommission = tcommission
}

func (task *StaticValidatorByMonthTask) DoTask() error {
	datetime := time.Now().In(cstZone)
	//year, month := datetime.Year(), datetime.Month()-1
	//if datetime.Month() == time.January {
	//	year = datetime.Year() - 1
	//	month = time.December
	//}
	year, month := getTerminalYearMonth(datetime)
	starttime, _ := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-01T00:00:00", year, month), cstZone)
	terminaldate, err := document.Getdate(task.staticModel.Name(), starttime, datetime, "-"+document.ValidatorStaticFieldDate)
	if err != nil {
		return err
	}

	terminalData, err := task.staticModel.GetDataByDate(terminaldate)
	if err != nil {
		logger.Error("Get GetData ByDate fail",
			logger.String("date", terminaldate.String()),
			logger.String("err", err.Error()))
		return err
	}
	res := make([]document.ExStaticValidatorMonth, 0, len(terminalData))

	addressHeightMap, err := task.getCreateValidatorTx()
	if err != nil {
		return err
	}

	for _, val := range terminalData {
		one := task.getStaticValidator(val, addressHeightMap)
		res = append(res, one)
	}

	if len(res) == 0 {
		return nil
	}
	return task.mStaticModel.Batch(task.saveOps(res))
}

func (task *StaticValidatorByMonthTask) getStaticValidator(terminalval document.ExStaticValidator, addrHeightMap map[string]int64) document.ExStaticValidatorMonth {
	address := utils.Convert(conf.Get().Hub.Prefix.AccAddr, terminalval.OperatorAddress)

	year, month := getTerminalYearMonth(terminalval.Date)
	starttime, _ := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-01T00:00:00", year, month), cstZone)
	date, err := document.Getdate(task.staticModel.Name(), starttime, terminalval.Date, document.ValidatorStaticFieldDate)
	if err != nil {
		logger.Error("Getdate have error",
			logger.String("time", terminalval.Date.String()),
			logger.String("err", err.Error()))

	}

	validatestatic, err := task.staticModel.GetDataOneDay(date, terminalval.OperatorAddress)
	if err != nil {
		logger.Error("get Validator static failed",
			logger.String("func", "Get DataOneDay"),
			logger.String("err", err.Error()))
	}

	delegation := lcd.GetDelegationsFromValAddrByDelAddr(types.FoundationDelegatorAddr, terminalval.OperatorAddress)

	item := document.ExStaticValidatorMonth{
		Address:                 address,
		OperatorAddress:         terminalval.OperatorAddress,
		Date:                    fmt.Sprintf("%d.%02d", terminalval.Date.Year(), terminalval.Date.Month()),
		TerminalDelegation:      terminalval.Delegations,
		TerminalDelegatorN:      terminalval.DelegatorNum,
		TerminalSelfBond:        terminalval.SelfBond,
		IncrementDelegation:     task.getIncrementDelegation(terminalval, validatestatic),
		IncrementDelegatorN:     terminalval.DelegatorNum - validatestatic.DelegatorNum,
		IncrementSelfBond:       task.getIncrementSelfBond(terminalval, validatestatic),
		FoundationDelegateT:     delegation.Shares,
		FoundationDelegateIncre: task.getFoundationDelegateIncre(terminalval, delegation.Shares),
		CommissionRateMax:       task.getCommissionRate(terminalval, validatestatic, "-"+document.ValidatorCommissionRateTag),
		CommissionRateMin:       task.getCommissionRate(terminalval, validatestatic, document.ValidatorCommissionRateTag),
	}
	if height, ok := addrHeightMap[address]; ok {
		item.CreateValidatorHeight = height
	}

	if tcommission, ok := task.AddressTerminalCommission[address]; ok {
		item.TerminalCommission = tcommission
		item.IncrementCommission = task.getIncrementCommission(address, item.TerminalCommission)
	}

	if desp, ok := service.ValidatorsDescriptionMap[terminalval.OperatorAddress]; ok {
		item.Moniker = desp.Moniker
	}

	return item
}

func getTerminalYearMonth(datetime time.Time) (year, month int) {
	year = datetime.Year()
	month = int(datetime.Month() - 1)
	if datetime.Month() == time.January {
		year = datetime.Year() - 1
		month = int(time.December)
	}
	//fmt.Println(year,month)
	return year, month
}

func (task *StaticValidatorByMonthTask) getIncrementDelegation(terminal, begin document.ExStaticValidator) string {
	subValue := funcSubStr(terminal.Delegations, begin.Delegations)
	if subValue != nil {
		return subValue.FloatString(18)
	}
	return ""
}

func (task *StaticValidatorByMonthTask) getIncrementSelfBond(terminal, begin document.ExStaticValidator) string {
	subValue := funcSubStr(terminal.SelfBond, begin.SelfBond)
	if subValue != nil {
		return subValue.FloatString(18)
	}
	return ""
}

func (task *StaticValidatorByMonthTask) getFoundationDelegateIncre(terminal document.ExStaticValidator, shares string) string {

	year, month := terminal.Date.Year(), terminal.Date.Month()-1
	if terminal.Date.Month() == time.January {
		year = terminal.Date.Year() - 1
		month = time.December
	}
	data, err := task.mStaticModel.GetValidatorStaticByMonth(fmt.Sprintf("%v.%v", year, month-1), terminal.OperatorAddress)
	if err != nil {
		logger.Error(err.Error())
	}

	subValue := funcSubStr(shares, data.FoundationDelegateT)
	if subValue != nil {
		return subValue.FloatString(18)
	}
	return ""
}

func (task *StaticValidatorByMonthTask) getIncrementCommission(address string, terminalCommission document.Coin) (IncreCommission document.Coin) {
	bcommission, ok1 := task.AddressBeginCommission[address]
	pcommission, ok2 := task.AddressPeriodCommission[address]
	if ok1 && ok2 {
		IncreCommission.Amount = terminalCommission.Amount - bcommission.Amount + pcommission.Amount
		IncreCommission.Denom = terminalCommission.Denom
	}
	return
}

func (task *StaticValidatorByMonthTask) getCommissionRate(terminal, begin document.ExStaticValidator, sorts string) (string) {
	cond := bson.M{
		document.ExStaticValidatorMonthDateTag: bson.M{
			"$gte": begin.Date,
			"$lt":  terminal.Date,
		},
	}
	selector := bson.M{
		document.ValidatorCommissionRateTag: 1,
	}
	data, err := task.staticModel.GetCommissionRate(selector, cond, sorts)
	if err != nil {
		logger.Error(err.Error())
		return ""
	}
	return data.Commission.Rate
}

func (task *StaticValidatorByMonthTask) getCreateValidatorTx() (map[string]int64, error) {
	txs, err := task.txModel.GetTxsByType(types.TxTypeStakeCreateValidator, types.Success)
	if err != nil {
		return nil, err
	}

	addressHeightMap := make(map[string]int64, len(txs))
	for _, val := range txs {
		addressHeightMap[val.From] = val.Height
	}

	return addressHeightMap, nil
}

func (task *StaticValidatorByMonthTask) saveOps(datas []document.ExStaticValidatorMonth) ([]txn.Op) {
	var ops = make([]txn.Op, 0, len(datas))
	for _, val := range datas {
		val.Id = bson.NewObjectId()
		val.CreateAt = time.Now().Unix()
		val.UpdateAt = time.Now().Unix()
		op := txn.Op{
			C:      task.mStaticModel.Name(),
			Id:     bson.NewObjectId(),
			Insert: val,
		}
		ops = append(ops, op)
	}
	return ops
}
