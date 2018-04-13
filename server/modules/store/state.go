package store

import (
	"errors"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

const (
	MgoUrl = "mgo-url"
)

var (
	session *mgo.Session
)

var docs []Docs

func RegisterDocs(d Docs) {
	docs = append(docs, d)
}

func Init() {
	url := viper.GetString(MgoUrl)
	log.Printf("Mgo start on %s", url)
	var err error
	session, err = mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	index()
}

func InitWithAuth(addrs []string, username, password string) {
	dialInfo := &mgo.DialInfo{
		Addrs:     addrs, //[]string{"192.168.6.122"}
		Direct:    false,
		Timeout:   time.Second * 1,
		Database:  DbIrisExp,
		Username:  username,
		Password:  password,
		PoolLimit: 4096, // Session.SetPoolLimit
	}

	session, err := mgo.DialWithInfo(dialInfo)
	session.SetMode(mgo.Monotonic, true)
	if nil != err {
		panic(err)
	}
	index()
}

func getSession() *mgo.Session {
	//最大连接池默认为4096
	return session.Clone()
}

//公共方法，获取collection对象
func ExecCollection(collection string, s func(*mgo.Collection) error) error {
	session := getSession()
	defer session.Close()
	c := session.DB(DbIrisExp).C(collection)
	return s(c)
}

func index() {
	if len(docs) == 0 {
		return
	}
	for _, h := range docs {
		indexKey := func(c *mgo.Collection) error {
			return c.EnsureIndex(h.Index())
		}
		ExecCollection(h.Name(), indexKey)
	}
}

func Save(h Docs) error {
	save := func(c *mgo.Collection) error {
		//先按照关键字查询，如果存在，直接返回
		n, _ := c.Find(h.PkKvPair()).Count()
		if n >= 1 {
			return errors.New("record existed")
		}
		log.Printf("insert %s  %+v ", h.Name(), h)
		return c.Insert(h)
	}

	return ExecCollection(h.Name(), save)
}

func SaveOrUpdate(h Docs) error {
	save := func(c *mgo.Collection) error {
		//先按照关键字查询，如果存在，直接返回
		n, err := c.Find(h.PkKvPair()).Count()
		log.Printf("Count:%d err:%+v ", n, err)
		if n >= 1 {
			return Update(h)
		}
		log.Printf("insert %s  %+v ", h.Name(), h)
		return c.Insert(h)
	}

	return ExecCollection(h.Name(), save)
}

//
func Update(h Docs) error {
	update := func(c *mgo.Collection) error {
		key := h.PkKvPair()
		log.Printf("update %s set %+v where %+v", h.Name(), h, key)
		return c.Update(h.PkKvPair(), h)
	}
	return ExecCollection(h.Name(), update)
}

/**
 * 执行查询
 * [SearchPerson description]
 * @param {[type]} cName 		  string [description]
 * @param {[type]} query          *Query [description]
 * @param {[type]} sort           string [description]
 * @param {[type]} skip           int    [description]
 * @param {[type]} limit          int)   (results      []interface{}, err error [description]
 */
func SelectByPage(cName string, query *Query, sort string, page int, size int) (results []interface{}, err error) {
	exop := func(c *mgo.Collection) error {
		skip := (page - 1) * size
		if len(sort) > 0 {
			return c.Find(query.Get()).Sort(sort).Skip(skip).Limit(size).All(&results)
		}
		log.Printf("select * from %s where %+v and skip=%d and limit=%d order by %s", cName, query, skip, size, sort)
		return c.Find(query.Get()).Skip(skip).Limit(size).All(&results)
	}
	return results, ExecCollection(cName, exop)
}

func Select(cName string, query *Query, sort string) (results []interface{}, err error) {
	exop := func(c *mgo.Collection) error {
		if len(sort) > 0 {
			return c.Find(query.Get()).Sort(sort).All(&results)
		}
		log.Printf("select * from %s where %+v order by %s", cName, query, sort)
		return c.Find(query.Get()).All(&results)
	}
	return results, ExecCollection(cName, exop)
}
