package orm

import (
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/irishub-sync/logger"
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

func QueryList(collation string, data interface{}, m map[string]interface{}, sort string, page, size int) model.Page {
	c := GetDatabase().C(collation)
	defer c.Database.Session.Close()
	count, err := c.Find(m).Count()
	if err != nil {
		return model.Page{Count: 0, Data: nil}
	}
	err = c.Find(m).Skip((page - 1) * size).Limit(size).Sort(sort).All(data)
	if err != nil {
		return model.Page{Count: count, Data: nil}
	} else {
		return model.Page{Count: count, Data: data}
	}
}

func QueryOne(collation string, data interface{}, m map[string]interface{}) error {
	c := GetDatabase().C(collation)
	defer c.Database.Session.Close()
	return c.Find(m).One(data)
}
