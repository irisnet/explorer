package document

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	CollectionNmValidator = "ex_validator"

	ValidatorFieldVotingPower     = "voting_power"
	ValidatorFieldJailed          = "jailed"
	ValidatorFieldStatus          = "status"
	ValidatorFieldOperatorAddress = "operator_address"
	ValidatorFieldDescription     = "description"

	ValidatorStatusValUnbonded  = 0
	ValidatorStatusValUnbonding = 1
	ValidatorStatusValBonded    = 2
)

type Validator struct {
	ID              bson.ObjectId `bson:"_id"`
	OperatorAddress string        `bson:"operator_address" json:"operator_address"`
	ConsensusPubkey string        `bson:"consensus_pubkey" json:"consensus_pubkey"`
	Jailed          bool          `bson:"jailed" json:"jailed"`
	Status          int           `bson:"status" json:"status"`
	Tokens          string        `bson:"tokens" json:"tokens"`
	DelegatorShares string        `bson:"delegator_shares" json:"delegator_shares"`
	Description     Description   `bson:"description" json:"description"`
	BondHeight      string        `bson:"bond_height" json:"bond_height"`
	UnbondingHeight string        `bson:"unbonding_height" json:"unbonding_height"`
	UnbondingTime   time.Time     `bson:"unbonding_time" json:"unbonding_time"`
	Commission      Commission    `bson:"commission" json:"commission"`
	Uptime          float32       `bson:"uptime" json:"uptime"`
	SelfBond        string        `bson:"self_bond" json:"self_bond"`
	DelegatorNum    int           `bson:"delegator_num" json:"delegator_num"`
	ProposerAddr    string        `bson:"proposer_addr" json:"proposer_addr"`
	VotingPower     int64         `bson:"voting_power" json:"voting_power"`
}

type Description struct {
	Moniker  string `bson:"moniker" json:"moniker"`
	Identity string `bson:"identity" json:"identity"`
	Website  string `bson:"website" json:"website"`
	Details  string `bson:"details" json:"details"`
}
type Commission struct {
	Rate          string    `bson:"rate" json:"rate"`
	MaxRate       string    `bson:"max_rate" json:"max_rate"`
	MaxChangeRate string    `bson:"max_change_rate" json:"max_change_rate"`
	UpdateTime    time.Time `bson:"update_time" json:"update_time"`
}

func (v Validator) Name() string {
	return CollectionNmValidator
}

func (v Validator) PkKvPair() map[string]interface{} {
	return bson.M{"operator_address": v.OperatorAddress}
}
