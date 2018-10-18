package rest

import (
	"encoding/json"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/irishub-sync/store/document"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func queryValidators(w http.ResponseWriter, r *http.Request) {
	var candidates []document.Candidate
	page, size := utils.TransferPage(r)
	db := utils.GetDatabase()
	defer db.Session.Close()
	cs := db.C("stake_role_candidate")
	cb := db.C("block")

	query := bson.M{}
	query["revoked"] = false
	query["status"] = types.TypeValStatusBonded

	cs.Find(query).Sort("-voting_power").Skip((page - 1) * size).Limit(size).All(&candidates)
	votePipe := cs.Pipe(
		[]bson.M{
			{"$match": query},
			{"$group": bson.M{
				"_id":   "voting_power",
				"count": bson.M{"$sum": "$voting_power"},
			}},
		},
	)
	var voteCount count
	votePipe.One(&voteCount)

	validatorsCount, _ := cs.Find(query).Count()
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
