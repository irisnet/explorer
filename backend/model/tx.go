package model

import (
	"fmt"
	"time"

	"github.com/irisnet/explorer/backend/utils"
	"gopkg.in/mgo.v2/bson"
)

type MsgSubmitProposal struct {
	Title          string      `json:"title"`          //  Title of the proposal
	Description    string      `json:"description"`    //  Description of the proposal
	ProposalType   string      `json:"proposalType"`   //
	Proposer       string      `json:"proposer"`       //  Address of the proposer
	InitialDeposit utils.Coins `json:"initialDeposit"` //  Initial deposit paid by sender. Must be strictly positive.
	Params         []Param     `json:"params"`
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
	ProposalID uint64      `json:"proposal_id"` // ID of the proposal
	Depositer  string      `json:"depositer"`   // Address of the depositer
	Amount     utils.Coins `json:"amount"`      // Coins to add to the proposal's deposit
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
	TransCnt       int `json:"trans_cnt"`
	StakeCnt       int `json:"stake_cnt"`
	DeclarationCnt int `json:"declaration_cnt"`
	GovCnt         int `json:"gov_cnt"`
}

func (s TxStatisticsVo) String() string {
	return fmt.Sprintf(`
		TransCnt       :%v
		StakeCnt       :%v
		DeclarationCnt :%v
		GovCnt         :%v
		`, s.TransCnt, s.StakeCnt, s.DeclarationCnt, s.GovCnt)
}

type TxNumGroupByDayVo struct {
	Date string `json:"date"`
	Num  int64  `json:"num"`
}

func (txCount TxNumGroupByDayVo) String() string {
	return fmt.Sprintf("Date: %v Num: %v \n", txCount.Date, txCount.Num)
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
	Hash          string          `json:"hash"`
	ActualFee     utils.ActualFee `json:"actual_fee"`
	Signer        string          `json:"Signer"`
	TxType        string          `json:"tx_type"`
	Status        string          `json:"status"`
	ProposalId    uint64          `json:"proposal_id"`
	ProposalType  string          `json:"proposal_type"`
	ProposalTitle string          `json:"proposal_title"`
}

type Tx struct {
	Moniker   string          `json:"moniker"`
	Hash      string          `json:"hash"`
	From      string          `json:"from"`
	To        string          `json:"to"`
	Amount    utils.Coins     `json:"amount"`
	ActualFee utils.ActualFee `json:"actual_fee"`
	Signer    string          `json:"Signer"`
	Type      string          `json:"type"`
	Status    string          `json:"status"`
	Timestamp time.Time       `json:"timestamp"`
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
	Fee         utils.ActualFee
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
	Amount utils.Coins
}

type StakeTx struct {
	TransTx
	Source      string
	FromMoniker string `json:"from_moniker"`
	ToMoniker   string `json:"to_moniker"`
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
	Amount       utils.Coins `json:"Amount"`
	OperatorAddr string      `json:"OperatorAddr"`
	Owner        string
	Moniker      string
	Pubkey       string
	Identity     string
	SelfBond     utils.Coins
	Website      string
	Details      string
}

type GovTx struct {
	BaseTx
	From         string
	ProposalId   uint64
	Description  string
	Amount       utils.Coins
	Option       string
	Title        string
	ProposalType string
}

type RecentTx struct {
	Fee    utils.Coin `json:"actual_fee"`
	Time   time.Time  `json:"time"`
	TxHash string     `json:"tx_hash"`
	Type   string     `json:"type"`
}

func (tx RecentTx) String() string {
	return fmt.Sprintf(`
		Fee    :%v
		Time   :%v
		TxHash :%v
		Type   :%v
		`, tx.Fee, tx.Time, tx.TxHash, tx.Type)
}

type Signer struct {
	AddrHex    string `bson:"addr_hex"`
	AddrBech32 string `bson:"addr_bech32"`
}

func (s Signer) String() string {
	return fmt.Sprintf("AddrHex: %v   AddrBech32: %v \n", s.AddrHex, s.AddrBech32)
}

type CommonTx struct {
	Time                 time.Time            `bson:"time"`
	Height               int64                `bson:"height"`
	TxHash               string               `bson:"tx_hash"`
	From                 string               `bson:"from"`
	To                   string               `bson:"to"`
	Amount               utils.Coins          `bson:"amount"`
	Type                 string               `bson:"type"`
	Fee                  utils.Fee            `bson:"fee"`
	Memo                 string               `bson:"memo"`
	Status               string               `bson:"status"`
	Code                 uint32               `bson:"code"`
	Log                  string               `bson:"log"`
	GasUsed              int64                `bson:"gas_used"`
	GasPrice             float64              `bson:"gas_price"`
	ActualFee            utils.ActualFee      `bson:"actual_fee"`
	ProposalId           uint64               `bson:"proposal_id"`
	Tags                 map[string]string    `bson:"tags"`
	StakeCreateValidator StakeCreateValidator `bson:"stake_create_validator"`
	StakeEditValidator   StakeEditValidator   `bson:"stake_edit_validator"`
	Msg                  Msg                  `bson:"-"`
	Signers              []Signer             `bson:"signers"`
}

func (tx CommonTx) String() string {
	return fmt.Sprintf(`
		Time                 :%v
		Height               :%v
		TxHash               :%v
		From                 :%v
		To                   :%v
		Amount               :%v
		Type                 :%v
		Fee                  :%v
		Memo                 :%v
		Status               :%v
		Code                 :%v
		Log                  :%v
		GasUsed              :%v
		GasPrice             :%v
		ActualFee            :%v
		ProposalId           :%v
		Tags                 :%v
		StakeCreateValidator :%v
		StakeEditValidator   :%v
		Msg                  :%v
		Signers              :%v
		`, tx.Time, tx.Height, tx.TxHash, tx.From, tx.To, tx.Amount, tx.Type, tx.Fee, tx.Memo, tx.Status, tx.Code, tx.Log,
		tx.GasUsed, tx.GasPrice, tx.ActualFee, tx.ProposalId, tx.Tags, tx.StakeCreateValidator, tx.StakeEditValidator, tx.Msg, tx.Signers)

}

type Msg interface {
	Type() string
	String() string
}

type StakeCreateValidator struct {
	PubKey      string         `bson:"pub_key"`
	Description ValDescription `bson:"description"`
}

type StakeEditValidator struct {
	Description ValDescription `bson:"description"`
}

// Description
type ValDescription struct {
	Moniker  string `bson:"moniker"`
	Identity string `bson:"identity"`
	Website  string `bson:"website"`
	Details  string `bson:"details"`
}

func (tx CommonTx) PrintHashTypeFromToAmount() string {

	return fmt.Sprintf(`
	hash:   %v
	type:   %v
	from:   %v
	to:     %v
	amount: %v
		`, tx.TxHash, tx.Type, tx.From, tx.To, tx.Amount)
}

type CountVo struct {
	Id    bson.ObjectId `bson:"_id,omitempty"`
	Count float64
}
