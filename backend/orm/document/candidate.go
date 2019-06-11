package document

import (
	"fmt"
	"strconv"
	"time"

	"github.com/irisnet/explorer/backend/orm"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/irishub-sync/logger"
	"github.com/irisnet/irishub-sync/store/document"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
)

const (
	CollectionNmStakeRoleCandidate = "stake_role_candidate"

	Candidate_Field_Address         = "address"
	Candidate_Field_PubKey          = "pub_key"
	Candidate_Field_PubKeyAddr      = "pub_key_addr"
	Candidate_Field_Jailed          = "jailed"
	Candidate_Field_Tokens          = "tokens"
	Candidate_Field_OriginalTokens  = "original_tokens"
	Candidate_Field_DelegatorShares = "delegator_shares"
	Candidate_Field_VotingPower     = "voting_power"
	Candidate_Field_Description     = "description"
	Candidate_Field_BondHeight      = "bond_height"
	Candidate_Field_Status          = "status"
)

type (
	Candidate struct {
		Address         string         `bson:"address"` // owner, identity key
		PubKey          string         `bson:"pub_key"`
		PubKeyAddr      string         `bson:"pub_key_addr"`
		Jailed          bool           `bson:"jailed"` // has the validator been revoked from bonded status
		Tokens          float64        `bson:"tokens"`
		OriginalTokens  string         `bson:"original_tokens"`
		DelegatorShares float64        `bson:"delegator_shares"`
		VotingPower     float64        `bson:"voting_power"` // Voting power if pubKey is a considered a validator
		Description     ValDescription `bson:"description"`  // Description terms for the candidate
		BondHeight      int64          `bson:"bond_height"`
		Status          string         `bson:"status"`
		Rank            int            `bson:"rank,omitempty"`
	}

	UptimeChangeVo struct {
		Address string
		Time    string
		Uptime  float64
	}

	ValVotingPowerChangeVo struct {
		Height  int64
		Address string
		Power   int64
		Time    time.Time
		Change  string
	}

	ValUpTimeVo struct {
		Time   string `bson:"_id,omitempty"`
		Uptime float64
	}

	CountVo struct {
		Id    bson.ObjectId `bson:"_id,omitempty"`
		Count float64
	}
)

func (d Candidate) String() string {
	return fmt.Sprintf(`
		Address         :%v
		PubKey          :%v
		PubKeyAddr      :%v
		Jailed          :%v
		Tokens          :%v
		OriginalTokens  :%v
		DelegatorShares :%v
		VotingPower     :%v
		Description     :%v
		BondHeight      :%v
		Status          :%v
		Rank            :%v
		`, d.Address, d.PubKey, d.PubKeyAddr, d.Jailed, d.Tokens, d.OriginalTokens, d.DelegatorShares, d.VotingPower, d.Description, d.BondHeight, d.Status, d.Rank)
}

func (d Candidate) Name() string {
	return CollectionNmStakeRoleCandidate
}

func (d Candidate) PkKvPair() map[string]interface{} {
	return bson.M{Candidate_Field_Address: d.Address}
}

func (_ Candidate) GetValidatorListByPage(typ string, page, size int) (int, []Validator, error) {

	var query = orm.NewQuery()
	defer query.Release()
	var validators []Validator
	condition := bson.M{}
	switch typ {
	case types.RoleValidator:
		condition[ValidatorFieldJailed] = false
		condition[ValidatorFieldStatus] = types.Bonded
		break
	case types.RoleCandidate:
		condition[ValidatorFieldJailed] = false
		condition[Candidate_Field_Status] = bson.M{
			"$in": []int{types.Unbonded, types.Unbonding},
		}
		break
	case types.RoleJailed:
		condition[ValidatorFieldJailed] = true
		break
	default:
	}

	query.SetCollection(CollectionNmValidator).
		SetCondition(condition).
		SetSort(desc(ValidatorFieldVotingPower)).
		SetPage(page).
		SetSize(size).
		SetResult(&validators)

	count, err := query.ExecPage()

	return count, validators, err
}

