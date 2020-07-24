package vo

import "time"

type AssetsRespond struct {
	Total     int        `json:"total"`
	Txs       []AssetsVo `json:"txs"`
	AssetType string     `json:"asset_type"`
}

type AssetsVo struct {
	Type          string    `json:"type"`
	Owner         string    `json:"owner"`
	Symbol        string    `json:"symbol"`
	InitialSupply int64     `json:"initial_supply,string"`
	MaxSupply     int64     `json:"max_supply,string"`
	Mintable      bool      `json:"mintable,string"`
	SymbolMin     string    `json:"symbol_min"`
	Name          string    `json:"name"`
	Decimal       uint32    `json:"scale"`
	MintTo        string    `json:"mint_to"`
	Amount        int64     `json:"amount"`
	SrcOwner      string    `json:"src_owner"`
	DstOwner      string    `json:"dst_owner"`
	Height        int64     `json:"height"`
	TxHash        string    `json:"tx_hash"`
	TxFee         ActualFee `json:"tx_fee"`
	TxStatus      string    `json:"tx_status"`
	Timestamp     time.Time `json:"timestamp"`
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

//type AssetGateways struct {
//	Owner    string `json:"owner"`
//	Moniker  string `json:"moniker"`
//	Identity string `json:"identity"`
//	Details  string `json:"details"`
//	Website  string `json:"website"`
//	Icons    string `json:"icons"`
//}

type AssetTokens struct {
	Owner         string `json:"owner"`
	TotalSupply   string `json:"total_supply"`
	InitialSupply string `json:"initial_supply"`
	MaxSupply     string `json:"max_supply"`
	Mintable      bool   `json:"mintable,string"`
	Symbol        string `json:"symbol"`
	Name          string `json:"name"`
	MinUnitAlias  string `json:"min_unit_alias"`
	Decimal       int    `json:"decimal"`
}

type AssetTokensRespond []AssetTokens
