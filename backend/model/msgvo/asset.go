package msgvo

import "github.com/irisnet/explorer/backend/orm/document/msg"

type (
	TxMsgIssueToken struct {
		Family          string           `json:"family"`
		Source          string           `json:"source"`
		Gateway         string           `json:"gateway"`
		Symbol          string           `json:"symbol"`
		CanonicalSymbol string           `json:"canonical_symbol"`
		Name            string           `json:"name"`
		Decimal         int32            `json:"decimal"`
		MinUnitAlias    string           `json:"min_unit_alias"`
		InitialSupply   int64            `json:"initial_supply"`
		MaxSupply       int64            `json:"max_supply"`
		Mintable        bool             `json:"mintable"`
		Owner           string           `json:"owner"`
		UdInfo          AssetTokenUdInfo `json:"ud_info"`
	}

	TxMsgEditToken struct {
		TokenId         string           `json:"token_id"`         //  id of token
		Owner           string           `json:"owner"`            //  owner of token
		CanonicalSymbol string           `json:"canonical_symbol"` //  canonical_symbol of token
		MinUnitAlias    string           `json:"min_unit_alias"`   //  min_unit_alias of token
		MaxSupply       int64            `json:"max_supply"`
		Mintable        bool             `json:"mintable"` //  mintable of token
		Name            string           `json:"name"`
		UdInfo          AssetTokenUdInfo `json:"ud_info"`
	}

	TxMsgMintToken struct {
		TokenId string           `json:"token_id"` // the unique id of the token
		Owner   string           `json:"owner"`    // the current owner address of the token
		To      string           `json:"to"`       // address of mint token to
		Amount  int64            `json:"amount"`   // amount of mint token
		UdInfo  AssetTokenUdInfo `json:"ud_info"`
	}

	TxMsgTransferTokenOwner struct {
		SrcOwner string           `json:"src_owner"` // the current owner address of the token
		DstOwner string           `json:"dst_owner"` // the new owner
		TokenId  string           `json:"token_id"`
		UdInfo   AssetTokenUdInfo `json:"ud_info"`
	}

	TxMsgCreateGateway struct {
		Owner    string `json:"owner"`    //  the owner address of the gateway
		Moniker  string `json:"moniker"`  //  the globally unique name of the gateway
		Identity string `json:"identity"` //  the identity of the gateway
		Details  string `json:"details"`  //  the description of the gateway
		Website  string `json:"website"`  //  the external website of the gateway
	}

	TxMsgEditGateway struct {
		Owner    string `json:"owner"`    // Owner of the gateway
		Moniker  string `json:"moniker"`  // Moniker of the gateway
		Identity string `json:"identity"` // Identity of the gateway
		Details  string `json:"details"`  // Details of the gateway
		Website  string `json:"website"`  // Website of the gateway
	}

	TxMsgTransferGatewayOwner struct {
		Owner   string `json:"owner"`   // the current owner address of the gateway
		Moniker string `json:"moniker"` // the unique name of the gateway to be transferred
		To      string `json:"to"`      // the new owner to which the gateway ownership will be transferred
	}

	AssetTokenUdInfo struct {
		Source  string `json:"source"`
		Gateway string `json:"gateway"`
		Symbol  string `json:"symbol"`
	}
)

func (vo *TxMsgIssueToken) BuildIssueTokenMsgVOFromDoc(msgData msg.TxMsgIssueToken) TxMsgIssueToken {
	return TxMsgIssueToken{
		Family:          msgData.Family,
		Source:          msgData.Source,
		Gateway:         msgData.Gateway,
		Symbol:          msgData.Symbol,
		CanonicalSymbol: msgData.CanonicalSymbol,
		Name:            msgData.Name,
		Decimal:         msgData.Decimal,
		MinUnitAlias:    msgData.MinUnitAlias,
		InitialSupply:   msgData.InitialSupply,
		MaxSupply:       msgData.MaxSupply,
		Mintable:        msgData.Mintable,
		Owner:           msgData.Owner,
		UdInfo: AssetTokenUdInfo{
			Source:  msgData.UdInfo.Source,
			Gateway: msgData.UdInfo.Gateway,
			Symbol:  msgData.UdInfo.Symbol,
		},
	}
}

func (vo *TxMsgEditToken) BuildEditTokenMsgVOFromDoc(msgData msg.TxMsgEditToken) TxMsgEditToken {
	return TxMsgEditToken{
		TokenId:         msgData.TokenId,
		Owner:           msgData.Owner,
		CanonicalSymbol: msgData.CanonicalSymbol,
		MinUnitAlias:    msgData.MinUnitAlias,
		MaxSupply:       msgData.MaxSupply,
		Mintable:        msgData.Mintable,
		Name:            msgData.Name,
		UdInfo: AssetTokenUdInfo{
			Source:  msgData.UdInfo.Source,
			Gateway: msgData.UdInfo.Gateway,
			Symbol:  msgData.UdInfo.Symbol,
		},
	}
}

func (vo *TxMsgMintToken) BuildMintTokenMsgVOFromDoc(msgData msg.TxMsgMintToken) TxMsgMintToken {
	return TxMsgMintToken{
		TokenId: msgData.TokenId,
		Owner:   msgData.Owner,
		To:      msgData.To,
		Amount:  msgData.Amount,
		UdInfo: AssetTokenUdInfo{
			Source:  msgData.UdInfo.Source,
			Gateway: msgData.UdInfo.Gateway,
			Symbol:  msgData.UdInfo.Symbol,
		},
	}
}

func (vo *TxMsgTransferTokenOwner) BuildTransferTokenOwnerMsgVOFromDoc(
	msgData msg.TxMsgTransferTokenOwner) TxMsgTransferTokenOwner {
	return TxMsgTransferTokenOwner{
		SrcOwner: msgData.SrcOwner,
		DstOwner: msgData.DstOwner,
		TokenId:  msgData.TokenId,
		UdInfo: AssetTokenUdInfo{
			Source:  msgData.UdInfo.Source,
			Gateway: msgData.UdInfo.Gateway,
			Symbol:  msgData.UdInfo.Symbol,
		},
	}
}

func (vo *TxMsgCreateGateway) BuildCreateGatewayMsgVOFromDoc(msgData msg.TxMsgCreateGateway) TxMsgCreateGateway {
	return TxMsgCreateGateway{
		Owner:    msgData.Owner,
		Moniker:  msgData.Moniker,
		Identity: msgData.Identity,
		Details:  msgData.Details,
		Website:  msgData.Website,
	}
}

func (vo *TxMsgEditGateway) BuildEditGatewayMsgVOFromDoc(msgData msg.TxMsgEditGateway) TxMsgEditGateway {
	return TxMsgEditGateway{
		Owner:    msgData.Owner,
		Moniker:  msgData.Moniker,
		Identity: msgData.Identity,
		Details:  msgData.Details,
		Website:  msgData.Website,
	}
}

func (vo *TxMsgTransferGatewayOwner) BuildTransferGatewayOwnerMsgVOFromDoc(msgData msg.TxMsgTransferGatewayOwner) TxMsgTransferGatewayOwner {
	return TxMsgTransferGatewayOwner{
		Owner:   msgData.Owner,
		Moniker: msgData.Moniker,
		To:      msgData.To,
	}
}
