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

	cs.Find(query).Sort("-voting_power").Skip((page - 1) * size).Limit(size).All(&candidates)
	var result []CandidateVo
	if len(candidates) > 0 {
		q := bson.M{}
		q["type"] = types.TypeCreateValidator

		for _, ca := range candidates {
			var tx document.CommonTx
			q["from"] = ca.Address
			txDoc.Find(q).Sort("height").One(tx)

			result = append(result, CandidateVo{
				Address:     ca.Address,
				Name:        ca.Description.Moniker,
				VotingPower: ca.VotingPower,
				Height:      tx.Height,
			})
		}
	}
	resultByte, _ := json.Marshal(result)
	w.Write(resultByte)

}
