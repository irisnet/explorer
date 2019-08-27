package txmsg

type TxMsgSetMemoRegexp struct {
	Owner      string `bson:"owner"`
	MemoRegexp string `bson:"memo_regexp"`
}
