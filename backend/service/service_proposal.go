package service

import (
	"encoding/hex"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/explorer/backend/vo"
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

func (service *ProposalService) QueryProposalsByHeight(height int64) []vo.ProposalInfoVo {

	resp := []vo.ProposalInfoVo{}

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

func (service *ProposalService) QueryDepositAndVotingProposalList() []vo.ProposalNewStyle {
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

	if len(data) < 1 {
		return nil
	}

	proposalStatusDepositData := []document.Proposal{}
	proposalStatusVotingData := []document.Proposal{}

	systemVotingPower, err := service.GetSystemVotingPower()
	if err != nil {
		logger.Error("get systemVotingPower fail", logger.String("err", err.Error()))
	}

	for _, v := range data {
		if v.Status == document.ProposalStatusDeposit {
			proposalStatusDepositData = append(proposalStatusDepositData, v)
		}
		if v.Status == document.ProposalStatusVoting {
			proposalStatusVotingData = append(proposalStatusVotingData, v)
		}
	}

	proposals := make([]vo.ProposalNewStyle, 0, len(data))

	proposalStatusDepositDatas := formatProposalStatusDepositData(service, proposalStatusDepositData)
	proposalStatusVotingDatas := formatProposalStatusVotingData(service, proposalStatusVotingData, systemVotingPower)

	proposals = append(proposals, proposalStatusDepositDatas...)
	proposals = append(proposals, proposalStatusVotingDatas...)

	return proposals
}

func (service *ProposalService) QueryDeposit(id int) vo.ProposalNewStyle {
	data, err := document.Proposal{}.QueryProposalById(id)
	if err != nil {
		logger.Error("QueryProposalById have error", logger.String("err", err.Error()))
		panic(types.CodeNotFound)
	}

	proposal := vo.ProposalNewStyle{
		ProposalId:     data.ProposalId,
		Title:          data.Title,
		Type:           data.Type,
		Status:         data.Status,
		DepositEndTime: data.DepositEndTime.UTC(),
	}

	tx, err := document.CommonTx{}.QueryProposalInitAmountTxById(id)
	if err != nil {
		logger.Error("QueryProposalInitAmountTxById have error", logger.String("err", err.Error()))
		panic(types.CodeNotFound)
	}
	proposal.InitialDeposit = vo.Coin{
		Denom:  tx.Amount[0].Denom,
		Amount: tx.Amount[0].Amount,
	}

	l := vo.Level{}
	name, err := lcd.GetProposalLevelByType(data.Type)
	if err != nil {
		logger.Error("get proposal level by type", logger.String("err", err.Error()), logger.String("param", data.Type))
	}
	l.Name = name

	coinAsDoc, err := lcd.GetMinDepositByProposalType(data.Type)
	if err != nil {
		logger.Error("get min deposit", logger.String("err", err.Error()), logger.String("param", data.Type))
	}
	l.GovParam = vo.GovParam{
		MinDeposit: vo.Coin{
			Amount: coinAsDoc.Amount,
			Denom:  coinAsDoc.Denom,
		},
	}
	proposal.Level = l

	if len(data.TotalDeposit) > 0 {
		proposal.TotalDeposit.Amount = data.TotalDeposit[0].Amount
		proposal.TotalDeposit.Denom = data.TotalDeposit[0].Denom
	}

	return proposal
}

func (service *ProposalService) QueryVoting(id int) vo.ProposalNewStyle {
	data, err := document.Proposal{}.QueryProposalById(id)
	if err != nil {
		logger.Error("QueryProposalById have error", logger.String("err", err.Error()))
		panic(types.CodeNotFound)
	}

	systemVotingPower, err := service.GetSystemVotingPower()
	if err != nil {
		logger.Error("get systemVotingPower fail", logger.String("err", err.Error()))
	}

	tmpVoteArr := make([]vo.VoteWithVoterInfo, 0, len(data.Votes))

	allBondedValidators, err := document.Validator{}.GetBondedValidatorsSharesTokens()
	if err != nil {
		logger.Error("query allBondedValidators", logger.String("err", err.Error()))
	}
	curValidators := make(map[string]*vo.ValidatorGovInfo, len(allBondedValidators))

	for _, v := range allBondedValidators {
		delegatorShares, _ := utils.ParseStringToFloat(v.DelegatorShares)
		tokens, _ := utils.ParseStringToFloat(v.Tokens)
		curValidators[v.OperatorAddress] = &vo.ValidatorGovInfo{
			Address:   v.OperatorAddress,
			DelShares: delegatorShares,
			Tokens:    tokens,
		}
	}
	curDelegators := make(map[string]*vo.DelegatorGovInfo, len(data.Votes))

	for _, v := range data.Votes {
		getDelegationsByVoter(curValidators, curDelegators, v)
	}
	computedValidatorsPower(curValidators, curDelegators)

	for _, v := range curDelegators {
		tmpVote := vo.VoteWithVoterInfo{
			Voter:          v.Address,
			Option:         v.Option,
			Time:           v.Time,
			DelVotingPower: v.DelVotingPower,
			ValVotingPower: v.ValVotingPower,
		}
		tmpVoteArr = append(tmpVoteArr, tmpVote)
	}

	tmp := vo.ProposalNewStyle{
		ProposalId:       data.ProposalId,
		Title:            data.Title,
		Type:             data.Type,
		Status:           data.Status,
		SubmitTime:       data.SubmitTime,
		DepositEndTime:   data.DepositEndTime.UTC(),
		VotingEndTime:    data.VotingEndTime,
		Votes:            tmpVoteArr,
		TotalVotingPower: systemVotingPower,
	}

	l := vo.Level{}
	name, err := lcd.GetProposalLevelByType(data.Type)
	if err != nil {
		logger.Error("get proposal level by type", logger.String("err", err.Error()), logger.String("param", data.Type))
	}
	l.Name = name
	passThreshold, vetoThreshold, participation, _, err := lcd.GetPassVetoThresholdAndParticipationMinDeposit(data.Type)
	if err != nil {
		logger.Error("GetThresholdAndParticipationMinDeposit", logger.String("err", err.Error()), logger.String("param", data.Type))
	}
	l.GovParam = vo.GovParam{
		PassThreshold: passThreshold,
		VetoThreshold: vetoThreshold,
		Participation: participation,
	}
	tmp.Level = l

	return tmp
}

func formatProposalStatusVotingData(service *ProposalService, proposalStatusVotingData []document.Proposal, systemVotingPower int64) []vo.ProposalNewStyle {
	proposals := make([]vo.ProposalNewStyle, 0, len(proposalStatusVotingData))

	if len(proposalStatusVotingData) < 1 {
		return proposals
	}

	unique_set := make(map[string]bool)
	allBondedValidators, err := document.Validator{}.GetBondedValidatorsSharesTokens()
	if err != nil {
		logger.Error("query allBondedValidators", logger.String("err", err.Error()))
	}

	for _, v := range proposalStatusVotingData {
		for _, addr := range v.Votes {
			unique_set[addr.Voter] = true
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

	curValidators := make(map[string]*vo.ValidatorGovInfo, len(allBondedValidators))
	for _, v := range allBondedValidators {
		delegatorShares, _ := utils.ParseStringToFloat(v.DelegatorShares)
		tokens, _ := utils.ParseStringToFloat(v.Tokens)
		curValidators[v.OperatorAddress] = &vo.ValidatorGovInfo{
			Address:   v.OperatorAddress,
			DelShares: delegatorShares,
			Tokens:    tokens,
		}
	}

	for _, propo := range proposalStatusVotingData {
		tmpVoteArr := make([]vo.VoteWithVoterInfo, 0, len(propo.Votes))

		for _, v := range curValidators {
			v.DelDeductionShares = 0
		}
		curDelegators := make(map[string]*vo.DelegatorGovInfo, len(propo.Votes))

		for _, v := range propo.Votes {
			getDelegationsByVoter(curValidators, curDelegators, v)
		}

		computedValidatorsPower(curValidators, curDelegators)

		for _, v := range curDelegators {
			var voterVotingPower float64
			voterInfo := addrAsMultiTypeMap[v.Address]

			// only bonded validator will calculate voting power
			//if voterInfo.Status == document.ValidatorStatusValBonded {
			//	voterVotingPower = voterInfo.VotingPower
			//}
			voterVotingPower = v.DelVotingPower + v.ValVotingPower

			tmpVote := vo.VoteWithVoterInfo{
				Voter:        v.Address,
				Option:       v.Option,
				Time:         v.Time,
				VotingPower:  voterVotingPower,
				VoterMoniker: voterInfo.Moniker,
			}
			tmpVoteArr = append(tmpVoteArr, tmpVote)
		}

		tmp := vo.ProposalNewStyle{
			ProposalId:       propo.ProposalId,
			Title:            propo.Title,
			Type:             propo.Type,
			Status:           propo.Status,
			Votes:            tmpVoteArr,
			TotalVotingPower: systemVotingPower,
		}

		l := vo.Level{}
		name, err := lcd.GetProposalLevelByType(propo.Type)
		if err != nil {
			logger.Error("get proposal level by type", logger.String("err", err.Error()), logger.String("param", propo.Type))
		}
		l.Name = name

		passThreshold, vetoThreshold, participation, _, err := lcd.GetPassVetoThresholdAndParticipationMinDeposit(propo.Type)

		if err != nil {
			logger.Error("GetThresholdAndParticipationMinDeposit", logger.String("err", err.Error()), logger.String("param", propo.Type))
		}
		l.GovParam = vo.GovParam{
			PassThreshold: passThreshold,
			VetoThreshold: vetoThreshold,
			Participation: participation,
		}

		tmp.Level = l

		proposals = append(proposals, tmp)
	}
	return proposals
}

func getDelegationsByVoter(curValidators map[string]*vo.ValidatorGovInfo, curDelegators map[string]*vo.DelegatorGovInfo, vote document.PVote) {
	curDelegators[vote.Voter] = &vo.DelegatorGovInfo{
		Option:      vote.Option,
		Address:     vote.Voter,
		Time:        vote.Time,
		IsValidator: false,
	}
	curDelegator := curDelegators[vote.Voter]
	valAddr := utils.Convert(conf.Get().Hub.Prefix.ValAddr, vote.Voter)
	delegations := lcd.GetDelegationsByDelAddr(vote.Voter)
	for _, delegation := range delegations {
		if delegation.ValidatorAddr == valAddr {
			curDelegator.IsValidator = true
			curDelegator.ValAddr = valAddr
		}
		var votingPower float64
		delegationShares, err := utils.ParseStringToFloat(delegation.Shares)
		if err != nil {
			logger.Error("ParseInt delegationShares have error")
		} else {
			votingPower = computedVotingPower(delegationShares, curValidators, delegation.ValidatorAddr)
		}
		curDelegator.DelVotingPower = curDelegator.DelVotingPower + votingPower
		curValidator := curValidators[delegation.ValidatorAddr]
		curValidator.DelDeductionShares = curValidator.DelDeductionShares + delegationShares
	}
}

func computedVotingPower(delegationShares float64, curValidators map[string]*vo.ValidatorGovInfo, valAddr string) float64 {
	if v, ok := curValidators[valAddr]; ok {
		tokens := v.Tokens
		delegatorShares := v.DelShares
		if delegatorShares != 0 {
			votingPower := (tokens / delegatorShares) * delegationShares
			return votingPower
		}
	}
	return 0
}

func computedValidatorsPower(curValidators map[string]*vo.ValidatorGovInfo, curDelegators map[string]*vo.DelegatorGovInfo) {
	for _, curDelegator := range curDelegators {
		if curDelegator.IsValidator {
			valAddr := curDelegator.ValAddr
			if curValidator, ok := curValidators[valAddr]; ok {
				var votingPower float64
				delShares := curValidator.DelShares
				if delShares != 0 {
					tokens := curValidator.Tokens
					votingPower = (delShares - curValidator.DelDeductionShares) * (tokens / delShares)
					curDelegator.ValVotingPower = curDelegator.ValVotingPower + votingPower
				}
			}
		}
	}
}

func formatProposalStatusDepositData(service *ProposalService, proposalStatusDepositData []document.Proposal) []vo.ProposalNewStyle {
	proposals := make([]vo.ProposalNewStyle, 0, len(proposalStatusDepositData))

	if len(proposalStatusDepositData) < 1 {
		return proposals
	}

	depositProposalIdArr := []uint64{}
	for _, v := range proposalStatusDepositData {
		depositProposalIdArr = append(depositProposalIdArr, v.ProposalId)
	}

	depositInitmap, err := service.GetDepositProposalInitAmount(depositProposalIdArr)
	if err != nil {
		logger.Error("query GetDepositProposalInitAmount", logger.String("err", err.Error()), logger.Any("depositProposalIdArr", depositProposalIdArr))
	}

	for _, propo := range proposalStatusDepositData {
		tmp := vo.ProposalNewStyle{
			ProposalId: propo.ProposalId,
			Title:      propo.Title,
			Type:       propo.Type,
			Status:     propo.Status,
		}

		l := vo.Level{}
		name, err := lcd.GetProposalLevelByType(propo.Type)
		if err != nil {
			logger.Error("get proposal level by type", logger.String("err", err.Error()), logger.String("param", propo.Type))
		}
		l.Name = name

		coinAsDoc, err := lcd.GetMinDepositByProposalType(propo.Type)
		if err != nil {
			logger.Error("get min deposit", logger.String("err", err.Error()), logger.String("param", propo.Type))
		}
		l.GovParam = vo.GovParam{
			MinDeposit: vo.Coin{
				Amount: coinAsDoc.Amount,
				Denom:  coinAsDoc.Denom,
			},
		}

		tmp.InitialDeposit = depositInitmap[tmp.ProposalId]

		if len(propo.TotalDeposit) > 0 {
			tmp.TotalDeposit.Amount = propo.TotalDeposit[0].Amount
			tmp.TotalDeposit.Denom = propo.TotalDeposit[0].Denom
		}

		tmp.Level = l
		proposals = append(proposals, tmp)
	}
	return proposals
}

func (service *ProposalService) QueryList(page, size int, istotal bool) (resp vo.PageVo) {

	cnt, data, err := document.Proposal{}.QueryProposalByPage(page, size, istotal)
	if err != nil {
		logger.Error("query proposal by page ", logger.String("err", err.Error()))
		return
	}

	proposals := make([]vo.ProposalNewStyle, 0, len(data))
	unique_set := make(map[string]bool)

	for _, v := range data {
		if v.Status == document.ProposalStatusVoting {
			for _, addr := range v.Votes {
				unique_set[addr.Voter] = true
			}
		}
	}

	totalVotingPower, err := service.GetSystemVotingPower()
	if err != nil {
		logger.Error("get systemVotingPower fail", logger.String("err", err.Error()))
	}

	var allBondedValidators []document.Validator
	if len(unique_set) > 0 {
		allBondedValidators, err = document.Validator{}.GetBondedValidatorsSharesTokens()
		if err != nil {
			logger.Error("query allBondedValidators", logger.String("err", err.Error()))
		}
	}

	curValidators := make(map[string]*vo.ValidatorGovInfo, len(allBondedValidators))
	for _, v := range allBondedValidators {
		delegatorShares, _ := utils.ParseStringToFloat(v.DelegatorShares)
		tokens, _ := utils.ParseStringToFloat(v.Tokens)
		curValidators[v.OperatorAddress] = &vo.ValidatorGovInfo{
			Address:   v.OperatorAddress,
			DelShares: delegatorShares,
			Tokens:    tokens,
		}
	}

	for _, propo := range data {
		tmpVoteArr := make([]vo.VoteWithVoterInfo, 0, len(propo.Votes))
		finalVotes := vo.FinalVotes{}

		if propo.Status == document.ProposalStatusPassed || propo.Status == document.ProposalStatusRejected {
			finalVotes = vo.FinalVotes{
				Yes:        propo.FinalVotes.Yes,
				No:         propo.FinalVotes.No,
				NoWithVeto: propo.FinalVotes.NoWithVeto,
				Abstain:    propo.FinalVotes.Abstain,
			}
		}

		if propo.Status == document.ProposalStatusVoting {
			for _, v := range curValidators {
				v.DelDeductionShares = 0
			}
			curDelegators := make(map[string]*vo.DelegatorGovInfo, len(propo.Votes))

			for _, v := range propo.Votes {
				getDelegationsByVoter(curValidators, curDelegators, v)
			}

			computedValidatorsPower(curValidators, curDelegators)

			for _, v := range curDelegators {
				var voterVotingPower float64
				// only bonded validator will calculate voting power
				//if voterInfo.Status == document.ValidatorStatusValBonded {
				//	voterVotingPower = voterInfo.VotingPower
				//}
				voterVotingPower = v.DelVotingPower + v.ValVotingPower

				tmpVote := vo.VoteWithVoterInfo{
					Voter:       v.Address,
					Option:      v.Option,
					Time:        v.Time,
					VotingPower: voterVotingPower,
				}
				tmpVoteArr = append(tmpVoteArr, tmpVote)
			}
		}

		var l vo.Level
		name, err := lcd.GetProposalLevelByType(propo.Type)
		if err != nil {
			logger.Error("get proposal level by type", logger.String("err", err.Error()), logger.String("param", propo.Type))
		}
		l.Name = name

		tmp := vo.ProposalNewStyle{
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

func (service *ProposalService) Query(id int) (resp vo.ProposalInfoVo) {

	data, err := document.Proposal{}.QueryProposalById(id)
	if err != nil {
		logger.Error("QueryProposalById have error", logger.String("err", err.Error()))
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

	passThreshold, vetoThreshold, participation, penalty, err := lcd.GetPassVetoThresholdAndParticipationMinDeposit(data.Type)
	if err != nil {
		logger.Error("GetThresholdAndParticipationMinDeposit", logger.String("err", err.Error()), logger.String("param", data.Type))
	}

	level, err := lcd.GetProposalLevelByType(data.Type)
	if err != nil {
		logger.Error("get proposal level by type", logger.String("err", err.Error()), logger.String("param", data.Type))
	}

	proposal := vo.Proposal{
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
		YesThreshold:    passThreshold,
		VetoThreshold:   vetoThreshold,
		Participation:   participation,
		Penalty:         penalty,
		Level:           level,
	}

	if data.Status == document.ProposalStatusPassed || data.Status == document.ProposalStatusRejected {
		systemVotingPower, err := service.GetSystemVotingPower()
		if err != nil {
			logger.Error("get systemVotingPower fail", logger.String("err", err.Error()))
		}
		systemVotingPowerFloat, _ := utils.ParseStringToFloat(strconv.FormatInt(systemVotingPower, 10))

		var votedNum float64
		var noWithVeto float64
		if yes, err := utils.ParseStringToFloat(data.FinalVotes.Yes); err == nil {
			votedNum += yes
		} else {
			logger.Error("ParseStringToFloat yes fail", logger.String("err", err.Error()))
		}
		if no, err := utils.ParseStringToFloat(data.FinalVotes.No); err == nil {
			votedNum += no
		} else {
			logger.Error("ParseStringToFloat no fail", logger.String("err", err.Error()))
		}
		if noWithVeto, err = utils.ParseStringToFloat(data.FinalVotes.NoWithVeto); err == nil {
			votedNum += noWithVeto
		} else {
			logger.Error("ParseStringToFloat noWithVeto fail", logger.String("err", err.Error()))
		}
		if abstain, err := utils.ParseStringToFloat(data.FinalVotes.Abstain); err == nil {
			votedNum += abstain
		} else {
			logger.Error("ParseStringToFloat abstain fail", logger.String("err", err.Error()))
		}
		participationFloat, _ := utils.ParseStringToFloat(participation)
		vetoThresholdFloat, _ := utils.ParseStringToFloat(vetoThreshold)
		isParticipation := bool((votedNum / systemVotingPowerFloat) >= participationFloat)
		isRejectVote := bool(isParticipation && bool((noWithVeto/votedNum) >= vetoThresholdFloat))

		burnPercent, err := lcd.GetProposalBurnPercentByResult(data.Status, isRejectVote)
		if err != nil {
			logger.Error("GetProposalBurnPercentByResult fail", logger.String("err", err.Error()))
		} else {
			proposal.BurnPercent = burnPercent
		}
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
			var msg vo.MsgSubmitSoftwareUpgradeProposal
			if err := json.Unmarshal([]byte(txMsg.Content), &msg); err == nil {
				proposal.Parameters = msg.Params
				proposal.Version = msg.Version
				proposal.Software = msg.Software
				proposal.SwitchHeight = msg.SwitchHeight
				proposal.Threshold = msg.Threshold
			}
		}
	}

	if proposal.Type == "CommunityTaxUsage" {
		txMsg, err := document.TxMsg{}.QueryTxMsgByHash(proposal.TxHash)
		if err != nil {
			logger.Error("query tx msg by hash ", logger.String("err", err.Error()))
		} else {
			var msg map[string]interface{}
			if err := json.Unmarshal([]byte(txMsg.Content), &msg); err == nil {
				if v, ok := msg["usage"].(string); ok {
					proposal.Usage = v
				}
			}
		}
	}

	resp.Proposal = proposal

	var result vo.VoteResult
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

func (_ ProposalService) GetDepositProposalInitAmount(idArr []uint64) (map[uint64]vo.Coin, error) {

	if len(idArr) == 0 {
		return nil, nil
	}

	txs, err := document.CommonTx{}.QueryProposalTxListById(idArr)
	if err != nil {
		return nil, err
	}

	res := map[uint64]vo.Coin{}

	for _, tx := range txs {
		tmp := vo.Coin{}
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

func (s *ProposalService) GetVoteTxs(proposalId int64, page, size int, istotal bool) vo.GetVoteTxResponse {
	var (
		txs        []document.CommonTx
		txHashs    []string
		txMsgs     []document.TxMsg
		valAddrs   []string
		validators []document.Validator
		res        vo.GetVoteTxResponse
	)
	accAddrValAddrMap := make(map[string]string)

	num, txList, err := document.CommonTx{}.QueryProposalTxById(proposalId, page, size, istotal)

	if err != nil {
		logger.Error("QueryProposalTxById have error", logger.String("err", err.Error()))
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

func (s *ProposalService) GetDepositTxs(proposalId int64, page, size int, istotal bool) vo.TxPage {
	var res vo.TxPage

	num, txs, err := document.CommonTx{}.QueryProposalTxByIdWithSubmitOrDepositType(proposalId, page, size, istotal)

	if err != nil {
		panic(err)
	}
	items := s.buildTx(txs)
	res.Total = num
	res.Items = items

	return res
}

func (s *ProposalService) buildTx(txs []document.CommonTx) []vo.Tx {
	var (
		res []vo.Tx
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

		tx := vo.Tx{
			Hash:      v.TxHash,
			From:      v.From,
			Amount:    coinsAsUtils,
			Type:      v.Type,
			Timestamp: v.Time,
		}
		valaddr := utils.Convert(conf.Get().Hub.Prefix.ValAddr, v.From)
		validator, err := lcd.Validator(valaddr)
		if err != nil {
			logger.Error("lcd.Validator have error", logger.String("valaddr", valaddr), logger.String("err", err.Error()))
		}

		if moniker := validator.Description.Moniker; len(moniker) > 0 {
			tx.Moniker = moniker
		}

		res = append(res, tx)
	}

	return res
}

func (s *ProposalService) buildVoteTxs(txs []document.CommonTx, msgs []document.TxMsg,
	validators []document.Validator, accAddrValAddrMap map[string]string) []vo.VoteTx {
	var (
		voteTxs []vo.VoteTx
		msgVote vo.MsgVote
	)
	if len(txs) == 0 {
		return voteTxs
	}
	for _, tx := range txs {
		voteTx := vo.VoteTx{
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
