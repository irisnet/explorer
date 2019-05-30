package model

import (
	"fmt"
	"time"

	"github.com/irisnet/explorer/backend/orm/document"
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

type TxNumGroupByDayVo struct {
	Date string `json:"date"`
	Num  int64  `json:"num"`
}

type TxPage struct {
	Total int  `json:"total"`
	Items []Tx `json:"items"`
}

type ProposalPage struct {
	Total int            `json:"total"`
	Items []ProposalInfo `json:"items"`
}

type ProposalInfo struct {
	Hash          string             `json:"hash"`
	ActualFee     document.ActualFee `json:"actual_fee"`
	Signer        string             `json:"Signer"`
	TxType        string             `json:"tx_type"`
	Status        string             `json:"status"`
	ProposalId    uint64             `json:"proposal_id"`
	ProposalType  string             `json:"proposal_type"`
	ProposalTitle string             `json:"proposal_title"`
}

type Tx struct {
	Hash      string             `json:"hash"`
	From      string             `json:"from"`
	To        string             `json:"to"`
	Amount    document.Coins     `json:"amount"`
	ActualFee document.ActualFee `json:"actual_fee"`
	Signer    string             `json:"Signer"`
	Type      string             `json:"type"`
	Status    string             `json:"status"`
	Timestamp time.Time          `json:"timestamp"`
}

func (t Tx) PrintHashFromToAmount() string {

	return fmt.Sprintf(`
		Hash:   %v
		Type:   %v
		From:   %v
		To:     %v
		Amount: %v
		`, t.Hash, t.Type, t.From, t.To, t.Amount)
}

type BaseTx struct {
	Signer      string `json:"Signer,omitempty"`
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

func (s StakeTx) PrintHashFromToAmount() string {
	return fmt.Sprintf(`
Hash  : %v
From  : %v
To    : %v
Amount: %v
`, s.Hash, s.From, s.To, s.Amount)
}

type DeclarationTx struct {
	BaseTx
	Amount       document.Coins `json:"Amount"`
	OperatorAddr string         `json:"OperatorAddr"`
	Owner        string
	Moniker      string
	Pubkey       string
	Identity     string
	SelfBond     document.Coins
	Website      string
	Details      string
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

type RecentTx struct {
	Fee    Coin      `json:"actual_fee"`
	Time   time.Time `json:"time"`
	TxHash string    `json:"tx_hash"`
	Type   string    `json:"type"`
}

type Coin struct {
	Amount float64 `json:"amount"`
	Denom  string  `json:"denom"`
}
