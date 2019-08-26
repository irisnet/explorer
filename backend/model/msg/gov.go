package msg

import (
	"github.com/irisnet/explorer/backend/utils"
)

type TxMsgSubmitProposal struct {
	Title          string      `json:"title"`          //  Title of the proposal
	Description    string      `json:"description"`    //  Description of the proposal
	Proposer       string      `json:"proposer"`       //  Address of the proposer
	InitialDeposit utils.Coins `json:"initialDeposit"` //  Initial deposit paid by sender. Must be strictly positive.
	ProposalType   string      `json:"proposalType"`   //  Initial deposit paid by sender. Must be strictly positive.
}

type TxMsgSubmitTokenAdditionProposal struct {
	TxMsgSubmitProposal
	Symbol          string `json:"symbol"`
	CanonicalSymbol string `json:"canonical_symbol"`
	Name            string `json:"name"`
	Decimal         uint8  `json:"decimal"`
	MinUnitAlias    string `json:"min_unit_alias"`
	InitialSupply   uint64 `json:"initial_supply"`
}
