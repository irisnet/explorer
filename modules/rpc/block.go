package rpc

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/cast"
	sdk "github.com/cosmos/cosmos-sdk"
	"github.com/cosmos/cosmos-sdk/client/commands"
	"github.com/irisnet/iris-explorer/modules/tools"
)

func queryBlock(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	height := args["height"]

	node := commands.GetNode()
	h := cast.ToInt64(height)
	block, err := node.Block(&h)
	if err != nil {
		sdk.WriteError(w, err)
		return
	}
	if err := tools.FmtOutPutResult(w, block); err != nil {
		sdk.WriteError(w, err)
	}
}

// mux.Router registrars

func registerQueryBlock(r *mux.Router) error {
	r.HandleFunc("/block/{height}", queryBlock).Methods("GET")
	return nil
}

func RegisterBlock(r *mux.Router) error {
	funcs := []func(*mux.Router) error{
		registerQueryBlock,
	}

	for _, fn := range funcs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}