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
	"math"
)

type StaticValidatorByMonthTask struct {
	mStaticModel              document.ExStaticValidatorMonth
	staticModel               document.ExStaticValidator
	txModel                   document.CommonTx
	AddressCoin               map[string]document.Coin
	AddressPeriodCommission   map[string]document.Coin
	AddressTerminalCommission map[string]document.Coin
	//AddressBeginCommission    map[string]document.Coin
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

func (task *StaticValidatorByMonthTask) SetAddressCoinMapData(rewards, pcommission, tcommission map[string]document.Coin) {
	task.AddressCoin = rewards
	task.AddressPeriodCommission = pcommission
	//task.AddressBeginCommission = bcommission
	task.AddressTerminalCommission = tcommission
}

func (task *StaticValidatorByMonthTask) DoTask() error {
	datetime := time.Now().In(cstZone)
	year, month := getStartYearMonth(datetime)
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

	for i, val := range terminalData {
		year, month := getStartYearMonth(val.Date)
		starttime, _ := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-01T00:00:00", year, month), cstZone)
		startdate, err := document.Getdate(task.staticModel.Name(), starttime, val.Date, document.ValidatorStaticFieldDate)
		if err != nil {
			logger.Error("Getdate have error",
				logger.String("time", val.Date.String()),
				logger.String("err", err.Error()))
			continue
		}
		one, errf := task.getStaticValidator(startdate, val, addressHeightMap)
		if errf != nil {
			logger.Error(errf.Error())
			continue
		}
		one.Rank = i + 1
		res = append(res, one)
	}

	if len(res) == 0 {
		return nil
	}
	return task.mStaticModel.Batch(task.saveOps(res))
}

func (task *StaticValidatorByMonthTask) getStaticValidator(startdate time.Time, terminalval document.ExStaticValidator, addrHeightMap map[string]int64) (document.ExStaticValidatorMonth, error) {
	address := utils.Convert(conf.Get().Hub.Prefix.AccAddr, terminalval.OperatorAddress)

	//validatestatic, err := task.staticModel.GetDataOneDay(date, terminalval.OperatorAddress)
	//if err != nil {
	//	logger.Error("get Validator static failed",
	//		logger.String("func", "Get DataOneDay"),
	//		logger.String("err", err.Error()))
	//	return document.ExStaticValidatorMonth{}, err
	//}
	//if validatestatic.OperatorAddress == "" {
	//	validatestatic.OperatorAddress = terminalval.OperatorAddress
	//	validatestatic.Date = date
	//}

	latestone, err := task.mStaticModel.GetLatest(terminalval.OperatorAddress)
	if err != nil {
		logger.Error("get latest one failed", logger.String("func", "get IncrementCommission"),
			logger.String("err", err.Error()))
	}
	if latestone.OperatorAddress == "" {
		latestone.OperatorAddress = terminalval.OperatorAddress
	}

	delegation := lcd.GetDelegationsFromValAddrByDelAddr(types.FoundationDelegatorAddr, terminalval.OperatorAddress)

	item := document.ExStaticValidatorMonth{
		Address:            address,
		OperatorAddress:    terminalval.OperatorAddress,
		Date:               fmt.Sprintf("%d.%02d", startdate.Year(), startdate.Month()),
		CaculateDate:       fmt.Sprintf("%d.%02d.%02d", terminalval.Date.Year(), terminalval.Date.Month(), terminalval.Date.Day()),
		TerminalDelegation: terminalval.Delegations,
		TerminalDelegatorN: terminalval.DelegatorNum,
		TerminalSelfBond:   terminalval.SelfBond,
		//Tokens:                  terminalval.Tokens,
		IncrementDelegation:     task.getIncrementDelegation(terminalval.Delegations, latestone.TerminalDelegation),
		IncrementDelegatorN:     terminalval.DelegatorNum - latestone.TerminalDelegatorN,
		IncrementSelfBond:       task.getIncrementSelfBond(terminalval.SelfBond, latestone.TerminalSelfBond),
		FoundationDelegateT:     delegation.Shares,
		FoundationDelegateIncre: task.getFoundationDelegateIncre(delegation.Shares, latestone.FoundationDelegateT),
		CommissionRateMax:       task.getCommissionRate(startdate, terminalval.Date, "-"+document.ValidatorCommissionRateTag),
		CommissionRateMin:       task.getCommissionRate(startdate, terminalval.Date, document.ValidatorCommissionRateTag),
	}
	if height, ok := addrHeightMap[address]; ok {
		item.CreateValidatorHeight = height
	}

	pcommission, ok := task.AddressPeriodCommission[address]
	if ok {
		if pcommission.Denom == types.IRISAttoUint {
			item.PeriodCommission = document.Coin{
				Denom:  types.IRISUint,
				Amount: pcommission.Amount / math.Pow10(18),
			}
		} else {
			item.PeriodCommission = pcommission
		}
	}

	tcommission, ok := task.AddressTerminalCommission[address]
	if ok {
		if tcommission.Denom == types.IRISAttoUint {
			item.TerminalCommission = document.Coin{
				Denom:  types.IRISUint,
				Amount: tcommission.Amount / math.Pow10(18),
			}
		} else {
			item.TerminalCommission = tcommission
		}
	}
	item.IncrementCommission = task.getIncrementCommission(pcommission, tcommission, latestone.TerminalCommission)

	if desp, ok := service.ValidatorsDescriptionMap[terminalval.OperatorAddress]; ok {
		item.Moniker = desp.Moniker
	}
	if item.IncrementCommission.Denom == types.IRISAttoUint {
		item.IncrementCommission = document.Coin{
			Denom:  types.IRISUint,
			Amount: item.IncrementCommission.Amount / math.Pow10(18),
		}
	}

	return item, nil
}

func getStartYearMonth(datetime time.Time) (year, month int) {
	year = datetime.Year()
	month = int(datetime.Month() - 1)
	if datetime.Month() == time.January {
		year = datetime.Year() - 1
		month = int(time.December)
	}
	//fmt.Println(year,month)
	return year, month
}

func (task *StaticValidatorByMonthTask) getIncrementDelegation(terminalDelegations, latestTerminalDelegation string) string {
	subValue := funcSubStr(terminalDelegations, latestTerminalDelegation)
	if subValue != nil {
		return subValue.FloatString(18)
	}
	return terminalDelegations
}

func (task *StaticValidatorByMonthTask) getIncrementSelfBond(terminalSelfBond, latestTerminalSelfBond string) string {
	subValue := funcSubStr(terminalSelfBond, latestTerminalSelfBond)
	if subValue != nil {
		return subValue.FloatString(18)
	}
	return terminalSelfBond
}

func (task *StaticValidatorByMonthTask) getFoundationDelegateIncre(foundationDelegateT, foundationDelegateLatest string) string {

	subValue := funcSubStr(foundationDelegateT, foundationDelegateLatest)
	if subValue != nil {
		return subValue.FloatString(18)
	}
	return foundationDelegateT
}

func (task *StaticValidatorByMonthTask) getIncrementCommission(pcommission, terminalCommission,
latestoneCommission document.Coin) (IncreCommission document.Coin) {
	//Rcx = Rcn - Rcn-1 + Rcw
	if latestoneCommission.Denom == types.IRISUint {
		latestoneCommission.Amount = latestoneCommission.Amount * math.Pow10(18)
	}
	IncreCommission.Amount = terminalCommission.Amount - latestoneCommission.Amount + pcommission.Amount
	IncreCommission.Denom = terminalCommission.Denom

	return
}

func (task *StaticValidatorByMonthTask) getCommissionRate(starttime, endtime time.Time, sorts string) (string) {
	cond := bson.M{
		document.ExStaticValidatorMonthDateTag: bson.M{
			"$gte": starttime,
			"$lt":  endtime,
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
