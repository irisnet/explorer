package msgvo

import (
	"encoding/json"
	"github.com/irisnet/explorer/backend/utils"
)

type TxMsgSetWithdrawAddress struct {
	DelegatorAddr string `json:"delegator_addr"`
	WithdrawAddr  string `json:"withdraw_addr"`
}

type TxMsgWithdrawDelegatorReward struct {
	DelegatorAddr string `json:"delegator_addr"`
	ValidatorAddr string `json:"validator_addr"`
}

type TxTypeMsgWithdrawValidatorCommission struct {
	ValidatorAddr string `json:"validator_addr"`
}

// msg struct for validator withdraw
type TxTypeMsgFundCommunityPool struct {
	Amount    utils.Coins `json:"amount"`
	Depositor string      `json:"depositor"`
}

func (vo *TxMsgSetWithdrawAddress) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}

func (vo *TxMsgWithdrawDelegatorReward) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}

func (vo *TxTypeMsgWithdrawValidatorCommission) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}

func (vo *TxTypeMsgFundCommunityPool) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}
