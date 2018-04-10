package store

import (
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"errors"
	"time"
	"github.com/spf13/viper"
)
const (
	MgoUrl   = "mgo-url"
)


var (
	session *mgo.Session
)

var docs []Docs

func RegisterDocs(d Docs){
	docs = append(docs,d)
}

func Init() {
	url := viper.GetString(MgoUrl)
	log.Printf("state :Mgo on %s", url)
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

func Find(collection string,query interface{}) *mgo.Query {
	session := getSession()
	defer session.Close()
	c := session.DB(DbIrisExp).C(collection)
	return c.Find(query)
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
 * 执行查询，此方法可拆分做为公共方法
 * [SearchPerson description]
 * @param {[type]} collectionName string [description]
 * @param {[type]} query          bson.M [description]
 * @param {[type]} sort           bson.M [description]
 * @param {[type]} fields         bson.M [description]
 * @param {[type]} skip           int    [description]
 * @param {[type]} limit          int)   (results      []interface{}, err error [description]
 */
func Query(collectionName string, query bson.M, sort string, fields bson.M, skip int, limit int) (results []interface{}, err error) {
	exop := func(c *mgo.Collection) error {
		return c.Find(query).Sort(sort).Select(fields).Skip(skip).Limit(limit).All(&results)
	}
	return results,ExecCollection(collectionName, exop)
}