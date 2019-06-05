package service

import (
	"encoding/json"

	"fmt"
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
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
		SetCondition(bson.M{document.Tx_Field_Type: types.TxTypeSubmitProposal, document.Proposal_Field_ProposalId: id}).
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

func (service *ProposalService) QueryTypeAndTitleByIds(ids []uint64) ([]document.Proposal, error) {

	proposalDocArr := []document.Proposal{}
	selector := bson.M{
		document.Proposal_Field_ProposalId: 1,
		document.Proposal_Field_Title:      1,
		document.Proposal_Field_Type:       1,
	}
	condition := bson.M{
		document.Proposal_Field_ProposalId: bson.M{"$in": ids},
	}
	err := queryAll(document.CollectionNmProposal, selector, condition, "", 0, &proposalDocArr)

	return proposalDocArr, err
}

func (s *ProposalService) GetVoteTxs(proposalId int64, page, size int) model.GetVoteTxResponse {
	var (
		txs        []document.CommonTx
		txHashs    []string
		txMsgs     []document.TxMsg
		valAddrs   []string
		validators []document.Validator
		res        model.GetVoteTxResponse
	)
	accAddrValAddrMap := make(map[string]string)

	// query vote txs by given proposal id
	selector := bson.M{
		document.Tx_Field_Height: 1,
		document.Tx_Field_Time:   1,
		document.Tx_Field_Hash:   1,
		document.Tx_Field_From:   1,
	}
	condition := bson.M{
		document.Tx_Field_Status:     types.TxTypeStatus,
		document.Tx_Field_ProposalId: proposalId,
		document.Tx_Field_Type:       types.TxTypeVote,
	}
	sort := fmt.Sprintf("-%v", document.Tx_Field_Height)

	if num, err := pageQuery(document.CollectionNmCommonTx, selector, condition, sort, page, size, &txs); err != nil {
		panic(types.CodeNotFound)
	} else {
		for _, v := range txs {
			txHashs = append(txHashs, v.TxHash)
			valAddr := utils.Convert(conf.Get().Hub.Prefix.ValAddr, v.From)
			if valAddr != "" {
				valAddrs = append(valAddrs, valAddr)
				accAddrValAddrMap[v.From] = valAddr
			}
		}

		// query tx msg by tx hash
		if len(txHashs) > 0 {
			selector := bson.M{
				document.TxMsg_Field_Hash:    1,
				document.TxMsg_Field_Content: 1,
			}
			condition := bson.M{
				document.TxMsg_Field_Hash: bson.M{
					"$in": txHashs,
				},
			}
			err := queryAll(document.CollectionNmTxMsg, selector, condition, "", 10000, &txMsgs)
			if err != nil {
				logger.Error("query tx msg fail", logger.String("err", err.Error()))
			}
		}

		// query validator info
		if len(valAddrs) > 0 {
			selector = bson.M{
				document.ValidatorFieldOperatorAddress: 1,
				document.ValidatorFieldDescription:     1,
			}
			condition = bson.M{
				document.ValidatorFieldOperatorAddress: bson.M{
					"$in": valAddrs,
				},
			}
			err = queryAll(document.CollectionNmValidator, selector, condition, "", 1000, &validators)
			if err != nil {
				logger.Error("query validator fail", logger.String("err", err.Error()))
			}
		}

		voteTxs := s.buildVoteTxs(txs, txMsgs, validators, accAddrValAddrMap)
		res.Total = num
		res.Items = voteTxs
	}

	return res
}

func (s *ProposalService) GetDepositTxs(proposalId int64, page, size int) model.TxPage {
	var (
		txs []document.CommonTx
		res model.TxPage
	)
	selector := bson.M{
		document.Tx_Field_Hash:   1,
		document.Tx_Field_From:   1,
		document.Tx_Field_Amount: 1,
		document.Tx_Field_Type:   1,
		document.Tx_Field_Time:   1,
	}
	condition := bson.M{
		document.Tx_Field_Status:     types.TxTypeStatus,
		document.Tx_Field_ProposalId: proposalId,
		document.Tx_Field_Type: bson.M{
			"$in": []string{types.TxTypeSubmitProposal, types.TxTypeDeposit},
		},
	}
	sort := fmt.Sprintf("-%v", document.Tx_Field_Height)
	if num, err := pageQuery(document.CollectionNmCommonTx, selector, condition, sort, page, size, &txs); err != nil {
		logger.Error("query tx fail", logger.String("err", err.Error()))
		panic(types.CodeNotFound)
	} else {
		items := s.buildTx(txs)

		res.Total = num
		res.Items = items
	}

	return res
}

func (s *ProposalService) buildTx(txs []document.CommonTx) []model.Tx {
	var (
		res []model.Tx
	)
	if len(txs) == 0 {
		return res
	}
	for _, v := range txs {
		tx := model.Tx{
			Hash:      v.TxHash,
			From:      v.From,
			Amount:    v.Amount,
			Type:      v.Type,
			Timestamp: v.Time,
		}

		res = append(res, tx)
	}

	return res
}

func (s *ProposalService) buildVoteTxs(txs []document.CommonTx, msgs []document.TxMsg,
	validators []document.Validator, accAddrValAddrMap map[string]string) []model.VoteTx {
	var (
		voteTxs []model.VoteTx
		msgVote model.MsgVote
	)
	if len(txs) == 0 {
		return voteTxs
	}
	for _, tx := range txs {
		voteTx := model.VoteTx{
			Voter:     tx.From,
			TxHash:    tx.TxHash,
			Timestamp: tx.Time,
		}

		for _, msg := range msgs {
			if tx.TxHash == msg.Hash {
				// marshal json
				if err := json.Unmarshal([]byte(msg.Content), &msgVote); err != nil {
					logger.Error("unmarshal json fail", logger.String("err", err.Error()))
					continue
				} else {
					voteTx.Option = msgVote.Option
				}
				break
			}
		}

		for _, v := range validators {
			if accAddrValAddrMap[tx.From] == v.OperatorAddress {
				voteTx.Voter = v.OperatorAddress
				voteTx.Moniker = v.Description.Moniker
				break
			}
		}

		voteTxs = append(voteTxs, voteTx)
	}

	return voteTxs
}
