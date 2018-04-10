package m

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"github.com/irisnet/irisplorer.io/server/modules/store"
	"log"
	"errors"
)

const (
	DocsNmCandidate = "candidate"
)

type Candidate struct {
	Address string `json:"address" bson:"address"`
	PubKey  string `json:"pub_key" bson:"pub_key"`
	Shares  int64  `json:"shares" bson:"shares"`
}

func (d Candidate) Name() string {
	return DocsNmCandidate
}

func (d Candidate) PkKvPair() (map[string]interface{}) {
	return bson.M{"address":d.Address,"pub_key":d.PubKey}
}

func (d Candidate) Index() mgo.Index {
	return mgo.Index{
		Key:        []string{"address"}, // 索引字段， 默认升序,若需降序在字段前加-
		Unique:     false,               // 唯一索引 同mysql唯一索引
		DropDups:   false,               // 索引重复替换旧文档,Unique为true时失效
		Background: true,                // 后台创建索引
	}
}

func QueryCandidateByAddressAndPubkey(address string,pubKey string) (Candidate, error) {
	var result Candidate
	query := func(c *mgo.Collection) error {
		err := c.Find(bson.M{"address": address,"pub_key":pubKey}).Sort("-shares").One(&result)
		return err
	}

	if store.ExecCollection(DocsNmDelegator, query) != nil {
		log.Printf("delegator is Empty")
		return result, errors.New("delegator is Empty")
	}

	return result, nil
}