func (_ Candidate) GetCandidatesTopN() ([]Candidate, string, map[string]int, error) {
	var candidates []Candidate
	var query = orm.NewQuery()
	defer query.Release()

	condition := bson.M{}
	condition[Candidate_Field_Jailed] = false
	condition[Candidate_Field_Status] = types.TypeValStatusBonded

	query.SetCollection(CollectionNmStakeRoleCandidate).
		SetCondition(condition).
		SetSort(desc(Candidate_Field_VotingPower)).SetSize(10).
		SetResult(&candidates)

	err := query.Exec()

	var allPower CountVo
	query.SetResult(&allPower)
	query.PipeQuery(
		[]bson.M{
			{"$match": condition},
			{"$group": bson.M{
				"_id":   Candidate_Field_VotingPower,
				"count": bson.M{"$sum": "$tokens"},
			}},
		},
	)

	power := strconv.FormatFloat(allPower.Count, 'f', 10, 64)

	upTimeMap := getValUpTime(query)

	return candidates, power, upTimeMap, err
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

func (_ Candidate) GetCandidatePubKeyAddrByAddr(addr string) (string, error) {
	db := getDb()
	c := db.C(CollectionNmStakeRoleCandidate)
	//u := db.C(CollectionNmUptimeChange)
	defer db.Session.Close()
	var candidate Candidate
	err := c.Find(bson.M{Candidate_Field_Address: addr}).One(&candidate)
	return candidate.PubKeyAddr, err
}

func (_ Candidate) QueryCandidateUptimeWithHour(addr string) ([]UptimeChangeVo, error) {

	db := getDb()
	u := db.C(CollectionNmUptimeChange)

	var upChanges []UptimeChangeVo
	now := time.Now()
	endTime := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location())
	d, _ := time.ParseDuration("-24h")
	startTime := endTime.Add(d)
	startStr := startTime.UTC().Format("2006-01-02 15")
	endStr := endTime.UTC().Format("2006-01-02 15")
	err := u.Find(bson.M{"address": addr, "time": bson.M{"$gte": startStr, "$lt": endStr}}).All(&upChanges)
	if err != nil {
		return nil, err
	}

	upChangeMap := make(map[string]float64)
	for _, upChange := range upChanges {
		upChangeMap[upChange.Time] = upChange.Uptime
	}
	d1, _ := time.ParseDuration("1h")
	result := []UptimeChangeVo{}
	for startTime.Before(endTime) {
		startStr := startTime.UTC().Format("2006-01-02 15")
		var uptime = float64(-1)
		if _, ok := upChangeMap[startStr]; ok {
			uptime = upChangeMap[startStr]
		}
		result = append(result, UptimeChangeVo{Address: addr, Uptime: uptime, Time: startStr})
		startTime = startTime.Add(d1)
	}

	return result, nil
}

func (_ Candidate) QueryCandidateUptimeByWeekOrMonth(addr, category string) ([]UptimeChangeVo, error) {

	db := getDb()
	u := db.C(CollectionNmUptimeChange)

	var upChanges []ValUpTimeVo
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
				"address": addr,
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
	err := pipe.All(&upChanges)
	if err != nil {
		return nil, err
	}

	upChangeMap := make(map[string]float64)
	for _, upChange := range upChanges {
		upChangeMap[upChange.Time] = upChange.Uptime
	}
	d1, _ := time.ParseDuration("24h")
	result := []UptimeChangeVo{}
	for startTime.Before(endTime) {
		startStr := startTime.UTC().Format("2006-01-02")
		var uptime = float64(-1)
		if _, ok := upChangeMap[startStr]; ok {
			uptime = upChangeMap[startStr]
		}
		result = append(result, UptimeChangeVo{Address: addr, Uptime: uptime, Time: startStr})
		startTime = startTime.Add(d1)
	}

	return result, nil
}

