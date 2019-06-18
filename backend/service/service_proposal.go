package service

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"gopkg.in/mgo.v2/bson"
)

type AddrAsMultiType struct {
	Moniker         string
	Va              string
	ConsensusPubKey string
	ConsensusHex    string
	Status          int
	VotingPower     int64
}

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

func (service *ProposalService) QueryDepositAndVotingProposalList() []model.ProposalNewStyle {

	var data []document.Proposal
	selector := bson.M{
		document.Proposal_Field_ProposalId:        1,
		document.Proposal_Field_Title:             1,
		document.Proposal_Field_Type:              1,
		document.Proposal_Field_Status:            1,
		document.Proposal_Field_Votes:             1,
		document.Proposal_Field_VotingStartHeight: 1,
		document.Proposal_Field_TotalDeposit:      1,
	}

	condition := bson.M{"status": bson.M{"$in": []string{document.ProposalStatusDeposit, document.ProposalStatusVoting}}}

	var query = orm.NewQuery()
	defer query.Release()
	query.SetCollection(document.CollectionNmProposal).
		SetCondition(condition).
		SetSelector(selector).
		SetSort(document.Proposal_Field_VotingEndTime, document.Proposal_Field_DepositEndTime).
		SetSize(0).
		SetResult(&data)

	err := query.Exec()
	if err != nil {
		logger.Error("query proposal collection", logger.String("err", err.Error()))
		return nil
	}

	depositProposalIdArr := []uint64{}
	unique_set := make(map[string]bool)

	for _, v := range data {
		if v.Status == document.ProposalStatusDeposit {
			depositProposalIdArr = append(depositProposalIdArr, v.ProposalId)
		}

		if v.Status == document.ProposalStatusVoting {
			for _, addr := range v.Votes {
				unique_set[addr.Voter] = true
			}
		}
	}

	voterAddrArr := make([]string, 0, len(unique_set))
	for x := range unique_set {
		voterAddrArr = append(voterAddrArr, x)
	}

	addrAsMultiTypeMap, err := service.GetValidatorPublicKeyMonikerFromProposalVoter(voterAddrArr)
	if err != nil {
		logger.Error("query GetValidatorPublicKeyMonikerFromProposalVoter", logger.String("err", err.Error()), logger.Any("voterAddrArr", voterAddrArr))
	}

	depositInitmap, err := service.GetDepositProposalInitAmount(depositProposalIdArr)
	if err != nil {
		logger.Error("query GetDepositProposalInitAmount", logger.String("err", err.Error()), logger.Any("depositProposalIdArr", depositProposalIdArr))
	}

	proposals := make([]model.ProposalNewStyle, 0, len(data))
	for _, propo := range data {
		var totalVotingPower int64
		tmpVoteArr := make([]model.VoteWithVoterInfo, 0, len(propo.Votes))

		for _, v := range propo.Votes {
			var voterVotingPower int64
			voterInfo := addrAsMultiTypeMap[v.Voter]

			// only bonded validator will calculate voting power
			if voterInfo.Status == document.ValidatorStatusValBonded {
				voterVotingPower = voterInfo.VotingPower
				totalVotingPower += voterInfo.VotingPower
			}

			tmpVote := model.VoteWithVoterInfo{
				Voter:        v.Voter,
				Option:       v.Option,
				Time:         v.Time,
				VotingPower:  voterVotingPower,
				VoterMoniker: voterInfo.Moniker,
			}
			tmpVoteArr = append(tmpVoteArr, tmpVote)
		}

		tmp := model.ProposalNewStyle{
			ProposalId:       propo.ProposalId,
			Title:            propo.Title,
			Type:             propo.Type,
			Status:           propo.Status,
			Votes:            tmpVoteArr,
			TotalVotingPower: totalVotingPower,
		}

		l := model.Level{}
		name, err := lcd.GetProposalLevelByType(propo.Type)
		if err != nil {
			logger.Error("get proposal level by type", logger.String("err", err.Error()), logger.String("param", propo.Type))
		}
		l.Name = name
		if propo.Status == document.ProposalStatusDeposit {
			coinAsDoc, err := lcd.GetMinDepositByProposalType(propo.Type)
			if err != nil {
				logger.Error("get min deposit", logger.String("err", err.Error()), logger.String("param", propo.Type))
			}
			l.GovParam = model.GovParam{
				MinDeposit: model.Coin{
					Amount: coinAsDoc.Amount,
					Denom:  coinAsDoc.Denom,
				},
			}

			tmp.InitialDeposit = depositInitmap[tmp.ProposalId]

			if len(propo.TotalDeposit) > 0 {
				tmp.TotalDeposit.Amount = propo.TotalDeposit[0].Amount
				tmp.TotalDeposit.Denom = propo.TotalDeposit[0].Denom
			}
		}

		if propo.Status == document.ProposalStatusVoting {
			passThreshold, vetoThreshold, participation, err := lcd.GetPassVetoThresholdAndParticipationMinDeposit(propo.Type)

			if err != nil {
				logger.Error("GetThresholdAndParticipationMinDeposit", logger.String("err", err.Error()), logger.String("param", propo.Type))
			}
			l.GovParam = model.GovParam{
				PassThreshold: passThreshold,
				VetoThreshold: vetoThreshold,
				Participation: participation,
			}
		}

		tmp.Level = l

		proposals = append(proposals, tmp)

	}

	return proposals
}

