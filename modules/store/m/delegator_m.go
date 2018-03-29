package m

import "gopkg.in/mgo.v2"

const (
	DocsNmDelegator = "delegator"
)

type Delegator struct {
	Address string `json:"address" bson:"address"`
	PubKey  string `json:"pub_key" bson:"pub_key"`
	Shares  int64  `json:"shares" bson:"shares"`
}

func (d Delegator) Name() string {
	return DocsNmDelegator
}

func (d Delegator) PkKvPair() (string, interface{}) {
	return "address", d.Address
}

func (d Delegator) Index() mgo.Index {
	return mgo.Index{
		Key:        []string{"address"}, // 索引字段， 默认升序,若需降序在字段前加-
		Unique:     false,               // 唯一索引 同mysql唯一索引
		DropDups:   false,               // 索引重复替换旧文档,Unique为true时失效
		Background: true,                // 后台创建索引
	}
}
