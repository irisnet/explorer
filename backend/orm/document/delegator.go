package document

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionNmStakeRoleDelegator = "stake_role_delegator"

	Delegator_Field_Addres              = "address"
	Delegator_Field_ValidatorAddr       = "validator_addr"
	Delegator_Field_Shares              = "shares"
	Delegator_Field_original_shares     = "original_shares"
	Delegator_Field_BondedHeight        = "height"
	Delegator_Field_UnbondingDelegation = "unbonding_delegation"
)

type Delegator struct {
	Address        string  `bson:"address"`
	ValidatorAddr  string  `bson:"validator_addr"` // validatorAddr
	Shares         float64 `bson:"shares"`
	OriginalShares string  `bson:"original_shares"`
	BondedHeight   int64   `bson:"height"`

	UnbondingDelegation UnbondingDelegation `bson:"unbonding_delegation"`
}

// UnbondingDelegation reflects a delegation's passive unbonding queue.
type UnbondingDelegation struct {
	CreationHeight int64 `bson:"creation_height"` // height which the unbonding took place
	MinTime        int64 `bson:"min_time"`        // unix time for unbonding completion
	InitialBalance Coins `bson:"initial_balance"` // atoms initially scheduled to receive at completion
	Balance        Coins `bson:"balance"`         // atoms to receive at completion
}

func (d Delegator) Name() string {
	return CollectionNmStakeRoleDelegator
}

func (d Delegator) PkKvPair() map[string]interface{} {
	return bson.M{Delegator_Field_Addres: d.Address, Delegator_Field_ValidatorAddr: d.ValidatorAddr}
}
