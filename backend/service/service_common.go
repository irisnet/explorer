package service

import (
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"gopkg.in/mgo.v2/bson"
)

type CommonService struct {
	BaseService
	genesis lcd.GenesisVo
}

func (service *CommonService) GetModule() Module {
	return Common
}

func (service CommonService) QueryText(text string) []model.ResultVo {
	db := getDb()
	defer db.Session.Close()

	var result []model.ResultVo
	i, isUint := utils.ParseUint(text)
	if isUint {
		//查询block信息
		var block document.Block
		blockStore := db.C(document.CollectionNmBlock)
		err := blockStore.Find(bson.M{document.Block_Field_Height: i}).One(&block)
		if err == nil {
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
		var proposal document.Proposal
		proposalStore := db.C(document.CollectionNmProposal)
		err = proposalStore.Find(bson.M{document.Proposal_Field_ProposalId: i}).One(&proposal)
		if err == nil {
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
		}
	}
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
	dbm := getDb()
	defer dbm.Session.Close()

	var configs []document.Config
	configStore := dbm.C(document.CollectionNmConfig)
	if err := configStore.Find(nil).All(&configs); err != nil {
		panic(types.ErrForEmpty("config document is not set"))
	}
	return configs
}
