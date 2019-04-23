package service

import (
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
	"strconv"
	"sync"
	"time"
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
		var query = orm.NewQuery()
		defer query.Release()
		var blackList = service.QueryBlackList(query.GetDb())
		var validators []document.Validator
		condition := bson.M{}
		switch typ {
		case types.RoleValidator:
			condition[document.ValidatorFieldJailed] = false
			condition[document.ValidatorFieldStatus] = types.Bonded
			break
		case types.RoleCandidate:
			condition[document.ValidatorFieldJailed] = false
			condition[document.Candidate_Field_Status] = bson.M{
				"$in": []int{types.Unbonded, types.Unbonding},
			}
			break
		case types.RoleJailed:
			condition[document.ValidatorFieldJailed] = true
			break
		default:
		}

		query.SetCollection(document.CollectionNmValidator).
			SetCondition(condition).
			SetSort(desc(document.ValidatorFieldTokens)).
			SetPage(page).
			SetSize(size).
			SetResult(&validators)

		count, err := query.ExecPage()

		if err != nil || count <= 0 {
			panic(types.CodeNotFound)
		}

		var totalVotingPower = getTotalVotingPower()
		for i, v := range validators {
			if desc, ok := blackList[v.OperatorAddress]; ok {
				validators[i].Description.Moniker = desc.Moniker
				validators[i].Description.Identity = desc.Identity
				validators[i].Description.Website = desc.Website
				validators[i].Description.Details = desc.Details
			}
			var validator lcd.ValidatorVo
			if err := utils.Copy(validators[i], &validator); err != nil {
				panic(types.CodeSysFailed)
			}
			power, _ := types.NewDecFromStr(v.Tokens)
			validator.VotingRate = float32(power.RoundInt64()) / float32(totalVotingPower)
			result = append(result, validator)
		}

		return model.PageVo{
			Data:  result,
			Count: count,
		}
	}
	return service.queryValForRainbow(typ, page, size)
}

func (service *CandidateService) GetValidator(address string) lcd.ValidatorVo {
	var validator, err = lcd.Validator(address)
	if err != nil {
		panic(types.CodeNotFound)
	}
	var query = orm.NewQuery()
	defer query.Release()
	var blackList = service.QueryBlackList(query.GetDb())
	if desc, ok := blackList[validator.OperatorAddress]; ok {
		validator.Description.Moniker = desc.Moniker
		validator.Description.Identity = desc.Identity
		validator.Description.Website = desc.Website
		validator.Description.Details = desc.Details
	}
	return validator
}

func (service *CandidateService) QueryValidatorByConAddr(address string) document.Candidate {
	var query = orm.NewQuery()
	defer query.Release()

	var result document.Candidate
	condition := bson.M{}
	condition[document.Candidate_Field_PubKeyAddr] = address

	query.SetCollection(document.CollectionNmStakeRoleCandidate).
		SetResult(&result).
		SetCondition(condition).
		SetSize(1)

	if err := query.Exec(); err != nil {
		logger.Error("not found validator by conAddr", logger.String("conAddr", address))
	}
	return result
}

func (service *CandidateService) QueryCandidate(address string) model.CandidatesInfoVo {
	//var candidate document.Candidate

	validator, err := lcd.Validator(address)
	if err != nil {
		panic(types.CodeNotFound)
	}

	var query = orm.NewQuery()
	defer query.Release()

	var moniker = validator.Description.Moniker
	var identity = validator.Description.Identity
	var website = validator.Description.Website
	var details = validator.Description.Details
	var blackList = service.QueryBlackList(query.GetDb())
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

	condition := bson.M{}
	condition[document.Candidate_Field_Jailed] = false
	condition[document.Candidate_Field_Status] = types.TypeValStatusBonded

	var count model.CountVo
	query.Reset().
		SetResult(&count).
		SetCollection(document.CollectionNmStakeRoleCandidate)
	query.PipeQuery([]bson.M{
		{"$match": condition},
		{"$group": bson.M{
			"_id":   document.Candidate_Field_VotingPower,
			"count": bson.M{"$sum": "$tokens"},
		}},
	})
	power := strconv.FormatFloat(count.Count, 'f', 10, 64)
	result := model.CandidatesInfoVo{
		PowerAll:  getVotingPowerFromToken(power),
		Validator: val,
	}
	return result
}

