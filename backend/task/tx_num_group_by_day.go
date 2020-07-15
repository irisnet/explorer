package task

import (
	"time"

	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/service"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/shopspring/decimal"
)

type TxNumGroupByDayTask struct{}

func (task TxNumGroupByDayTask) Name() string {
	return "tx_num_stat_task"
}
func (task TxNumGroupByDayTask) Start() {
	taskName := task.Name()
	timeInterval := conf.Get().Server.CronTimeTxNumByDay

	if err := tcService.runTask(taskName, timeInterval, task.DoTask); err != nil {
		logger.Error(err.Error())
	}
}

func (task TxNumGroupByDayTask) DoTask(fn func(string) chan bool) error {
	stop := fn(task.Name())
	defer HeartQuit(stop)
	today := utils.TruncateTime(time.Now().In(cstZone), utils.Day)
	yesterday := today.Add(-24 * time.Hour)

	total, err := document.CommonTx{}.GetTxCountByDuration(yesterday, today)
	if err != nil {
		return err
	}

	// account info statistics
	var (
		totalAccNum  int64
		delegatorNum int64
	)
	accModel := document.Account{}
	if v, err := accModel.CountAll(); err != nil {
		logger.Error("CountAllAccNum fail", logger.String("err", err.Error()))
	} else {
		totalAccNum = int64(v)
	}
	if v, err := accModel.CountDelegatorNum(); err != nil {
		logger.Error("CountDelegatorNum fail", logger.String("err", err.Error()))
	} else {
		delegatorNum = int64(v)
	}

	// token stat statistics
	tokenStat := task.getTokenStat()

	txNumStat := document.TxNumStat{
		Date:         utils.FmtTime(yesterday, utils.DateFmtYYYYMMDD),
		Num:          int64(total),
		TotalAccNum:  totalAccNum,
		DelegatorNum: delegatorNum,
		TokenStat:    tokenStat,
		CreateTime:   time.Now(),
	}

	if err := txNumStat.Insert(); err != nil {
		return err
	}

	return nil
}

func (task TxNumGroupByDayTask) getTokenStat() document.TokenStat {
	var (
		res              document.TokenStat
		totalSupply      string
		circulation      string
		bonded           string
		foundationBonded string
	)
	tokenStatService := service.TokenStatsService{}
	if v, err := tokenStatService.QueryTokenStats(); err != nil {
		return res
	} else {
		totalSupply = v.TotalsupplyTokens.Amount
		circulation = v.CirculationTokens.Amount
		bondedAmtIrisAtto := v.DelegatedTokens.Amount

		if d, err := decimal.NewFromString(bondedAmtIrisAtto); err != nil {
			bonded = "0"
		} else {
			bonded = d.Shift(-18).String()
		}

		foundationBonded = v.FoundationBonded.Amount
	}

	res.TotalSupply = totalSupply
	res.Circulation = circulation
	res.Bonded = bonded
	res.FoundationBonded = foundationBonded

	return res
}

// init ex_tx_num_stat document
func (task TxNumGroupByDayTask) init() {

	now := utils.TruncateTime(time.Now().In(cstZone), utils.Day)
	skip := time.Duration(-14 * 24 * time.Hour)
	beginDate := now.Add(skip)
	endDate := now.Add(-24 * time.Hour)

	cnt, err := document.TxNumStat{}.Count()
	if err != nil {
		logger.Error("tx num statistic count ", logger.String("err", err.Error()))
	}

	if cnt != 0 {
		return
	}

	result, err := document.TxNumStat{}.QueryByDuration(beginDate, endDate)

	txNumStatList := []document.TxNumStat{}

	if err != nil {
		logger.Error("tx num statistic query by duration", logger.String("err", err.Error()))
	} else {

		var dayMap = make(map[string]int64)
		for _, r := range result {
			dayMap[r.Date] = int64(r.Num)
		}
		var tmp = beginDate
		for tmp.Unix() < endDate.Unix() {
			fmtTmp := utils.FmtTime(tmp, utils.DateFmtYYYYMMDD)
			txNumStatList = append(txNumStatList, document.TxNumStat{
				Date:       fmtTmp,
				Num:        dayMap[fmtTmp],
				CreateTime: time.Now(),
			})
			tmp = tmp.Add(24 * time.Hour)
		}

		if len(txNumStatList) == 0 {
			return
		}

		err := document.TxNumStat{}.InsertList(txNumStatList)
		if err != nil {
			logger.Error("txNumStatTask failed", logger.String("err", err.Error()))
		}
	}
}

type txNumGroup struct {
	Date string `bson:"_id"`
	Num  int64  `bson:"count"`
}
