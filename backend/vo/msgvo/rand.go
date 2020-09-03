package msgvo

import (
	"encoding/json"
	"github.com/irisnet/explorer/backend/utils"
)

type TxMsgRequestRand struct {
	Consumer      string      `json:"consumer"`       // request address
	BlockInterval int64       `json:"block-interval"` // block interval after which the requested random number will be generated
	Oracle        bool        `json:"oracle"`
	ServiceFeeCap utils.Coins `json:"service_fee_cap"`
}

func (vo *TxMsgRequestRand) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}
