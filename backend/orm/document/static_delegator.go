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
	CollectionNameExStaticDelegator = "ex_static_rewards"

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

func Getdate(collectionName string, starttime, endtime time.Time, sort string) (time.Time, error) {
	var res struct {
		Date time.Time `bson:"date"`
	}
	var query = orm.NewQuery()
	defer query.Release()

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

func (d ExStaticDelegator) GetDataByDate(date time.Time) (map[string]ExStaticDelegator, error) {
	var res []ExStaticDelegator
	cond := bson.M{
		ExStaticDelegatorDateTag: date,
	}

	limit := 100
	offset := 0
	for {
		var ret []ExStaticDelegator
		if err := querylistByOffsetAndSize(d.Name(), nil, cond, "-date", offset, limit, &ret); err != nil {
			return nil, err
		}
		length := len(ret)
		res = append(res, ret...)
		if length < limit {
			break
		}
		offset += limit
	}

	dataMap := make(map[string]ExStaticDelegator, len(res))

	for _, one := range res {
		dataMap[one.Address] = one
	}

	return dataMap, nil
}

func (d ExStaticDelegator) GetAddressByPeriod(lastcaculatetime, caculatetime time.Time) (map[string]string, error) {
	type Result struct {
		Address string `bson:"address"`
	}
	var (
		res   []Result
		addrs = make(map[string]string, 1024)
	)
	cond := bson.M{
		ExStaticDelegatorDateTag: bson.M{
			"$gte": lastcaculatetime,
			"$lt":  caculatetime,
		},
	}
	filter := bson.M{
		ExStaticDelegatorAddressTag: 1,
	}

	limit := 100
	offset := 0
	for {
		var ret []Result
		if err := querylistByOffsetAndSize(d.Name(), filter, cond, "-date", offset, limit, &ret); err != nil {
			return nil, err
		}
		length := len(ret)
		res = append(res, ret...)
		if length < limit {
			break
		}
		offset += limit
	}

	for _, one := range res {
		addrs[one.Address] = one.Address
	}

	return addrs, nil
}

func (d ExStaticDelegator) GetDataOneDay(date time.Time, address string) (ExStaticDelegator, error) {
	var res ExStaticDelegator
	cond := bson.M{
		ExStaticDelegatorDateTag:    date,
		ExStaticDelegatorAddressTag: address,
	}
	if err := queryOne(d.Name(), nil, cond, &res); err != nil && err != mgo.ErrNotFound {
		return res, err
	}

	return res, nil
}
