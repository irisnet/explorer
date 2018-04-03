package store

import (
	"gopkg.in/mgo.v2"
)

const (
	DbIrisExp = "irisplorer"
	PageSize  = 10
)

type Docs interface {
	Name() string
	PkKvPair() (string, interface{})
	Index() mgo.Index
}
