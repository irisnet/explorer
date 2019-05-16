package document

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	CollectionTxNumStat        = "ex_tx_num_stat"
	TxNumStat_Field_Date       = "date"
	TxNumStat_Field_Num        = "num"
	TxNumStat_Field_CreateTime = "create_time"
)

type TxNumStat struct {
	Date       string    `bson:"date"`
	Num        int64     `bson:"num"`
	CreateTime time.Time `bson:"create_time"`
}

func (m TxNumStat) Name() string {
	return CollectionTxNumStat
}

func (m TxNumStat) PkKvPair() map[string]interface{} {
	return bson.M{TxNumStat_Field_Date: m.Date}
}
