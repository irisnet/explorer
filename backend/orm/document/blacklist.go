package document

import (
	"fmt"

	"github.com/irisnet/explorer/backend/orm"
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionNmBlackList = "ex_val_black_list"

	BlackListFieldValAddr = "val_addr"
)

type BlackList struct {
	OperatorAddr string `bson:"operator_addr"`
	Moniker      string `bson:"moniker"`
	Identity     string `bson:"identity"`
	Website      string `bson:"website"`
	Details      string `bson:"details"`
}

func (b BlackList) String() string {
	return fmt.Sprintf(`

	OperatorAddr : %v
	Moniker      : %v
	Identity     : %v
	Website      : %v
	Details      : %v
`, b.OperatorAddr, b.Moniker, b.Identity, b.Website, b.Details)
}

func (d BlackList) Name() string {
	return CollectionNmBlackList
}

func (d BlackList) PkKvPair() map[string]interface{} {
	return bson.M{BlackListFieldValAddr: d.OperatorAddr}
}

func (b BlackList) QueryBlackList() map[string]BlackList {

	database := orm.NewQuery().GetDb()
	var blackListStore = database.C(CollectionNmBlackList)
	var blackList []BlackList
	var blackListMap = make(map[string]BlackList)
	if err := blackListStore.Find(nil).All(&blackList); err == nil {
		for _, v := range blackList {
			blackListMap[v.OperatorAddr] = v
		}
	}
	return blackListMap
}
