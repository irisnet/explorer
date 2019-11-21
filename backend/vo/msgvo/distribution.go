package msgvo

import "encoding/json"

type TxMsgSetWithdrawAddress struct {
	DelegatorAddr string `json:"delegator_addr"`
	WithdrawAddr  string `json:"withdraw_addr"`
}

type TxMsgWithdrawDelegatorReward struct {
	DelegatorAddr string `json:"delegator_addr"`
	ValidatorAddr string `json:"validator_addr"`
}

type TxMsgWithdrawDelegatorRewardsAll struct {
	DelegatorAddr string `json:"delegator_addr"`
}

// msg struct for validator withdraw
type TxMsgWithdrawValidatorRewardsAll struct {
	ValidatorAddr string `json:"validator_addr"`
}

func (vo *TxMsgSetWithdrawAddress) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}

func (vo *TxMsgWithdrawDelegatorReward) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}

func (vo *TxMsgWithdrawDelegatorRewardsAll) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}

func (vo *TxMsgWithdrawValidatorRewardsAll) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}
