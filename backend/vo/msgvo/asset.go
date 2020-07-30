package msgvo

import (
	"encoding/json"
)

type (
	TxMsgIssueToken struct {
		Symbol        string `json:"symbol"`
		Name          string `json:"name"`
		Scale         uint32 `json:"scale"`
		MinUnit       string `json:"min_unit"`
		InitialSupply int64  `json:"initial_supply"`
		MaxSupply     int64  `json:"max_supply"`
		Mintable      bool   `json:"mintable"`
		Owner         string `json:"owner"`
	}

	TxMsgEditToken struct {
		Owner     string `json:"owner"` //  owner of token
		MaxSupply int64  `json:"max_supply"`
		Mintable  bool   `json:"mintable"` //  mintable of token
		Symbol    string `json:"symbol"`
		Name      string `json:"name"`
	}

	TxMsgMintToken struct {
		Symbol string `json:"symbol"`
		Owner  string `json:"owner"`  // the current owner address of the token
		To     string `json:"to"`     // address of mint token to
		Amount int64  `json:"amount"` // amount of mint token
	}

	TxMsgTransferTokenOwner struct {
		SrcOwner string `json:"src_owner"` // the current owner address of the token
		DstOwner string `json:"dst_owner"` // the new owner
		Symbol   string `json:"symbol"`
	}

	//TxMsgCreateGateway struct {
	//	Owner    string `json:"owner"`    //  the owner address of the gateway
	//	Moniker  string `json:"moniker"`  //  the globally unique name of the gateway
	//	Identity string `json:"identity"` //  the identity of the gateway
	//	Details  string `json:"details"`  //  the description of the gateway
	//	Website  string `json:"website"`  //  the external website of the gateway
	//}
	//
	//TxMsgEditGateway struct {
	//	Owner    string `json:"owner"`    // Owner of the gateway
	//	Moniker  string `json:"moniker"`  // Moniker of the gateway
	//	Identity string `json:"identity"` // Identity of the gateway
	//	Details  string `json:"details"`  // Details of the gateway
	//	Website  string `json:"website"`  // Website of the gateway
	//}
	//
	//TxMsgTransferGatewayOwner struct {
	//	Owner   string `json:"owner"`   // the current owner address of the gateway
	//	Moniker string `json:"moniker"` // the unique name of the gateway to be transferred
	//	To      string `json:"to"`      // the new owner to which the gateway ownership will be transferred
	//}
	//
	//AssetTokenUdInfo struct {
	//	Source  string `json:"source"`
	//	Gateway string `json:"gateway"`
	//	Symbol  string `json:"symbol"`
	//}
)

func (vo *TxMsgIssueToken) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}

func (vo *TxMsgEditToken) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}

func (vo *TxMsgMintToken) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}

func (vo *TxMsgTransferTokenOwner) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}

//
//func (vo *TxMsgCreateGateway) BuildMsgByUnmarshalJson(data []byte) error {
//	return json.Unmarshal(data, vo)
//}
//
//func (vo *TxMsgEditGateway) BuildMsgByUnmarshalJson(data []byte) error {
//	return json.Unmarshal(data, vo)
//}
//
//func (vo *TxMsgTransferGatewayOwner) BuildMsgByUnmarshalJson(data []byte) error {
//	return json.Unmarshal(data, vo)
//}
