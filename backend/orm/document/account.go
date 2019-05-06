package document

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionNmAccount          = "account"
	AccountFieldAddress          = "address"
	AccountFieldAccountNumber    = "account_number"
	AccountFieldCoinIrisUpdateAt = "coin_iris_update_at"
)

type Account struct {
	Address          string `bson:"address"`
	AccountNumber    uint64 `bson:"account_number"`
	CoinIris         Coin   `bson:"coin_iris"`
	CoinIrisUpdateAt int64  `bson:"coin_iris_update_at"`
}

func (a Account) Name() string {
	return CollectionNmAccount
}

func (a Account) PkKvPair() map[string]interface{} {
	return bson.M{AccountFieldAddress: a.Address}
}
