package document

import "time"

const (
	CollectionNmTokenFlow = "token_flow"
)

type TokenFlow struct {
	BlockHeight int64     `bson:"block_height" json:"block_height"`
	BlockHash   string    `bson:"block_hash" json:"block_hash"`
	TxHash      string    `bson:"tx_hash" json:"tx_hash"`
	From        string    `bson:"from" json:"from"`
	To          string    `bson:"to" json:"to"`
	Amount      Coin      `bson:"amount" json:"amount"`
	Fee         Fee       `bson:"fee" json:"fee"`
	TxInitiator string    `bson:"tx_initiator" json:"tx_initiator"`
	TxType      string    `bson:"tx_type" json:"tx_type"`
	FlowType    string    `bson:"flow_type" json:"flow_type"`
	Status      string    `bson:"status" json:"status"`
	Timestamp   time.Time `bson:"timestamp" json:"timestamp"`
}
