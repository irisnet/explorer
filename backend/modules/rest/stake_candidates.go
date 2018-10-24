package rest

import (
	"encoding/json"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/irishub-sync/store/document"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func queryCandidates(w http.ResponseWriter, r *http.Request) {
	var candidates []document.Candidate
	page, size := utils.TransferPage(r)
	db := utils.GetDatabase()
	defer db.Session.Close()
	cs := db.C("stake_role_candidate")
	txDoc := db.C("tx_common")

	query := bson.M{}
	query["revoked"] = false
	query["status"] = bson.M{
		"$in": []string{types.TypeValStatusUnbonded, types.TypeValStatusUnbonding},
	}

	validatorsCount, _ := cs.Find(query).Count()
	cs.Find(query).Sort("-voting_power").Skip((page - 1) * size).Limit(size).All(&candidates)
	var result []CandidateAll
	var resp Candidates
	if len(candidates) > 0 {
		q := bson.M{}
		q["type"] = types.TypeCreateValidator

		for _, ca := range candidates {
			var tx document.CommonTx
			q["from"] = ca.Address
			txDoc.Find(q).Sort("height").One(tx)

			ca.BondHeight = tx.Height
			result = append(result, CandidateAll{
				Candidate: ca,
			})
		}

		resp = Candidates{
			Count:      validatorsCount,
			Candidates: result,
		}
	}
	resultByte, _ := json.Marshal(resp)
	w.Write(resultByte)

}
