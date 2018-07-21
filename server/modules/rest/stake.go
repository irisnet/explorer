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
	PowerAll   int64
	Candidates []CandidateAll
}

type CandidateAll struct {
	document.Candidate
	CandidateStatus
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
	var result document.Candidate
	err := c.Find(bson.M{"address": address}).Sort("-update_time").One(&result)
	if err == nil {
		resultByte, _ := json.Marshal(result)
		w.Write(resultByte)
	}
}

func queryCandidateStatus(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	address := args["address"]
	c := utils.GetDatabase().C("block")
	defer c.Database.Session.Close()
	var result []document.Block
	var uptime int
	c.Find(nil).Limit(100).Sort("height").All(&result)
	for _, block := range result {
		for _, pre := range block.Block.LastCommit.Precommits {
			if pre.ValidatorAddress == address {
				uptime++
				break
			}
		}
	}
	precommitCount, _ := c.Find(bson.M{"block.last_commit.precommits":
	bson.M{"$elemMatch": bson.M{"validator_address": address}}}).Count()
	resp := CandidateStatus{
		Uptime:         uptime,
		TotalBlock:     len(result),
		PrecommitCount: precommitCount,
	}
	resultByte, _ := json.Marshal(resp)
	w.Write(resultByte)
}

func queryCandidates(w http.ResponseWriter, r *http.Request) {
	var data []document.Candidate
	w.Write(utils.QueryList("stake_role_candidate", &data, nil, "-update_time", r))
}

func queryCandidatesTop(w http.ResponseWriter, r *http.Request) {
	var data []document.Candidate
	cs := utils.GetDatabase().C("stake_role_candidate")
	defer cs.Database.Session.Close()
	cb := utils.GetDatabase().C("block")
	defer cb.Database.Session.Close()
	cs.Find(nil).Sort("-voting_power").All(&data)
	var powerAll int64
	for _, candidate := range data {
		powerAll += candidate.VotingPower
	}
	length := len(data)
	if length > 10 {
		length = 10
	}
	candidates := data[:length]
	var candidatesAll []CandidateAll
	for _, candidate := range candidates {
		var result []document.Block
		var uptime int
		cb.Find(nil).Limit(100).Sort("height").All(&result)
		for _, block := range result {
			for _, pre := range block.Block.LastCommit.Precommits {
				if pre.ValidatorAddress == candidate.Address {
					uptime++
					break
				}
			}
		}
		precommitCount, _ := cb.Find(bson.M{"block.last_commit.precommits":
		bson.M{"$elemMatch": bson.M{"validator_address": candidate.Address}}}).Count()
		status := CandidateStatus{
			Uptime:         uptime,
			TotalBlock:     len(result),
			PrecommitCount: precommitCount,
		}
		candidateAll := CandidateAll{
			Candidate:candidate,
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
