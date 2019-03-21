package service

import (
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"math"
	"time"
)

type CandidateService struct {
	BaseService
}

func (service *CandidateService) GetModule() Module {
	return Candidate
}

func (service *CandidateService) QueryValidators(page, pageSize int) model.ValDetailVo {
	var candidates []document.Candidate
	db := getDb()
	defer db.Session.Close()
	cs := db.C(document.CollectionNmStakeRoleCandidate)
	cb := db.C(document.CollectionNmBlock)

	query := bson.M{}
	query[document.Candidate_Field_Jailed] = false
	query[document.Candidate_Field_Status] = types.TypeValStatusBonded

	err := cs.Find(query).Sort(desc(document.Candidate_Field_VotingPower)).Skip((page - 1) * pageSize).Limit(pageSize).All(&candidates)
	if err != nil {
		panic(types.CodeNotFound)
	}

	votePipe := cs.Pipe(
		[]bson.M{
			{"$match": query},
			{"$group": bson.M{
				"_id":   document.Candidate_Field_VotingPower,
				"count": bson.M{"$sum": "$tokens"},
			}},
		},
	)
	var voteCount model.CountVo
	votePipe.One(&voteCount)

	validatorsCount, _ := cs.Find(query).Count()
	var validators []model.Validator
	upTimeMap := make(map[string]int)
	var result []document.Block
	cb.Find(nil).Limit(100).Sort(desc(document.Block_Field_Height)).All(&result)
	for _, block := range result {
		for _, pre := range block.Block.LastCommit.Precommits {
			upTimeMap[pre.ValidatorAddress]++
		}
	}
	for _, candidate := range candidates {
		validator := service.convert(db, candidate)
		validator.Uptime = upTimeMap[candidate.PubKeyAddr]
		validator.TotalBlock = len(result)
		validators = append(validators, validator)
	}
	resp := model.ValDetailVo{
		Count:      validatorsCount,
		PowerAll:   getVotingPowerFromToken(voteCount.Count),
		Validators: validators,
	}
	return resp
}

