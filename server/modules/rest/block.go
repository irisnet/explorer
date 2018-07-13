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
	if err == nil {
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

func RegisterBlock(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryBlock,
		registerQueryValidators,
		registerQueryRecentBlock,
		registerQueryBlocks,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}
