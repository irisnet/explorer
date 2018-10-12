package types

import (
	"github.com/irisnet/irishub-sync/store"
	"time"
)

type BaseTx struct {
	Hash        string
	BlockHeight int64
	Type        string
	Fee         store.ActualFee
	GasLimit    int64
	GasUsed     int64
	GasPrice    float64
	Memo        string
	Timestamp   time.Time
}

type TransTx struct {
	BaseTx
	From   string
	To     string
	Amount store.Coins
}

type StakeTx struct {
	TransTx
	Source string
}

type DeclarationTx struct {
	BaseTx
	Owner    string
	Moniker  string
	Pubkey   string
	Identity string
	SelfBond store.Coins
	Website  string
	Details  string
}

type GovTx struct {
	BaseTx
	From         string
	ProposalId   int64
	Description  string
	Amount       store.Coins
	Option       string
	Title        string
	ProposalType string
}
