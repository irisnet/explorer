package document

import (
	"gopkg.in/mgo.v2"
)

type (
	Doc interface {
		Name() string
		PkKvPair() map[string]interface{}
		EnsureIndexes() []mgo.Index
	}
)

var (
	TaskControlModel TaskControl

	Docs = []Doc{
		TaskControlModel,
		new(Config),
		new(GovParams),
		new(AssetToken),
		new(BlackList),
		new(Proposal),
		new(Validator),
	}
)
