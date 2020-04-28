package task

import (
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/logger"
	"time"
	"github.com/irisnet/explorer/backend/types"
	"fmt"
	"github.com/irisnet/explorer/backend/utils"
	"math"
	"strings"
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/conf"
	"sync"
	"strconv"
)

//caculatetime format [ 2020-04-03T00:00:00 ]
type StaticDelegatorByMonthTask struct {
	mStaticModel           document.ExStaticDelegatorMonth
	staticModel            document.ExStaticDelegator
	txModel                document.CommonTx
	AddressCoin            map[string]document.Coin
	AddrPeriodCommission   map[string]document.Coin
	AddrTerminalCommission map[string]document.Coin
	//AddrBeginCommission    map[string]document.Coin

	startTime time.Time
	endTime   time.Time
	address   string
	isSetTime bool
}

func (task *StaticDelegatorByMonthTask) Name() string {
	return "static_delegator_month"
}
func (task *StaticDelegatorByMonthTask) Start() {
	taskName := task.Name()
	timeInterval := conf.Get().Server.CronTimeStaticDelegatorMonth

	if err := tcService.runTask(taskName, timeInterval, task.DoTask); err != nil {
		logger.Error(err.Error())
	}
}

func (task *StaticDelegatorByMonthTask) SetCaculateScope(start, end time.Time) {
	task.isSetTime = true
	task.startTime = start.In(cstZone)
	task.endTime = end.In(cstZone)
}

func (task *StaticDelegatorByMonthTask) SetCaculateAddress(address string) {
	task.address = address
}

func (task *StaticDelegatorByMonthTask) DoTask() error {

	res, err := task.caculateWork()
	if err != nil {
		return err
	}
	if len(res) == 0 {
		return nil
	}
	for _, one := range res {
		if err := task.mStaticModel.Save(one); err != nil {
			logger.Error("save static delegator month data error",
				logger.String("address", one.Address),
				logger.String("date", one.Date),
				logger.String("err", err.Error()))
		}
	}

	return nil
}

func (task *StaticDelegatorByMonthTask) caculateWork() ([]document.ExStaticDelegatorMonth, error) {
	datetime := time.Now().In(cstZone)

	year, month := getStartYearMonth(datetime)
	starttime, _ := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-01T00:00:00", year, month), cstZone)
	if task.isSetTime {
		datetime = task.endTime
		starttime = task.startTime
	}
	if conf.Get().Server.CaculateDebug {
		arr := strings.Split(conf.Get().Server.CronTimeFormatStaticMonth, " ")
		minutedata := strings.Split(arr[0], "/")
		intervalstr := minutedata[1]
		if len(minutedata) == 1 {
			intervalstr = minutedata[0]
		}
		interval, err := strconv.ParseInt(intervalstr, 10, 64)
		if err != nil {
			logger.Error(err.Error())
		}
		hour, minute := datetime.Hour(), datetime.Minute()
		if int64(minute) < interval {
			if hour < 1 {
				//no work
				return nil, fmt.Errorf("time hour is smaller than 1")
			} else {
				hour --
				minute = minute - int(interval) + 60
			}
		} else {
			minute = minute - int(interval)
		}
		starttime, _ = time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT%02d:%02d:00", datetime.Year(), datetime.Month(), datetime.Day(), hour, minute), cstZone)
	}
	//find last date
	endtime, err := document.Getdate(task.staticModel.Name(), starttime, datetime, "-"+document.ExStaticDelegatorDateTag)
	if err != nil {
		return nil, err
	}

	var terminalData []document.ExStaticDelegator
	if task.address != "" {
		data, err := task.staticModel.GetDataOneDay(endtime, task.address)
		if err != nil {
			return nil, err
		}
		terminalData = append(terminalData, data)
	} else {
		terminalData, err = task.staticModel.GetDataByDate(endtime)
		if err != nil {
			logger.Error("Get GetData ByDate fail",
				logger.String("date", datetime.String()),
				logger.String("err", err.Error()))
			return nil, err
		}
	}

	res := make([]document.ExStaticDelegatorMonth, 0, len(terminalData))

	txs, err := task.getPeriodTxByAddress(starttime, datetime, task.address) //default is all address txs
	if err != nil {
		return nil, err
	}
	logger.Debug(fmt.Sprintf("total delegator:%v", len(terminalData)))

	for i, val := range terminalData {
		one, err := task.getStaticDelegator(starttime, val, txs)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		one.CaculateDate = fmt.Sprintf("%d.%02d.%02d", datetime.Year(), datetime.Month(), datetime.Day())
		if conf.Get().Server.CaculateDebug {
			one.Date = fmt.Sprintf("%d.%02d.%02d %02d:%02d:%02d", starttime.Year(),
				starttime.Month(), starttime.Day(), starttime.Hour(), starttime.Minute(), starttime.Second())
			one.CaculateDate = fmt.Sprintf("%d.%02d.%02d %02d:%02d:%02d", datetime.Year(),
				datetime.Month(), datetime.Day(), datetime.Hour(), datetime.Minute(), datetime.Second())
		}
		res = append(res, one)
		logger.Debug(fmt.Sprintf("caculate delegator index:%v", i))
	}
	return res, nil
}

