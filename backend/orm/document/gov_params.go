package document

import (
	"encoding/json"
	"fmt"
	"github.com/irisnet/explorer/backend/utils"

	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
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
	Module       string      `bson:"module" json:"module"`
	Key          string      `bson:"key" json:"key"`
	Type         string      `bson:"type" json:"type"`
	Range        string      `bson:"range" json:"range"`
	InitialValue string      `bson:"initial_value" json:"initial_value"`
	GenesisValue string      `bson:"genesis_value" json:"genesis_value"`
	CurrentValue interface{} `bson:"current_value" json:"current_value"`
	Description  string      `bson:"description" json:"description"`
	Note         string      `bson:"note" json:"note"`
}

type AmountCurrentValue struct {
	Amount string `bson:"amount" json:"amount"`
	Denom  string `bson:"denom" json:"denom"`
}

func (v *AmountCurrentValue) BuildAmountCurrentValueByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, v)
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

func (_ GovParams) QueryAll() ([]GovParams, error) {

	var params []GovParams

	var query = orm.NewQuery()
	defer query.Release()

	query.SetCollection(CollectionNmGovParams).
		SetCondition(nil).
		SetSelector(nil).
		SetSort(asc(GovParamsFieldModule), asc(GovParamsFieldKey)).
		SetSize(0).
		SetResult(&params)

	err := query.Exec()

	return params, err
}

func (_ GovParams) Batch(txs []txn.Op) error {
	return orm.Batch(txs)
}

func (_ GovParams) UpdateCurrentModuleParamValue(kv map[string]interface{}) error {
	collection := getDb().C(CollectionNmGovParams)
	defer collection.Database.Session.Close()
	bulk := collection.Bulk()

	for k, v := range kv {
		vStr := ""
		update := bson.M{}
		switch vType := v.(type) {
		case string:
			vStr = vType
			update["$set"] = bson.M{GovParamsFieldCurrentValue: vStr}
		case map[string]interface{}:
			currentValueMap := AmountCurrentValue{}
			if err := currentValueMap.BuildAmountCurrentValueByUnmarshalJson(utils.MarshalJsonIgnoreErr(vType)); err != nil {
				logger.Error("BuildAmountCurrentValueByUnmarshalJson have error", logger.String("err", err.Error()))
			}
			update["$set"] = bson.M{GovParamsFieldCurrentValue: currentValueMap}
		default:
			vStr = fmt.Sprintf("%v", vType)
			update["$set"] = bson.M{GovParamsFieldCurrentValue: vStr}
		}
		sel := bson.M{GovParamsFieldKey: k}

		bulk.Update(sel, update)

	}

	updateRes, err := bulk.Run()

	if err != nil {
		return err
	}
	logger.Info("batch upsert reesult", logger.Any("bulk res", *updateRes))

	return nil

}
