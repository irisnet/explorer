package store

import (
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"errors"
	"time"
)

var (
	session *mgo.Session
)

var docs []Docs

func RegisterDocs(d Docs){
	docs = append(docs,d)
}

func Init(url string) {
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
		k, v := h.PkKvPair()
		n, _ := c.Find(bson.M{k: v}).Count()
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
		k, v := h.PkKvPair()
		n, err := c.Find(bson.M{k: v}).Count()
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
		k, v := h.PkKvPair()
		log.Printf("update %s set %+v where %s=%v", h.Name(), h, k, v)
		return c.Update(bson.M{k: v}, h)
	}
	return ExecCollection(h.Name(), update)
}
