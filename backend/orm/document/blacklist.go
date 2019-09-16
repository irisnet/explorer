package document

import (
	"fmt"

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
	var (
		blackList []BlackList
	)
	var blackListMap = make(map[string]BlackList)

	projection := bson.M{
		"operator_addr": 1,
		"moniker":       1,
		"identity":      1,
		"website":       1,
		"details":       1,
	}
	if err := queryAll(CollectionNmBlackList, projection, nil, "", 100, &blackList); err == nil {
		for _, v := range blackList {
			blackListMap[v.OperatorAddr] = v
		}
	}
	return blackListMap
}
