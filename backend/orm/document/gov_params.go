package document

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionNmGovParams = "gov_params"
	GovParamsFieldModule  = "module"
	GovParamsFieldKey     = "key"
)

type GovParams struct {
	Module      string `bson:"module" json:"module"`
	Key         string `bson:"key" json:"key"`
	Value       string `bson:"value" json:"value"`
	Type        string `bson:"type" json:"type"`
	Range       Range  `bson:"range" json:"range"`
	Description string `bson:"description" json:"description"`
	Note        string `bson:"note" json:"note"`
}

type Range struct {
	Minimum string `bson:"minimum" json:"minimum"`
	Maximum string `bson:"maximum" json:"maximum"`
}

func (g GovParams) Name() string {
	return CollectionNmGovParams
}

func (g GovParams) PkKvPair() map[string]interface{} {
	return bson.M{GovParamsFieldModule: g.Module, GovParamsFieldKey: g.Key}
}
