package document

import (
	"time"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"github.com/irisnet/explorer/backend/orm"
	"gopkg.in/mgo.v2/txn"
)

const (
	CollectionNameExStaticRewards = "ex_static_rewards"

	ExStaticRewardsAddressTag = "address"
	ExStaticRewardsDateTag    = "date"
)

type ExStaticRewards struct {
	Id                bson.ObjectId        `bson:"_id"`
	Address           string               `bson:"address"`
	Date              time.Time            `bson:"date"`
	Total             []Rewards            `bson:"total"`
	DelegationsDetail []DelegationsRewards `bson:"delegations_detail"`
	Delegations       []Rewards            `bson:"delegations"`
	Commission        []Rewards            `bson:"commission"`
}

type Rewards struct {
	IrisAtto string  `bson:"iris_atto"`
	Iris     float64 `bson:"iris"`
}

type DelegationsRewards struct {
	Validator string    `bson:"validator"`
	Reward    []Rewards `bson:"reward"`
}

func (d ExStaticRewards) Name() string {
	return CollectionNameExStaticRewards
}

func (d ExStaticRewards) PkKvPair() map[string]interface{} {
	return bson.M{ExStaticRewardsAddressTag: d.Address, ExStaticRewardsDateTag: d.Date}
}

func (d ExStaticRewards) EnsureIndexes() []mgo.Index {
	indexes := []mgo.Index{
		{
			Key:        []string{ExStaticRewardsAddressTag, ExStaticRewardsDateTag},
			Unique:     true,
			Background: true,
		},
	}

	return indexes
}

func (_ ExStaticRewards) Batch(txs []txn.Op) error {
	return orm.Batch(txs)
}
