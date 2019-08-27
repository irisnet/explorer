package msgvo

import "encoding/json"

type TxMsgSetMemoRegexp struct {
	Owner      string `json:"owner"`
	MemoRegexp string `json:"memo_regexp"`
}

func (vo *TxMsgSetMemoRegexp) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}
