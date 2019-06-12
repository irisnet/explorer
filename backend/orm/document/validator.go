package document

import (
	"fmt"
	"time"

	"github.com/irisnet/explorer/backend/orm"
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionNmValidator = "ex_validator"

	ValidatorFieldVotingPower     = "voting_power"
	ValidatorFieldJailed          = "jailed"
	ValidatorFieldStatus          = "status"
	ValidatorFieldOperatorAddress = "operator_address"
	ValidatorFieldDescription     = "description"
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

func (v Validator) String() string {

	return fmt.Sprintf(`
		ID              :%v
		OperatorAddress :%v
		ConsensusPubkey :%v
		Jailed          :%v
		Status          :%v
		Tokens          :%v
		DelegatorShares :%v
		Description     :%v
		BondHeight      :%v
		UnbondingHeight :%v
		UnbondingTime   :%v
		Commission      :%v
		Uptime          :%v
		SelfBond        :%v
		DelegatorNum    :%v
		ProposerAddr    :%v
		VotingPower     :%v
		`, v.ID, v.OperatorAddress, v.ConsensusPubkey, v.Jailed, v.Status, v.Tokens, v.DelegatorShares, v.Description, v.BondHeight, v.UnbondingHeight, v.UnbondingTime, v.Commission, v.Uptime, v.SelfBond, v.DelegatorNum, v.ProposerAddr, v.VotingPower)
}

func (v Validator) GetValidatorList() ([]Validator, error) {
	var validatorsDocArr []Validator
	var selector = bson.M{"description.moniker": 1, "operator_address": 1, "consensus_pubkey": 1, "proposer_addr": 1}
	err := queryAll(CollectionNmValidator, selector, nil, "", 0, &validatorsDocArr)

	return validatorsDocArr, err
}

func (v Validator) GetValidatorByProposerAddr(addr string) (Validator, error) {

	var selector = bson.M{"description.moniker": 1, "operator_address": 1}
	err := queryOne(CollectionNmValidator, selector, bson.M{"proposer_addr": addr}, &v)

	return v, err
}

type Description struct {
	Moniker  string `bson:"moniker" json:"moniker"`
	Identity string `bson:"identity" json:"identity"`
	Website  string `bson:"website" json:"website"`
	Details  string `bson:"details" json:"details"`
}

func (d Description) String() string {
	return fmt.Sprintf(`Moniker  :%v  Identity :%v Website  :%v Details  :%v`, d.Moniker, d.Identity, d.Website, d.Details)
}

type Commission struct {
	Rate          string    `bson:"rate" json:"rate"`
	MaxRate       string    `bson:"max_rate" json:"max_rate"`
	MaxChangeRate string    `bson:"max_change_rate" json:"max_change_rate"`
	UpdateTime    time.Time `bson:"update_time" json:"update_time"`
}

func (c Commission) String() string {
	return fmt.Sprintf(`Rate: %v  		MaxRate: %v 		MaxChangeRate: %v 		UpdateTime: %v   `, c.Rate, c.MaxRate, c.MaxChangeRate, c.UpdateTime)
}

func (v Validator) Name() string {
	return CollectionNmValidator
}

func (v Validator) PkKvPair() map[string]interface{} {
	return bson.M{"operator_address": v.OperatorAddress}
}

func (_ Validator) GetAllValidator() ([]Validator, error) {
	var validators []Validator
	var query = orm.NewQuery()
	defer query.Release()
	query.SetCollection(CollectionNmValidator).
		SetResult(&validators)

	err := query.Exec()

	return validators, err
}

func (v Validator) QueryValidatorMonikerOpAddrConsensusPubkey(addrArrAsVa []string) ([]Validator, error) {
	var validators []Validator
	var selector = bson.M{"description.moniker": 1, "operator_address": 1, "consensus_pubkey": 1}

	err := queryAll(CollectionNmValidator, selector, bson.M{"operator_address": bson.M{"$in": addrArrAsVa}}, "", 0, &validators)
	return validators, err
}