func (service *CandidateService) QueryCandidatesTopN() model.ValDetailVo {
	var candidates []document.Candidate
	var query = orm.NewQuery()
	defer query.Release()

	condition := bson.M{}
	condition[document.Candidate_Field_Jailed] = false
	condition[document.Candidate_Field_Status] = types.TypeValStatusBonded

	query.SetCollection(document.CollectionNmStakeRoleCandidate).
		SetCondition(condition).
		SetSort(desc(document.Candidate_Field_VotingPower)).SetSize(10).
		SetResult(&candidates)
	if err := query.Exec(); err != nil {
		panic(types.CodeNotFound)
	}

	var allPower model.CountVo
	query.SetResult(&allPower)
	query.PipeQuery(
		[]bson.M{
			{"$match": condition},
			{"$group": bson.M{
				"_id":   document.Candidate_Field_VotingPower,
				"count": bson.M{"$sum": "$tokens"},
			}},
		},
	)

	power := strconv.FormatFloat(allPower.Count, 'f', 10, 64)

	var validators []model.Validator
	var upTimeMap = getValUpTime(query)
	for _, candidate := range candidates {
		validator := service.convert(query.GetDb(), candidate)
		validator.Uptime = upTimeMap[candidate.PubKeyAddr]
		validators = append(validators, validator)
	}
	resp := model.ValDetailVo{
		PowerAll:   getVotingPowerFromToken(power),
		Validators: validators,
	}
	return resp
}

func (service *CandidateService) QueryCandidateUptime(address, category string) (result []model.UptimeChangeVo) {
	db := getDb()
	c := db.C(document.CollectionNmStakeRoleCandidate)
	u := db.C("uptime_change")
	defer db.Session.Close()
	var candidate document.Candidate
	err := c.Find(bson.M{document.Candidate_Field_Address: address}).One(&candidate)
	address = candidate.PubKeyAddr
	if err != nil || address == "" {
		panic(types.CodeNotFound)
	}

	switch category {
	case "hour":
		var upChanges []model.UptimeChangeVo
		now := time.Now()
		endTime := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location())
		d, _ := time.ParseDuration("-24h")
		startTime := endTime.Add(d)
		startStr := startTime.UTC().Format("2006-01-02 15")
		endStr := endTime.UTC().Format("2006-01-02 15")
		u.Find(bson.M{"address": address, "time": bson.M{"$gte": startStr, "$lt": endStr}}).All(&upChanges)
		upChangeMap := make(map[string]float64)
		for _, upChange := range upChanges {
			upChangeMap[upChange.Time] = upChange.Uptime
		}
		d1, _ := time.ParseDuration("1h")
		for startTime.Before(endTime) {
			startStr := startTime.UTC().Format("2006-01-02 15")
			var uptime = float64(-1)
			if _, ok := upChangeMap[startStr]; ok {
				uptime = upChangeMap[startStr]
			}
			result = append(result, model.UptimeChangeVo{Address: address, Uptime: uptime, Time: startStr})
			startTime = startTime.Add(d1)
		}
		break
	case "week", "month":
		var upChanges []model.ValUpTimeVo
		agoStr := "-336h"
		if category == "month" {
			agoStr = "-720h"
		}
		now := time.Now()
		endTime := now
		d, _ := time.ParseDuration(agoStr)
		startTime := endTime.Add(d)
		startStr := startTime.Format("2006-01-02 15")
		endStr := endTime.Format("2006-01-02 15")
		pipe := u.Pipe(
			[]bson.M{
				{"$match": bson.M{
					"address": address,
					"time":    bson.M{"$gte": startStr, "$lt": endStr},
				}},
				{"$project": bson.M{
					"day":    bson.M{"$substr": []interface{}{"$time", 0, 10}},
					"uptime": "$uptime",
				}},
				{"$group": bson.M{
					"_id":    "$day",
					"uptime": bson.M{"$avg": "$uptime"},
				}},
				{"$sort": bson.M{
					"_id": 1,
				}},
			},
		)
		err = pipe.All(&upChanges)

		upChangeMap := make(map[string]float64)
		for _, upChange := range upChanges {
			upChangeMap[upChange.Time] = upChange.Uptime
		}
		d1, _ := time.ParseDuration("24h")
		for startTime.Before(endTime) {
			startStr := startTime.UTC().Format("2006-01-02")
			var uptime = float64(-1)
			if _, ok := upChangeMap[startStr]; ok {
				uptime = upChangeMap[startStr]
			}
			result = append(result, model.UptimeChangeVo{Address: address, Uptime: uptime, Time: startStr})
			startTime = startTime.Add(d1)
		}
	}
	return result
}

