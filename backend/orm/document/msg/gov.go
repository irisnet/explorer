package msg

import "github.com/irisnet/explorer/backend/utils"

type TxMsgSubmitProposal struct {
	Title          string      `bson:"title"`          //  Title of the proposal
	Description    string      `bson:"description"`    //  Description of the proposal
	Proposer       string      `bson:"proposer"`       //  Address of the proposer
	InitialDeposit utils.Coins `bson:"initialDeposit"` //  Initial deposit paid by sender. Must be strictly positive.
	ProposalType   string      `bson:"proposalType"`   //  Initial deposit paid by sender. Must be strictly positive.
}

type TxMsgSubmitTokenAdditionProposal struct {
	TxMsgSubmitProposal
	Symbol          string `bson:"symbol"`
	CanonicalSymbol string `bson:"canonical_symbol"`
	Name            string `bson:"name"`
	Decimal         uint8  `bson:"decimal"`
	MinUnitAlias    string `bson:"min_unit_alias"`
	InitialSupply   uint64 `bson:"initial_supply"`
}

func (msg *TxMsgSubmitProposal) Nil() {
}

func (msg *TxMsgSubmitTokenAdditionProposal) Nil() {
}

