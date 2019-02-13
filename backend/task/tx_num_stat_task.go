package task

import (
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/utils"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type TxNumStatTask struct{}

func (task TxNumStatTask) Name() string {
	return "tx_num_stat_task"
}
func (task TxNumStatTask) Start() {
	task.init()
	utils.RunTimer(24*time.Hour, utils.Day, func() {
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

func (task TxNumStatTask) init() {
	db := orm.GetDatabase()
	defer db.Session.Close()
	var txNumStatList []interface{}

	today := utils.TruncateTime(time.Now(), utils.Day)
	yesterday := today.Add(-24 * time.Hour)

	txNumStatStore := db.C(document.CollectionTxNumStat)
	txStore := db.C(document.CollectionNmCommonTx)
	if cnt, _ := txNumStatStore.Count(); cnt == 0 {
		blockStore := db.C(document.CollectionNmBlock)
		var block document.Block
		err := blockStore.Find(bson.M{document.Block_Field_Height: int64(1)}).Select(bson.M{"time": 1}).One(&block)
		if err != nil {
			panic("block not exist with height equal 1")
		}
		beginDate := utils.TruncateTime(block.Time, utils.Day)
		endDate := yesterday

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

		var result []model.TxDayVo

		if err := pipe.All(&result); err == nil {
			var dayMap = make(map[string]int64)
			for _, r := range result {
				dayMap[r.Time] = int64(r.Count)
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

			if err := txNumStatStore.Insert(txNumStatList...); err != nil {
				logger.Error("txNumStatTask failed", logger.String("err", err.Error()))
			}
		}
	}
	return
}
