package rest

import (
	"github.com/gorilla/mux"
	"net/http"

	"github.com/irisnet/explorer/server/utils"
	"github.com/irisnet/irishub-sync/model/store/document"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
)

func RegisterStake(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryCandidatesByAccount,
		registerQueryCandidates,
		registerQueryCandidate,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

func registerQueryCandidatesByAccount(r *mux.Router) error {
	r.HandleFunc("/stake/candidates/{address}", queryCandidatesByAccount).Methods("GET")
	return nil
}

func registerQueryCandidates(r *mux.Router) error {
	r.HandleFunc("/stake/candidates/{page}/{size}", queryCandidates).Methods("GET")
	return nil
}

func registerQueryCandidate(r *mux.Router) error {
	r.HandleFunc("/stake/candidate/{address}", queryCandidate).Methods("GET")
	return nil
}

func queryCandidatesByAccount(w http.ResponseWriter, r *http.Request) {

}

func queryCandidate(w http.ResponseWriter, r *http.Request)  {
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

func queryCandidates(w http.ResponseWriter, r *http.Request) {
	var data []document.Candidate
	w.Write(utils.QueryList("stake_role_candidate", &data, nil, "-update_time", r))
}
