package msgvo

import "github.com/irisnet/explorer/backend/orm/document/msg"

type TxMsgSetMemoRegexp struct {
	Owner      string `json:"owner"`
	MemoRegexp string `json:"memo_regexp"`
}

func (vo *TxMsgSetMemoRegexp) BuildSetMemoRegexpMsgVOFromDoc(msgData msg.TxMsgSetMemoRegexp) TxMsgSetMemoRegexp {
	return TxMsgSetMemoRegexp{
		Owner:      msgData.Owner,
		MemoRegexp: msgData.MemoRegexp,
	}
}
