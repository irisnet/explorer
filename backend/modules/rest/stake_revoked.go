package rest

import (
	"encoding/json"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/irishub-sync/store/document"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func queryJailedValidator(w http.ResponseWriter, r *http.Request) {
	var candidates []document.Candidate
	page, size := utils.TransferPage(r)
	db := utils.GetDatabase()
	defer db.Session.Close()
	cs := db.C("stake_role_candidate")

	query := bson.M{}
	query["jailed"] = true

	validatorsCount, _ := cs.Find(query).Count()

	cs.Find(query).Sort("-voting_power").Skip((page - 1) * size).Limit(size).All(&candidates)

	var result []CandidateAll
	for _, ca := range candidates {
		result = append(result, CandidateAll{
			Candidate: ca,
		})
	}

	resp := Candidates{
		Count:      validatorsCount,
		Candidates: result,
	}
	resultByte, _ := json.Marshal(resp)
	w.Write(resultByte)
}
