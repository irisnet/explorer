package rest

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/irisnet/explorer/server/utils"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"time"
)

type count struct {
	Id    bson.ObjectId `bson:"_id,omitempty"`
	Count float64
}

type chainStatus struct {
	ValidatorsCount int
	TxCount         int
	VotingPower     float64
	Tps             float64
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
	defer db.Session.Close()
	cs := db.C("stake_role_candidate")
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

	validatorsCount, _ := cs.Find(bson.M{"revoked": false}).Count()
	cc := db.C("tx_common")
	txCount, _ := cc.Count()

	t := time.Now().Add(-1 * time.Minute)
	txs, _ := cc.Find(bson.M{"time": bson.M{"$gte": t}}).Count()

	resp := chainStatus{
		ValidatorsCount: validatorsCount,
		TxCount:         txCount,
		VotingPower:     count.Count,
		Tps:             float64(txs) / 60,
	}
	resultByte, _ := json.Marshal(resp)
	w.Write(resultByte)
}
