package service

import (
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
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
	var query = orm.NewQuery()
	defer query.Release()

	condition := bson.M{}
	condition[document.Candidate_Field_Jailed] = false
	condition[document.Candidate_Field_Status] = types.TypeValStatusBonded

	query.SetCollection(document.CollectionNmStakeRoleCandidate).
		SetCondition(condition).
		SetSort(desc(document.Candidate_Field_VotingPower)).
		SetPage(page).
		SetSize(pageSize).
		SetResult(&candidates)
	count, err := query.ExecPage()

	if err != nil {
		panic(types.CodeNotFound)
	}

	var voteCount model.CountVo
	query.SetResult(&voteCount)
	var pipeline = []bson.M{
		{"$match": condition},
		{"$group": bson.M{
			"_id":   document.Candidate_Field_VotingPower,
			"count": bson.M{"$sum": "$tokens"},
		}},
	}

	if err = query.PipeQuery(pipeline); err != nil {
		panic(types.CodeNotFound)
	}

	var upTimeMap = getValUpTime(query)
	var validators []model.Validator
	for _, candidate := range candidates {
		validator := convert(candidate)
		validator.Uptime = upTimeMap[candidate.PubKeyAddr]
		validators = append(validators, validator)
	}
	resp := model.ValDetailVo{
		Count:      count,
		PowerAll:   getVotingPowerFromToken(voteCount.Count),
		Validators: validators,
	}
	return resp
}

func (service *CandidateService) QueryRevokedValidator(page, size int) model.ValDetailVo {
	var candidates []document.Candidate
	var query = orm.NewQuery()
	defer query.Release()
	query.Reset().
		SetCollection(document.CollectionNmStakeRoleCandidate).
		SetCondition(bson.M{document.Candidate_Field_Jailed: true}).
		SetPage(page).
		SetSize(size).
		SetSort(desc(document.Candidate_Field_VotingPower)).
		SetResult(&candidates)

	count, err := query.ExecPage()
	if err != nil {
		logger.Info("QueryRevokedValidator error", logger.String("err", err.Error()))
	}

	var validators []model.Validator
	for _, ca := range candidates {
		validators = append(validators, convert(ca))
	}

	resp := model.ValDetailVo{
		Count:      count,
		Validators: validators,
	}
	return resp
}

func (service *CandidateService) QueryCandidates(page, size int) model.ValDetailVo {
	var candidates []document.Candidate
	var query = orm.NewQuery()
	defer query.Release()

	var condition = bson.M{}
	condition[document.Candidate_Field_Jailed] = false
	condition[document.Candidate_Field_Status] = bson.M{
		"$in": []string{types.TypeValStatusUnbonded, types.TypeValStatusUnbonding},
	}

	query.Reset().
		SetCollection(document.CollectionNmStakeRoleCandidate).
		SetCondition(condition).
		SetPage(page).
		SetSize(size).
		SetSort(desc(document.Candidate_Field_VotingPower)).
		SetResult(&candidates)

	count, err := query.ExecPage()
	if err != nil {
		logger.Info("QueryRevokedValidator error", logger.String("err", err.Error()))
	}

	var validators []model.Validator
	var resp model.ValDetailVo
	if len(candidates) > 0 {
		var tx document.CommonTx
		var condition = bson.M{}
		var selector = bson.M{"height": 1}
		condition[document.Tx_Field_Type] = types.TypeCreateValidator
		query.Reset().
			SetCollection(document.CollectionNmCommonTx).
			SetSelector(selector).
			SetSort(desc(document.Tx_Field_Height)).
			SetResult(&tx)
		for _, ca := range candidates {
			condition[document.Tx_Field_From] = ca.Address
			query.SetCondition(condition)
			if err := query.Exec(); err == nil {
				ca.BondHeight = tx.Height
			}
			validators = append(validators, convert(ca))
		}

		resp = model.ValDetailVo{
			Count:      count,
			Validators: validators,
		}
	}
	return resp
}

func (service *CandidateService) QueryCandidate(address string) model.CandidatesInfoVo {
	var candidate document.Candidate
	var query = orm.NewQuery()
	defer query.Release()

	query.SetCollection(document.CollectionNmStakeRoleCandidate).
		SetCondition(bson.M{document.Candidate_Field_Address: address}).
		SetResult(&candidate)

	if err := query.Exec(); err != nil {
		panic(types.CodeNotFound)
	}

	condition := bson.M{}
	condition[document.Candidate_Field_Jailed] = false
	condition[document.Candidate_Field_Status] = types.TypeValStatusBonded

	var count model.CountVo
	query.SetResult(&count)
	query.PipeQuery([]bson.M{
		{"$match": condition},
		{"$group": bson.M{
			"_id":   document.Candidate_Field_VotingPower,
			"count": bson.M{"$sum": "$tokens"},
		}},
	})

	result := model.CandidatesInfoVo{
		PowerAll:  getVotingPowerFromToken(count.Count),
		Validator: convert(candidate),
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

	var validators []model.Validator
	var upTimeMap = getValUpTime(query)
	for _, candidate := range candidates {
		validator := convert(candidate)
		validator.Uptime = upTimeMap[candidate.PubKeyAddr]
		validators = append(validators, validator)
	}
	resp := model.ValDetailVo{
		PowerAll:   getVotingPowerFromToken(allPower.Count),
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

	resp := model.ChainStatusVo{
		ValidatorsCount: activeValidatorsCnt,
		TxCount:         txCount,
		VotingPower:     getVotingPowerFromToken(count.Count),
		Tps:             float64(txCnt) / 60,
	}
	return resp
}

func convert(candidate document.Candidate) model.Validator {
	return model.Validator{
		Address:        candidate.Address,
		PubKey:         candidate.PubKey,
		Jailed:         candidate.Jailed,
		Status:         candidate.Status,
		BondHeight:     candidate.BondHeight,
		OriginalTokens: candidate.OriginalTokens,
		VotingPower:    getVotingPowerFromToken(candidate.Tokens),
		Description: model.Description{
			Moniker:  candidate.Description.Moniker,
			Identity: candidate.Description.Identity,
			Website:  candidate.Description.Website,
			Details:  candidate.Description.Details,
		},
	}
}

func getVotingPowerFromToken(token float64) int64 {
	return utils.Round(token / math.Pow10(18))
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
