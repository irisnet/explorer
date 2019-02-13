package orm

import (
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/model"
	"time"

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

// TODO will replace with `One`
func QueryRow(collation string, data interface{}, m map[string]interface{}) error {
	c := GetDatabase().C(collation)
	defer c.Database.Session.Close()
	return c.Find(m).One(data)
}

func All(query MQuery) error {
	c := GetDatabase().C(query.C)
	defer c.Database.Session.Close()
	q := buildQuery(c, query)
	return q.All(query.Result)
}

func One(query MQuery) error {
	c := GetDatabase().C(query.C)
	defer c.Database.Session.Close()
	q := buildQuery(c, query)
	return q.One(query.Result)
}

func AllWithCount(query MQuery) (int, error) {
	c := GetDatabase().C(query.C)
	defer c.Database.Session.Close()
	count, err := c.Find(query.Q).Count()
	if err != nil {
		return count, err
	}
	q := buildQuery(c, query)
	err = q.All(query.Result)
	return count, err
}

type MQuery struct {
	C        string
	Result   interface{}
	Q        map[string]interface{}
	Sort     string
	Page     int
	Size     int
	Selector interface{}
}

func buildQuery(c *mgo.Collection, query MQuery) *mgo.Query {
	var q = c.Find(query.Q)
	if query.Selector != nil {
		q = q.Select(query.Selector)
	}
	if query.Size != 0 {
		q = q.Limit(query.Size)
	}
	if query.Page != 0 {
		q = q.Skip((query.Page - 1) * query.Size)
	}
	if query.Sort != "" {
		q = q.Sort(query.Sort)
	}
	return q
}