func (service *CandidateService) GetValidators(page, size int) []lcd.ValidatorVo {
	var validators = lcd.Validators(page, size)
	db := getDb()
	defer db.Session.Close()
	var blackList = service.QueryBlackList(db)
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

func (service *CandidateService) GetValidator(address string) lcd.ValidatorVo {
	var validator, err = lcd.Validator(address)
	if err != nil {
		panic(types.CodeNotFound)
	}
	db := getDb()
	defer db.Session.Close()
	var blackList = service.QueryBlackList(db)
	if desc, ok := blackList[validator.OperatorAddress]; ok {
		validator.Description.Moniker = desc.Moniker
		validator.Description.Identity = desc.Identity
		validator.Description.Website = desc.Website
		validator.Description.Details = desc.Details
	}
	return validator
}

func (service *CandidateService) QueryRevokedValidator(page, size int) model.ValDetailVo {
	var candidates []document.Candidate
	db := getDb()
	defer db.Session.Close()
	cs := db.C(document.CollectionNmStakeRoleCandidate)

	query := bson.M{}
	query[document.Candidate_Field_Jailed] = true

	validatorsCount, _ := cs.Find(query).Count()

	err := cs.Find(query).Sort(desc(document.Candidate_Field_VotingPower)).Skip((page - 1) * size).Limit(size).All(&candidates)
	if err != nil {
		panic(types.CodeNotFound)
	}

	var validators []model.Validator
	for _, ca := range candidates {
		validators = append(validators, service.convert(db, ca))
	}

	resp := model.ValDetailVo{
		Count:      validatorsCount,
		Validators: validators,
	}
	return resp
}

func (service *CandidateService) QueryCandidates(page, size int) model.ValDetailVo {
	var candidates []document.Candidate
	db := getDb()
	defer db.Session.Close()
	cs := db.C(document.CollectionNmStakeRoleCandidate)
	txDoc := db.C(document.CollectionNmCommonTx)

	query := bson.M{}
	query[document.Candidate_Field_Jailed] = false
	query[document.Candidate_Field_Status] = bson.M{
		"$in": []string{types.TypeValStatusUnbonded, types.TypeValStatusUnbonding},
	}

	validatorsCount, _ := cs.Find(query).Count()
	err := cs.Find(query).Sort(desc(document.Candidate_Field_VotingPower)).Skip((page - 1) * size).Limit(size).All(&candidates)
	if err != nil {
		panic(types.CodeNotFound)
	}
	var validators []model.Validator
	var resp model.ValDetailVo
	if len(candidates) > 0 {
		q := bson.M{}
		q[document.Tx_Field_Type] = types.TypeCreateValidator

		for _, ca := range candidates {
			var tx document.CommonTx
			q[document.Tx_Field_From] = ca.Address
			txDoc.Find(q).Sort(document.Tx_Field_Height).One(&tx)

			ca.BondHeight = tx.Height
			validators = append(validators, service.convert(db, ca))
		}

		resp = model.ValDetailVo{
			Count:      validatorsCount,
			Validators: validators,
		}
	}
	return resp
}

func (service *CandidateService) QueryCandidate(address string) model.CandidatesInfoVo {
	db := getDb()
	defer db.Session.Close()
	c := db.C(document.CollectionNmStakeRoleCandidate)

	validator, err := lcd.Validator(address)
	if err != nil {
		panic(types.CodeNotFound)
	}
	var moniker = validator.Description.Moniker
	var identity = validator.Description.Identity
	var website = validator.Description.Website
	var details = validator.Description.Details
	var blackList = service.QueryBlackList(db)
	if desc, ok := blackList[validator.OperatorAddress]; ok {
		moniker = desc.Moniker
		identity = desc.Identity
		website = desc.Website
		details = desc.Details
	}
	var val = model.Validator{
		Address:        validator.OperatorAddress,
		PubKey:         validator.ConsensusPubkey,
		Jailed:         validator.Jailed,
		Status:         BondStatusToString(validator.Status),
		BondHeight:     utils.ParseIntWithDefault(validator.BondHeight, 0),
		OriginalTokens: utils.RoundToString(validator.Tokens, 0),
		VotingPower:    utils.RoundString(validator.Tokens),
		Description: model.Description{
			Moniker:  moniker,
			Identity: identity,
			Website:  website,
			Details:  details,
		},
		Rate: validator.Commission.Rate,
	}

	query := bson.M{}
	query[document.Candidate_Field_Jailed] = false
	query[document.Candidate_Field_Status] = types.TypeValStatusBonded
	pipe := c.Pipe(
		[]bson.M{
			{"$match": query},
			{"$group": bson.M{
				"_id":   document.Candidate_Field_VotingPower,
				"count": bson.M{"$sum": "$tokens"},
			}},
		},
	)
	var count model.CountVo
	pipe.One(&count)
	result := model.CandidatesInfoVo{
		PowerAll:  getVotingPowerFromToken(count.Count),
		Validator: val,
	}
	return result
}

func (service *CandidateService) QueryCandidatesTopN() model.ValDetailVo {
	var candidates []document.Candidate

	db := getDb()
	defer db.Session.Close()
	cs := db.C(document.CollectionNmStakeRoleCandidate)
	cb := db.C(document.CollectionNmBlock)

	query := bson.M{}
	query[document.Candidate_Field_Jailed] = false
	query[document.Candidate_Field_Status] = types.TypeValStatusBonded

	err := cs.Find(query).Sort(desc(document.Candidate_Field_VotingPower)).Limit(10).All(&candidates)
	if err != nil {
		panic(types.CodeNotFound)
	}
	votePipe := cs.Pipe(
		[]bson.M{
			{"$match": query},
			{"$group": bson.M{
				"_id":   document.Candidate_Field_VotingPower,
				"count": bson.M{"$sum": "$tokens"},
			}},
		},
	)
	var voteCount model.CountVo
	votePipe.One(&voteCount)

	var validators []model.Validator
	upTimeMap := make(map[string]int)
	var result []document.Block
	cb.Find(nil).Limit(100).Sort(desc(document.Block_Field_Height)).All(&result)
	for _, block := range result {
		for _, pre := range block.Block.LastCommit.Precommits {
			upTimeMap[pre.ValidatorAddress]++
		}
	}
	for _, candidate := range candidates {
		validator := service.convert(db, candidate)
		validator.Uptime = upTimeMap[candidate.PubKeyAddr]
		validator.TotalBlock = len(result)
		validators = append(validators, validator)
	}
	resp := model.ValDetailVo{
		PowerAll:   getVotingPowerFromToken(voteCount.Count),
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
	db := getDb()
	defer db.Session.Close()
	cs := db.C(document.CollectionNmStakeRoleCandidate)
	cb := db.C(document.CollectionNmBlock)
	var candidate document.Candidate
	err := cs.Find(bson.M{document.Candidate_Field_Address: address}).One(&candidate)
	if err != nil {
		panic(types.CodeNotFound)
	}

	var upTime int
	var result []document.Block
	cb.Find(nil).Limit(100).Sort(desc(document.Block_Field_Height)).All(&result)
	validatorAddress := ""
	for _, block := range result {
		var index1 int
		for _, validator := range block.Validators {
			if validatorAddress == "" && validator.Address == candidate.PubKeyAddr {
				validatorAddress = validator.Address
			}
			if index1+1 <= len(block.Block.LastCommit.Precommits) && block.Block.LastCommit.Precommits[index1].ValidatorAddress == validator.Address {
				if validator.Address == candidate.PubKeyAddr {
					upTime++
				}
				index1++
			}
		}
	}
	precommitCount, _ := cb.Find(bson.M{"block.last_commit.precommits": bson.M{"$elemMatch": bson.M{"validator_address": validatorAddress}}}).Count()
	resp = model.ValStatus{
		Uptime:         upTime,
		TotalBlock:     len(result),
		PrecommitCount: float64(precommitCount),
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

	resp := model.ChainStatusVo{
		ValidatorsCount: activeValidatorsCnt,
		TxCount:         txCount,
		VotingPower:     getVotingPowerFromToken(count.Count),
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
		Jailed:         candidate.Jailed,
		Status:         candidate.Status,
		BondHeight:     candidate.BondHeight,
		OriginalTokens: candidate.OriginalTokens,
		VotingPower:    getVotingPowerFromToken(candidate.Tokens),
		Description: model.Description{
			Moniker:  moniker,
			Identity: identity,
			Website:  website,
			Details:  details,
		},
	}
}

func getVotingPowerFromToken(token float64) int64 {
	return utils.Round(token / math.Pow10(18))
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
