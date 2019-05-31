package document

import (
	"github.com/irisnet/explorer/backend/utils"
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionNmAccount       = "account"
	AccountFieldAddress       = "address"
	AccountFieldAccountNumber = "account_number"
	AccountFieldTotalUpdateAt = "total_update_at"
)

type Account struct {
	Address                         string     `bson:"address"`
	AccountNumber                   uint64     `bson:"account_number"`
	Total                           utils.Coin `bson:"total"`
	TotalUpdateHeight               int64      `bson:"total_update_height"`
	TotalUpdateAt                   int64      `bson:"total_update_at"`
	CoinIris                        utils.Coin `bson:"coin_iris"`
	CoinIrisUpdateHeight            int64      `bson:"coin_iris_update_height"`
	CoinIrisUpdateAt                int64      `bson:"coin_iris_update_at"`
	Delegation                      utils.Coin `bson:"delegation"`
	DelegationUpdateHeight          int64      `bson:"delegation_update_height"`
	DelegationUpdateAt              int64      `bson:"delegation_update_at"`
	UnbondingDelegation             utils.Coin `bson:"unbonding_delegation"`
	UnbondingDelegationUpdateHeight int64      `bson:"unbonding_delegation_update_height"`
	UnbondingDelegationUpdateAt     int64      `bson:"unbonding_delegation_update_at"`
}

func (a Account) Name() string {
	return CollectionNmAccount
}

func (a Account) PkKvPair() map[string]interface{} {
	return bson.M{AccountFieldAddress: a.Address}
}
