package document

import (
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/irishub-sync/logger"
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
	CreationHeight int64       `bson:"creation_height"` // height which the unbonding took place
	MinTime        int64       `bson:"min_time"`        // unix time for unbonding completion
	InitialBalance utils.Coins `bson:"initial_balance"` // atoms initially scheduled to receive at completion
	Balance        utils.Coins `bson:"balance"`         // atoms to receive at completion
}

func (d Delegator) Name() string {
	return CollectionNmStakeRoleDelegator
}

func (d Delegator) PkKvPair() map[string]interface{} {
	return bson.M{Delegator_Field_Addres: d.Address, Delegator_Field_ValidatorAddr: d.ValidatorAddr}
}

func (_ Delegator) GetValidatorTokenAndSelfBond(addr string) (string, string, error) {

	validator := Validator{}
	selector := bson.M{
		"tokens":    1,
		"self_bond": 1}
	err := queryOne(CollectionNmValidator, selector, bson.M{ValidatorFieldOperatorAddress: addr}, &validator)
	if err != nil {
		logger.Error("validator not found", logger.Any("err", err.Error()))
		return "", "", err
	}
	return validator.Tokens, validator.SelfBond, err
}

func (_ Delegator) GetDepositValidator(valAddrArr []string) ([]Validator, error) {
	validators := []Validator{}
	selector := bson.M{
		ValidatorFieldOperatorAddress: 1,
		"tokens":                      1,
		"delegator_shares":            1}
	condition := bson.M{
		ValidatorFieldOperatorAddress: bson.M{"$in": valAddrArr},
	}
	err := queryAll(CollectionNmValidator, selector, condition, "", 0, &validators)
	if err != nil {
		logger.Error("validator not found", logger.Any("err", err.Error()))
		return nil, err
	}
	return validators, nil
}
