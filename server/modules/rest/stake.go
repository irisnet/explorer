package rest

import (
	"github.com/gorilla/mux"
	"net/http"

	"github.com/irisnet/explorer/server/utils"
	"github.com/irisnet/irishub-sync/store/document"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
)

func RegisterStake(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryCandidatesByAccount,
		registerQueryCandidates,
		registerQueryCandidate,
		registerQueryCandidateStatus,
		registerQueryCandidatesTop,
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
	PrecommitCount int
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

func queryCandidateStatus(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	address := args["address"]
	db:=utils.GetDatabase()
	cs := db.C("stake_role_candidate")
	cb := db.C("block")
	defer db.Session.Close()
	//var result []document.Block
	//var uptime int
	//c.Find(nil).Limit(100).Sort("height").All(&result)
	//for _, block := range result {
	//	for _, pre := range block.Block.LastCommit.Precommits {
	//		if pre.ValidatorAddress == address {
	//			uptime++
	//			break
	//		}
	//	}
	//}
	var candidate document.Candidate
	err := cs.Find(bson.M{"address": address}).Sort("-update_time").One(&candidate)
	upTimeMap := make(map[string]int)
	var result []document.Block
	cb.Find(nil).Limit(100).Sort("-height").All(&result)
	validatorAddress :=""
	for _, block := range result {
		var index1 int
		for _, validator := range block.Validators {
			if validatorAddress == "" && validator.PubKey == candidate.PubKey{
				validatorAddress = validator.Address
			}
			if block.Block.LastCommit.Precommits[index1].ValidatorAddress == validator.Address {
				upTimeMap[validator.PubKey] ++
				index1++
			}
		}
	}
	precommitCount, _ := cb.Find(bson.M{"block.last_commit.precommits":
	bson.M{"$elemMatch": bson.M{"validator_address": validatorAddress}}}).Count()
	resp := CandidateStatus{
		Uptime:         upTimeMap[candidate.PubKey],
		TotalBlock:     len(result),
		PrecommitCount: precommitCount,
	}
	if err == nil {
		resultByte, _ := json.Marshal(resp)
		w.Write(resultByte)
	}
}

func queryCandidates(w http.ResponseWriter, r *http.Request) {
	var data []document.Candidate
	page, size := utils.TransferPage(r)
	cs := utils.GetDatabase().C("stake_role_candidate")
	defer cs.Database.Session.Close()
	cb := utils.GetDatabase().C("block")
	defer cb.Database.Session.Close()
	cs.Find(nil).Sort("-voting_power").All(&data)
	var powerAll float64
	for _, candidate := range data {
		powerAll += candidate.VotingPower
	}
	length := len(data)
	start := (page - 1) * size
	end := page * size
	if end > length {
		end = length
	}
	candidates := data[start:end]
	var candidatesAll []CandidateAll
	upTimeMap := make(map[string]int)
	var result []document.Block
	cb.Find(nil).Limit(100).Sort("-height").All(&result)
	for _, block := range result {
		var index1 int
		for _, validator := range block.Validators {
			if block.Block.LastCommit.Precommits[index1].ValidatorAddress == validator.Address {
				upTimeMap[validator.PubKey] ++
				index1++
			}
		}
	}
	for _, candidate := range candidates {
		//var result []document.Block
		//var uptime int
		//cb.Find(nil).Limit(100).Sort("-height").All(&result)
		//for _, block := range result {
		//	for _, pre := range block.Block.LastCommit.Precommits {
		//		if pre.ValidatorAddress == candidate.Address {
		//			uptime++
		//			break
		//		}
		//	}
		//}
		status := CandidateStatus{
			Uptime:     upTimeMap[candidate.PubKey],
			TotalBlock: len(result),
		}
		candidateAll := CandidateAll{
			Candidate:       candidate,
			CandidateStatus: status,
		}
		candidatesAll = append(candidatesAll, candidateAll)
	}
	resp := Candidates{
		Count:      length,
		PowerAll:   powerAll,
		Candidates: candidatesAll,
	}
	resultByte, _ := json.Marshal(resp)
	w.Write(resultByte)
}

func queryCandidatesTop(w http.ResponseWriter, r *http.Request) {
	var data []document.Candidate
	cs := utils.GetDatabase().C("stake_role_candidate")
	defer cs.Database.Session.Close()
	cb := utils.GetDatabase().C("block")
	defer cb.Database.Session.Close()
	cs.Find(nil).Sort("-voting_power").All(&data)
	var powerAll float64
	for _, candidate := range data {
		powerAll += candidate.VotingPower
	}
	length := len(data)
	if length > 10 {
		length = 10
	}
	candidates := data[:length]
	var candidatesAll []CandidateAll
	upTimeMap := make(map[string]int)
	var result []document.Block
	cb.Find(nil).Limit(100).Sort("-height").All(&result)
	for _, block := range result {
		var index1 int
		for _, validator := range block.Validators {
			if block.Block.LastCommit.Precommits[index1].ValidatorAddress == validator.Address {
				upTimeMap[validator.PubKey] ++
				index1++
			}
		}
	}
	for _, candidate := range candidates {
		//var result []document.Block
		//var uptime int
		//cb.Find(nil).Limit(100).Sort("-height").All(&result)
		//for _, block := range result {
		//	var index1 int
		//	for _, validator := range block.Validators {
		//		if block.Block.LastCommit.Precommits[index1].ValidatorAddress == validator.Address {
		//			index1++
		//		}
		//		if validator.PubKey == candidate.PubKey {
		//			upTimeMap[validator.PubKey] ++
		//			break
		//		}
		//	}
		//}
		precommitCount, _ := cb.Find(bson.M{"block.last_commit.precommits":
		bson.M{"$elemMatch": bson.M{"validator_address": candidate.Address}}}).Count()
		status := CandidateStatus{
			Uptime:         upTimeMap[candidate.PubKey],
			TotalBlock:     len(result),
			PrecommitCount: precommitCount,
		}
		candidateAll := CandidateAll{
			Candidate:       candidate,
			CandidateStatus: status,
		}
		candidatesAll = append(candidatesAll, candidateAll)
	}
	resp := CandidatesTop{
		PowerAll:   powerAll,
		Candidates: candidatesAll,
	}
	resultByte, _ := json.Marshal(resp)
	w.Write(resultByte)
}
