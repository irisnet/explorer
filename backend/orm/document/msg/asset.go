package msg

type (
	TxMsgIssueToken struct {
		Family          string           `bson:"family"`
		Source          string           `bson:"source"`
		Gateway         string           `bson:"gateway"`
		Symbol          string           `bson:"symbol"`
		CanonicalSymbol string           `bson:"canonical_symbol"`
		Name            string           `bson:"name"`
		Decimal         int32            `bson:"decimal"`
		MinUnitAlias    string           `bson:"min_unit_alias"`
		InitialSupply   int64            `bson:"initial_supply"`
		MaxSupply       int64            `bson:"max_supply"`
		Mintable        bool             `bson:"mintable"`
		Owner           string           `bson:"owner"`
		UdInfo          AssetTokenUdInfo `bson:"ud_info"`
	}

	TxMsgEditToken struct {
		TokenId         string           `bson:"token_id"`         //  id of token
		Owner           string           `bson:"owner"`            //  owner of token
		CanonicalSymbol string           `bson:"canonical_symbol"` //  canonical_symbol of token
		MinUnitAlias    string           `bson:"min_unit_alias"`   //  min_unit_alias of token
		MaxSupply       int64            `bson:"max_supply"`
		Mintable        bool             `bson:"mintable"` //  mintable of token
		Name            string           `bson:"name"`
		UdInfo          AssetTokenUdInfo `bson:"ud_info"`
	}

	TxMsgMintToken struct {
		TokenId string           `bson:"token_id"` // the unique id of the token
		Owner   string           `bson:"owner"`    // the current owner address of the token
		To      string           `bson:"to"`       // address of mint token to
		Amount  int64            `bson:"amount"`   // amount of mint token
		UdInfo  AssetTokenUdInfo `bson:"ud_info"`
	}

	TxMsgTransferTokenOwner struct {
		SrcOwner string           `bson:"src_owner"` // the current owner address of the token
		DstOwner string           `bson:"dst_owner"` // the new owner
		TokenId  string           `bson:"token_id"`
		UdInfo   AssetTokenUdInfo `bson:"ud_info"`
	}

	TxMsgCreateGateway struct {
		Owner    string `bson:"owner"`    //  the owner address of the gateway
		Moniker  string `bson:"moniker"`  //  the globally unique name of the gateway
		Identity string `bson:"identity"` //  the identity of the gateway
		Details  string `bson:"details"`  //  the description of the gateway
		Website  string `bson:"website"`  //  the external website of the gateway
	}

	TxMsgEditGateway struct {
		Owner    string `bson:"owner"`    // Owner of the gateway
		Moniker  string `bson:"moniker"`  // Moniker of the gateway
		Identity string `bson:"identity"` // Identity of the gateway
		Details  string `bson:"details"`  // Details of the gateway
		Website  string `bson:"website"`  // Website of the gateway
	}

	TxMsgTransferGatewayOwner struct {
		Owner   string `bson:"owner"`   // the current owner address of the gateway
		Moniker string `bson:"moniker"` // the unique name of the gateway to be transferred
		To      string `bson:"to"`      // the new owner to which the gateway ownership will be transferred
	}

	AssetTokenUdInfo struct {
		Source  string `bson:"source"`
		Gateway string `bson:"gateway"`
		Symbol  string `bson:"symbol"`
	}
)

func (msg TxMsgIssueToken) Nil() {
}

func (msg TxMsgEditToken) Nil() {
}

func (msg TxMsgMintToken) Nil() {
}

func (msg TxMsgTransferTokenOwner) Nil() {
}

func (msg TxMsgCreateGateway) Nil() {
}

func (msg TxMsgEditGateway) Nil() {
}

func (msg TxMsgTransferGatewayOwner) Nil() {
}
