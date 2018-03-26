package rpc

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/cast"
	sdk "github.com/cosmos/cosmos-sdk"
	"github.com/irisnet/iris-explorer/modules/tools"
)

func queryBlock(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	height := args["height"]
	node := tools.GetNode()
	defer node.Release()
	h := cast.ToInt64(height)
	block, err := node.Client.Block(&h)
	if err != nil {
		sdk.WriteError(w, err)
		return
	}
	if err := tools.FmtOutPutResult(w, block); err != nil {
		sdk.WriteError(w, err)
	}
}

func queryValidators(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	height := args["height"]

	node := tools.GetNode()
	defer node.Release()

	h := cast.ToInt64(height)
	block, err := node.Client.Validators(&h)
	if err != nil {
		sdk.WriteError(w, err)
		return
	}
	if err := tools.FmtOutPutResult(w, block); err != nil {
		sdk.WriteError(w, err)
	}
}

func queryRecentBlocks(w http.ResponseWriter, r *http.Request) {
	node := tools.GetNode()
	defer node.Release()

	blocks, err := node.Client.BlockchainInfo(0,0)
	if err != nil {
		sdk.WriteError(w, err)
		return
	}
	if err := tools.FmtOutPutResult(w, blocks); err != nil {
		sdk.WriteError(w, err)
	}
}

// mux.Router registrars

func registerQueryBlock(r *mux.Router) error {
	r.HandleFunc("/block/{height}", queryBlock).Methods("GET")
	return nil
}

func registerQueryValidators(r *mux.Router) error {
	r.HandleFunc("/validators/{height}", queryValidators).Methods("GET")
	return nil
}

func registerQueryRecentBlocks(r *mux.Router) error {
	r.HandleFunc("/blocks/recent", queryRecentBlocks).Methods("GET")
	return nil
}

func RegisterBlock(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryBlock,
		registerQueryValidators,
		registerQueryRecentBlocks,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}