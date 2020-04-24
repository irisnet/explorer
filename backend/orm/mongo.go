package orm

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/logger"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"

	"gopkg.in/mgo.v2"
)

func init() {

	dialInfo := &mgo.DialInfo{
		Addrs:     conf.Get().Db.Addrs,
		Database:  conf.Get().Db.Database,
		Username:  conf.Get().Db.UserName,
		Password:  conf.Get().Db.Password,
		Direct:    false,
		Timeout:   time.Second * 20,
		PoolLimit: conf.Get().Db.PoolLimit,
	}

	var err error
	session, err = mgo.DialWithInfo(dialInfo)
	if err != nil {
		logger.Error("start mongo client failed", logger.String("err", err.Error()))
	}
	session.SetMode(mgo.Nearest, true)
}

var session *mgo.Session

func GetDatabase() *mgo.Database {
	return session.Clone().DB(conf.Get().Db.Database)
}

type Query struct {
	db         *mgo.Database
	collection string
	result     interface{}
	condition  bson.M
	sort       []string
	page       int
	offset     int
	size       int
	selector   interface{}
}

func NewQuery() *Query {
	var q = &Query{}
	q.db = GetDatabase()
	return q
}

func Save(collectionname string, doc interface{}) error {
	db := GetDatabase()
	defer db.Session.Close()
	c := db.C(collectionname)
	return c.Insert(doc)
}

func Update(collectionname string, selecter, update interface{}) error {
	db := GetDatabase()
	defer db.Session.Close()

	c := db.C(collectionname)
	return c.Update(selecter, update)
}

func (query *Query) GetDb() *mgo.Database {
	return query.db
}

func (query *Query) Exec() error {
	q := query.buildQuery()

	vType := reflect.ValueOf(query.result)
	switch vType.Elem().Kind() {
	case reflect.Slice:
		return q.All(query.result)
	default:
		return q.One(query.result)
	}
}
func (query *Query) Count() (cnt int, err error) {
	q := query.buildQuery()
	return q.Count()
}

func (query *Query) PipeQuery(pipeline interface{}) error {
	var c = query.db.C(query.collection)
	var pipe = c.Pipe(pipeline)
	vType := reflect.ValueOf(query.result)
	switch vType.Elem().Kind() {
	case reflect.Slice:
		return pipe.All(query.result)
	default:
		return pipe.One(query.result)
	}
}

func (query *Query) ExecPage(total bool) (cnt int, err error) {
	var c = query.db.C(query.collection)
	var q = c.Find(query.condition)
	if total {
		cnt, err = q.Count()
		if err != nil {
			return cnt, errors.New(fmt.Sprintf("query error:%v,collection:%s,condition:%+v",
				err.Error(), query.collection, query.condition))
		}
	}
	if cnt == 0 && total {
		return cnt, errors.New("not found")
	}
	q = query.buildQuery()
	return cnt, q.All(query.result)
}

func (query *Query) Reset() *Query {
	query.collection = ""
	query.result = nil
	query.condition = nil
	query.sort = nil
	query.page = 0
	query.size = 0
	query.selector = nil
	return query
}

func (query *Query) SetCollection(collection string) *Query {
	query.collection = collection
	return query
}
func (query *Query) SetResult(result interface{}) *Query {
	query.result = result
	return query
}
func (query *Query) SetCondition(condition bson.M) *Query {
	query.condition = condition
	return query
}
func (query *Query) SetSort(sort ...string) *Query {

	strArr := []string{}
	for _, v := range sort {
		if v != "" {
			strArr = append(strArr, v)
		}
	}
	query.sort = strArr
	return query
}
func (query *Query) SetPage(page int) *Query {
	query.page = page
	return query
}

func (query *Query) SetOffset(offset int) *Query {
	query.offset = offset
	return query
}

func (query *Query) SetSize(size int) *Query {
	query.size = size
	return query
}

func (query *Query) SetSelector(selector interface{}) *Query {
	query.selector = selector
	return query
}

func (query *Query) Release() {
	if query.db == nil {
		return
	}
	query.db.Session.Close()
}

func (query *Query) buildQuery() *mgo.Query {
	var c = query.db.C(query.collection)
	var q = c.Find(query.condition)
	if query.selector != nil {
		q = q.Select(query.selector)
	}
	if query.size != 0 {
		q = q.Limit(query.size)
	}
	if query.page != 0 {
		q = q.Skip((query.page - 1) * query.size)
	}

	if query.offset != 0 {
		q = q.Skip(query.offset)
	}

	if query.sort != nil && len(query.sort) > 0 {
		q = q.Sort(query.sort...)
	}
	return q
}

// mgo transaction method
// detail to see: https://godoc.org/gopkg.in/mgo.v2/txn
func Batch(ops []txn.Op) error {
	if len(ops) == 0 {
		return nil
	}
	c := GetDatabase().C("mgo_txn")
	defer c.Database.Session.Close()
	runner := txn.NewRunner(c)

	txObjectId := bson.NewObjectId()
	err := runner.Run(ops, txObjectId, nil)
	if err != nil {
		if err == txn.ErrAborted {
			err = runner.Resume(txObjectId)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}
