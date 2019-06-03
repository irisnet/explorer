package service

import (
	"encoding/json"

	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
)

type ProposalService struct {
	BaseService
}

func (service *ProposalService) GetModule() Module {
	return Proposal
}

func (service *ProposalService) QueryProposalsByHeight(height int64) []model.ProposalInfoVo {

	resp := []model.ProposalInfoVo{}

	data, err := document.Proposal{}.QuerySubmitProposalByHeight(height)
	if err != nil {
		logger.Error("query proposal err", logger.String("error", err.Error()), service.GetTraceLog())
		return resp
	}

	for _, v := range data {
		resp = append(resp, service.Query(int(v.ProposalId)))
	}

	return resp
}

func (service *ProposalService) QueryList(page, size int) (resp model.PageVo) {
	total, data, err := document.Proposal{}.QueryByPage(page, size)
	if err != nil {
		logger.Error("query proposal by page", logger.String("err", err.Error()))
		panic(types.CodeNotFound)
	}

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
	resp.Count = total
	return resp
}

func (service *ProposalService) Query(id int) (resp model.ProposalInfoVo) {

	data, err := document.Proposal{}.QueryProposalById(id)
	if err != nil {
		panic(types.CodeNotFound)
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

	from, to, err := document.Proposal{}.QueryTxFromToByTypeAndProposalId(id)

	if err != nil {
		logger.Error("query tx by proposal type and id ", logger.String("err", err.Error()))
	}
	proposal.Proposer = from
	proposal.TxHash = to

	if proposal.Type == "ParameterChange" || proposal.Type == "SoftwareUpgrade" {
		txMsg, err := document.TxMsg{}.QueryTxMsgByHash(proposal.TxHash)
		if err != nil {
			logger.Error("query tx msg by hash ", logger.String("err", err.Error()))
		} else {
			var msg model.MsgSubmitSoftwareUpgradeProposal
			if err := json.Unmarshal([]byte(txMsg.Content), &msg); err == nil {
				proposal.Parameters = msg.Params
				proposal.Version = msg.Version
				proposal.Software = msg.Software
				proposal.SwitchHeight = msg.SwitchHeight
				proposal.Threshold = msg.Threshold
			}
		}
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
