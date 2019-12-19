package vo

import (
	"fmt"
	"time"

	"github.com/irisnet/explorer/backend/utils"
	"gopkg.in/mgo.v2/bson"
)

type Param struct {
	Subspace string `json:"subspace"`
	Key      string `json:"key"`
	Value    string `json:"value"`
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

type TxNumGroupByDayVoRespond []TxNumGroupByDayVo
type QueryTxTypeRespond []string

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
	Signer      string            `json:"signer,omitempty"`
	Hash        string            `json:"hash"`
	BlockHeight int64             `json:"block_height"`
	Type        string            `json:"type"`
	Fee         utils.ActualFee   `json:"fee"`
	Amount      utils.Coins       `json:"amount"`
	Status      string            `json:"status"`
	GasLimit    int64             `json:"gas_limit"`
	GasUsed     int64             `json:"gas_used"`
	GasWanted   int64             `json:"gas_wanted"`
	GasPrice    float64           `json:"gas_price"`
	Memo        string            `json:"memo"`
	Log         string            `json:"log"`
	Timestamp   time.Time         `json:"timestamp"`
	Tags        map[string]string `json:"tags"`
}

type TransTx struct {
	BaseTx
	From        string `json:"from"`
	To          string `json:"to"`
	FromMoniker string `json:"from_moniker"`
	ToMoniker   string `json:"to_moniker"`
	//Amount utils.Coins `json:"amount"`
}

type StakeTx struct {
	TransTx
	Source      string            `json:"source"`
	FromMoniker string            `json:"from_moniker"`
	ToMoniker   string            `json:"to_moniker"`
	Msgs        []MsgItem         `json:"msgs"`
	Monikers    map[string]string `json:"monikers"`
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
	OperatorAddr string      `json:"operator_addr"`
	Moniker      string      `json:"moniker"`
	SelfBond     utils.Coins `json:"self_bond"`
}

type GovTx struct {
	BaseTx
	From         string      `json:"from"`
	ProposalId   uint64      `json:"proposal_id"`
	Amount       utils.Coins `json:"amount"`
	Title        string      `json:"title"`
	ProposalType string      `json:"proposal_type"`
}

type RecentTx struct {
	Fee    utils.Coin `json:"actual_fee"`
	Time   time.Time  `json:"time"`
	TxHash string     `json:"tx_hash"`
	Type   string     `json:"type"`
}

type RecentTxRespond []RecentTx
type AssetTx struct {
	BaseTx
	From string `json:"from"`
	To   string `json:"to"`
	//Amount utils.Coins       `json:"amount"`
	Tags map[string]string `json:"tags"`
}
type GuardianTx struct {
	BaseTx
	From string `json:"from"`
	To   string `json:"to"`
	//Amount utils.Coins       `json:"amount"`
	Tags map[string]string `json:"tags"`
}

type HtlcTx struct {
	BaseTx
	From         string            `json:"from"`
	To           string            `json:"to"`
	FromMoniker  string            `json:"from_moniker"`
	ToMoniker    string            `json:"to_moniker"`
	ExpireHeight int64             `json:"expire_height,string"`
	Tags         map[string]string `json:"tags"`
}

type CoinswapTx struct {
	BaseTx
	From string `json:"from"`
	To   string `json:"to"`
	//Amount utils.Coins       `json:"amount"`
	Tags map[string]string `json:"tags"`
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
	AddrHex    string `json:"addr_hex"`
	AddrBech32 string `json:"addr_bech32"`
}

func (s Signer) String() string {
	return fmt.Sprintf("AddrHex: %v   AddrBech32: %v \n", s.AddrHex, s.AddrBech32)
}

type MsgItem struct {
	Type    string      `json:"type"`
	MsgData interface{} `json:"msg"`
}

type UdInfo struct {
	Source  string `json:"source"`
	Gateway string `json:"gateway"`
	Symbol  string `json:"symbol"`
}

type CommonTx struct {
	Time       time.Time         `json:"time"`
	Height     int64             `json:"height"`
	TxHash     string            `json:"tx_hash"`
	From       string            `json:"from"`
	To         string            `json:"to"`
	Amount     utils.Coins       `json:"amount"`
	Type       string            `json:"type"`
	Fee        utils.Fee         `json:"fee"`
	Memo       string            `json:"memo"`
	Status     string            `json:"status"`
	Code       uint32            `json:"code"`
	Log        string            `json:"log"`
	GasUsed    int64             `json:"gas_used"`
	GasWanted  int64             `json:"gas_wanted"`
	GasPrice   float64           `json:"gas_price"`
	ActualFee  utils.ActualFee   `json:"actual_fee"`
	ProposalId uint64            `json:"proposal_id"`
	Tags       map[string]string `json:"tags"`
	Msgs       []MsgItem         `json:"msgs"`
	Signers    []Signer          `json:"signers"`
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
		Msgs                 :%v
		Signers              :%v
		`, tx.Time, tx.Height, tx.TxHash, tx.From, tx.To, tx.Amount, tx.Type, tx.Fee, tx.Memo, tx.Status, tx.Code, tx.Log,
		tx.GasUsed, tx.GasPrice, tx.ActualFee, tx.ProposalId, tx.Tags, tx.Msgs, tx.Signers)

}

type Msg interface {
	Type() string
	String() string
}

type StakeCreateValidator struct {
	PubKey      string         `json:"pub_key"`
	Description ValDescription `json:"description"`
	Commission  CommissionMsg  `json:"commission"`
}

type CommissionMsg struct {
	Rate          string `json:"rate"`            // the commission rate charged to delegators
	MaxRate       string `json:"max_rate"`        // maximum commission rate which validator can ever charge
	MaxChangeRate string `json:"max_change_rate"` // maximum daily increase of the validator commission
}

type StakeEditValidator struct {
	Description    ValDescription `json:"description"`
	CommissionRate string         `json:"commission_rate"`
}

// Description
type ValDescription struct {
	Moniker  string `json:"moniker"`
	Identity string `json:"identity"`
	Website  string `json:"website"`
	Details  string `json:"details"`
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
	Id    bson.ObjectId `json:"_id,omitempty"`
	Count float64
}
