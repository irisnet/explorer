package model

import (
	"github.com/irisnet/explorer/backend/orm/document"
	"time"
)

type MsgSubmitProposal struct {
	Title          string         `json:"title"`          //  Title of the proposal
	Description    string         `json:"description"`    //  Description of the proposal
	ProposalType   string         `json:"proposalType"`   //
	Proposer       string         `json:"proposer"`       //  Address of the proposer
	InitialDeposit document.Coins `json:"initialDeposit"` //  Initial deposit paid by sender. Must be strictly positive.
	Params         []Param        `json:"params"`
}

type MsgSubmitSoftwareUpgradeProposal struct {
	MsgSubmitProposal
	Version      uint64 `json:"version"`
	Software     string `json:"software"`
	SwitchHeight uint64 `json:"switch_height"`
	Threshold    string `json:"threshold"`
}

type Param struct {
	Subspace string `json:"subspace"`
	Key      string `json:"key"`
	Value    string `json:"value"`
}

type MsgDeposit struct {
	ProposalID uint64         `json:"proposal_id"` // ID of the proposal
	Depositer  string         `json:"depositer"`   // Address of the depositer
	Amount     document.Coins `json:"amount"`      // Coins to add to the proposal's deposit
}

type MsgVote struct {
	ProposalID uint64 `json:"proposal_id"`
	Voter      string `json:"voter"`
	Option     string `json:"option"`
}

type MsgBeginRedelegate struct {
	DelegatorAddr    string `json:"delegator_addr"`
	ValidatorSrcAddr string `json:"validator_src_addr"`
	ValidatorDstAddr string `json:"validator_dst_addr"`
	SharesAmount     string `json:"shares_amount"`
}

type TxStatisticsVo struct {
	TransCnt       int
	StakeCnt       int
	DeclarationCnt int
	GovCnt         int
	Data           []document.CommonTx
}

type TxDayVo struct {
	Time  string `bson:"_id,omitempty"`
	Count int
}

type BaseTx struct {
	Hash        string
	BlockHeight int64
	Type        string
	Fee         document.ActualFee
	Status      string
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
	Amount document.Coins
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
	SelfBond document.Coins
	Website  string
	Details  string
}

type GovTx struct {
	BaseTx
	From         string
	ProposalId   uint64
	Description  string
	Amount       document.Coins
	Option       string
	Title        string
	ProposalType string
}
