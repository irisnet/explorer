package document

import (
	"fmt"

	"github.com/irisnet/explorer/backend/orm"
	"github.com/irisnet/explorer/backend/utils"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
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
	Rewards                         utils.Coin `bson:"rewards"`
}

func (a Account) String() string {
	return fmt.Sprintf(
		`
Address                         :%v
AccountNumber                   :%v
Total                           :%v
TotalUpdateHeight               :%v
TotalUpdateAt                   :%v
CoinIris                        :%v
CoinIrisUpdateHeight            :%v
CoinIrisUpdateAt                :%v
Delegation                      :%v
DelegationUpdateHeight          :%v
DelegationUpdateAt              :%v
UnbondingDelegation             :%v
UnbondingDelegationUpdateHeight :%v
UnbondingDelegationUpdateAt     :%v
`, a.Address, a.AccountNumber, a.Total, a.TotalUpdateHeight, a.TotalUpdateAt, a.CoinIris, a.CoinIrisUpdateHeight, a.CoinIrisUpdateAt, a.Delegation, a.DelegationUpdateAt, a.DelegationUpdateAt,
		a.UnbondingDelegation, a.UnbondingDelegationUpdateHeight, a.UnbondingDelegationUpdateAt)
}

func (a Account) GetAccountList() ([]Account, error) {
	var result []Account

	var query = orm.NewQuery()
	defer query.Release()

	query.SetCollection(CollectionNmAccount).
		SetSort(desc("total.amount"), AccountFieldTotalUpdateAt, AccountFieldAccountNumber).
		SetSize(100).
		SetResult(&result)
	err := query.Exec()
	return result, err
}

func (a Account) GetDelegatores(offset, size int) ([]Account, error) {
	var result []Account

	query := orm.NewQuery()
	defer query.Release()

	condition := bson.M{
		"delegation.amount": bson.M{
			"$gt": 0,
		},
	}

	query.SetCollection(CollectionNmAccount).SetCondition(condition).
		SetSort(desc("total.amount"), AccountFieldTotalUpdateAt, AccountFieldAccountNumber).
		SetOffset(offset).
		SetSize(size).
		SetResult(&result)
	err := query.Exec()
	return result, err
}

func (a Account) GetAllAccount() ([]Account, error) {
	var result []Account

	var query = orm.NewQuery()
	defer query.Release()

	query.SetCollection(CollectionNmAccount).
		SetSort(desc("total.amount"), AccountFieldTotalUpdateAt, AccountFieldAccountNumber).
		SetResult(&result)
	err := query.Exec()
	return result, err
}

func (a Account) Name() string {
	return CollectionNmAccount
}

func (a Account) PkKvPair() map[string]interface{} {
	return bson.M{AccountFieldAddress: a.Address}
}

func (a Account) Batch(txs []txn.Op) error {
	return orm.Batch(txs)
}

func (a Account) CountAll() (int, error) {
	query := orm.NewQuery()
	defer query.Release()
	return query.SetCollection(a.Name()).Count()
}

func (a Account) Update(account Account) error {
	query := orm.NewQuery()
	defer query.Release()
	c := query.GetDb().C(a.Name())
	return c.Update(account.PkKvPair(), account)
}

func (a Account) CountDelegatorNum() (int, error) {
	query := orm.NewQuery()
	defer query.Release()

	condition := bson.M{
		"delegation.amount": bson.M{
			"$gt": 0,
		},
	}
	return query.SetCollection(a.Name()).SetCondition(condition).Count()
}
