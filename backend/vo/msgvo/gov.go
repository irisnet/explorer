package msgvo

import (
	"github.com/irisnet/explorer/backend/utils"
	"encoding/json"
)

type TxMsgSubmitProposal struct {
	Title          string      `json:"title"`          //  Title of the proposal
	Description    string      `json:"description"`    //  Description of the proposal
	Proposer       string      `json:"proposer"`       //  Address of the proposer
	InitialDeposit utils.Coins `json:"initialDeposit"` //  Initial deposit paid by sender. Must be strictly positive.
	ProposalType   string      `json:"proposalType"`   //  Initial deposit paid by sender. Must be strictly positive.
	Params         Params      `json:"params"`
}

type Param struct {
	Subspace string `json:"subspace"`
	Key      string `json:"key"`
	Value    string `json:"value"`
}

type Params []Param

func (vo *TxMsgSubmitProposal) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}

type TxMsgSubmitSoftwareUpgradeProposal struct {
	TxMsgSubmitProposal
	Version      int64  `json:"version"`
	Software     string `json:"software"`
	SwitchHeight int64  `json:"switch_height"`
	Threshold    string `json:"threshold"`
}

func (vo *TxMsgSubmitSoftwareUpgradeProposal) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}

type TxMsgSubmitCommunityTaxUsageProposal struct {
	TxMsgSubmitProposal
	Usage       string `json:"usage"`
	DestAddress string `json:"dest_address"`
	Percent     string `json:"percent"`
}

func (vo *TxMsgSubmitCommunityTaxUsageProposal) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
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

func (vo *TxMsgSubmitTokenAdditionProposal) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}

type TxMsgVote struct {
	ProposalID uint64 `json:"proposal_id"` // ID of the proposal
	Voter      string `json:"voter"`       //  address of the voter
	Option     string `json:"option"`      //  option from OptionSet chosen by the voter
}

func (vo *TxMsgVote) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}

type TxMsgDeposit struct {
	ProposalID uint64      `json:"proposal_id"` // ID of the proposal
	Depositor  string      `json:"depositor"`   // Address of the depositor
	Amount     utils.Coins `json:"amount"`      // Coins to add to the proposal's deposit
}

func (vo *TxMsgDeposit) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}
