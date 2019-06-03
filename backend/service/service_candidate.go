package service

import (
	"strconv"
	"sync"

	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
)

type CandidateService struct {
	BaseService
}

func (service *CandidateService) GetModule() Module {
	return Candidate
}

func (service *CandidateService) GetValidators(typ, origin string, page, size int) interface{} {
	if origin == "browser" {
		var result []lcd.ValidatorVo
		var blackList = service.QueryBlackList()

		total, validatorList, err := document.Candidate{}.GetValidatorListByPage(typ, page, size)

		if err != nil || total <= 0 {
			panic(types.CodeNotFound)
		}

		var totalVotingPower = getTotalVotingPower()
		for i, v := range validatorList {
			if desc, ok := blackList[v.OperatorAddress]; ok {
				validatorList[i].Description.Moniker = desc.Moniker
				validatorList[i].Description.Identity = desc.Identity
				validatorList[i].Description.Website = desc.Website
				validatorList[i].Description.Details = desc.Details
			}
			var validator lcd.ValidatorVo
			if err := utils.Copy(validatorList[i], &validator); err != nil {
				panic(types.CodeSysFailed)
			}
			validator.VotingRate = float32(v.VotingPower) / float32(totalVotingPower)
			result = append(result, validator)
		}

		return model.PageVo{
			Data:  result,
			Count: total,
		}
	}

	return service.queryValForRainbow(typ, page, size)
}

func (service *CandidateService) GetValidator(address string) lcd.ValidatorVo {
	var validator, err = lcd.Validator(address)
	if err != nil {
		panic(types.CodeNotFound)
	}

	var blackList = service.QueryBlackList()
	if desc, ok := blackList[validator.OperatorAddress]; ok {
		validator.Description.Moniker = desc.Moniker
		validator.Description.Identity = desc.Identity
		validator.Description.Website = desc.Website
		validator.Description.Details = desc.Details
	}
	return validator
}

func (service *CandidateService) QueryValidatorByConAddr(address string) document.Candidate {

	candadate, err := document.Candidate{}.QueryValidatorByConsensusAddr(address)
	if err != nil {
		logger.Error("not found validator by conAddr", logger.String("conAddr", address))
	}
	return candadate
}

func (service *CandidateService) QueryCandidate(address string) model.CandidatesInfoVo {
	//var candidate document.Candidate

	validator, err := lcd.Validator(address)
	if err != nil {
		panic(types.CodeNotFound)
	}

	var moniker = validator.Description.Moniker
	var identity = validator.Description.Identity
	var website = validator.Description.Website
	var details = validator.Description.Details
	var blackList = service.QueryBlackList()
	if desc, ok := blackList[validator.OperatorAddress]; ok {
		moniker = desc.Moniker
		identity = desc.Identity
		website = desc.Website
		details = desc.Details
	}
	var tokenDec, _ = types.NewDecFromStr(validator.Tokens)
	var val = model.Validator{
		Address:        validator.OperatorAddress,
		PubKey:         validator.ConsensusPubkey,
		Owner:          utils.Convert(conf.Get().Hub.Prefix.AccAddr, validator.OperatorAddress),
		Jailed:         validator.Jailed,
		Status:         BondStatusToString(validator.Status),
		BondHeight:     utils.ParseIntWithDefault(validator.BondHeight, 0),
		OriginalTokens: utils.RoundToString(validator.Tokens, 0),
		VotingPower:    tokenDec.RoundInt64(),
		Description: model.Description{
			Moniker:  moniker,
			Identity: identity,
			Website:  website,
			Details:  details,
		},
		Rate: validator.Commission.Rate,
	}

	result := model.CandidatesInfoVo{
		Validator: val,
	}
	count, err := document.Candidate{}.QueryPowerWithBonded()

	if err != nil {
		logger.Error("query candidate power with bonded ", logger.String("err", err.Error()))
		return result
	}

	result.PowerAll = getVotingPowerFromToken(strconv.FormatFloat(count, 'f', 10, 64))
	return result
}

func (service *CandidateService) QueryCandidatesTopN() model.ValDetailVo {

	candidateList, power, upTimeMap, err := document.Candidate{}.GetCandidatesTopN()

	if err != nil {
		panic(types.CodeNotFound)
	}

	var validators []model.Validator

	for _, candidate := range candidateList {
		validator := service.convert(candidate)
		validator.Uptime = upTimeMap[candidate.PubKeyAddr]
		validators = append(validators, validator)
	}
	resp := model.ValDetailVo{
		PowerAll:   getVotingPowerFromToken(power),
		Validators: validators,
	}
	return resp
}

