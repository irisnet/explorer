package vo

import "github.com/irisnet/explorer/backend/utils"

type AssetTxInfo struct {
	BaseTx
	Msgs []MsgItem `json:"msgs"`
}
type GuardianTxInfo struct {
	BaseTx
	IsProfiler bool      `json:"isProfiler"`
	Msgs       []MsgItem `json:"msgs"`
}

type HtlcTxInfo struct {
	BaseTx
	FromMoniker  string    `json:"from_moniker"`
	ToMoniker    string    `json:"to_moniker"`
	ExpireHeight int64     `json:"expire_height,string"`
	Msgs         []MsgItem `json:"msgs"`
}

type CoinswapTxInfo struct {
	BaseTx
	Msgs []MsgItem `json:"msgs"`
}

type GovTxInfo struct {
	BaseTx
	Msgs []MsgItem `json:"msgs"`
}

type StakeTxInfo = StakeTx

type TransTxInfo struct {
	BaseTx
	Msgs []MsgItem `json:"msgs"`
}

type DeclarationTxInfo struct {
	BaseTx
	SelfBond utils.Coins `json:"self_bond"`
	Msgs     []MsgItem   `json:"msgs"`
}
