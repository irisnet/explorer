package document

import (
	"fmt"

	"github.com/irisnet/explorer/backend/orm"
	"github.com/irisnet/explorer/backend/utils"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
	"time"
)

const (
	CollectionNmAccount       = "account"
	AccountFieldAddress       = "address"
	AccountFieldAccountNumber = "account_number"
	AccountFieldUpdateAt      = "update_at"
)

type Account struct {
	Address             string     `bson:"address"`
	AccountNumber       uint64     `bson:"account_number"`
	Total               utils.Coin `bson:"total"`
	CoinIris            utils.Coin `bson:"coin_iris"`
	Delegation          utils.Coin `bson:"delegation"`
	UnbondingDelegation utils.Coin `bson:"unbonding_delegation"`
	Rewards             utils.Coin `bson:"rewards"`
	UpdateAt            int64      `bson:"update_at"`
	CreateAt            int64      `bson:"create_at"`
}

func (a Account) String() string {
	return fmt.Sprintf(
		`
Address                         :%v
AccountNumber                   :%v
Total                           :%v
CoinIris                        :%v
Delegation                      :%v
UnbondingDelegation             :%v
UpdateAt                        :%v
CreateAt                        :%v
`, a.Address, a.AccountNumber, a.Total, a.CoinIris, a.Delegation, a.UnbondingDelegation, a.UpdateAt, a.CreateAt)
}

func (a Account) GetAccountList() ([]Account, error) {
	var result []Account

	var query = orm.NewQuery()
	defer query.Release()

	query.SetCollection(CollectionNmAccount).
		SetSort(desc("total.amount"), AccountFieldUpdateAt, AccountFieldAccountNumber).
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
		SetSort(desc("total.amount"), AccountFieldUpdateAt, AccountFieldAccountNumber).
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
		SetSort(desc("total.amount"), AccountFieldUpdateAt, AccountFieldAccountNumber).
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
	account.UpdateAt = time.Now().Unix()
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
