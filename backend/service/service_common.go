package service

import (
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/irishub-sync/logger"
)

type CommonService struct {
	BaseService
	genesis lcd.GenesisVo
}

func (service *CommonService) GetModule() Module {
	return Common
}

func (service CommonService) QueryText(text string) []model.ResultVo {
	var result []model.ResultVo
	i, isUint := utils.ParseInt(text)

	if !isUint {
		return result
	}

	//查询block信息
	if isUint {
		block, err := document.Block{}.QueryBlockHeightTimeHashByHeight(i)
		if err != nil {
			logger.Error("queryBlockHeightTimehashByHeight", logger.String("err", err.Error()))
		}
		vo := model.ResultVo{
			Type: "block",
			Data: model.SimpleBlockVo{
				Height:    block.Height,
				Timestamp: block.Time,
				Hash:      block.Hash,
			},
		}
		result = append(result, vo)
	}

	//查询proposal信息
	proposal, err := document.Proposal{}.QueryById(i)

	if err != nil {
		logger.Error("Query proposal By Id", logger.String("err", err.Error()))
	}

	vo := model.ResultVo{
		Type: "proposal",
		Data: model.SimpleProposalVo{
			ProposalId: proposal.ProposalId,
			Title:      proposal.Title,
			Type:       proposal.Type,
			Status:     proposal.Status,
			SubmitTime: proposal.SubmitTime,
		},
	}
	result = append(result, vo)

	return result
}

func (service CommonService) GetGenesis() lcd.GenesisVo {
	if len(service.genesis.Result.Genesis.ChainID) == 0 {
		result, err := lcd.Genesis()
		if err != nil {
			panic(err)
		}
		service.genesis = result
	}
	return service.genesis
}

func (service CommonService) GetConfig() []document.Config {

	configs, err := document.Config{}.GetConfig()

	if err != nil {
		panic(types.ErrForEmpty("config document is not set"))
	}
	return configs
}
