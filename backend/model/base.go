package model

import "gopkg.in/mgo.v2/bson"

type Count struct {
	Id    bson.ObjectId `bson:"_id,omitempty"`
	Count float64
}

type Page struct {
	Count int
	Data  interface{}
}
