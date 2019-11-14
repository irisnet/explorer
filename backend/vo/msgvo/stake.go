package msgvo

import (
	"github.com/irisnet/explorer/backend/vo"
	"github.com/irisnet/explorer/backend/utils"
	"encoding/json"
	"github.com/irisnet/irishub-sync/store"
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

type TxMsgUnjail struct {
	ValidatorAddr string `json:"address"` // address of the validator operator
}

func (vo *TxMsgUnjail) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}

type TxMsgBeginRedelegate struct {
	DelegatorAddr    string `json:"delegator_addr"`
	ValidatorSrcAddr string `json:"validator_src_addr"`
	ValidatorDstAddr string `json:"validator_dst_addr"`
	SharesAmount     string `json:"shares_amount"`
}

func (vo *TxMsgBeginRedelegate) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}

type TxMsgBeginUnbonding struct {
	DelegatorAddr string `json:"delegator_addr"`
	ValidatorAddr string `json:"validator_addr"`
	SharesAmount  string `json:"shares_amount"`
}

func (vo *TxMsgBeginUnbonding) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}

type TxMsgDelegate struct {
	DelegatorAddr string     `json:"delegator_addr"`
	ValidatorAddr string     `json:"validator_addr"`
	Delegation    store.Coin `json:"delegation"`
}

func (vo *TxMsgDelegate) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}