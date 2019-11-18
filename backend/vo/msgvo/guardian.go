package msgvo

import "encoding/json"

type TxMsgAddProfiler struct {
	AddGuardian AddGuardian `json:"addguardian"`
}

type TxMsgAddTrustee struct {
	AddGuardian AddGuardian `json:"addguardian"`
}

type AddGuardian struct {
	Description string `json:"description"`
	Address     string `json:"address"`  // address added
	AddedBy     string `json:"added_by"` // address that initiated the tx
}

type TxMsgDeleteProfiler struct {
	DeleteGuardian DeleteGuardian `json:"deleteguardian"`
}

type TxMsgDeleteTrustee struct {
	DeleteGuardian DeleteGuardian `json:"deleteguardian"`
}

type DeleteGuardian struct {
	Address   string `json:"address"`    // address deleted
	DeletedBy string `json:"deleted_by"` // address that initiated the tx
}

func (vo *TxMsgAddProfiler) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}

func (vo *TxMsgAddTrustee) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}

func (vo *TxMsgDeleteProfiler) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}

func (vo *TxMsgDeleteTrustee) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}
