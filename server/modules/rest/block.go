package rest

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/server/utils"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"github.com/irisnet/irishub-sync/store/document"
	"strconv"
)

func queryBlock(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	height := args["height"]

	c := utils.GetDatabase().C("block")
	defer c.Database.Session.Close()
	var iHeight int
	if iHeight, _ = strconv.Atoi(height); iHeight < 0 {
		iHeight = 0
	}
	var result document.Block
	err := c.Find(bson.M{"height": iHeight}).Sort("-time").One(&result)
	var pres []string
	if err == nil {
		for _, pre := range result.Block.LastCommit.Precommits {
			pres = append(pres, pre.ValidatorAddress)
		}
		if len(pres) > 0 {
			var candidates []document.Candidate
			ca := utils.GetDatabase().C("stake_role_candidate")
			defer ca.Database.Session.Close()
			err = ca.Find(bson.M{"pub_key_addr": bson.M{"$in": pres}}).All(&candidates)
			candidateMap := make(map[string]string)
			for _, candidate := range candidates {
				candidateMap[candidate.PubKeyAddr] = candidate.Address
			}
			for index, pre := range result.Block.LastCommit.Precommits {
				result.Block.LastCommit.Precommits[index].ValidatorAddress = candidateMap[pre.ValidatorAddress]
			}
		}
		resultByte, _ := json.Marshal(result)
		w.Write(resultByte)
	}
}

func queryValidators(w http.ResponseWriter, r *http.Request) {

}

func queryRecentBlock(w http.ResponseWriter, r *http.Request) {
}

func queryBlocks(w http.ResponseWriter, r *http.Request) {
	var data []document.Block
	w.Write(utils.QueryList("block", &data, nil, "-height", r))
}

func QueryBlocksPrecommits(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	address := args["address"]
	c := utils.GetDatabase().C("stake_role_candidate")
	defer c.Database.Session.Close()
	var candidate document.Candidate
	c.Find(bson.M{"address": address}).Sort("-bond_height").One(&candidate)
	if candidate.PubKeyAddr == "" {
		return
	}
	var data []document.Block
	w.Write(utils.QueryList("block", &data, bson.M{"block.last_commit.precommits":
	bson.M{"$elemMatch": bson.M{"validator_address": candidate.PubKeyAddr}}}, "-height", r))
}

// mux.Router registrars

func registerQueryBlock(r *mux.Router) error {
	r.HandleFunc("/api/block/{height}", queryBlock).Methods("GET")
	return nil
}

func registerQueryValidators(r *mux.Router) error {
	r.HandleFunc("/api/validators/{height}", queryValidators).Methods("GET")
	return nil
}

func registerQueryRecentBlock(r *mux.Router) error {
	r.HandleFunc("/api/blocks/recent", queryRecentBlock).Methods("GET")
	return nil
}

func registerQueryBlocks(r *mux.Router) error {
	r.HandleFunc("/api/blocks/{page}/{size}", queryBlocks).Methods("GET")
	return nil
}

func registerQueryBlocksPrecommits(r *mux.Router) error {
	r.HandleFunc("/api/blocks/precommits/{address}/{page}/{size}", QueryBlocksPrecommits).Methods("GET")
	return nil
}

func RegisterBlock(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryBlock,
		registerQueryValidators,
		registerQueryRecentBlock,
		registerQueryBlocks,
		registerQueryBlocksPrecommits,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}
