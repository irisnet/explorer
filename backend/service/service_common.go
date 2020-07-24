package service

import (
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/explorer/backend/vo"
)

type CommonService struct {
	BaseService
}

func (service CommonService) QueryText(text string) vo.ResultVo {
	result := vo.ResultVo{}
	i, isUint := utils.ParseInt(text)

	if !isUint {
		return result
	}

	//查询block信息
	if isUint {
		block, err := document.Block{}.QueryBlockHeightTimeHashByHeight(i)
		if err != nil {
			logger.Error("queryBlockHeightTimehashByHeight", logger.String("err", err.Error()))
		} else {

			blockAsModel := vo.SimpleBlockVo{
				Height:    block.Height,
				Timestamp: block.Time.String(),
				Hash:      block.Hash,
			}
			result.Block = blockAsModel
		}
	}

	//查询proposal信息
	proposal, err := document.Proposal{}.QueryById(i)

	if err != nil {
		logger.Error("Query proposal By Id", logger.String("err", err.Error()))
	} else {
		vo := vo.SimpleProposalVo{
			ProposalId: proposal.ProposalId,
			Title:      proposal.Title,
			Type:       proposal.Type,
			Status:     proposal.Status,
			SubmitTime: proposal.SubmitTime.String(),
		}
		result.Proposal = vo
	}

	return result
}

//func (service CommonService) GetGenesis() lcd.GenesisVo {
//	if len(service.genesis.Result.Genesis.ChainID) == 0 {
//		result, err := lcd.Genesis()
//		if err != nil {
//			logger.Error("lcd Genesis have error", logger.String("err", err.Error()))
//		} else {
//			service.genesis = result
//		}
//	}
//	return service.genesis
//}

func (service CommonService) GetConfig() []document.Config {

	configs, err := document.Config{}.GetConfig()

	if err != nil {
		panic(types.ErrForEmpty("config document is not set"))
	}
	return configs
}