func (service *ProposalService) QueryList(page, size int) (resp model.PageVo) {
	var data []document.Proposal
	sort := desc(document.Proposal_Field_ProposalId)
	selector := bson.M{
		document.Proposal_Field_ProposalId:        1,
		document.Proposal_Field_Title:             1,
		document.Proposal_Field_Type:              1,
		document.Proposal_Field_Status:            1,
		document.Proposal_Field_SubmitTime:        1,
		document.Proposal_Field_DepositEndTime:    1,
		document.Proposal_Field_VotingEndTime:     1,
		document.Proposal_Field_Votes:             1,
		document.Proposal_Field_VotingStartHeight: 1,
		document.Proposal_Field_Final_Votes:       1,
	}

	if cnt, err := pageQuery(document.CollectionNmProposal, selector, nil, sort, page, size, &data); err == nil {
		proposals := make([]model.ProposalNewStyle, 0, len(data))
		unique_set := make(map[string]bool)

		for _, v := range data {
			if v.Status == document.ProposalStatusVoting {
				for _, addr := range v.Votes {
					unique_set[addr.Voter] = true
				}
			}
		}

		voterAddrArr := make([]string, 0, len(unique_set))
		for x := range unique_set {
			voterAddrArr = append(voterAddrArr, x)
		}

		addrAsMultiTypeMap, err := service.GetValidatorPublicKeyMonikerFromProposalVoter(voterAddrArr)
		if err != nil {
			logger.Error("query GetValidatorPublicKeyMonikerFromProposalVoter", logger.String("err", err.Error()))
		}

		for _, propo := range data {
			var totalVotingPower int64
			tmpVoteArr := make([]model.VoteWithVoterInfo, 0, len(propo.Votes))
			finalVotes := model.FinalVotes{}

			if propo.Status == document.ProposalStatusPassed || propo.Status == document.ProposalStatusRejected {
				finalVotes = model.FinalVotes{
					Yes:        propo.FinalVotes.Yes,
					No:         propo.FinalVotes.No,
					NoWithVeto: propo.FinalVotes.NoWithVeto,
					Abstain:    propo.FinalVotes.Abstain,
				}
			}

			if propo.Status == document.ProposalStatusVoting {

				for _, v := range propo.Votes {
					var voterVotingPower int64
					voterInfo := addrAsMultiTypeMap[v.Voter]

					// only bonded validator will calculate voting power
					if voterInfo.Status == document.ValidatorStatusValBonded {
						voterVotingPower = voterInfo.VotingPower
						totalVotingPower += voterInfo.VotingPower
					}

					tmpVote := model.VoteWithVoterInfo{
						Voter:       v.Voter,
						Option:      v.Option,
						Time:        v.Time,
						VotingPower: voterVotingPower,
					}
					tmpVoteArr = append(tmpVoteArr, tmpVote)
				}
			}

			var l model.Level
			name, err := lcd.GetProposalLevelByType(propo.Type)
			if err != nil {
				logger.Error("get proposal level by type", logger.String("err", err.Error()), logger.String("param", propo.Type))
			}
			l.Name = name

			tmp := model.ProposalNewStyle{
				ProposalId:       propo.ProposalId,
				Title:            propo.Title,
				Type:             propo.Type,
				Status:           propo.Status,
				SubmitTime:       propo.SubmitTime,
				DepositEndTime:   propo.DepositEndTime,
				VotingEndTime:    propo.VotingEndTime,
				Votes:            tmpVoteArr,
				TotalVotingPower: totalVotingPower,
				FinalVotes:       finalVotes,
				Level:            l,
			}
			proposals = append(proposals, tmp)
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

	var result model.VoteResult
	for _, v := range data.Votes {

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

func (_ ProposalService) GetValidatorPublicKeyMonikerFromProposalVoter(addrArrAsAa []string) (map[string]AddrAsMultiType, error) {

	if len(addrArrAsAa) == 0 {
		return nil, nil
	}
	AddrAsMultiTypeMap := map[string]AddrAsMultiType{}

	addrArrAsVa := make([]string, 0, len(addrArrAsAa))

	for _, v := range addrArrAsAa {
		va := utils.Convert(conf.Get().Hub.Prefix.ValAddr, v)
		addrArrAsVa = append(addrArrAsVa, va)
		AddrAsMultiTypeMap[v] = AddrAsMultiType{
			Va: va,
		}
	}

	var validatorsDoc []document.Validator
	var selector = bson.M{
		"description.moniker": 1,
		"operator_address":    1,
		"consensus_pubkey":    1,
		"status":              1,
		"voting_power":        1,
	}

	err := queryAll(document.CollectionNmValidator, selector, bson.M{"operator_address": bson.M{"$in": addrArrAsVa}}, "", 0, &validatorsDoc)
	if err != nil {
		return nil, err
	}

	for _, validator := range validatorsDoc {
		for k, v := range AddrAsMultiTypeMap {
			if v.Va == validator.OperatorAddress {
				v.ConsensusPubKey = validator.ConsensusPubkey
				_, bytes, err := utils.DecodeAndConvert(validator.ConsensusPubkey)
				if err != nil {
					logger.Error("DecodeAndConvert", logger.String("err", err.Error()), logger.String("param", validator.ConsensusPubkey))
					continue
				}
				v.ConsensusHex = strings.ToUpper(hex.EncodeToString(bytes))
				v.Moniker = validator.Description.Moniker
				v.Status = validator.Status
				v.VotingPower = validator.VotingPower
				AddrAsMultiTypeMap[k] = v
			}
		}
	}

	return AddrAsMultiTypeMap, nil
}

func (_ ProposalService) GetDepositProposalInitAmount(idArr []uint64) (map[uint64]model.Coin, error) {

	if len(idArr) == 0 {
		return nil, nil
	}

	selector := bson.M{document.Tx_Field_Amount: 1, document.Tx_Field_ProposalId: 1}
	condition := bson.M{document.Tx_Field_Type: "SubmitProposal", document.Tx_Field_Status: "success", document.Tx_Field_ProposalId: bson.M{"$in": idArr}}
	var txs []document.CommonTx

	err := queryAll(document.CollectionNmCommonTx, selector, condition, desc(document.Tx_Field_Time), 0, &txs)
	if err != nil {
		return nil, err
	}

	res := map[uint64]model.Coin{}

	for _, tx := range txs {
		tmp := model.Coin{}
		if len(tx.Amount) > 0 {
			tmp.Amount = tx.Amount[0].Amount
			tmp.Denom = tx.Amount[0].Denom
		}

		res[tx.ProposalId] = tmp
	}

	return res, nil
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
			err := queryAll(document.CollectionNmTxMsg, selector, condition, "", 0, &txMsgs)
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
			err = queryAll(document.CollectionNmValidator, selector, condition, "", 0, &validators)
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
