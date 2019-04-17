package task

import (
	"fmt"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/utils"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type TxNumGroupByDayTask struct{}

func (task TxNumGroupByDayTask) Name() string {
	return "tx_num_stat_task"
}
func (task TxNumGroupByDayTask) Start() {
	task.init()
	utils.RunTimer(1, utils.Day, func() {
		db := orm.GetDatabase()
		defer db.Session.Close()

		today := utils.TruncateTime(time.Now(), utils.Day)
		yesterday := today.Add(-24 * time.Hour)

		txNumStatStore := db.C(document.CollectionTxNumStat)
		txStore := db.C(document.CollectionNmCommonTx)

		query := bson.M{}
		query["time"] = bson.M{"$gte": yesterday, "$lt": today}

		if cnt, err := txStore.Find(query).Count(); err == nil {
			txNumStat := document.TxNumStat{
				Date:       utils.FmtTime(yesterday, utils.DateFmtYYYYMMDD),
				Num:        int64(cnt),
				CreateTime: time.Now(),
			}

			if err := txNumStatStore.Insert(txNumStat); err != nil {
				logger.Error("txNumStatTask failed", logger.String("err", err.Error()))
			}
		}
	})
}

// init ex_tx_num_stat document
func (task TxNumGroupByDayTask) init() {
	db := orm.GetDatabase()
	defer db.Session.Close()
	var txNumStatList []interface{}

	now := utils.TruncateTime(time.Now(), utils.Day)

	skip := time.Duration(-14 * 24 * time.Hour)
	beginDate := now.Add(skip)
	endDate := now.Add(-24 * time.Hour)

	fmt.Println(now, beginDate, endDate)

	txNumStatStore := db.C(document.CollectionTxNumStat)
	txStore := db.C(document.CollectionNmCommonTx)
	if cnt, _ := txNumStatStore.Count(); cnt == 0 {
		pipe := txStore.Pipe(
			[]bson.M{
				{"$match": bson.M{
					"time": bson.M{
						"$gte": beginDate,
						"$lt":  endDate,
					},
				}},
				{"$project": bson.M{
					"day": bson.M{"$substr": []interface{}{"$time", 0, 10}},
				}},
				{"$group": bson.M{
					"_id":   "$day",
					"count": bson.M{"$sum": 1},
				}},
				{"$sort": bson.M{
					"_id": 1,
				}},
			},
		)

		var result []txNumGroup

		if err := pipe.All(&result); err == nil {
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

			if err := txNumStatStore.Insert(txNumStatList...); err != nil {
				logger.Error("txNumStatTask failed", logger.String("err", err.Error()))
			}
		}
	}
	return
}

type txNumGroup struct {
	Date string `bson:"_id"`
	Num  int64  `bson:"count"`
}