func parseCoinAmountAndUnitFromStr(s string) (document.Coin) {
	for _, denom := range rewardsDenom {
		if strings.HasSuffix(s, denom) {
			coin := utils.ParseCoin(s)
			return document.Coin{
				Denom:  coin.Denom,
				Amount: coin.Amount,
			}
		}
	}
	return document.Coin{}
}

func (task *StaticDelegatorByMonthTask) getStaticDelegator(starttime time.Time, terminalval document.ExStaticDelegator, txs []document.CommonTx) (document.ExStaticDelegatorMonth, error) {
	periodRewards := task.getPeriodRewards(terminalval.Address)
	delagation, err := task.staticModel.GetDataOneDay(starttime, terminalval.Address)
	if err != nil {
		logger.Error("get DelegationData failed",
			logger.String("func", "get StaticDelegator"),
			logger.String("err", err.Error()))
		return document.ExStaticDelegatorMonth{}, err
	}
	if delagation.Address == "" {
		delagation.Address = terminalval.Address
		delagation.Date = starttime
	}

	//if task.AddrBeginCommission == nil {
	//	task.AddrBeginCommission = make(map[string]document.Coin)
	//}
	if task.AddrTerminalCommission == nil {
		task.AddrTerminalCommission = make(map[string]document.Coin)
	}
	//if len(delagation.Commission) > 0 {
	//	task.AddrBeginCommission[delagation.Address] = document.Coin{
	//		Amount: delagation.Commission[0].Iris * math.Pow10(18),
	//		Denom:  types.IRISAttoUint,
	//	}
	//}
	if len(terminalval.Commission) > 0 {
		task.AddrTerminalCommission[terminalval.Address] = document.Coin{
			Amount: terminalval.Commission[0].Iris * math.Pow10(18),
			Denom:  types.IRISAttoUint,
		}
	}

	delagationlastmonth, err := task.mStaticModel.GetLatest(terminalval.Address)
	if err != nil {
		logger.Error("get GetLatest failed",
			logger.String("func", "get StaticDelegatorMonth"),
			logger.String("err", err.Error()))
		return document.ExStaticDelegatorMonth{}, err
	}
	if delagationlastmonth.Address == "" {
		delagationlastmonth.Address = terminalval.Address
	}

	delegationrewards := float64(0)
	if len(terminalval.Total) > 0 {
		if len(terminalval.Commission) > 0 {
			delegationrewards = terminalval.Total[0].Iris - terminalval.Commission[0].Iris
		} else {
			delegationrewards = terminalval.Total[0].Iris
		}
	}
	terminalRewards := document.Rewards{Iris: delegationrewards, IrisAtto: fmt.Sprint(delegationrewards * math.Pow10(18))}
	incrementRewards := task.getIncrementRewards(terminalRewards, delagationlastmonth, periodRewards)

	item := document.ExStaticDelegatorMonth{
		Address:                terminalval.Address,
		Date:                   fmt.Sprintf("%d.%02d.%02d", starttime.Year(), starttime.Month(), starttime.Day()),
		TerminalDelegation:     document.Coin{Denom: terminalval.Delegation.Denom, Amount: terminalval.Delegation.Amount},
		PeriodDelegationTimes:  task.getPeriodDelegationTimes(terminalval.Address, txs),
		PeriodWithdrawRewards:  periodRewards,
		IncrementDelegation:    task.getIncrementDelegation(terminalval.Delegation, delagation.Delegation),
		PeriodIncrementRewards: incrementRewards,
		TerminalRewards:        terminalRewards,
	}

	return item, nil
}
func (task *StaticDelegatorByMonthTask) getPeriodTxByAddress(starttime, endtime time.Time, address string) ([]document.CommonTx, error) {
	txs, err := task.txModel.GetTxsByDurationAddress(starttime, endtime, address)
	if err != nil {
		return nil, err
	}
	var coinflow []string
	group := sync.WaitGroup{}
	for _, tx := range txs {
		switch tx.Type {
		case types.TxTypeWithdrawDelegatorReward, types.TxTypeWithdrawDelegatorRewardsAll, types.TxTypeWithdrawValidatorRewardsAll,
			types.TxTypeBeginRedelegate, types.TxTypeStakeBeginUnbonding, types.TxTypeStakeDelegate:
			group.Add(1)
			go func(txHash string) {
				result := lcd.BlockCoinFlow(txHash)
				coinflow = append(coinflow, result.CoinFlow...)
				//fmt.Println(txHash)
				group.Done()
			}(tx.TxHash)
		}
	}
	group.Wait()
	task.getCoinflow(coinflow)

	return txs, nil
}

