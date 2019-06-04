package service

import (
	"fmt"

	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
)

type GovParamsService struct {
	BaseService
}

func (service *GovParamsService) GetModule() Module {
	return GovParams
}

func (service *GovParamsService) QueryAll() []document.GovParams {
	var params []document.GovParams
	err := queryAll(document.CollectionNmGovParams, nil, nil, desc(document.GovParamsFieldModule), 0, &params)
	if err != nil {
		panic(types.CodeNotFound)
	}
	return params
}

func (_ GovParamsService) GetGovAuthParamList(genesisMap, currentMap map[string]interface{}) ([]document.GovParams, error) {
	authKeyRangeMap := lcd.GetAuthKeyWithRangeMap()
	result := make([]document.GovParams, 0, len(authKeyRangeMap))
	for k, v := range authKeyRangeMap {
		genValueStr := ""
		switch genV := genesisMap[k].(type) {
		case string:
			genValueStr = genV
		}
		currentValueStr := ""
		switch curV := currentMap[k].(type) {
		case string:
			currentValueStr = curV
		}

		tmp := document.GovParams{
			Module:       lcd.GovModuleAuth,
			Key:          k,
			Range:        v.Range,
			Description:  v.Description,
			GenesisValue: genValueStr,
			CurrentValue: currentValueStr,
		}
		result = append(result, tmp)
	}
	return result, nil
}

func (_ GovParamsService) GetGovStakeParamList(genesisMap, currentMap map[string]interface{}) ([]document.GovParams, error) {
	stakeKeyRangeMap := lcd.GetStakeKeyWithRangeMap()
	result := make([]document.GovParams, 0, len(stakeKeyRangeMap))
	for k, v := range stakeKeyRangeMap {
		genValueStr := ""
		switch genV := genesisMap[k].(type) {
		case string:
			genValueStr = genV
		default:
			genValueStr = fmt.Sprintf("%v", genV)
		}
		currentValueStr := ""

		switch curV := currentMap[k].(type) {
		case string:
			currentValueStr = curV
		default:
			currentValueStr = fmt.Sprintf("%v", curV)
		}

		tmp := document.GovParams{
			Module:       lcd.GovModuleStake,
			Key:          k,
			Range:        v.Range,
			Description:  v.Description,
			GenesisValue: genValueStr,
			CurrentValue: currentValueStr,
		}
		result = append(result, tmp)
	}
	return result, nil
}

func (_ GovParamsService) GetGovDistrParamList(genesisMap, currentMap map[string]interface{}) ([]document.GovParams, error) {
	distrKeyRangeMap := lcd.GetDistrKeyWithRangeMap()
	result := make([]document.GovParams, 0, len(distrKeyRangeMap))
	for k, v := range distrKeyRangeMap {
		genValueStr := ""
		switch genV := genesisMap[k].(type) {
		case string:
			genValueStr = genV
		}
		currentValueStr := ""
		switch curV := currentMap[k].(type) {
		case string:
			currentValueStr = curV
		}

		tmp := document.GovParams{
			Module:       lcd.GovModuleDistr,
			Key:          k,
			Range:        v.Range,
			Description:  v.Description,
			GenesisValue: genValueStr,
			CurrentValue: currentValueStr,
		}
		result = append(result, tmp)
	}
	return result, nil
}

func (_ GovParamsService) GetGovMintParamList(genesisMap, currentMap map[string]interface{}) ([]document.GovParams, error) {
	mintKeyRangeMap := lcd.GetMintKeyWithRangeMap()
	result := make([]document.GovParams, 0, len(mintKeyRangeMap))
	for k, v := range mintKeyRangeMap {
		genValueStr := ""
		switch genV := genesisMap[k].(type) {
		case string:
			genValueStr = genV
		}
		currentValueStr := ""
		switch curV := currentMap[k].(type) {
		case string:
			currentValueStr = curV
		}

		tmp := document.GovParams{
			Module:       lcd.GovModuleMint,
			Key:          k,
			Range:        v.Range,
			Description:  v.Description,
			GenesisValue: genValueStr,
			CurrentValue: currentValueStr,
		}
		result = append(result, tmp)
	}
	return result, nil
}

func (_ GovParamsService) GetGovSlashingParamList(genesisMap, currentMap map[string]interface{}) ([]document.GovParams, error) {
	slashingKeyRangeMap := lcd.GetSlashingKeyWithRangeMap()
	result := make([]document.GovParams, 0, len(slashingKeyRangeMap))
	for k, v := range slashingKeyRangeMap {
		genValueStr := ""
		switch genV := genesisMap[k].(type) {
		case string:
			genValueStr = genV
		}
		currentValueStr := ""
		switch curV := currentMap[k].(type) {
		case string:
			currentValueStr = curV
		}

		tmp := document.GovParams{
			Module:       lcd.GovModuleSlashing,
			Key:          k,
			Range:        v.Range,
			Description:  v.Description,
			GenesisValue: genValueStr,
			CurrentValue: currentValueStr,
		}
		result = append(result, tmp)
	}
	return result, nil
}

func (gov GovParamsService) GetDbInitGovModuleParamList(genesisMap, currentMap map[string]interface{}) ([]document.GovParams, error) {

	authList, err := gov.GetGovAuthParamList(genesisMap, currentMap)
	if err != nil {
		return nil, err
	}

	stakeList, err := gov.GetGovStakeParamList(genesisMap, currentMap)
	if err != nil {
		return nil, err
	}

	stakeList = append(stakeList, authList...)

	distrList, err := gov.GetGovDistrParamList(genesisMap, currentMap)
	if err != nil {
		return nil, err
	}

	distrList = append(distrList, stakeList...)

	mintList, err := gov.GetGovMintParamList(genesisMap, currentMap)
	if err != nil {
		return nil, err
	}

	mintList = append(mintList, distrList...)

	slashingList, err := gov.GetGovSlashingParamList(genesisMap, currentMap)
	if err != nil {
		return nil, err
	}
	return append(slashingList, mintList...), nil
}

func (gov GovParamsService) UpdateCurrentValueByKey(kv map[string]interface{}) error {

	bulk := getDb().C(document.CollectionNmGovParams).Bulk()

	for k, v := range kv {
		vStr := ""
		switch vType := v.(type) {
		case string:
			vStr = vType
		default:
			vStr = fmt.Sprintf("%v", vType)
		}
		sel := bson.M{document.GovParamsFieldKey: k}
		update := bson.M{"$set": bson.M{document.GovParamsFieldCurrentValue: vStr}}
		bulk.Upsert(sel, update)
	}

	upsertRes, err := bulk.Run()

	if err != nil {
		return err
	}
	logger.Info("batch upsert reesult", logger.Any("bulk res", *upsertRes))

	return nil
}

func init() {
	var initParams = func() {
		var ops []txn.Op
		genGovModuleMap, _ := lcd.GetGenesisGovModuleParamMap()
		currentParamMap, _ := lcd.GetAllGovModuleParam()
		govParamList, _ := GovParamsService{}.GetDbInitGovModuleParamList(genGovModuleMap, currentParamMap)

		for _, v := range govParamList {
			ops = append(ops, txn.Op{
				C:      document.CollectionNmGovParams,
				Id:     bson.NewObjectId(),
				Insert: v,
			})
		}

		if err := orm.Batch(ops); err != nil {
			logger.Error("init gov_params data error", logger.String("err", err.Error()))
		}
	}
	if data := govParamsService.QueryAll(); len(data) == 0 {
		initParams()
	}
}
