package document

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/txn"
	"github.com/irisnet/explorer/backend/orm"
)

const (
	CollectionNameExValidatorStatic = "ex_validator_static"

	ValidatorStaticFieldOperatorAddress = "operator_address"
	ValidatorStaticFieldDate            = "date"
)

type ExValidatorStatic struct {
	Id              bson.ObjectId `bson:"_id"`
	OperatorAddress string        `bson:"operator_address"`
	Date            time.Time     `bson:"date"`
	Tokens          string        `bson:"tokens"`
	DelegatorShares string        `bson:"delegator_shares"`
	Delegations     string        `bson:"delegations"` //tokens - self_bond
	SelfBond        string        `bson:"self_bond"`
	Commission      Commission    `bson:"commission"`
}

func (d ExValidatorStatic) Name() string {
	return CollectionNameExValidatorStatic
}

func (d ExValidatorStatic) PkKvPair() map[string]interface{} {
	return bson.M{ValidatorStaticFieldOperatorAddress: d.OperatorAddress, ValidatorStaticFieldDate: d.Date}
}

func (d ExValidatorStatic) EnsureIndexes() []mgo.Index {
	indexes := []mgo.Index{
		{
			Key:        []string{ValidatorStaticFieldOperatorAddress, ValidatorStaticFieldDate},
			Unique:     true,
			Background: true,
		},
	}

	return indexes
}

func (_ ExValidatorStatic) Batch(txs []txn.Op) error {
	return orm.Batch(txs)
}
