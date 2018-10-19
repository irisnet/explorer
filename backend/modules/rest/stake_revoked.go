package rest

import (
	"encoding/json"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/irishub-sync/store/document"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func queryRevokedValidator(w http.ResponseWriter, r *http.Request) {
	var candidates []document.Candidate
	page, size := utils.TransferPage(r)
	db := utils.GetDatabase()
	defer db.Session.Close()
	cs := db.C("stake_role_candidate")

	query := bson.M{}
	query["revoked"] = true

	cs.Find(query).Sort("-voting_power").Skip((page - 1) * size).Limit(size).All(&candidates)

	var result []CandidateVo
	for _, ca := range candidates {
		result = append(result, CandidateVo{
			Address:     ca.Address,
			Name:        ca.Description.Moniker,
			VotingPower: ca.VotingPower,
			Height:      ca.BondHeight, //TODO
		})
	}

	resultByte, _ := json.Marshal(result)
	w.Write(resultByte)
}

type CandidateVo struct {
	Address     string
	Name        string
	VotingPower float64
	Height      int64
}
