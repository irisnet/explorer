package service

import (
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
	"gopkg.in/mgo.v2/bson"
)

type ProposalService struct {
	BaseService
}

func (service *ProposalService) GetModule() Module {
	return Proposal
}

func (service *ProposalService) QueryList(page, size int) (resp model.PageVo) {
	var data []document.Proposal
	sort := desc(document.Proposal_Field_SubmitTime)
	resp = queryPage(document.CollectionNmProposal, &data, nil, sort, page, size)

	var proposals []model.Proposal
	for _, propo := range data {
		mP := model.Proposal{
			Title:           propo.Title,
			ProposalId:      propo.ProposalId,
			Type:            propo.Type,
			Description:     propo.Description,
			Status:          propo.Status,
			SubmitTime:      propo.SubmitTime.UTC(),
			DepositEndTime:  propo.DepositEndTime.UTC(),
			VotingStartTime: propo.VotingStartTime.UTC(),
			VotingEndTime:   propo.VotingEndTime.UTC(),
			TotalDeposit:    propo.TotalDeposit,
		}
		proposals = append(proposals, mP)
	}
	resp.Data = proposals
	return resp
}

func (service *ProposalService) Query(id int) (resp model.ProposalInfoVo) {
	var data document.Proposal
	db := getDb()
	defer db.Session.Close()
	propoStore := db.C(document.CollectionNmProposal)
	txStore := db.C(document.CollectionNmCommonTx)

	if err := propoStore.Find(bson.M{document.Proposal_Field_ProposalId: id}).One(&data); err != nil {
		panic(types.CodeNotFound)
		return
	}

	proposal := model.Proposal{
		Title:           data.Title,
		ProposalId:      data.ProposalId,
		Type:            data.Type,
		Description:     data.Description,
		Status:          data.Status,
		SubmitTime:      data.SubmitTime.UTC(),
		DepositEndTime:  data.DepositEndTime.UTC(),
		VotingStartTime: data.VotingStartTime.UTC(),
		VotingEndTime:   data.VotingEndTime.UTC(),
		TotalDeposit:    data.TotalDeposit,
	}

	var tx document.CommonTx
	if err := txStore.Find(bson.M{document.Tx_Field_Type: types.TypeSubmitProposal, document.Proposal_Field_ProposalId: id}).One(&tx); err == nil {
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
