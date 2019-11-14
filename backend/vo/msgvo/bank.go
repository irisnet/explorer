package msgvo

import (
	"encoding/json"
	"github.com/irisnet/explorer/backend/utils"
)

type TxMsgSetMemoRegexp struct {
	Owner      string `json:"owner"`
	MemoRegexp string `json:"memo_regexp"`
}

type TxMsgBurn struct {
	Owner string      `json:"owner"`
	Coins utils.Coins `json:"coins"`
}

type TxMsgSend struct {
	Inputs  []Data `json:"inputs"`
	Outputs []Data `json:"outputs"`
}

type Data struct {
	Address string      `json:"address"`
	Coins   utils.Coins `json:"coins"`
}

func (vo *TxMsgSetMemoRegexp) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}

func (vo *TxMsgBurn) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}

func (vo *TxMsgSend) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}
