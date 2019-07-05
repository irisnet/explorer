package service

import (
	"encoding/hex"
	"encoding/json"
	"strings"

	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
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

func (service *ProposalService) QueryDepositAndVotingProposalList() []model.ProposalNewStyle {
	var (
		proposalDocument document.Proposal
	)
	status := []string{document.ProposalStatusDeposit, document.ProposalStatusVoting}
	sorts := []string{document.Proposal_Field_VotingEndTime, document.Proposal_Field_DepositEndTime}

	data, err := proposalDocument.GetProposalsByStatus(status, sorts)
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

	totalVotingPower, err := service.GetSystemVotingPower()
	if err != nil {
		logger.Error("get systemVotingPower fail", logger.String("err", err.Error()))
	}
	proposals := make([]model.ProposalNewStyle, 0, len(data))
	for _, propo := range data {
		tmpVoteArr := make([]model.VoteWithVoterInfo, 0, len(propo.Votes))

		for _, v := range propo.Votes {
			var voterVotingPower int64
			voterInfo := addrAsMultiTypeMap[v.Voter]

			// only bonded validator will calculate voting power
			if voterInfo.Status == document.ValidatorStatusValBonded {
				voterVotingPower = voterInfo.VotingPower
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

	cnt, data, err := document.Proposal{}.QueryProposalByPage(page, size)
	if err != nil {
		logger.Error("query proposal by page ", logger.String("err", err.Error()))
		return
	}

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

	totalVotingPower, err := service.GetSystemVotingPower()
	if err != nil {
		logger.Error("get systemVotingPower fail", logger.String("err", err.Error()))
	}

	for _, propo := range data {
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
	return resp
}

func (service *ProposalService) Query(id int) (resp model.ProposalInfoVo) {

	data, err := document.Proposal{}.QueryProposalById(id)
	if err != nil {
		panic(types.CodeNotFound)
	}

	coinsAsUtils := make(utils.Coins, 0, len(data.TotalDeposit))

	for _, v := range data.TotalDeposit {
		tmp := utils.Coin{
			Amount: v.Amount,
			Denom:  v.Denom,
		}
		coinsAsUtils = append(coinsAsUtils, tmp)
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
		TotalDeposit:    coinsAsUtils,
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

	validatorsDoc, err := document.Validator{}.QueryValidatorMonikerOpAddrConsensusPubkey(addrArrAsVa)

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

	txs, err := document.CommonTx{}.QueryProposalTxListById(idArr)
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

// get systemVotingPower by get sum of validator votingPower which status is bonded
func (_ ProposalService) GetSystemVotingPower() (int64, error) {
	var (
		validatorDocument document.Validator
		totalVotingPower  int64
	)

	if validators, err := validatorDocument.GetBondedValidators(); err == nil {
		for _, v := range validators {
			totalVotingPower += v.VotingPower
		}
	}

	return totalVotingPower, nil
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

	num, txList, err := document.CommonTx{}.QueryProposalTxById(proposalId, page, size)

	if err != nil {
		panic(types.CodeNotFound)
	}
	txs = txList

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
		txMsgList, err := document.TxMsg{}.QueryTxMsgListByHashList(txHashs)

		if err != nil {
			logger.Error("query tx msg fail", logger.String("err", err.Error()))
		} else {
			txMsgs = txMsgList
		}
	}

	// query validator info
	if len(valAddrs) > 0 {

		validatorList, err := document.Validator{}.QueryValidatorMonikerOpAddrConsensusPubkey(valAddrs)
		if err != nil {
			logger.Error("query validator fail", logger.String("err", err.Error()))
		} else {
			validators = validatorList
		}

	}

	voteTxs := s.buildVoteTxs(txs, txMsgs, validators, accAddrValAddrMap)
	res.Total = num
	res.Items = voteTxs

	return res
}

func (s *ProposalService) GetDepositTxs(proposalId int64, page, size int) model.TxPage {
	var res model.TxPage

	num, txs, err := document.CommonTx{}.QueryProposalTxByIdWithSubmitOrDepositType(proposalId, page, size)

	if err != nil {
		panic(err)
	}
	items := s.buildTx(txs)
	res.Total = num
	res.Items = items

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

		coinsAsUtils := utils.Coins{}
		for _, v := range v.Amount {
			coinAsUtils := utils.Coin{
				Denom:  v.Denom,
				Amount: v.Amount,
			}
			coinsAsUtils = append(coinsAsUtils, coinAsUtils)
		}

		tx := model.Tx{
			Hash:      v.TxHash,
			From:      v.From,
			Amount:    coinsAsUtils,
			Type:      v.Type,
			Timestamp: v.Time,
		}
		validator := lcd.GetValidatorsByDelegator(v.From)
		if moniker := validator[0].Description.Moniker; len(moniker) > 0 {
			tx.Moniker = moniker
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
