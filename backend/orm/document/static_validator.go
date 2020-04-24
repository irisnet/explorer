package document

import (
	"github.com/irisnet/explorer/backend/orm"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
	"time"
)

const (
	CollectionNameExValidatorStatic = "ex_static_validator"

	ValidatorStaticFieldOperatorAddress = "operator_address"
	ValidatorStaticFieldDate            = "date"
	ValidatorCommissionRateTag          = "commission.rate"
)

type ExStaticValidator struct {
	Id              bson.ObjectId `bson:"_id"`
	OperatorAddress string        `bson:"operator_address"`
	Status          string        `bson:"status"`
	Date            time.Time     `bson:"date"`
	Tokens          string        `bson:"tokens"`
	DelegatorShares string        `bson:"delegator_shares"`
	Delegations     string        `bson:"delegations"` //tokens - self_bond
	SelfBond        string        `bson:"self_bond"`
	Commission      Commission    `bson:"commission"`
	DelegatorNum    int           `bson:"delegator_num"`
	CreateAt        int64         `bson:"create_at"`
}

func (d ExStaticValidator) Name() string {
	return CollectionNameExValidatorStatic
}

func (d ExStaticValidator) PkKvPair() map[string]interface{} {
	return bson.M{ValidatorStaticFieldOperatorAddress: d.OperatorAddress, ValidatorStaticFieldDate: d.Date}
}

func (d ExStaticValidator) EnsureIndexes() []mgo.Index {
	indexes := []mgo.Index{
		{
			Key:        []string{ValidatorStaticFieldOperatorAddress, ValidatorStaticFieldDate},
			Unique:     true,
			Background: true,
		},
	}

	return indexes
}

func (d ExStaticValidator) Batch(txs []txn.Op) error {
	return orm.Batch(txs)
}

func (d ExStaticValidator) GetDataByDate(date time.Time) ([]ExStaticValidator, error) {
	var res []ExStaticValidator
	cond := bson.M{
		ValidatorStaticFieldDate: date,
	}

	limit := 100
	offset := 0
	for {
		var ret []ExStaticValidator
		if err := querylistByOffsetAndSize(d.Name(), nil, cond, "-tokens", offset, limit, &ret); err != nil {
			return res, err
		}
		length := len(ret)
		res = append(res, ret...)
		if length < limit {
			break
		}
		offset += limit
	}

	return res, nil
}

func (d ExStaticValidator) GetDataOneDay(date time.Time, operatoraddr string) (ExStaticValidator, error) {
	var res ExStaticValidator
	cond := bson.M{
		ValidatorStaticFieldDate:            date,
		ValidatorStaticFieldOperatorAddress: operatoraddr,
	}
	if err := queryOne(d.Name(), nil, cond, &res); err != nil && err != mgo.ErrNotFound {
		return res, err
	}

	return res, nil
}

func (d ExStaticValidator) GetCommissionRate(selector, cond bson.M, sorts string) (ExStaticValidator, error) {
	var query = orm.NewQuery()
	defer query.Release()
	var result ExStaticValidator

	query.SetCollection(d.Name()).
		SetCondition(cond).
		SetSelector(selector).
		SetSize(1).
		SetSort(sorts).
		SetResult(&result)

	err := query.Exec()
	if err != nil {
		return result, err
	}
	return result, nil
}
