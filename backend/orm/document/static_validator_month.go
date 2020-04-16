package document

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/txn"
	"github.com/irisnet/explorer/backend/orm"
)

const (
	CollectionNameExStaticValidatorMonth = "ex_static_delegator_month"

	ExStaticValidatorMonthAddressTag = "address"
	ExStaticValidatorMonthDateTag    = "date"
)

type ExStaticValidatorMonth struct {
	Id                      bson.ObjectId `bson:"_id"`
	Address                 string        `bson:"address"`
	OperatorAddress         string        `bson:"operator_address"`
	CreateValidatorHeight   int64         `bson:"create_validator_height"`
	Date                    string        `bson:"date"`
	TerminalCommission      Coin          `bson:"terminal_commission"`
	PeriodCommission        Coin          `bson:"period_commission"`
	IncrementCommission     Coin          `bson:"increment_commission"`
	TerminalDelegation      string        `bson:"terminal_delegation"`
	IncrementDelegation     string        `bson:"increment_delegation"`
	Tokens                  string        `bson:"tokens"` //权重排名用
	TerminalDelegatorN      int           `bson:"terminal_delegator_n"`
	IncrementDelegatorN     int           `bson:"increment_delegator_n"`
	TerminalSelfBond        string        `bson:"terminal_self_bond"`
	IncrementSelfBond       string        `bson:"increment_self_bond"`
	CommissionRateMax       string        `bson:"commission_rate_max"`
	CommissionRateMin       string        `bson:"commission_rate_min"`
	FoundationDelegateT     string        `bson:"foundation_delegate_t"`
	FoundationDelegateIncre string        `bson:"foundation_delegate_incre"`
	Moniker                 string        `bson:"moniker"`
	CreateAt                int64         `bson:"create_at"`
	UpdateAt                int64         `bson:"update_at"`
}

func (d ExStaticValidatorMonth) Name() string {
	return CollectionNameExStaticValidatorMonth
}

func (d ExStaticValidatorMonth) PkKvPair() map[string]interface{} {
	return bson.M{ExStaticValidatorMonthAddressTag: d.Address, ExStaticValidatorMonthDateTag: d.Date}
}

func (d ExStaticValidatorMonth) EnsureIndexes() []mgo.Index {
	indexes := []mgo.Index{
		{
			Key:        []string{ExStaticValidatorMonthAddressTag, ExStaticValidatorMonthDateTag},
			Unique:     true,
			Background: true,
		},
	}

	return indexes
}

func (_ ExStaticValidatorMonth) Batch(txs []txn.Op) error {
	return orm.Batch(txs)
}
