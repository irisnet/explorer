package document

import (
	"github.com/irisnet/explorer/backend/orm"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
	"time"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/explorer/backend/types"
	"fmt"
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

func (d ExStaticDelegator) Batch(txs []txn.Op) error {
	return orm.Batch(txs)
}

func Getdate(collectionName string, year, month int, sort string, location *time.Location) (time.Time, error) {
	var res struct {
		Date time.Time `bson:"date"`
	}
	var query = orm.NewQuery()
	defer query.Release()
	starttime, _ := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-01T00:00:00", year, month), location)
	endtime, _ := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-01T00:00:00", year, month+1), location)
	cond := bson.M{
		"date": bson.M{
			"$gte": starttime,
			"$lt":  endtime,
		},
	}
	query.SetCollection(collectionName).
		SetSelector(bson.M{"date": 1}).
		SetCondition(cond).
		SetSort(sort).
		SetSize(1).
		SetResult(&res)

	err := query.Exec()
	if err != nil {
		return time.Time{}, err
	}

	return res.Date, nil
}

func (d ExStaticDelegator) GetDataByDate(date time.Time) ([]ExStaticDelegator, error) {
	var res []ExStaticDelegator
	cond := bson.M{
		ExStaticDelegatorDateTag: date,
	}

	limit := 100
	for {
		var ret []ExStaticDelegator
		if err := queryAll(d.Name(), nil, cond, "-date", 100, &res); err != nil {
			return res, err
		}
		length := len(ret)
		res = append(res, ret...)
		if length < limit {
			break
		}
	}

	return res, nil
}

func (d ExStaticDelegator) GetDataOneDay(date time.Time, address string) (ExStaticDelegator, error) {
	var res ExStaticDelegator
	cond := bson.M{
		ExStaticDelegatorDateTag:    date,
		ExStaticDelegatorAddressTag: address,
	}
	if err := queryOne(d.Name(), nil, cond, &res); err != nil {
		return res, err
	}

	return res, nil
}
