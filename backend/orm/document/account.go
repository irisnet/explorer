package document

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionNmAccount = "account"
	AccountFieldAddress = "address"
	AccountFieldBalance = "balance"
)

type Account struct {
	Address         string  `bson:"address"`
	Balance         float64 `bson:"balance"`
	BalanceUpdateAt int64   `bson:"balance_update_at"`
}

func (a Account) Name() string {
	return CollectionNmAccount
}

func (a Account) PkKvPair() map[string]interface{} {
	return bson.M{AccountFieldAddress: a.Address}
}