func (task *StaticDelegatorByMonthTask) getPeriodRewards(address string) document.Rewards {
	var rewards document.Rewards
	if data, ok := task.AddressCoin[address]; ok {
		rewards.Iris += data.Amount / math.Pow10(18)
	}
	rewards.IrisAtto = fmt.Sprint(rewards.Iris * math.Pow10(18))
	return rewards
}

func (task *StaticDelegatorByMonthTask) getPeriodDelegationTimes(address string, txs []document.CommonTx) (total int) {
	for _, val := range txs {
		if len(val.Signers) == 0 {
			continue
		}
		if address != val.Signers[0].AddrBech32 {
			continue
		}
		if val.Type == types.TxTypeStakeDelegate ||
			val.Type == types.TxTypeBeginRedelegate ||
			val.Type == types.TxTypeStakeBeginUnbonding {
			total++
		}
	}
	return
}

func (task *StaticDelegatorByMonthTask) getIncrementRewards(delegaterewards document.Rewards,
	delagationlastmonth document.ExStaticDelegatorMonth,
	periodRewards document.Rewards) (document.Rewards) {
	var rewards document.Rewards
	rewards.Iris = delegaterewards.Iris

	//Rdx = Rdn - Rdn-1  + Rdw
	rewards.Iris -= delagationlastmonth.TerminalRewards.Iris

	rewards.Iris += periodRewards.Iris
	rewards.IrisAtto = fmt.Sprint(rewards.Iris * math.Pow10(18))
	return rewards
}

func (task *StaticDelegatorByMonthTask) getCoinflow(coinFlow []string) {

	if length := len(coinFlow); length > 0 {
		if task.AddressCoin == nil {
			task.AddressCoin = make(map[string]document.Coin, length)
		}
		if task.AddrPeriodCommission == nil {
			task.AddrPeriodCommission = make(map[string]document.Coin, length)
		}
		for _, val := range coinFlow {
			values := strings.Split(val, "::")
			if len(values) != 6 {
				continue
			}
			coinflowAmount := parseCoinAmountAndUnitFromStr(values[2])
			if strings.HasPrefix(values[3], types.DelegatorRewardTag) {
				if data, ok := task.AddressCoin[values[1]]; ok {
					data.Amount += coinflowAmount.Amount
					task.AddressCoin[values[1]] = data
				} else {
					task.AddressCoin[values[1]] = coinflowAmount
				}
			} else if strings.HasPrefix(values[3], types.ValidatorRewardTag) {
				address := utils.Convert(conf.Get().Hub.Prefix.AccAddr, values[0])
				if data, ok := task.AddressCoin[address]; ok {
					data.Amount += coinflowAmount.Amount
					task.AddressCoin[address] = data
				} else {
					task.AddressCoin[address] = coinflowAmount
				}
			} else if strings.HasPrefix(values[3], types.ValidatorCommissionTag) {
				address := utils.Convert(conf.Get().Hub.Prefix.AccAddr, values[0])
				if data, ok := task.AddrPeriodCommission[address]; ok {
					data.Amount += coinflowAmount.Amount
					task.AddrPeriodCommission[address] = data
				} else {
					task.AddrPeriodCommission[address] = coinflowAmount
				}
			}
		}

	}
	return
}

func (task *StaticDelegatorByMonthTask) getIncrementDelegation(terminal, delagation utils.Coin) document.Coin {
	amount := terminal.Amount - delagation.Amount
	return document.Coin{
		Denom:  terminal.Denom,
		Amount: amount,
	}

}

//func (task *StaticDelegatorByMonthTask) getDelegationData(caculatetime string) ([]document.ExStaticDelegator, error) {
//	date, err := time.ParseInLocation(types.TimeLayout, caculatetime, cstZone)
//	if err != nil {
//		return nil, err
//	}
//	data, err := task.staticModel.GetDataByDate(date)
//	if err != nil {
//		logger.Error("Get GetData ByDate fail",
//			logger.String("date", date.String()),
//			logger.String("err", err.Error()))
//		return nil, err
//	}
//	return data, nil
//}

//func (task *StaticDelegatorByMonthTask) saveOps(datas []document.ExStaticDelegatorMonth) ([]txn.Op) {
//	var ops = make([]txn.Op, 0, len(datas))
//	for _, val := range datas {
//		val.Id = bson.NewObjectId()
//		val.CreateAt = time.Now().Unix()
//		val.UpdateAt = time.Now().Unix()
//		op := txn.Op{
//			C:      task.mStaticModel.Name(),
//			Id:     bson.NewObjectId(),
//			Insert: val,
//		}
//		ops = append(ops, op)
//	}
//	return ops
//}