func (service *CandidateService) QueryCandidateUptime(address, category string) []model.UptimeChangeVo {

	address, err := document.Candidate{}.GetCandidatePubKeyAddrByAddr(address)

	if err != nil || address == "" {
		panic(types.CodeNotFound)
	}

	switch category {
	case "hour":
		resultAsDoc, err := document.Candidate{}.QueryCandidateUptimeWithHour(address)
		if err != nil {
			panic(types.CodeNotFound)
		}
		result := make([]model.UptimeChangeVo, 0, len(resultAsDoc))

		for _, v := range resultAsDoc {
			result = append(result, model.UptimeChangeVo{
				Address: v.Address,
				Time:    v.Time,
				Uptime:  v.Uptime,
			})
		}

		return result
	case "week", "month":

		resultAsDoc, err := document.Candidate{}.QueryCandidateUptimeByWeekOrMonth(address, category)
		if err != nil {
			panic(types.CodeNotFound)
		}
		result := make([]model.UptimeChangeVo, 0, len(resultAsDoc))

		for _, v := range resultAsDoc {
			result = append(result, model.UptimeChangeVo{
				Address: v.Address,
				Time:    v.Time,
				Uptime:  v.Uptime,
			})
		}
		return result
	}
	return nil
}

func (service *CandidateService) QueryCandidatePower(address, category string) []model.ValVotingPowerChangeVo {

	var agoStr string
	switch category {
	case "week":
		agoStr = "-336h"
		break
	case "month":
		agoStr = "-720h"
		break
	case "months":
		agoStr = "-1440h"
		break
	}

	candidatePowerArr, err := document.Candidate{}.QueryCandidatePower(address, agoStr)

	if err != nil {
		panic(types.CodeNotFound)
	}

	result := make([]model.ValVotingPowerChangeVo, 0, len(candidatePowerArr))

	for _, v := range candidatePowerArr {

		result = append(result, model.ValVotingPowerChangeVo{
			Height:  v.Height,
			Address: v.Address,
			Power:   v.Power,
			Time:    v.Time,
			Change:  v.Change,
		})
	}

	return result
}

func (service *CandidateService) QueryCandidateStatus(address string) (resp model.ValStatus) {

	preCommitCount, uptime, err := document.Candidate{}.QueryCandidateStatus(address)

	if err != nil {
		logger.Error("query candidate status", logger.String("err", err.Error()))
		panic(types.CodeNotFound)
	}

	resp = model.ValStatus{
		Uptime:         uptime,
		PrecommitCount: float64(preCommitCount),
	}

	return resp
}

func (service *CandidateService) convert(candidate document.Candidate) model.Validator {
	var moniker = candidate.Description.Moniker
	var identity = candidate.Description.Identity
	var website = candidate.Description.Website
	var details = candidate.Description.Details
	var blackList = service.QueryBlackList()
	if desc, ok := blackList[candidate.Address]; ok {
		moniker = desc.Moniker
		identity = desc.Identity
		website = desc.Website
		details = desc.Details
	}
	return model.Validator{
		Address:        candidate.Address,
		PubKey:         utils.Convert(conf.Get().Hub.Prefix.ConsPub, candidate.PubKey),
		Owner:          utils.Convert(conf.Get().Hub.Prefix.AccAddr, candidate.Address),
		Jailed:         candidate.Jailed,
		Status:         candidate.Status,
		BondHeight:     candidate.BondHeight,
		OriginalTokens: candidate.OriginalTokens,
		VotingPower:    getVotingPowerFromToken(candidate.OriginalTokens),
		Description: model.Description{
			Moniker:  moniker,
			Identity: identity,
			Website:  website,
			Details:  details,
		},
	}
}

func getVotingPowerFromToken(token string) int64 {
	tokenPrecision := types.NewIntWithDecimal(1, 18)
	power, err := types.NewDecFromStr(token)
	if err != nil {
		logger.Error("invalid token", logger.String("token", token))
		return 0
	}
	return power.QuoInt(tokenPrecision).RoundInt64()
}

func getTotalVotingPower() int64 {
	var total = int64(0)
	var set = lcd.LatestValidatorSet()
	for _, v := range set.Validators {
		votingPower := utils.ParseIntWithDefault(v.VotingPower, 0)
		total += votingPower
	}
	return total
}

func BondStatusToString(b int) string {
	switch b {
	case 0:
		return types.TypeValStatusUnbonded
	case 1:
		return types.TypeValStatusUnbonding
	case 2:
		return types.TypeValStatusBonded
	default:
		panic("improper use of BondStatusToString")
	}
}

func (service *CandidateService) queryValForRainbow(typ string, page, size int) interface{} {
	var validators = lcd.Validators(page, size)

	var blackList = service.QueryBlackList()
	for i, v := range validators {
		if desc, ok := blackList[v.OperatorAddress]; ok {
			validators[i].Description.Moniker = desc.Moniker
			validators[i].Description.Identity = desc.Identity
			validators[i].Description.Website = desc.Website
			validators[i].Description.Details = desc.Details
		}
	}
	return validators
}

