package store

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	DbIrisExp = "irisplorer"
	PageSize  = 10
)

type Docs interface {
	Name() string
	PkKvPair() map[string]interface{}
	Index() mgo.Index
}

type M = []bson.M

type Query struct {
	Q bson.M
}

func NewQuery() *Query {
	return &Query{
		make(bson.M),
	}
}

func CreateQuery(key string, value interface{}) *Query {
	return &Query{
		bson.M{key: value},
	}
}

func (q *Query) And(query M) *Query {
	q.Q["$and"] = query
	return q
}

func (q *Query) Or(query M) *Query {
	q.Q["$or"] = query
	return q
}

func (q *Query) Get() bson.M {
	return q.Q
}
