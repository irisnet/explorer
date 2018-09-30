package rest

import (
	"net/http"

	"github.com/gorilla/mux"

	"encoding/json"

	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/irishub-sync/store/document"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func RegisterStake(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryCandidatesByAccount,
		registerQueryCandidates,
		registerQueryCandidate,
		registerQueryCandidateStatus,
		registerQueryCandidatesTop,
		registerQueryCandidateUptime,
		registerQueryCandidatePower,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

type CandidateStatus struct {
	Uptime         int
	TotalBlock     int
	PrecommitCount float64
}

type CandidatesTop struct {
	PowerAll   float64
	Candidates []CandidateAll
}

type Candidates struct {
	Count      int
	PowerAll   float64
	Candidates []CandidateAll
}

type CandidateAll struct {
	document.Candidate
	CandidateStatus
}

type CandidateWithPower struct {
	PowerAll float64
	document.Candidate
}

func registerQueryCandidatesByAccount(r *mux.Router) error {
	r.HandleFunc("/api/stake/candidates/{address}", queryCandidatesByAccount).Methods("GET")
	return nil
}

func registerQueryCandidates(r *mux.Router) error {
	r.HandleFunc("/api/stake/candidates/{page}/{size}", queryCandidates).Methods("GET")
	return nil
}

func registerQueryCandidatesTop(r *mux.Router) error {
	r.HandleFunc("/api/stake/candidatesTop", queryCandidatesTop).Methods("GET")
	return nil
}

func registerQueryCandidate(r *mux.Router) error {
	r.HandleFunc("/api/stake/candidate/{address}", queryCandidate).Methods("GET")
	return nil
}

func registerQueryCandidateUptime(r *mux.Router) error {
	r.HandleFunc("/api/stake/candidate/{address}/uptime/{category}", QueryCandidateUptime).Methods("GET")
	return nil
}

func registerQueryCandidatePower(r *mux.Router) error {
	r.HandleFunc("/api/stake/candidate/{address}/power/{category}", QueryCandidatePower).Methods("GET")
	return nil
}

func registerQueryCandidateStatus(r *mux.Router) error {
	r.HandleFunc("/api/stake/candidate/{address}/status", queryCandidateStatus).Methods("GET")
	return nil
}

func queryCandidatesByAccount(w http.ResponseWriter, r *http.Request) {

}

func queryCandidate(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	address := args["address"]

	c := utils.GetDatabase().C("stake_role_candidate")
	defer c.Database.Session.Close()
	var candidate document.Candidate
	err := c.Find(bson.M{"address": address}).Sort("-update_time").One(&candidate)
	pipe := c.Pipe(
		[]bson.M{
			{"$match": bson.M{"revoked": false}},
			{"$group": bson.M{
				"_id":   "voting_power",
				"count": bson.M{"$sum": "$voting_power"},
			}},
		},
	)
	var count count
	pipe.One(&count)
	result := CandidateWithPower{
		PowerAll:  count.Count,
		Candidate: candidate,
	}
	if err == nil {
		resultByte, _ := json.Marshal(result)
		w.Write(resultByte)
	}
}

func QueryCandidateUptime(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	address := args["address"]
	category := args["category"]

	db := utils.GetDatabase()
	c := db.C("stake_role_candidate")
	u := db.C("uptime_change")
	defer db.Session.Close()
	var candidate document.Candidate
	err := c.Find(bson.M{"address": address}).Sort("-update_time").One(&candidate)
	address = candidate.PubKeyAddr
	if err != nil || address == "" {
		return
	}

	switch category {
	case "hour":
		var upChanges []types.UptimeChange
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
		var result []types.UptimeChange
		d1, _ := time.ParseDuration("1h")
		d2, _ := time.ParseDuration("-1h")
		endTimeHour := endTime.Add(d2)
		for startTime.Before(endTime) {
			startStr := startTime.UTC().Format("2006-01-02 15")
			var uptime = float64(-1)
			if _, ok := upChangeMap[startStr]; ok {
				uptime = upChangeMap[startStr]
			} else if startTime.Equal(endTimeHour) {
				startTime = startTime.Add(d1)
				continue
			}
			result = append(result, types.UptimeChange{Address: address, Uptime: uptime, Time: startStr})
			startTime = startTime.Add(d1)
		}
		resultByte, _ := json.Marshal(result)
		w.Write(resultByte)
		break
	case "week", "month":
		var upChanges []types.CandidateUptime
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
		var result []types.UptimeChange
		d1, _ := time.ParseDuration("24h")
		d2, _ := time.ParseDuration("-24h")
		endTimeDay := endTime.Add(d2)
		for startTime.Before(endTime) {
			startStr := startTime.UTC().Format("2006-01-02")
			var uptime = float64(-1)
			if _, ok := upChangeMap[startStr]; ok {
				uptime = upChangeMap[startStr]
			} else if startTime.Equal(endTimeDay) {
				startTime = startTime.Add(d1)
				continue
			}
			result = append(result, types.UptimeChange{Address: address, Uptime: uptime, Time: startStr})
			startTime = startTime.Add(d1)
		}
		resultByte, _ := json.Marshal(result)
		w.Write(resultByte)
	}
}

func QueryCandidatePower(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	address := args["address"]
	category := args["category"]

	db := utils.GetDatabase()
	c := db.C("stake_role_candidate")
	p := db.C("power_change")
	defer db.Session.Close()
	var candidate document.Candidate
	err := c.Find(bson.M{"address": address}).Sort("-update_time").One(&candidate)
	address = candidate.PubKeyAddr
	if err != nil || address == "" {
		return
	}
	var powers []types.PowerChange
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

	var result []types.PowerChange
	var power types.PowerChange
	p.Find(bson.M{"address": address, "time": bson.M{"$lt": startTime}}).Sort("-time").One(&power)
	if power.Address != "" {
		result = []types.PowerChange{power}
	} else {
		result = []types.PowerChange{{Address: address, Time: startTime}}
	}
	result = append(result, powers...)
	result = append(result, types.PowerChange{Address: address, Time: endTime, Power: result[len(result)-1].Power})
	resultByte, _ := json.Marshal(result)
	w.Write(resultByte)
}

func queryCandidateStatus(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	address := args["address"]
	db := utils.GetDatabase()
	defer db.Session.Close()
	cs := db.C("stake_role_candidate")
	cb := db.C("block")
	var candidate document.Candidate
	err := cs.Find(bson.M{"address": address}).Sort("-update_time").One(&candidate)
	var upTime int
	var result []document.Block
	cb.Find(nil).Limit(100).Sort("-height").All(&result)
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
	resp := CandidateStatus{
		Uptime:         upTime,
		TotalBlock:     len(result),
		PrecommitCount: float64(precommitCount),
	}
	if err == nil {
		resultByte, _ := json.Marshal(resp)
		w.Write(resultByte)
	}
}

func queryCandidates(w http.ResponseWriter, r *http.Request) {
	var candidates []document.Candidate
	page, size := utils.TransferPage(r)
	db := utils.GetDatabase()
	defer db.Session.Close()
	cs := db.C("stake_role_candidate")
	cb := db.C("block")
	cs.Find(bson.M{"revoked": false}).Sort("-voting_power").Skip((page - 1) * size).Limit(size).All(&candidates)
	votePipe := cs.Pipe(
		[]bson.M{
			//{"$match": bson.M{"revoked": false}},
			{"$group": bson.M{
				"_id":   "voting_power",
				"count": bson.M{"$sum": "$voting_power"},
			}},
		},
	)
	var voteCount count
	votePipe.One(&voteCount)

	validatorsCount, _ := cs.Find(bson.M{"revoked": false}).Count()
	var candidatesAll []CandidateAll
	upTimeMap := make(map[string]int)
	var result []document.Block
	cb.Find(nil).Limit(100).Sort("-height").All(&result)
	for _, block := range result {
		for _, pre := range block.Block.LastCommit.Precommits {
			upTimeMap[pre.ValidatorAddress]++
		}
	}
	for _, candidate := range candidates {
		status := CandidateStatus{
			Uptime:     upTimeMap[candidate.PubKeyAddr],
			TotalBlock: len(result),
		}
		candidateAll := CandidateAll{
			Candidate:       candidate,
			CandidateStatus: status,
		}
		candidatesAll = append(candidatesAll, candidateAll)
	}
	resp := Candidates{
		Count:      validatorsCount,
		PowerAll:   voteCount.Count,
		Candidates: candidatesAll,
	}
	resultByte, _ := json.Marshal(resp)
	w.Write(resultByte)
}

func queryCandidatesTop(w http.ResponseWriter, r *http.Request) {
	var candidates []document.Candidate

	db := utils.GetDatabase()
	defer db.Session.Close()
	cs := db.C("stake_role_candidate")
	cb := db.C("block")
	cs.Find(bson.M{"revoked": false}).Sort("-voting_power").Limit(10).All(&candidates)
	votePipe := cs.Pipe(
		[]bson.M{
			{"$match": bson.M{"revoked": false}},
			{"$group": bson.M{
				"_id":   "voting_power",
				"count": bson.M{"$sum": "$voting_power"},
			}},
		},
	)
	var voteCount count
	votePipe.One(&voteCount)

	var candidatesAll []CandidateAll
	upTimeMap := make(map[string]int)
	var result []document.Block
	cb.Find(nil).Limit(100).Sort("-height").All(&result)
	for _, block := range result {
		for _, pre := range block.Block.LastCommit.Precommits {
			upTimeMap[pre.ValidatorAddress]++
		}
	}
	//prePipe := cb.Pipe(
	//	[]bson.M{
	//		{"$unwind": "$block.last_commit.precommits"},
	//		{"$group": bson.M{
	//			"_id":   "$block.last_commit.precommits.validator_address",
	//			"count": bson.M{"$sum": 1},
	//		}},
	//	},
	//)
	//var preCount []count
	//prePipe.All(&preCount)
	//preMap := make(map[string]float64)
	//for _, pre := range preCount {
	//	preMap[pre.Id.String()] = pre.Count
	//}
	for _, candidate := range candidates {
		status := CandidateStatus{
			Uptime:     upTimeMap[candidate.PubKeyAddr],
			TotalBlock: len(result),
		}
		candidateAll := CandidateAll{
			Candidate:       candidate,
			CandidateStatus: status,
		}
		candidatesAll = append(candidatesAll, candidateAll)
	}
	resp := CandidatesTop{
		PowerAll:   voteCount.Count,
		Candidates: candidatesAll,
	}
	resultByte, _ := json.Marshal(resp)
	w.Write(resultByte)
}
