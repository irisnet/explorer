package document

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

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

func (tm TxMsg) String() string {
	return fmt.Sprintf(`

Hash:    %v
Type:    %v
Content: %v

	  `, tm.Hash, tm.Type, tm.Content)
}

func (m TxMsg) Name() string {
	return CollectionNmTxMsg
}

func (m TxMsg) PkKvPair() map[string]interface{} {
	return bson.M{TxMsg_Field_Hash: m.Hash}
}
