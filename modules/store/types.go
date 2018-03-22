package store

import (
	"time"
	"github.com/cosmos/cosmos-sdk/modules/coin"
)

const (
	DbCosmosTxn = "cosmos_txn"
	TbNmCoinTx = "coin_tx"
	TbNmStakeTx = "stake_tx"
	TbNmSyncBlock = "sync_block"
	PageSize = 10
)

type TxHander interface {
	TbNm() string
	KvPair() (string,string)
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

func(c CoinTx) TbNm() string{
	return TbNmCoinTx
}

func(c CoinTx) KvPair() (string,string){
	return "tx_hash",c.TxHash
}

func(c StakeTx) TbNm() string{
	return TbNmStakeTx
}

func(c StakeTx) KvPair() (string,string){
	return "tx_hash",c.TxHash
}

func(c SyncBlock) TbNm() string{
	return TbNmSyncBlock
}

func(c SyncBlock) KvPair() (string,string){
	return "height",string(c.Height)
}