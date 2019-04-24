package service

import (
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
)

type GovParamsService struct {
	BaseService
}

func (service *GovParamsService) GetModule() Module {
	return GovParams
}

func (service *GovParamsService) QueryAll() interface{} {
	var params []document.GovParams
	err := queryAll(document.CollectionNmGovParams, nil, nil, desc(document.GovParamsFieldModule), 0, &params)
	if err != nil {
		panic(types.CodeNotFound)
	}
	return params
}
