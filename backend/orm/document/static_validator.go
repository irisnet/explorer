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
	Id                    bson.ObjectId `bson:"_id"`
	OperatorAddress       string        `bson:"operator_address"`
	Status                string        `bson:"status"`
	Date                  time.Time     `bson:"date"`
	Tokens                string        `bson:"tokens"`
	DelegatorShares       string        `bson:"delegator_shares"`
	Delegations           string        `bson:"delegations"` //tokens - self_bond
	SelfBond              string        `bson:"self_bond"`
	FoundationDelegations string        `bson:"foundation_delegations"`
	Commission            Commission    `bson:"commission"`
	DelegatorNum          int           `bson:"delegator_num"`
	CreateAt              int64         `bson:"create_at"`
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

func (_ ExStaticValidator) Batch(txs []txn.Op) error {
	return orm.Batch(txs)
}
