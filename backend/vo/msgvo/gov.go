package msgvo

import (
	"github.com/irisnet/explorer/backend/utils"
	"encoding/json"
)

type TxMsgSubmitProposal struct {
	Proposer       string      `json:"proposer"`       //  Address of the proposer
	InitialDeposit utils.Coins `json:"initialDeposit"` //  Initial deposit paid by sender. Must be strictly positive.
	Content        Any         `bson:"content"`
}

type Any struct {
	// nolint
	TypeUrl string `bson:"type_url"`
	// Must be a valid serialized protocol buffer of the above specified type.
	Value string `bson:"value"`

	//cachedValue interface{} `bson:"cached_value"`

	//compat anyCompat
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
