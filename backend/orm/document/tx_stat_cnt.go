package document

import (
	"fmt"
	"time"

	"github.com/irisnet/explorer/backend/orm"
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionTxNumStat        = "ex_tx_num_stat"
	TxNumStat_Field_Date       = "date"
	TxNumStat_Field_Num        = "num"
	TxNumStat_Field_CreateTime = "create_time"
)

type TxNumStat struct {
	Date         string    `bson:"date"`
	Num          int64     `bson:"num"`
	TotalAccNum  int64     `bson:"total_acc_num"`
	DelegatorNum int64     `bson:"delegator_num"`
	CreateTime   time.Time `bson:"create_time"`
}

func (m TxNumStat) Name() string {
	return CollectionTxNumStat
}

func (m TxNumStat) PkKvPair() map[string]interface{} {
	return bson.M{TxNumStat_Field_Date: m.Date}
}

func (self TxNumStat) Insert() error {
	db := orm.GetDatabase()
	txNumStatStore := db.C(CollectionTxNumStat)
	return txNumStatStore.Insert(self)
}

func (self TxNumStat) Count() (int, error) {
	db := orm.GetDatabase()
	defer db.Session.Close()

	txNumStatStore := db.C(CollectionTxNumStat)

	return txNumStatStore.Count()

}

func (_ TxNumStat) QueryByDuration(startTime, endTime time.Time) ([]TxNumGroup, error) {

	db := orm.GetDatabase()
	defer db.Session.Close()
	txStore := db.C(CollectionNmCommonTx)

	pipe := txStore.Pipe(
		[]bson.M{
			{"$match": bson.M{
				"time": bson.M{
					"$gte": startTime,
					"$lt":  endTime,
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

	result := []TxNumGroup{}
	err := pipe.All(&result)

	return result, err
}

func (_ TxNumStat) InsertList(list []TxNumStat) error {
	db := orm.GetDatabase()
	defer db.Session.Close()

	txNumStatStore := db.C(CollectionTxNumStat)

	return txNumStatStore.Insert(list)
}

type TxNumGroup struct {
	Date string `bson:"_id"`
	Num  int64  `bson:"count"`
}

func (group TxNumStat) String() string {
	return fmt.Sprintf(`
		Date :%v
		Num  :%v
		`, group.Date, group.Num)
}
