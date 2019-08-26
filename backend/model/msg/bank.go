package msg

type TxMsgSetMemoRegexp struct {
	Owner      string `json:"owner"`
	MemoRegexp string `json:"memo_regexp"`
}
