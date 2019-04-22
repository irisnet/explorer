package document

import "github.com/irisnet/irishub-sync/store"

const (
	CollectionNmTokenFlow = "token_flow"
)

type TokenFlow struct {
	BlockHeight int64      `bson:"block_height" json:"block_height"`
	BlockHash   string     `bson:"block_hash" json:"block_hash"`
	TxHash      string     `bson:"tx_hash" json:"tx_hash"`
	From        string     `bson:"from" json:"from"`
	To          string     `bson:"to" json:"to"`
	Amount      store.Coin `bson:"amount" json:"amount"`
	Fee         store.Fee  `bson:"fee" json:"fee"`
	TxInitiator string     `bson:"tx_initiator" json:"tx_initiator"`
	TxType      string     `bson:"tx_type" json:"tx_type"`
	FlowType    string     `bson:"flow_type" json:"flow_type"`
	Status      string     `bson:"status" json:"status"`
	Timestamp   string     `bson:"timestamp" json:"timestamp"`
}
