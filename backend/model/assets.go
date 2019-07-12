package model

import (
	"time"
)

type AssetsRespond struct {
	Total  int        `json:"total"`
	Assets []AssetsVo `json:"assets"`
	//TokenType string     `json:"token_type"`
	AssetType string `json:"asset_type"`
}

type AssetsVo struct {
	TokenId        string    `json:"token_id"`
	Type           string    `json:"type"`
	Owner          string    `json:"owner"`
	Gateway        string    `json:"gateway,omitempty"`
	Symbol         string    `json:"symbol"`
	InitialSupply  int64     `json:"initial_supply,string"`
	MaxSupply      int64     `json:"max_supply,string"`
	Mintable       bool      `json:"mintable,string"`
	Decimal        int32     `json:"decimal,string"`
	SymbolAtSource string    `json:"symbol_at_source,omitempty"`
	SymbolMin      string    `json:"symbol_min"`
	Name           string    `json:"name"`
	MintTo         string    `json:"mint_to"`
	Amount         Coins     `json:"amount"`
	SrcOwner       string    `json:"src_owner"`
	DstOwner       string    `json:"dst_owner"`
	Height         int64     `json:"height"`
	TxHash         string    `json:"tx_hash"`
	TxFee          Fee       `json:"tx_fee"`
	TxStatus       string    `json:"tx_status"`
	Timestamp      time.Time `json:"timestamp"`
}

type Coins []Coin

type Fee struct {
	Amount Coins `json:"amount"`
	Gas    int64 `json:"gas"`
}

//
//type IssueTokenVo struct {
//	Owner          string    `json:"owner"`
//	Gateway        string    `json:"gateway,omitempty"`
//	Symbol         string    `json:"symbol"`
//	InitialSupply  string    `json:"initial_supply"`
//	MaxSupply      string    `json:"max_supply"`
//	Mintable       string    `json:"mintable"`
//	Decimal        string    `json:"decimal"`
//	SymbolAtSource string    `json:"symbol_at_source,omitempty"`
//	SymbolMin      string    `json:"symbol_min"`
//	Name           string    `json:"name"`
//	Height         int64     `json:"height"`
//	TxHash         string    `json:"tx_hash"`
//	TxFee          Coin      `json:"tx_fee"`
//	TxStatus       string    `json:"tx_status"`
//	Timestamp      time.Time `json:"timestamp"`
//}
//
//type EditTokenVo struct {
//	TokenId        string    `json:"token_id"`
//	Owner          string    `json:"owner"`
//	MaxSupply      string    `json:"max_supply"`
//	Mintable       string    `json:"mintable"`
//	Decimal        string    `json:"decimal"`
//	SymbolAtSource string    `json:"symbol_at_source,omitempty"`
//	SymbolMin      string    `json:"symbol_min"`
//	Name           string    `json:"name"`
//	Height         int64     `json:"height"`
//	TxHash         string    `json:"tx_hash"`
//	TxFee          Coin      `json:"tx_fee"`
//	TxStatus       string    `json:"tx_status"`
//	Timestamp      time.Time `json:"timestamp"`
//}
//
//type MintTokenVo struct {
//	TokenId   string    `json:"token_id"`
//	Owner     string    `json:"owner"`
//	MintTo    string    `json:"mint_to"`
//	Height    int64     `json:"height"`
//	Amount    Coin      `json:"amount"`
//	TxHash    string    `json:"tx_hash"`
//	TxFee     Coin      `json:"tx_fee"`
//	TxStatus  string    `json:"tx_status"`
//	Timestamp time.Time `json:"timestamp"`
//}
//
//type TransferOwnerVo struct {
//	TokenId   string    `json:"token_id"`
//	SrcOwner  string    `json:"src_owner"`
//	DstOwner  string    `json:"dst_owner"`
//	Height    int64     `json:"height"`
//	TxHash    string    `json:"tx_hash"`
//	TxFee     Coin      `json:"tx_fee"`
//	TxStatus  string    `json:"tx_status"`
//	Timestamp time.Time `json:"timestamp"`
//}
