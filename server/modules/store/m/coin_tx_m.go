package m

import (
	"time"
	"github.com/cosmos/cosmos-sdk/modules/coin"
	"github.com/irisnet/irisplorer.io/server/modules/store"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

const (
	DocsNmCoinTx = "coin_tx"
)

//Coin交易
type CoinTx struct {
	TxHash string     `json:"tx_hash" bson:"tx_hash"`
	Time   time.Time  `json:"time" bson:"time"`
	Height int64      `json:"height" bson:"height"`
	From   string     `json:"from" bson:"from"`
	To     string     `json:"to" bson:"to"`
	Amount coin.Coins `json:"coins" bson:"amount"`
}

func (c CoinTx) Name() string {
	return DocsNmCoinTx
}

func (c CoinTx) PkKvPair() (map[string]interface{}) {
	return bson.M{"tx_hash":c.TxHash}
}

func (c CoinTx) Index() mgo.Index {
	return mgo.Index{
		Key:        []string{"from", "to"}, // 索引字段， 默认升序,若需降序在字段前加-
		Unique:     false,                  // 唯一索引 同mysql唯一索引
		DropDups:   false,                  // 索引重复替换旧文档,Unique为true时失效
		Background: true,                   // 后台创建索引
	}
}

func QueryAllCoinTx() []CoinTx {
	result := []CoinTx{}
	query := func(c *mgo.Collection) error {
		err := c.Find(nil).Sort("-time").All(&result)
		return err
	}

	if store.ExecCollection(DocsNmCoinTx, query) != nil {
		log.Printf("CoinTx is Empry")
	}
	return result

}

//
func QueryCoinTxsByAccount(account string) ([]CoinTx) {
	result := []CoinTx{}
	query := func(c *mgo.Collection) error {
		err := c.Find(bson.M{"$or": []bson.M{bson.M{"from": account}, bson.M{"to": account}}}).Sort("-time").All(&result)
		return err
	}

	if store.ExecCollection(DocsNmCoinTx, query) != nil {
		log.Printf("CoinTx is Empry")
	}
	return result
}

//
func QueryPageCoinTxsByAccount(from string, page,pageSize int ) ([]CoinTx) {
	result := []CoinTx{}
	query := func(c *mgo.Collection) error {
		skip := (page - 1) * pageSize
		err := c.Find(bson.M{"from": from}).Sort("-time").Skip(skip).Limit(pageSize).All(&result)
		return err
	}

	if store.ExecCollection(DocsNmCoinTx, query) != nil {
		log.Printf("CoinTx is Empry")
	}

	return result
}
func QueryAllPageCoinTxs(page,pageSize int ) ([]CoinTx) {
	result := []CoinTx{}
	query := func(c *mgo.Collection) error {
		skip := (page - 1) * pageSize
		err := c.Find(nil).Sort("-time").Skip(skip).Limit(pageSize).All(&result)
		return err
	}

	if store.ExecCollection(DocsNmCoinTx, query) != nil {
		log.Printf("CoinTx is Empry")
	}

	return result
}
