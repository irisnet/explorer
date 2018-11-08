package rest

import (
	"encoding/json"
	"github.com/irisnet/explorer/backend/types"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/irishub-sync/store/document"
	"gopkg.in/mgo.v2/bson"
)

type BlockRsp struct {
	CandidateMap map[string]string
	Counter      Counter
	document.Block
}

type Counter struct {
	TxCnt    int
	PropoCnt int
}

func queryBlock(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	height := args["height"]

	dbm := utils.GetDatabase()
	defer dbm.Session.Close()
	var iHeight int
	if iHeight, _ = strconv.Atoi(height); iHeight < 0 {
		iHeight = 0
	}
	var block document.Block
	var result BlockRsp
	err := dbm.C("block").Find(bson.M{"height": iHeight}).Sort("-time").One(&block)
	var pres []string
	if err == nil {
		for _, pre := range block.Block.LastCommit.Precommits {
			pres = append(pres, pre.ValidatorAddress)
		}
		if len(pres) > 0 {
			var candidates []document.Candidate
			//ca := utils.GetDatabase().C("stake_role_candidate")
			//defer ca.Database.Session.Close()
			err = dbm.C("stake_role_candidate").Find(bson.M{"pub_key_addr": bson.M{"$in": pres}}).All(&candidates)
			candidateMap := make(map[string]string)
			for _, candidate := range candidates {
				candidateMap[candidate.PubKeyAddr] = candidate.Address
			}
			result.CandidateMap = candidateMap
		}
		result.Block = block

		//query txs from block height
		//txCommonC := utils.GetDatabase().C("tx_common")
		//defer txCommonC.Database.Session.Clone()
		var txs []document.CommonTx
		err = dbm.C("tx_common").Find(bson.M{"height": iHeight}).All(&txs)
		if err == nil {
			var counter Counter
			for _, tx := range txs {
				if tx.Type == "SubmitProposal" || tx.Type == "Deposit" || tx.Type == "Vote" {
					counter.PropoCnt = counter.PropoCnt + 1
				} else {
					counter.TxCnt = counter.TxCnt + 1
				}

			}
			result.Counter = counter
		}

		resultByte, _ := json.Marshal(result)
		w.Write(resultByte)
	}
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
	w.Write(utils.QueryList("block", &data, bson.M{"block.last_commit.precommits": bson.M{"$elemMatch": bson.M{"validator_address": candidate.PubKeyAddr}}}, "-height", r))
}

// mux.Router registrars

func registerQueryBlock(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryBlock, queryBlock).Methods("GET")
	return nil
}

func registerQueryBlocks(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryBlocks, queryBlocks).Methods("GET")
	return nil
}

func registerQueryBlocksPrecommits(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryBlocksPrecommits, QueryBlocksPrecommits).Methods("GET")
	return nil
}

func RegisterBlock(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryBlock,
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
