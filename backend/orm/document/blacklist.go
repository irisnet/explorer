package document

import "gopkg.in/mgo.v2/bson"

const (
	CollectionNmBlackList = "val_black_list"

	BlackListFieldValAddr = "val_addr"
)

type BlackList struct {
	OperatorAddr string `bson:"operator_addr"`
	Moniker      string `bson:"moniker"`
	Identity     string `bson:"identity"`
	Website      string `bson:"website"`
	Details      string `bson:"details"`
}

func (d BlackList) Name() string {
	return CollectionNmBlackList
}

func (d BlackList) PkKvPair() map[string]interface{} {
	return bson.M{BlackListFieldValAddr: d.OperatorAddr}
}
