package msgvo

import "github.com/irisnet/explorer/backend/orm/document/msg"

type TxMsgRequestRand struct {
	Consumer      string `json:"consumer"`       // request address
	BlockInterval int64  `json:"block_interval"` // block interval after which the requested random number will be generated
}

func (vo *TxMsgRequestRand) BuildRequestRandMsgVOFromDoc(msgData msg.TxMsgRequestRand) TxMsgRequestRand {
	return TxMsgRequestRand{
		Consumer:      msgData.Consumer,
		BlockInterval: msgData.BlockInterval,
	}
}
