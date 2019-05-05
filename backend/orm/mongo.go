package orm

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/model"
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
		Timeout:   time.Second * 10,
		PoolLimit: conf.Get().Db.PoolLimit,
	}

	var err error
	session, err = mgo.DialWithInfo(dialInfo)
	if err != nil {
		logger.Error("start mongo client failed", logger.String("err", err.Error()))
	}
	session.SetMode(mgo.Monotonic, true)
}

var session *mgo.Session

func GetDatabase() *mgo.Database {
	return session.Clone().DB(conf.Get().Db.Database)
}

// TODO will replace with `AllWithCount`
func QueryRows(collation string, data interface{}, m map[string]interface{}, sort string, page, size int) model.PageVo {
	c := GetDatabase().C(collation)
	defer c.Database.Session.Close()
	count, err := c.Find(m).Count()
	if err != nil {
		logger.Error("QueryRows Count failed", logger.String("err", err.Error()))
		return model.PageVo{Count: 0, Data: nil}
	}
	err = c.Find(m).Skip((page - 1) * size).Limit(size).Sort(sort).All(data)
	if err != nil {
		logger.Error("QueryRows Find failed", logger.String("err", err.Error()))
		return model.PageVo{Count: count, Data: nil}
	} else {
		return model.PageVo{Count: count, Data: data}
	}
}

type Query struct {
	db         *mgo.Database
	collection string
	result     interface{}
	condition  bson.M
	sort       string
	page       int
	size       int
	selector   interface{}
}

func NewQuery() *Query {
	var q = &Query{}
	q.db = GetDatabase()
	return q
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

func (query *Query) ExecPage() (cnt int, err error) {
	var c = query.db.C(query.collection)
	var q = c.Find(query.condition)
	cnt, err = q.Count()
	if err != nil || cnt == 0 {
		return cnt, errors.New(fmt.Sprintf("query error,collection:%s,condition:%+v",
			query.collection, query.condition))
	}
	q = query.buildQuery()
	return cnt, q.All(query.result)
}

func (query *Query) Reset() *Query {
	query.collection = ""
	query.result = nil
	query.condition = nil
	query.sort = ""
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
func (query *Query) SetSort(sort string) *Query {
	query.sort = sort
	return query
}
func (query *Query) SetPage(page int) *Query {
	query.page = page
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
	if query.sort != "" {
		q = q.Sort(query.sort)
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
