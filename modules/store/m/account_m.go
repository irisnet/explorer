package m

import (
	"github.com/cosmos/cosmos-sdk/modules/coin"
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"errors"
	"github.com/irisnet/irisplorer.io/modules/store"
)

const (
	DocsNmAccount = "account"

)

//账户信息
type Account struct {
	Address string     `json:"address" bson:"address"`
	Amount  coin.Coins `json:"amount" bson:"amount"`
	Time    time.Time  `json:"time" bson:"time"`
	Height  int64      `json:"height" bson:"height"`
}

func (a Account) Name() string {
	return DocsNmAccount
}

func (a Account) PkKvPair() (string, interface{}) {
	return "address", a.Address
}

func (a Account) Index() mgo.Index {
	return mgo.Index{
		Key:        []string{"address"}, // 索引字段， 默认升序,若需降序在字段前加-
		Unique:     true,               // 唯一索引 同mysql唯一索引
		DropDups:   false,               // 索引重复替换旧文档,Unique为true时失效
		Background: true,                // 后台创建索引
	}
}

//Account
func QueryAccount(address string) (Account, error) {
	var result Account
	query := func(c *mgo.Collection) error {
		err := c.Find(bson.M{"address": address}).Sort("-amount").One(&result)
		return err
	}

	if store.ExecCollection(DocsNmAccount, query) != nil {
		log.Printf("Account is Empry")
		return result, errors.New("Account is Empry")
	}

	return result, nil
}

func QueryAccountByPage(page int) []Account {
	result := []Account{}
	query := func(c *mgo.Collection) error {
		skip := (page - 1) * PageSize
		err := c.Find(nil).Sort("-amount").Skip(skip).Limit(PageSize).All(&result)
		return err
	}

	if store.ExecCollection(DocsNmAccount, query) != nil {
		log.Printf("Account is Empry")
	}
	return result
}