func (_ Candidate) QueryCandidatePower(address, agoStr string) ([]ValVotingPowerChangeVo, error) {

	db := getDb()
	c := db.C(CollectionNmStakeRoleCandidate)
	p := db.C(CollectionNmPowerChange)
	defer db.Session.Close()
	var candidate Candidate
	err := c.Find(bson.M{Candidate_Field_Address: address}).One(&candidate)
	if err != nil {
		return nil, err
	}
	address = candidate.PubKeyAddr

	var powers []ValVotingPowerChangeVo

	now := time.Now()
	endTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	d, _ := time.ParseDuration(agoStr)
	startTime := endTime.Add(d)
	p.Find(bson.M{"address": address, "time": bson.M{"$gte": startTime, "$lt": endTime}}).All(&powers)

	var power ValVotingPowerChangeVo
	p.Find(bson.M{"address": address, "time": bson.M{"$lt": startTime}}).Sort("-time").One(&power)

	result := []ValVotingPowerChangeVo{}
	if power.Address != "" {
		power.Time = startTime
		result = []ValVotingPowerChangeVo{power}
	} else {
		result = []ValVotingPowerChangeVo{{Address: address, Time: startTime}}
	}
	result = append(result, powers...)
	result = append(result, ValVotingPowerChangeVo{Address: address, Time: endTime, Power: result[len(result)-1].Power})
	return result, nil
}

func (_ Candidate) QueryCandidateStatus(addr string) (int, int, error) {

	var query = orm.NewQuery()
	defer query.Release()

	var candidate Candidate
	query.SetCollection(CollectionNmStakeRoleCandidate).
		SetCondition(bson.M{Candidate_Field_Address: addr}).
		SetResult(&candidate)
	if err := query.Exec(); err != nil {
		return 0, 0, err
	}

	upTimeMap := getValUpTime(query)
	query.Reset().SetCollection(CollectionNmBlock).
		SetCondition(bson.M{"block.last_commit.precommits": bson.M{"$elemMatch": bson.M{"validator_address": candidate.PubKeyAddr}}})

	var preCommitCount int
	if cnt, err := query.Count(); err == nil {
		preCommitCount = cnt
	}

	return preCommitCount, upTimeMap[candidate.PubKeyAddr], nil
}

func (_ Candidate) QueryCandidateListByAddrList(addrs []string) ([]Candidate, error) {
	candidateArr := []Candidate{}

	canCondition := bson.M{
		Candidate_Field_Address: bson.M{"$in": addrs},
	}

	err := queryAll(CollectionNmStakeRoleCandidate, nil, canCondition, "", 0, &candidateArr)

	return candidateArr, err
}

func (_ Candidate) QueryValidatorByConsensusAddr(addr string) (Candidate, error) {
	var query = orm.NewQuery()
	defer query.Release()

	var result Candidate
	condition := bson.M{}
	condition[Candidate_Field_PubKeyAddr] = addr

	query.SetCollection(CollectionNmStakeRoleCandidate).
		SetResult(&result).
		SetCondition(condition).
		SetSize(1)
	err := query.Exec()

	return result, err
}

func (_ Candidate) QueryPowerWithBonded() (float64, error) {

	var query = orm.NewQuery()
	defer query.Release()

	condition := bson.M{}
	condition[Candidate_Field_Jailed] = false
	condition[Candidate_Field_Status] = types.TypeValStatusBonded

	var count CountVo
	query.Reset().
		SetResult(&count).
		SetCollection(CollectionNmStakeRoleCandidate)
	err := query.PipeQuery([]bson.M{
		{"$match": condition},
		{"$group": bson.M{
			"_id":   Candidate_Field_VotingPower,
			"count": bson.M{"$sum": "$tokens"},
		}},
	})

	return count.Count, err
}

func (_ Candidate) Batch(txs []txn.Op) error {
	return orm.Batch(txs)
}
