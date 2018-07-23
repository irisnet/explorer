package rest

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/irisnet/explorer/server/utils"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
)

type count struct {
	Id    bson.ObjectId `bson:"_id,omitempty"`
	Count int
}

type chainStatus struct {
	ValidatorsCount int
	TxCount         int
	VotingPower     int
}

func RegisterChain(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		RegisterQueryChain,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

func RegisterQueryChain(r *mux.Router) error {
	r.HandleFunc("/api/chain/status", queryChainStatus).Methods("GET")
	return nil
}

func queryChainStatus(w http.ResponseWriter, r *http.Request) {
	db := utils.GetDatabase()
	cs := db.C("stake_role_candidate")
	defer db.Session.Close()
	pipe := cs.Pipe(
		[]bson.M{
			{"$group": bson.M{
				"_id":   "voting_power",
				"count": bson.M{"$sum": "$voting_power"},
			}},
		},
	)
	var count count
	pipe.One(&count)

	validatorsCount, _ := cs.Count()
	cc := db.C("tx_common")
	txCount, _ := cc.Count()

	resp := chainStatus{
		ValidatorsCount: validatorsCount,
		TxCount:         txCount,
		VotingPower:     count.Count,
	}
	resultByte, _ := json.Marshal(resp)
	w.Write(resultByte)
}
