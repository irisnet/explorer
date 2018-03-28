package store

import (
	"time"
	"github.com/cosmos/cosmos-sdk/modules/coin"
	"gopkg.in/mgo.v2"
)

const (
	DbIrisExp       = "iris-exp"
	DocsNmCoinTx    = "coin_tx"
	DocsNmStakeTx   = "stake_tx"
	DocsNmSyncTask  = "sync_task"
	DocsNmAccount   = "account"
	DocsNmDelegator = "Delegator"
	PageSize        = 10
)

type DocsHander interface {
	Name() string
	KvPair() (string,string)
	Index() mgo.Index
}

//Coin交易
type CoinTx struct {
	TxHash string `json:"tx_hash" bson:"tx_hash"`
	Time time.Time `json:"time" bson:"time"`
	Height int64 `json:"height" bson:"height"`
	From string `json:"from" bson:"from"`
	To string `json:"to" bson:"to"`
	Amount coin.Coins `json:"coins" bson:"amount"`
}
//Stake交易
type StakeTx struct {
	TxHash string `json:"tx_hash" bson:"tx_hash"`
	Time time.Time `json:"time" bson:"time"`
	Height int64 `json:"height" bson:"height"`
	From string `json:"from" bson:"from"`
	PubKey string `json:"pub_key" bson:"pub_key"`
	Type string `json:"type" bson:"type"`
	Amount coin.Coin `json:"amount" bson:"amount"`
}
//同步信息
type SyncTask struct {
	ChainID string    `json:"chain_id" bson:"chain_id"`
	Height  int64     `json:"height" bson:"height"`
	Time    time.Time `json:"time" bson:"time"`
}
//账户信息
type Account struct {
	Address string    `json:"address" bson:"address"`
	Amount coin.Coins `json:"amount" bson:"amount"`
	Time    time.Time `json:"time" bson:"time"`
	Height  int64     `json:"height" bson:"height"`
}

type Delegator struct {
	Address string    `json:"address" bson:"address"`
	PubKey string `json:"pub_key" bson:"pub_key"`
	Shares  int64 `json:"shares" bson:"shares"`
}

func(c CoinTx) Name() string{
	return DocsNmCoinTx
}

func(c CoinTx) KvPair() (string,string){
	return "tx_hash",c.TxHash
}

func(c CoinTx) Index () mgo.Index{
	return mgo.Index{
		Key:        []string{"from","to"}, // 索引字段， 默认升序,若需降序在字段前加-
		Unique:     false,             // 唯一索引 同mysql唯一索引
		DropDups:   false,             // 索引重复替换旧文档,Unique为true时失效
		Background: true,             // 后台创建索引
	}
}

func(c StakeTx) Name() string{
	return DocsNmStakeTx
}

func(c StakeTx) KvPair() (string,string){
	return "tx_hash",c.TxHash
}

func(c StakeTx) Index () mgo.Index{
	return mgo.Index{
		Key:        []string{"from"}, // 索引字段， 默认升序,若需降序在字段前加-
		Unique:     false,             // 唯一索引 同mysql唯一索引
		DropDups:   false,             // 索引重复替换旧文档,Unique为true时失效
		Background: true,             // 后台创建索引
	}
}

func(c SyncTask) Name() string{
	return DocsNmSyncTask
}

func(c SyncTask) KvPair() (string,string){
	return "chain_id",string(c.ChainID)
}

func(c SyncTask) Index () mgo.Index{
	return mgo.Index{
		Key:        []string{"chain_id"}, // 索引字段， 默认升序,若需降序在字段前加-
		Unique:     false,             // 唯一索引 同mysql唯一索引
		DropDups:   false,             // 索引重复替换旧文档,Unique为true时失效
		Background: true,             // 后台创建索引
	}
}

func(a Account) Name() string{
	return DocsNmAccount
}

func(a Account) KvPair() (string,string){
	return "address",a.Address
}

func(a Account) Index () mgo.Index{
	return mgo.Index{
		Key:        []string{"address"}, // 索引字段， 默认升序,若需降序在字段前加-
		Unique:     false,             // 唯一索引 同mysql唯一索引
		DropDups:   false,             // 索引重复替换旧文档,Unique为true时失效
		Background: true,             // 后台创建索引
	}
}

func(d Delegator) Name() string{
	return DocsNmDelegator
}

func(d Delegator) KvPair() (string,string){
	return "address",d.Address
}

func(d Delegator) Index () mgo.Index{
	return mgo.Index{
		Key:        []string{"address"}, // 索引字段， 默认升序,若需降序在字段前加-
		Unique:     false,             // 唯一索引 同mysql唯一索引
		DropDups:   false,             // 索引重复替换旧文档,Unique为true时失效
		Background: true,             // 后台创建索引
	}
}