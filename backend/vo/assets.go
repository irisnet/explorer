package vo

import (
	"time"
)

type AssetsRespond struct {
	Total     int        `json:"total"`
	Txs       []AssetsVo `json:"txs"`
	AssetType string     `json:"asset_type"`
}

type AssetsVo struct {
	TokenId         string    `json:"token_id"`
	Type            string    `json:"type"`
	Owner           string    `json:"owner"`
	Gateway         string    `json:"gateway"`
	Symbol          string    `json:"symbol"`
	InitialSupply   int64     `json:"initial_supply,string"`
	MaxSupply       int64     `json:"max_supply,string"`
	Mintable        bool      `json:"mintable,string"`
	Decimal         int32     `json:"decimal,string"`
	CanonicalSymbol string    `json:"canonical_symbol"`
	SymbolMin       string    `json:"symbol_min"`
	Name            string    `json:"name"`
	MintTo          string    `json:"mint_to"`
	Amount          int64     `json:"amount"`
	SrcOwner        string    `json:"src_owner"`
	DstOwner        string    `json:"dst_owner"`
	Height          int64     `json:"height"`
	TxHash          string    `json:"tx_hash"`
	TxFee           ActualFee `json:"tx_fee"`
	TxStatus        string    `json:"tx_status"`
	Timestamp       time.Time `json:"timestamp"`
}

type Coins []Coin

type Fee struct {
	Amount Coins `json:"amount"`
	Gas    int64 `json:"gas"`
}

type ActualFee struct {
	Denom  string  `json:"denom"`
	Amount float64 `json:"amount"`
}

type AssetGateways struct {
	Owner    string `json:"owner"`
	Moniker  string `json:"moniker"`
	Identity string `json:"identity"`
	Details  string `json:"details"`
	Website  string `json:"website"`
	Icons    string `json:"icons"`
}

type AssetTokens struct {
	TokenId         string         `json:"token_id"`
	Owner           string         `json:"owner"`
	Gateway         string         `json:"gateway"`
	Family          string         `json:"family"`
	TotalSupply     string         `json:"total_supply"`
	InitialSupply   string         `json:"initial_supply"`
	MaxSupply       string         `json:"max_supply"`
	Mintable        bool           `json:"mintable,string"`
	Decimal         int            `json:"decimal,string"`
	Symbol          string         `json:"symbol"`
	CanonicalSymbol string         `json:"canonical_symbol"`
	Name            string         `json:"name"`
	MinUnitAlias    string         `json:"min_unit_alias"`
	Source          string         `json:"source"`
	AssetGateway    *AssetGateways `json:"asset_gateway"`
}

type AssetTokensRespond []AssetTokens