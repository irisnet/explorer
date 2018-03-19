package store

import (
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"errors"
)
type MgoBackend struct {
	Session *mgo.Session
}


var Mgo = MgoBackend{}

func (m *MgoBackend) Init(url string){
	log.Printf("state :Mgo on %s",url)
	var session, err = mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	m.Session = session
	m.index()
}

func (m *MgoBackend) index(){
	c := m.Session.DB(DbCosmosTxn).C(TbNmCoinTx)

	index := mgo.Index{
		Key:        []string{"from"}, // 索引字段， 默认升序,若需降序在字段前加-
		Unique:     false,             // 唯一索引 同mysql唯一索引
		DropDups:   false,             // 索引重复替换旧文档,Unique为true时失效
		Background: true,             // 后台创建索引
	}

	c.EnsureIndex(index)

	c = m.Session.DB(DbCosmosTxn).C(TbNmStakeTx)
	c.EnsureIndex(index)
}

func (m *MgoBackend) Save(tx TxHander) error{
	c := m.Session.DB(DbCosmosTxn).C(tx.TbNm())
	//先按照关键字查询，如果存在，直接返回
	k,v := tx.KvPair()
	n,err := c.Find(bson.M{k: v}).Count()
	if(n >= 1){
		return errors.New("record existed")
	}

	err = c.Insert(tx)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (m *MgoBackend) QueryCoinTxs() ([]CoinTx)  {
	result := []CoinTx{}
	c := m.Session.DB(DbCosmosTxn).C(TbNmCoinTx)
	err := c.Find(nil).Sort("-time").All(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func (m *MgoBackend) QueryStakeTxs() ([]StakeTx)  {
	result := []StakeTx{}
	c := m.Session.DB(DbCosmosTxn).C(TbNmStakeTx)
	err := c.Find(nil).Sort("-time").All(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func (m *MgoBackend) QueryCoinTxsByFrom(from string) ([]CoinTx)  {
	result := []CoinTx{}
	c := m.Session.DB(DbCosmosTxn).C(TbNmCoinTx)
	err := c.Find(bson.M{"from": from}).Sort("-time").All(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func (m *MgoBackend) QueryCoinTxsByAccount(account string) ([]CoinTx)  {
	result := []CoinTx{}
	c := m.Session.DB(DbCosmosTxn).C(TbNmCoinTx)
	err := c.Find(bson.M{"$or": []bson.M{bson.M{"from": account}, bson.M{"to": account}}}).Sort("-time").All(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func (m *MgoBackend) QueryStakeTxsByAccount(account string) ([]StakeTx)  {
	result := []StakeTx{}
	c := m.Session.DB(DbCosmosTxn).C(TbNmStakeTx)
	err := c.Find(bson.M{"from": account}).Sort("-time").All(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func (m *MgoBackend) QueryPageCoinTxsByFrom(from string,page int)([]CoinTx)  {
	result := []CoinTx{}
	c := m.Session.DB(DbCosmosTxn).C(TbNmCoinTx)
	skip := (page-1) * PageSize
	err := c.Find(bson.M{"from": from}).Sort("-time").Skip(skip).Limit(PageSize).All(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func (m *MgoBackend) QueryLastedBlock()(SyncBlock,error){
	result := SyncBlock{}
	c := m.Session.DB(DbCosmosTxn).C(TbNmSyncBlock)
	err := c.Find(bson.M{}).One(&result)
	return result,err
}

func (m *MgoBackend) UpdateBlock(b SyncBlock) error{
	c := m.Session.DB(DbCosmosTxn).C(TbNmSyncBlock)
	return c.Update(nil,b)
}