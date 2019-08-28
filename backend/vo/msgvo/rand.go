package msgvo

import "encoding/json"

type TxMsgRequestRand struct {
	Consumer      string `json:"consumer"`       // request address
	BlockInterval int64  `json:"block-interval"` // block interval after which the requested random number will be generated
}

func (vo *TxMsgRequestRand) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}
