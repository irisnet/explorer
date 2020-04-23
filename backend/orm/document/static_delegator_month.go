package document

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/txn"
	"github.com/irisnet/explorer/backend/orm"
)

const (
	CollectionNameExStaticDelegatorMonth = "ex_static_delegator_month"

	ExStaticDelegatorMonthAddressTag = "address"
	ExStaticDelegatorMonthDateTag    = "date"
)

type ExStaticDelegatorMonth struct {
	Id                     bson.ObjectId `bson:"_id"`
	Address                string        `bson:"address"`
	Date                   string        `bson:"date"`
	TerminalRewards        Rewards       `bson:"terminal_rewards"`
	PeriodWithdrawRewards  Rewards       `bson:"period_withdraw_rewards"`
	PeriodIncrementRewards Rewards       `bson:"period_increment_rewards"`
	TerminalDelegation     Coin          `bson:"terminal_delegation"`
	IncrementDelegation    Coin          `bson:"increment_delegation"`
	PeriodDelegationTimes  int           `bson:"period_delegation_times"`
	CreateAt               int64         `bson:"create_at"`
	UpdateAt               int64         `bson:"update_at"`
}

func (d ExStaticDelegatorMonth) Name() string {
	return CollectionNameExStaticDelegatorMonth
}

func (d ExStaticDelegatorMonth) PkKvPair() map[string]interface{} {
	return bson.M{ExStaticDelegatorMonthAddressTag: d.Address, ExStaticDelegatorMonthDateTag: d.Date}
}

func (d ExStaticDelegatorMonth) EnsureIndexes() []mgo.Index {
	indexes := []mgo.Index{
		{
			Key:        []string{ExStaticDelegatorMonthAddressTag, ExStaticDelegatorMonthDateTag},
			Unique:     true,
			Background: true,
		},
	}

	return indexes
}

func (d ExStaticDelegatorMonth) GetLatest(address string) (ExStaticDelegatorMonth, error) {
	var res ExStaticDelegatorMonth
	cond := bson.M{}
	if address != "" {
		cond[ExStaticDelegatorMonthAddressTag] = address
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

func (d ExStaticDelegatorMonth) Batch(txs []txn.Op) error {
	return orm.Batch(txs)
}
