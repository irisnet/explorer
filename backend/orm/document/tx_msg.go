package document

import "gopkg.in/mgo.v2/bson"

const (
	CollectionNmTxMsg   = "tx_msg"
	TxMsg_Field_Hash    = "hash"
	TxMsg_Field_Content = "content"
	TxMsg_Field_Type    = "type"
)

type TxMsg struct {
	Hash    string `bson:"hash"`
	Type    string `bson:"type"`
	Content string `bson:"content"`
}

func (m TxMsg) Name() string {
	return CollectionNmTxMsg
}

func (m TxMsg) PkKvPair() map[string]interface{} {
	return bson.M{TxMsg_Field_Hash: m.Hash}
}
