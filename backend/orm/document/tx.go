package document

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionNmCommonTx = "tx_common"
	TxStatusSuccess      = "success"
	TxStatusFail         = "fail"

	Tx_Field_Time                 = "time"
	Tx_Field_Height               = "height"
	Tx_Field_Hash                 = "tx_hash"
	Tx_Field_From                 = "from"
	Tx_Field_To                   = "to"
	Tx_Field_Amount               = "amount"
	Tx_Field_Type                 = "type"
	Tx_Field_Fee                  = "fee"
	Tx_Field_Memo                 = "memo"
	Tx_Field_Status               = "status"
	Tx_Field_Code                 = "code"
	Tx_Field_Log                  = "log"
	Tx_Field_GasUsed              = "gas_used"
	Tx_Field_GasPrice             = "gas_price"
	Tx_Field_ActualFee            = "actual_fee"
	Tx_Field_ProposalId           = "proposal_id"
	Tx_Field_Tags                 = "tags"
	Tx_Field_StakeCreateValidator = "stake_create_validator"
	Tx_Field_StakeEditValidator   = "stake_edit_validator"
)

type CommonTx struct {
	Time       time.Time         `bson:"time"`
	Height     int64             `bson:"height"`
	TxHash     string            `bson:"tx_hash"`
	From       string            `bson:"from"`
	To         string            `bson:"to"`
	Amount     Coins             `bson:"amount"`
	Type       string            `bson:"type"`
	Fee        Fee               `bson:"fee"`
	Memo       string            `bson:"memo"`
	Status     string            `bson:"status"`
	Code       uint32            `bson:"code"`
	Log        string            `bson:"log"`
	GasUsed    int64             `bson:"gas_used"`
	GasPrice   float64           `bson:"gas_price"`
	ActualFee  ActualFee         `bson:"actual_fee"`
	ProposalId uint64            `bson:"proposal_id"`
	Tags       map[string]string `bson:"tags"`

	StakeCreateValidator StakeCreateValidator `bson:"stake_create_validator"`
	StakeEditValidator   StakeEditValidator   `bson:"stake_edit_validator"`
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

// Description
type ValDescription struct {
	Moniker  string `bson:"moniker"`
	Identity string `bson:"identity"`
	Website  string `bson:"website"`
	Details  string `bson:"details"`
}

type StakeCreateValidator struct {
	PubKey      string         `bson:"pub_key"`
	Description ValDescription `bson:"description"`
}

type StakeEditValidator struct {
	Description ValDescription `bson:"description"`
}

func (d CommonTx) Name() string {
	return CollectionNmCommonTx
}

func (d CommonTx) PkKvPair() map[string]interface{} {
	return bson.M{Tx_Field_Hash: d.TxHash}
}
