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
	CollectionNameExStaticDelegator = "ex_static_delegator"

	ExStaticDelegatorAddressTag = "address"
	ExStaticDelegatorDateTag    = "date"
)

type ExStaticDelegator struct {
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

//type DelegationsRewards struct {
//	Validator string    `bson:"validator"`
//	Reward    []Rewards `bson:"reward"`
//}

func (d ExStaticDelegator) Name() string {
	return CollectionNameExStaticDelegator
}

func (d ExStaticDelegator) PkKvPair() map[string]interface{} {
	return bson.M{ExStaticDelegatorAddressTag: d.Address, ExStaticDelegatorDateTag: d.Date}
}

func (d ExStaticDelegator) EnsureIndexes() []mgo.Index {
	indexes := []mgo.Index{
		{
			Key:        []string{ExStaticDelegatorAddressTag, ExStaticDelegatorDateTag},
			Unique:     true,
			Background: true,
		},
	}

	return indexes
}

func (_ ExStaticDelegator) Batch(txs []txn.Op) error {
	return orm.Batch(txs)
}
