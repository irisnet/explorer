package document

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionNmGovParams           = "ex_gov_params"
	GovParamsFieldModule            = "module"
	GovParamsFieldKey               = "key"
	GovParamsFieldCurrentValue      = "current_value"
	EQ                         Sign = "eq"
	NEQ                        Sign = "neq"
)

type GovParams struct {
	Module       string `bson:"module" json:"module"`
	Key          string `bson:"key" json:"key"`
	Type         string `bson:"type" json:"type"`
	Range        string `bson:"range" json:"range"`
	GenesisValue string `bson:"genesis_value" json:"genesis_value"`
	CurrentValue string `bson:"current_value" json:"current_value"`
	Description  string `bson:"description" json:"description"`
	Note         string `bson:"note" json:"note"`
}

type Range struct {
	Minimum Op `bson:"minimum" json:"minimum"`
	Maximum Op `bson:"maximum" json:"maximum"`
}

type Op struct {
	Sign Sign   `bson:"sign" json:"sign"`
	Data string `bson:"data" json:"data"`
}

type Sign string

func (g GovParams) Name() string {
	return CollectionNmGovParams
}

func (g GovParams) PkKvPair() map[string]interface{} {
	return bson.M{GovParamsFieldModule: g.Module, GovParamsFieldKey: g.Key}
}
