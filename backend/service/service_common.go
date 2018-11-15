package service

import (
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/irishub-sync/store/document"
	"gopkg.in/mgo.v2/bson"
)

type CommonService struct {
	*BaseService
}

func GetCommon() *CommonService {
	return commonService
}

func (service CommonService) QueryText(text string) []model.ResultVo {
	db := service.GetDb()
	defer db.Session.Close()

	var result []model.ResultVo
	i, isUint := utils.ParseUint(text)
	if isUint {
		//查询block信息
		var block document.Block
		blockStore := db.C(document.CollectionNmBlock)
		err := blockStore.Find(bson.M{"height": i}).One(&block)
		if err == nil {
			vo := model.ResultVo{
				Type: "block",
				Data: model.SimpleBlock{
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
		err = proposalStore.Find(bson.M{"proposal_id": i}).One(&proposal)
		if err == nil {
			vo := model.ResultVo{
				Type: "proposal",
				Data: model.SimpleProposal{
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
