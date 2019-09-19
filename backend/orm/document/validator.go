package document

import (
	"fmt"
	"math/big"
	"time"

	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/vo"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
)

const (
	CollectionNmValidator = "ex_validator"

	ValidatorFieldVotingPower      = "voting_power"
	ValidatorFieldJailed           = "jailed"
	ValidatorFieldStatus           = "status"
	ValidatorFieldOperatorAddress  = "operator_address"
	ValidatorFieldDescription      = "description"
	ValidatorFieldConsensusAddr    = "consensus_pubkey"
	ValidatorFieldProposerHashAddr = "proposer_addr"
	ValidatorFieldTokens           = "tokens"
	ValidatorFieldDelegatorShares  = "delegator_shares"
	ValidatorStatusValUnbonded     = 0
	ValidatorStatusValUnbonding    = 1
	ValidatorStatusValBonded       = 2
)

func (v Validator) GetValidatorStatus() string {

	if v.Jailed == false && v.Status == types.Bonded {
		return "Active"
	}

	if v.Status != types.Bonded && v.Jailed == false {
		return "Candidate"
	}

	return "Jailed"

}

func (v Validator) IsCandidatorWithStatus() bool {

	if v.Status != types.Bonded && v.Jailed == false {
		return true
	}

	return false

}

type (
	UptimeChangeVo struct {
		Address string
		Time    string
		Uptime  float64
	}

	ValVotingPowerChangeVo struct {
		Height  int64
		Address string
		Power   int64
		Time    time.Time
		Change  string
	}

	ValUpTimeVo struct {
		Time   string `bson:"_id,omitempty"`
		Uptime float64
	}

	CountVo struct {
		Id    bson.ObjectId `bson:"_id,omitempty"`
		Count float64
	}
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
	Icons           string        `bson:"icons" json:"icons"`
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
		Icons           :%v
		`, v.ID, v.OperatorAddress, v.ConsensusPubkey, v.Jailed, v.Status, v.Tokens, v.DelegatorShares, v.Description, v.BondHeight, v.UnbondingHeight, v.UnbondingTime, v.Commission, v.Uptime, v.SelfBond, v.DelegatorNum, v.ProposerAddr, v.VotingPower, v.Icons)
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
	var selector = bson.M{
		"description.moniker": 1,
		"operator_address":    1,
		"consensus_pubkey":    1,
		"status":              1,
		"voting_power":        1,
	}

	err := queryAll(CollectionNmValidator, selector, bson.M{"operator_address": bson.M{"$in": addrArrAsVa}}, "", 0, &validators)
	return validators, err
}

func (v Validator) QueryValidatorMonikerOpAddrByHashAddr(hashAddr []string) ([]Validator, error) {
	var validators []Validator
	var selector = bson.M{"description.moniker": 1, "operator_address": 1, "proposer_addr": 1}

	err := queryAll(CollectionNmValidator, selector, bson.M{"proposer_addr": bson.M{"$in": hashAddr}}, "", 0, &validators)
	return validators, err
}

func (_ Validator) GetValidatorListByPage(typ string, page, size int, ispage, total bool) (int, []Validator, error) {

	var query = orm.NewQuery()
	defer query.Release()
	var validators []Validator
	condition := bson.M{}
	switch typ {
	case types.RoleValidator:
		condition[ValidatorFieldJailed] = false
		condition[ValidatorFieldStatus] = types.Bonded
		break
	case types.RoleCandidate:
		condition[ValidatorFieldJailed] = false
		condition[ValidatorFieldStatus] = bson.M{
			"$in": []int{types.Unbonded, types.Unbonding},
		}
		break
	case types.RoleJailed:
		condition[ValidatorFieldJailed] = true
		break
	default:
	}

	if ispage {
		query.SetCollection(CollectionNmValidator).
			SetCondition(condition).
			SetSort(desc(ValidatorFieldVotingPower)).
			SetPage(page).
			SetSize(size).
			SetResult(&validators)
	} else {
		query.SetCollection(CollectionNmValidator).
			SetCondition(condition).
			SetSort(desc(ValidatorFieldVotingPower)).
			SetResult(&validators)
	}

	count, err := query.ExecPage(total)

	return count, validators, err
}

func (_ Validator) GetCandidatesTopN() ([]Validator, int64, map[string]int, error) {
	var validators []Validator
	var query = orm.NewQuery()
	defer query.Release()

	condition := bson.M{}
	condition[ValidatorFieldJailed] = false
	condition[ValidatorFieldStatus] = types.Bonded

	query.SetCollection(CollectionNmValidator).
		SetCondition(condition).
		SetSort(desc(ValidatorFieldVotingPower)).SetSize(10).
		SetResult(&validators)

	err := query.Exec()
	if err != nil {
		return nil, 0, nil, err
	}

	var allPower vo.CountVo
	query.SetResult(&allPower)
	query.PipeQuery(
		[]bson.M{
			{"$match": condition},
			{"$group": bson.M{
				"_id":   ValidatorFieldVotingPower,
				"count": bson.M{"$sum": "$voting_power"},
			}},
		},
	)

	upTimeMap := getValUpTime(query)

	return validators, int64(allPower.Count), upTimeMap, err
}

func (_ Validator) GetCandidatePubKeyAddrByAddr(addr string) (string, error) {
	validator, err := GetValidatorByAddr(addr)
	return validator.ConsensusPubkey, err
}

func GetValidatorByAddr(addr string) (Validator, error) {
	db := getDb()
	c := db.C(CollectionNmValidator)
	defer db.Session.Close()
	var validator Validator
	err := c.Find(bson.M{ValidatorFieldOperatorAddress: addr}).One(&validator)

	return validator, err
}

func (_ Validator) GetBondedValidators() ([]Validator, error) {
	var (
		validators []Validator
	)

	selector := bson.M{
		ValidatorFieldVotingPower: "1",
	}
	condition := bson.M{
		ValidatorFieldStatus: ValidatorStatusValBonded,
	}

	err := queryAll(CollectionNmValidator, selector, condition, "", 0, &validators)

	return validators, err
}

func (_ Validator) QueryPowerWithBonded() (int64, error) {

	var query = orm.NewQuery()
	defer query.Release()

	condition := bson.M{}
	condition[ValidatorFieldJailed] = false
	condition[ValidatorFieldStatus] = types.Bonded

	var count struct {
		Id    bson.ObjectId `bson:"_id,omitempty"`
		Count int64
	}
	query.Reset().
		SetResult(&count).
		SetCollection(CollectionNmValidator)
	err := query.PipeQuery([]bson.M{
		{"$match": condition},
		{"$group": bson.M{
			"_id":   ValidatorFieldVotingPower,
			"count": bson.M{"$sum": "$voting_power"},
		}},
	})
	return count.Count, err
}

func (_ Validator) QueryCandidateStatus(addr string) (int, int, error) {

	var query = orm.NewQuery()
	defer query.Release()

	var validator Validator
	query.SetCollection(CollectionNmValidator).
		SetCondition(bson.M{ValidatorFieldOperatorAddress: addr}).
		SetResult(&validator)
	if err := query.Exec(); err != nil {
		return 0, 0, err
	}

	upTimeMap := getValUpTime(query)
	query.Reset().SetCollection(CollectionNmBlock).
		SetCondition(bson.M{"block.last_commit.precommits": bson.M{"$elemMatch": bson.M{"validator_address": validator.OperatorAddress}}})

	var preCommitCount int
	if cnt, err := query.Count(); err == nil {
		preCommitCount = cnt
	}

	return preCommitCount, upTimeMap[validator.OperatorAddress], nil
}

func (_ Validator) QueryValidatorMonikerByAddrArr(addrs []string) (map[string]string, error) {
	validatorArr := []Validator{}

	valCondition := bson.M{
		ValidatorFieldOperatorAddress: bson.M{"$in": addrs},
	}

	selector := bson.M{ValidatorFieldDescription: 1, ValidatorFieldOperatorAddress: 1}

	if err := queryAll(CollectionNmValidator, selector, valCondition, "", 0, &validatorArr); err != nil {
		return nil, err
	}

	res := map[string]string{}

	for _, v := range validatorArr {
		res[v.OperatorAddress] = v.Description.Moniker
	}

	return res, nil

}

func (_ Validator) QueryValidatorListByAddrList(addrs []string) ([]Validator, error) {
	validatorArr := []Validator{}

	valCondition := bson.M{
		ValidatorFieldOperatorAddress: bson.M{"$in": addrs},
	}

	err := queryAll(CollectionNmValidator, nil, valCondition, "", 0, &validatorArr)

	return validatorArr, err
}

func (_ Validator) QueryMonikerAndValidatorAddrByHashAddr(addr string) (Validator, error) {

	selector := bson.M{ValidatorFieldOperatorAddress: 1, ValidatorFieldDescription: 1}
	condition := bson.M{ValidatorFieldProposerHashAddr: addr}
	var val Validator
	err := queryOne(CollectionNmValidator, selector, condition, &val)

	return val, err
}

func (_ Validator) QueryValidatorByConsensusAddr(addr string) (Validator, error) {
	var query = orm.NewQuery()
	defer query.Release()

	var result Validator
	condition := bson.M{}
	condition[ValidatorFieldConsensusAddr] = addr

	query.SetCollection(CollectionNmValidator).
		SetResult(&result).
		SetCondition(condition).
		SetSize(1)
	err := query.Exec()

	return result, err
}

func (_ Validator) QueryTokensAndShareRatioByValidatorAddrs(addrArrAsVa []string) (map[string]*big.Rat, error) {

	var validators []Validator
	var selector = bson.M{ValidatorFieldTokens: 1, ValidatorFieldOperatorAddress: 1, ValidatorFieldDelegatorShares: 1}

	err := queryAll(CollectionNmValidator, selector, bson.M{ValidatorFieldOperatorAddress: bson.M{"$in": addrArrAsVa}}, "", 0, &validators)

	if err != nil {
		return nil, err
	}

	result := map[string]*big.Rat{}

	for _, v := range validators {
		tokensAsRat, ok := new(big.Rat).SetString(v.Tokens)
		if !ok {
			logger.Error("convert validator token type (string -> big.Rat) err", logger.String("token str", v.Tokens))
			continue
		}

		delegatorShareAsRat, ok := new(big.Rat).SetString(v.DelegatorShares)
		if !ok {
			logger.Error("convert validator DelegatorShares type (string -> big.Rat) err", logger.String("DelegatorShares str", v.DelegatorShares))
			continue
		}
		result[v.OperatorAddress] = new(big.Rat).Quo(tokensAsRat, delegatorShareAsRat)
	}
	return result, nil
}

func (_ Validator) QueryValidatorDetailByOperatorAddr(opAddr string) (Validator, error) {

	validator := Validator{}

	valCondition := bson.M{
		ValidatorFieldOperatorAddress: opAddr,
	}

	err := queryOne(CollectionNmValidator, nil, valCondition, &validator)

	return validator, err
}

func (_ Validator) QueryTotalActiveValidatorVotingPower() (int64, error) {

	validators := []Validator{}
	condition := bson.M{ValidatorFieldJailed: false, ValidatorFieldStatus: types.Bonded}
	var selector = bson.M{ValidatorFieldVotingPower: 1}

	err := queryAll(CollectionNmValidator, selector, condition, "", 0, &validators)

	if err != nil {
		return 0, err
	}

	totalVotingPower := int64(0)
	for _, v := range validators {
		totalVotingPower += v.VotingPower
	}
	return totalVotingPower, nil
}

func (_ Validator) Batch(txs []txn.Op) error {
	return orm.Batch(txs)
}

func getValUpTime(query *orm.Query) map[string]int {
	var result []Block
	var upTimeMap = make(map[string]int)
	var selector = bson.M{"block.last_commit.precommits.validator_address": 1}
	query.Reset().
		SetCollection(CollectionNmBlock).
		SetSelector(selector).
		SetSize(100).
		SetSort(desc(Block_Field_Height)).
		SetResult(&result)

	if err := query.Exec(); err != nil {
		logger.Error("getValUpTime error", logger.String("err", err.Error()))
	}
	for _, block := range result {
		for _, pre := range block.Block.LastCommit.Precommits {
			upTimeMap[pre.ValidatorAddress]++
		}
	}
	return upTimeMap
}

// update document by primary key
func (_ Validator) UpdateByPk(validator Validator) error {
	db := orm.GetDatabase()
	defer db.Session.Close()

	c := db.C(CollectionNmValidator)
	return c.Update(validator.PkKvPair(), validator)
}
