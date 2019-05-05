package service

import (
	"encoding/json"

	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm"
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

func (service *ProposalService) QueryProposalsByHeight(height int64) []model.ProposalInfoVo {

	resp := []model.ProposalInfoVo{}

	var query = orm.NewQuery()
	defer query.Release()

	var data []document.CommonTx

	var selector = bson.M{"proposal_id": 1, "type": 1, "from": 1}

	query.SetCollection(document.CollectionNmCommonTx).
		SetSelector(selector).
		SetCondition(bson.M{document.Tx_Field_Height: height, document.Tx_Field_Type: "SubmitProposal"}).
		SetResult(&data)
	if err := query.Exec(); err != nil {
		logger.Error("query proposal err", logger.String("error", err.Error()), service.GetTraceLog())
	}

	for _, v := range data {
		resp = append(resp, service.Query(int(v.ProposalId)))
	}

	return resp
}

func (service *ProposalService) QueryList(page, size int) (resp model.PageVo) {
	var data []document.Proposal
	sort := desc(document.Proposal_Field_SubmitTime)
	if cnt, err := pageQuery(document.CollectionNmProposal, nil, nil, sort, page, size, &data); err == nil {
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
		resp.Count = cnt
	}
	return resp
}

func (service *ProposalService) Query(id int) (resp model.ProposalInfoVo) {
	var query = orm.NewQuery()
	defer query.Release()

	var data document.Proposal
	query.SetCollection(document.CollectionNmProposal).
		SetCondition(bson.M{document.Proposal_Field_ProposalId: id}).
		SetResult(&data)

	if err := query.Exec(); err != nil {
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

	var tx document.CommonTx
	query.Reset().SetCollection(document.CollectionNmCommonTx).
		SetCondition(bson.M{document.Tx_Field_Type: types.TypeSubmitProposal, document.Proposal_Field_ProposalId: id}).
		SetResult(&tx)
	if err := query.Exec(); err == nil {
		proposal.Proposer = tx.From
		proposal.TxHash = tx.TxHash
	}

	if proposal.Type == "ParameterChange" || proposal.Type == "SoftwareUpgrade" {
		var txMsg document.TxMsg
		query.Reset().SetCollection(document.CollectionNmTxMsg).
			SetCondition(bson.M{document.TxMsg_Field_Hash: proposal.TxHash}).
			SetResult(&txMsg)
		if err := query.Exec(); err == nil {
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