func (service *CandidateService) QueryCandidatePower(address, category string) (result []model.ValVotingPowerChangeVo) {
	db := getDb()
	c := db.C(document.CollectionNmStakeRoleCandidate)
	p := db.C("power_change")
	defer db.Session.Close()
	var candidate document.Candidate
	err := c.Find(bson.M{document.Candidate_Field_Address: address}).One(&candidate)
	address = candidate.PubKeyAddr
	if err != nil || address == "" {
		panic(types.CodeNotFound)
	}
	var powers []model.ValVotingPowerChangeVo
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
	now := time.Now()
	endTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	d, _ := time.ParseDuration(agoStr)
	startTime := endTime.Add(d)
	p.Find(bson.M{"address": address, "time": bson.M{"$gte": startTime, "$lt": endTime}}).All(&powers)

	var power model.ValVotingPowerChangeVo
	p.Find(bson.M{"address": address, "time": bson.M{"$lt": startTime}}).Sort("-time").One(&power)
	if power.Address != "" {
		power.Time = startTime
		result = []model.ValVotingPowerChangeVo{power}
	} else {
		result = []model.ValVotingPowerChangeVo{{Address: address, Time: startTime}}
	}
	result = append(result, powers...)
	result = append(result, model.ValVotingPowerChangeVo{Address: address, Time: endTime, Power: result[len(result)-1].Power})
	return result
}

func (service *CandidateService) QueryCandidateStatus(address string) (resp model.ValStatus) {

	var query = orm.NewQuery()
	defer query.Release()

	var candidate document.Candidate
	query.SetCollection(document.CollectionNmStakeRoleCandidate).
		SetCondition(bson.M{document.Candidate_Field_Address: address}).
		SetResult(&candidate)
	if err := query.Exec(); err != nil {
		panic(types.CodeNotFound)
	}

	var upTimeMap = getValUpTime(query)
	query.Reset().SetCollection(document.CollectionNmBlock).
		SetCondition(bson.M{"block.last_commit.precommits": bson.M{"$elemMatch": bson.M{"validator_address": candidate.PubKeyAddr}}})

	var preCommitCount int
	if cnt, err := query.Count(); err == nil {
		preCommitCount = cnt
	}

	resp = model.ValStatus{
		Uptime:         upTimeMap[candidate.PubKeyAddr],
		PrecommitCount: float64(preCommitCount),
	}

	return resp
}

func (service *CandidateService) QueryChainStatus() model.ChainStatusVo {
	store := getDb()
	defer store.Session.Close()
	candidateStore := store.C(document.CollectionNmStakeRoleCandidate)
	pipe := candidateStore.Pipe(
		[]bson.M{
			{"$group": bson.M{
				"_id":   "voting_power",
				"count": bson.M{"$sum": "$tokens"},
			}},
		},
	)
	var count model.CountVo
	pipe.One(&count)

	query := bson.M{}
	query[document.Candidate_Field_Status] = types.TypeValStatusBonded
	query[document.Candidate_Field_Jailed] = false
	activeValidatorsCnt, _ := candidateStore.Find(query).Count()

	blockStore := store.C(document.CollectionNmBlock)
	txStore := store.C(document.CollectionNmCommonTx)
	txCount, _ := txStore.Count()

	var recentBlock document.Block
	var txCnt int
	err := blockStore.Find(nil).Sort(desc(document.Block_Field_Height)).Limit(1).One(&recentBlock)
	if err == nil {
		startTime := recentBlock.Time.Add(-1 * time.Minute)
		txCnt, _ = txStore.Find(bson.M{document.Tx_Field_Time: bson.M{"$gte": startTime}}).Count()
	}
	power := strconv.FormatFloat(count.Count, 'f', 10, 64)
	resp := model.ChainStatusVo{
		ValidatorsCount: activeValidatorsCnt,
		TxCount:         txCount,
		VotingPower:     getVotingPowerFromToken(power),
		Tps:             float64(txCnt) / 60,
	}
	return resp
}

func (service *CandidateService) convert(database *mgo.Database, candidate document.Candidate) model.Validator {
	var moniker = candidate.Description.Moniker
	var identity = candidate.Description.Identity
	var website = candidate.Description.Website
	var details = candidate.Description.Details
	var blackList = service.QueryBlackList(database)
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

func getValUpTime(query *orm.Query) map[string]int {
	var result []document.Block
	var upTimeMap = make(map[string]int)
	var selector = bson.M{"block.last_commit.precommits.validator_address": 1}
	query.Reset().
		SetCollection(document.CollectionNmBlock).
		SetSelector(selector).
		SetSize(100).
		SetSort(desc(document.Block_Field_Height)).
		SetResult(&result)

	if err := query.Exec(); err != nil {
		logger.Error("getValUpTime error", logger.String("err", err.Error()))
	}
	for _, block := range result {
		for _, pre := range block.Block.LastCommit.Precommits {
			upTimeMap[pre.ValidatorAddress]++
		}
	}
	return upTimeMap
}

func (service *CandidateService) queryValForRainbow(typ string, page, size int) interface{} {
	var validators = lcd.Validators(page, size)
	var query = orm.NewQuery()
	defer query.Release()
	var blackList = service.QueryBlackList(query.GetDb())
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
	return orm.Batch(txs)
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
