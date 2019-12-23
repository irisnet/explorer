package vo

import "github.com/irisnet/explorer/backend/utils"

type AssetTxInfo struct {
	BaseTx
	Msgs     []MsgItem         `json:"msgs"`
	Monikers map[string]string `json:"monikers"`
}
type GuardianTxInfo struct {
	BaseTx
	IsProfiler bool              `json:"isProfiler"`
	Msgs       []MsgItem         `json:"msgs"`
	Monikers   map[string]string `json:"monikers"`
}

type HtlcTxInfo struct {
	BaseTx
	//FromMoniker  string            `json:"from_moniker"`
	//ToMoniker    string            `json:"to_moniker"`
	ExpireHeight int64             `json:"expire_height,string"`
	Msgs         []MsgItem         `json:"msgs"`
	Monikers     map[string]string `json:"monikers"`
}

type CoinswapTxInfo struct {
	BaseTx
	Msgs     []MsgItem         `json:"msgs"`
	Monikers map[string]string `json:"monikers"`
}

type GovTxInfo struct {
	BaseTx
	Msgs     []MsgItem         `json:"msgs"`
	Monikers map[string]string `json:"monikers"`
}

type StakeTxInfo = StakeTx

type TransTxInfo struct {
	BaseTx
	Msgs     []MsgItem         `json:"msgs"`
	Monikers map[string]string `json:"monikers"`
}

type DeclarationTxInfo struct {
	BaseTx
	SelfBond utils.Coins       `json:"self_bond"`
	Msgs     []MsgItem         `json:"msgs"`
	Monikers map[string]string `json:"monikers"`
}
