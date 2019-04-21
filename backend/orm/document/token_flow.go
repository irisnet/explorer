package document

import "github.com/irisnet/irishub-sync/store"

const (
	CollectionNmTokenFlow = "token_flow"
)

type TokenFlow struct {
	BlockHeight int64      `bson:"block_height"`
	BlockHash   string     `bson:"block_hash"`
	TxHash      string     `bson:"tx_hash"`
	From        string     `bson:"from"`
	To          string     `bson:"to"`
	Amount      store.Coin `bson:"amount"`
	Fee         store.Fee  `bson:"fee"`
	TxInitiator string     `bson:"tx_initiator"`
	TxType      string     `bson:"tx_type"`
	FlowType    string     `bson:"flow_type"`
	Status      string     `bson:"status"`
	Timestamp   string     `bson:"timestamp"`
}
