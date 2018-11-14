package model

import (
	"github.com/irisnet/irishub-sync/store"
	"github.com/irisnet/irishub-sync/store/document"
	"time"
)

type MsgSubmitProposal struct {
	Title          string      `json:"title"`          //  Title of the proposal
	Description    string      `json:"description"`    //  Description of the proposal
	ProposalType   string      `json:"proposalType"`   //  Type of proposal. Initial set {PlainTextProposal, SoftwareUpgradeProposal}
	Proposer       string      `json:"proposer"`       //  Address of the proposer
	InitialDeposit store.Coins `json:"initialDeposit"` //  Initial deposit paid by sender. Must be strictly positive.
	Params         []Param     `json:"params"`
}

type Param struct {
	Key   string `bson:"key"`
	Value string `bson:"value"`
	Op    string `bson:"op"`
}

type MsgDeposit struct {
	ProposalID int64       `json:"proposal_id"` // ID of the proposal
	Depositer  string      `json:"depositer"`   // Address of the depositer
	Amount     store.Coins `json:"amount"`      // Coins to add to the proposal's deposit
}

type MsgVote struct {
	ProposalID int64  `json:"proposal_id"`
	Voter      string `json:"voter"`
	Option     string `json:"option"`
}

type MsgBeginRedelegate struct {
	DelegatorAddr    string  `json:"delegator_addr"`
	ValidatorSrcAddr string  `json:"validator_src_addr"`
	ValidatorDstAddr string  `json:"validator_dst_addr"`
	SharesAmount     float64 `json:"shares_amount"`
}

type TxCounter struct {
	TransCnt       int
	StakeCnt       int
	DeclarationCnt int
	GovCnt         int
	Data           []document.CommonTx
}

type TxDay struct {
	Time  string `bson:"_id,omitempty"`
	Count int
}

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