func (service *CandidateService) UpdateValidators(vs []document.Validator) error {
	var vMap = make(map[string]document.Validator)
	for _, v := range vs {
		vMap[v.OperatorAddress] = v
	}

	var txs []txn.Op
	dstValidators := buildValidators()
	for _, v := range dstValidators {
		if v1, ok := vMap[v.OperatorAddress]; ok {
			if isDiffValidator(v1, v) {
				v.ID = v1.ID
				txs = append(txs, txn.Op{
					C:  document.CollectionNmValidator,
					Id: v1.ID,
					Update: bson.M{
						"$set": v,
					},
				})
			}
			delete(vMap, v.OperatorAddress)
		} else {
			v.ID = bson.NewObjectId()
			txs = append(txs, txn.Op{
				C:      document.CollectionNmValidator,
				Id:     bson.NewObjectId(),
				Insert: v,
			})
		}
	}
	if len(vMap) > 0 {
		for addr := range vMap {
			v := vMap[addr]
			txs = append(txs, txn.Op{
				C:      document.CollectionNmValidator,
				Id:     v.ID,
				Remove: true,
			})
		}
	}
	return document.Candidate{}.Batch(txs)
}

func buildValidators() []document.Validator {

	res := lcd.Validators(1, 100)
	if res2 := lcd.Validators(2, 100); len(res2) > 0 {
		res = append(res, res2...)
	}

	var result []document.Validator
	height := utils.ParseIntWithDefault(lcd.BlockLatest().BlockMeta.Header.Height, 0)

	var buildValidator = func(v lcd.ValidatorVo) (document.Validator, error) {
		var validator document.Validator
		if err := utils.Copy(v, &validator); err != nil {
			logger.Error("utils.copy validator failed")
			return validator, err
		}
		validator.Uptime = computeUptime(v.ConsensusPubkey, height)
		validator.SelfBond, validator.ProposerAddr, validator.DelegatorNum = queryDelegationInfo(v.OperatorAddress, v.ConsensusPubkey)

		votingPower, err := types.NewDecFromStr(v.Tokens)
		if err == nil {
			validator.VotingPower = votingPower.RoundInt64()
		}

		return validator, nil
	}
	var group sync.WaitGroup
	group.Add(len(res))
	for _, v := range res {
		var genValidator = func(va lcd.ValidatorVo, result *[]document.Validator) {
			defer group.Done()
			validator, err := buildValidator(va)
			if err != nil {
				logger.Error("utils.copy validator failed")
				panic(err)
			}
			*result = append(*result, validator)
		}
		go genValidator(v, &result)
	}
	group.Wait()
	return result
}

func computeUptime(valPub string, height int64) float32 {
	result := lcd.SignInfo(valPub)
	missedBlocksCounter := utils.ParseIntWithDefault(result.MissedBlocksCounter, 0)
	startHeight := utils.ParseIntWithDefault(result.StartHeight, 0)
	tmp := float32(missedBlocksCounter) / float32(height-startHeight+1)
	return 1 - tmp
}

func queryDelegationInfo(operatorAddress string, consensusPubkey string) (string, string, int) {
	delegations := lcd.DelegationByValidator(operatorAddress)
	var selfBond string
	for _, d := range delegations {
		addr := utils.Convert(conf.Get().Hub.Prefix.AccAddr, operatorAddress)
		if d.DelegatorAddr == addr {
			selfBond = d.Shares
			break
		}
	}
	proposerAddr := utils.GenHexAddrFromPubKey(consensusPubkey)
	delegatorNum := len(delegations)
	return selfBond, proposerAddr, delegatorNum
}

func isDiffValidator(src, dst document.Validator) bool {
	if src.OperatorAddress != dst.OperatorAddress ||
		src.ConsensusPubkey != dst.ConsensusPubkey ||
		src.Jailed != dst.Jailed ||
		src.Status != dst.Status ||
		src.Tokens != dst.Tokens ||
		src.DelegatorShares != dst.DelegatorShares ||
		src.BondHeight != dst.BondHeight ||
		src.UnbondingHeight != dst.UnbondingHeight ||
		src.UnbondingTime.Second() != dst.UnbondingTime.Second() ||
		src.Uptime != dst.Uptime ||
		src.SelfBond != dst.SelfBond ||
		src.DelegatorNum != dst.DelegatorNum ||
		src.VotingPower != dst.VotingPower ||
		src.ProposerAddr != dst.ProposerAddr ||
		src.Description.Moniker != dst.Description.Moniker ||
		src.Description.Identity != dst.Description.Identity ||
		src.Description.Website != dst.Description.Website ||
		src.Description.Details != dst.Description.Details ||
		src.Commission.Rate != dst.Commission.Rate ||
		src.Commission.MaxRate != dst.Commission.MaxRate ||
		src.Commission.MaxChangeRate != dst.Commission.MaxChangeRate ||
		src.Commission.UpdateTime.Second() != dst.Commission.UpdateTime.Second() {
		logger.Info("validator has changed", logger.String("OperatorAddress", src.OperatorAddress))
		return true
	}
	return false
}
