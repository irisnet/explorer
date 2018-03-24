package store

import (
	"time"
	"github.com/cosmos/cosmos-sdk/modules/coin"
	"gopkg.in/mgo.v2"
)

const (
	DbCosmosTxn     = "cosmos_txn"
	DocsNmCoinTx    = "coin_tx"
	DocsNmStakeTx   = "stake_tx"
	DocsNmSyncBlock = "sync_block"
	PageSize        = 10
)

type DocsHander interface {
	Name() string
	KvPair() (string,string)
	Index() mgo.Index
}

type CoinTx struct {
	TxHash string `json:"tx_hash"`
	Time time.Time `json:"time"`
	Height int64 `json:"height"`
	From string `json:"from"`
	To string `json:"to"`
	Amount coin.Coins `json:"coins"`
}

type StakeTx struct {
	TxHash string `json:"tx_hash"`
	Time time.Time `json:"time"`
	Height int64 `json:"height"`
	From string `json:"from"`
	PubKey string `json:"pub_key"`
	Type string `json:"type"`
	Amount coin.Coin `json:"amount"`
}

type SyncBlock struct {
	ChainID string    `json:"chain_id"`
	Height  int64     `json:"height"`
	Time    time.Time `json:"time"`
}

func(c CoinTx) Name() string{
	return DocsNmCoinTx
}

func(c CoinTx) KvPair() (string,string){
	return "tx_hash",c.TxHash
}

func(c CoinTx) Index () mgo.Index{
	return mgo.Index{
		Key:        []string{"from"}, // 索引字段， 默认升序,若需降序在字段前加-
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

func(c SyncBlock) Name() string{
	return DocsNmSyncBlock
}

func(c SyncBlock) KvPair() (string,string){
	return "height",string(c.Height)
}

func(c SyncBlock) Index () mgo.Index{
	return mgo.Index{
		Key:        []string{"Height"}, // 索引字段， 默认升序,若需降序在字段前加-
		Unique:     false,             // 唯一索引 同mysql唯一索引
		DropDups:   false,             // 索引重复替换旧文档,Unique为true时失效
		Background: true,             // 后台创建索引
	}
}