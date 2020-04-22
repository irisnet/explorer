package document

import (
	"github.com/irisnet/explorer/backend/orm"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
	"time"
	"github.com/irisnet/explorer/backend/utils"
)

const (
	CollectionNameExStaticRewards = "ex_static_rewards"

	ExStaticRewardsAddressTag = "address"
	ExStaticRewardsDateTag    = "date"
)

type ExStaticRewards struct {
	Id                 bson.ObjectId `bson:"_id"`
	Address            string        `bson:"address"`
	Date               time.Time     `bson:"date"`
	Total              []Rewards     `bson:"total"`
	Delegation         utils.Coin    `bson:"delegation"`
	DelegationsRewards []Rewards     `bson:"delegations_rewards"`
	Commission         []Rewards     `bson:"commission"`
	CreateAt           int64         `bson:"create_at"`
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
