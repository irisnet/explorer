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

func Init(url string){
	log.Printf("state :Mgo on %s",url)
	var err error
	session, err = mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	cNames := []DocsHander{
		new(CoinTx),new(StakeTx),new(SyncBlock),
	}
	index(cNames)
}

func InitWithAuth(addrs []string,username,password string){
	dialInfo := &mgo.DialInfo{
		Addrs:     addrs,//[]string{"192.168.6.122"}
		Direct:    false,
		Timeout:   time.Second * 1,
		Database:  DbCosmosTxn,
		Username:  username,
		Password:  password,
		PoolLimit: 4096, // Session.SetPoolLimit
	}

	session, err := mgo.DialWithInfo(dialInfo)
	session.SetMode(mgo.Monotonic, true)
	if nil != err {
		panic(err)
	}
	cNames := []DocsHander{
		new(CoinTx),new(StakeTx),new(SyncBlock),
	}
	index(cNames)
}

func getSession() *mgo.Session {
	//最大连接池默认为4096
	return session.Clone()
}

//公共方法，获取collection对象
func execCollection(collection string, s func(*mgo.Collection) error) error {
	session := getSession()
	defer session.Close()
	c := session.DB(DbCosmosTxn).C(collection)
	return s(c)
}

func index(docsHander []DocsHander){
	if len(docsHander) == 0 {
		return
	}
	for _,h := range docsHander {
		indexKey := func(c *mgo.Collection) error{
			return c.EnsureIndex(h.Index())
		}
		execCollection(h.Name(),indexKey)
	}
}

func Save(h DocsHander) error{
	save := func(c *mgo.Collection) error{
		//先按照关键字查询，如果存在，直接返回
		k,v := h.KvPair()
		n,_ := c.Find(bson.M{k: v}).Count()
		if n >= 1 {
			return errors.New("record existed")
		}
		return c.Insert(h)
	}

	return execCollection(h.Name(),save)
}

func QueryAllCoinTx() []CoinTx{
	result := []CoinTx{}
	query := func(c *mgo.Collection) error{
		err := c.Find(nil).Sort("-time").All(&result)
		return err
	}

	if execCollection(DocsNmCoinTx,query) != nil {
		log.Printf("CoinTx is Empry")
	}
	return result

}

func QueryAllStakeTx() ([]StakeTx)  {
	result := []StakeTx{}
	query := func(c *mgo.Collection) error{
		err := c.Find(nil).Sort("-time").All(&result)
		return err
	}

	if execCollection(DocsNmStakeTx,query) != nil {
		log.Printf("CoinTx is Empry")
	}
	return result
}
//
func QueryCoinTxsByFrom(from string) ([]CoinTx)  {
	result := []CoinTx{}
	query := func(c *mgo.Collection) error{
		err := c.Find(bson.M{"from": from}).Sort("-time").All(&result)
		return err
	}

	if execCollection(DocsNmCoinTx,query) != nil {
		log.Printf("CoinTx is Empry")
	}
	return result
}
//
func QueryCoinTxsByAccount(account string) ([]CoinTx)  {
	result := []CoinTx{}
	query := func(c *mgo.Collection) error{
		err := c.Find(bson.M{"$or": []bson.M{bson.M{"from": account}, bson.M{"to": account}}}).Sort("-time").All(&result)
		return err
	}

	if execCollection(DocsNmCoinTx,query) != nil {
		log.Printf("CoinTx is Empry")
	}
	return result
}
//
func QueryStakeTxsByAccount(account string) ([]StakeTx)  {
	result := []StakeTx{}
	query := func(c *mgo.Collection) error{
		err := c.Find(bson.M{"from": account}).Sort("-time").All(&result)
		return err
	}

	if execCollection(DocsNmStakeTx,query) != nil {
		log.Printf("StakeTx is Empry")
	}

	return result
}
//
func QueryPageCoinTxsByFrom(from string,page int)([]CoinTx)  {
	result := []CoinTx{}
	query := func(c *mgo.Collection) error{
		skip := (page-1) * PageSize
		err := c.Find(bson.M{"from": from}).Sort("-time").Skip(skip).Limit(PageSize).All(&result)
		return err
	}

	if execCollection(DocsNmCoinTx,query) != nil {
		log.Printf("StakeTx is Empry")
	}

	return result
}
//
func QueryLastedBlock()(SyncBlock,error){
	result := SyncBlock{}

	query := func(c *mgo.Collection) error{
		err := c.Find(bson.M{}).One(&result)
		return err
	}

	err := execCollection(DocsNmSyncBlock,query)
	return result,err
}
//
func UpdateBlock(b SyncBlock) error{
	update := func(c *mgo.Collection) error{
		return c.Update(nil,b)
	}
	return execCollection(DocsNmSyncBlock,update)
}