package document

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/txn"
	"github.com/irisnet/explorer/backend/orm"
	"time"
)

const (
	CollectionNameExStaticValidatorMonth = "ex_static_validator_month"

	ExStaticValidatorMonthAddressTag = "address"
	ExStaticValidatorMonthDateTag    = "date"
)

type ExStaticValidatorMonth struct {
	Id                      bson.ObjectId `bson:"_id"`
	Address                 string        `bson:"address"`
	OperatorAddress         string        `bson:"operator_address"`
	CreateValidatorHeight   int64         `bson:"create_validator_height"`
	Date                    string        `bson:"date"`
	CaculateDate            string        `bson:"caculate_date"`
	TerminalCommission      Coin          `bson:"terminal_commission"`
	PeriodCommission        Coin          `bson:"period_commission"`
	IncrementCommission     Coin          `bson:"increment_commission"`
	TerminalDelegation      string        `bson:"terminal_delegation"`
	IncrementDelegation     string        `bson:"increment_delegation"`
	Rank                    int           `bson:"rank"` //权重排名用
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

func (d ExStaticValidatorMonth) Batch(txs []txn.Op) error {
	return orm.Batch(txs)
}
func (d ExStaticValidatorMonth) Save(validatormonth ExStaticValidatorMonth) error {
	validatormonth.Id = bson.NewObjectId()
	validatormonth.CreateAt = time.Now().Unix()
	validatormonth.UpdateAt = time.Now().Unix()
	return orm.Save(d.Name(), validatormonth)
}

func (d ExStaticValidatorMonth) GetValidatorStaticByMonth(datestr, operatoraddr string) (ExStaticValidatorMonth, error) {
	var res ExStaticValidatorMonth
	cond := bson.M{
		ValidatorStaticFieldDate:            datestr,
		ValidatorStaticFieldOperatorAddress: operatoraddr,
	}
	if err := queryOne(d.Name(), nil, cond, &res); err != nil && err != mgo.ErrNotFound {
		return res, err
	}

	return res, nil
}

func (d ExStaticValidatorMonth) GetLatest(address string) (ExStaticValidatorMonth, error) {
	var res ExStaticValidatorMonth
	cond := bson.M{}
	if address != "" {
		cond[ValidatorStaticFieldOperatorAddress] = address
	}
	var query = orm.NewQuery()
	defer query.Release()
	query.SetCollection(d.Name()).
		SetCondition(cond).
		SetSort("-date").
		SetSize(1).
		SetResult(&res)

	err := query.Exec()
	if err != nil && err != mgo.ErrNotFound {
		return res, err
	}
	return res, nil
}

func (d ExStaticValidatorMonth) List(cond bson.M, pageNum, pageSize int, istotal bool) ([]ExStaticValidatorMonth, int, error) {
	var res []ExStaticValidatorMonth

	total, err := pageQuery(d.Name(), nil, cond, desc(ExStaticDelegatorDateTag), pageNum, pageSize, istotal, &res)
	if err != nil && err != mgo.ErrNotFound {
		return res, 0, err
	}
	return res, total, nil
}
