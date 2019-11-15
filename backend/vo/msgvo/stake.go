package msgvo

import (
	"github.com/irisnet/explorer/backend/vo"
	"github.com/irisnet/explorer/backend/utils"
	"encoding/json"
)

type TxMsgStakeCreate struct {
	ValDescription vo.Description   `json:"valdescription"`
	Commission     vo.CommissionMsg `json:"commission"`
	DelegatorAddr  string           `json:"delegator_address"`
	ValidatorAddr  string           `json:"validator_address"`
	PubKey         string           `json:"pubkey"`
	Delegation     utils.Coin       `json:"delegation"`
}

func (vo *TxMsgStakeCreate) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}

type TxMsgStakeEdit struct {
	ValDescription vo.Description `json:"valdescription"`
	ValidatorAddr  string         `json:"address"`
	CommissionRate string         `json:"commission_rate"`
}

func (vo *TxMsgStakeEdit) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}
