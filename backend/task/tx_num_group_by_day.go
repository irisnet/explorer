package task

import (
	"fmt"
	"time"

	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/utils"
)

type TxNumGroupByDayTask struct{}

func (task TxNumGroupByDayTask) Name() string {
	return "tx_num_stat_task"
}
func (task TxNumGroupByDayTask) Start() {
	task.init()
	utils.RunTimer(conf.Get().Server.CronTimeTxNumByDay, utils.Sec, func() {
		if err := task.DoTask(); err != nil {
			logger.Error(fmt.Sprintf("%s fail", task.Name()), logger.String("err", err.Error()))
		} else {
			logger.Info(fmt.Sprintf("%s success", task.Name()))
		}
	})
}

func (task TxNumGroupByDayTask) DoTask() error {
	today := utils.TruncateTime(time.Now(), utils.Day)
	yesterday := today.Add(-24 * time.Hour)

	total, err := document.CommonTx{}.GetTxCountByDuration(yesterday, today)
	if err != nil {
		return err
	}

	txNumStat := document.TxNumStat{
		Date:       utils.FmtTime(yesterday, utils.DateFmtYYYYMMDD),
		Num:        int64(total),
		CreateTime: time.Now(),
	}

	if err := txNumStat.Insert(); err != nil {
		return err
	}

	return nil
}

// init ex_tx_num_stat document
func (task TxNumGroupByDayTask) init() {

	now := utils.TruncateTime(time.Now(), utils.Day)
	skip := time.Duration(-14 * 24 * time.Hour)
	beginDate := now.Add(skip)
	endDate := now.Add(-24 * time.Hour)

	fmt.Println(now, beginDate, endDate)

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
