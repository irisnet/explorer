package service

import (
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/irishub-sync/store/document"
	"gopkg.in/mgo.v2/bson"
)

type ProposalService struct {
	*BaseService
}

func GetProposal() *ProposalService {
	return proposalService
}

func (service *ProposalService) QueryList(page, size int) (resp model.Page) {
	var data []document.Proposal
	resp = service.QueryPage(document.CollectionNmProposal, &data, nil, "-submit_block", page, size)

	var proposals []model.Proposal
	for _, propo := range data {
		mP := model.Proposal{
			Title:            propo.Title,
			ProposalId:       propo.ProposalId,
			Type:             propo.Type,
			Description:      propo.Description,
			Status:           propo.Status,
			SubmitBlock:      propo.SubmitBlock,
			SubmitTime:       propo.SubmitTime,
			VotingStartBlock: propo.VotingStartBlock,
		}
		proposals = append(proposals, mP)
	}
	resp.Data = proposals
	return resp
}

func (service *ProposalService) Query(id int) (resp model.ProposalInfo) {
	var data document.Proposal
	db := service.GetDb()
	defer db.Session.Close()
	propoStore := db.C(document.CollectionNmProposal)
	txStore := db.C(document.CollectionNmCommonTx)

	if err := propoStore.Find(bson.M{"proposal_id": id}).One(&data); err != nil {
		panic(types.ErrorCodeNotFound)
		return
	}

	proposal := model.Proposal{
		Title:            data.Title,
		ProposalId:       data.ProposalId,
		Type:             data.Type,
		Description:      data.Description,
		Status:           data.Status,
		SubmitBlock:      data.SubmitBlock,
		SubmitTime:       data.SubmitTime,
		TotalDeposit:     data.TotalDeposit,
		VotingStartBlock: data.VotingStartBlock,
	}

	var tx document.CommonTx
	if err := txStore.Find(bson.M{"type": types.TypeSubmitProposal, "proposal_id": id}).One(&tx); err == nil {
		proposal.Proposer = tx.From
		proposal.TxHash = tx.TxHash
	}

	resp.Proposal = proposal

	var votes []model.Vote
	var result model.VoteResult
	for _, v := range data.Votes {
		vote := model.Vote{
			Voter:  v.Voter,
			Option: v.Option,
			Time:   v.Time,
		}
		votes = append(votes, vote)

		switch v.Option {
		case "Yes":
			result.Yes++
		case "Abstain":
			result.Abstain++
		case "No":
			result.No++
		case "NoWithVeto":
			result.NoWithVeto++
		}
	}
	resp.Votes = votes
	resp.Result = result
	return
}
