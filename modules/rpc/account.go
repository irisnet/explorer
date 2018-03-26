package rpc

import (
	"github.com/gorilla/mux"
	"github.com/cosmos/cosmos-sdk"
	"github.com/cosmos/cosmos-sdk/modules/coin"
	"github.com/spf13/viper"
	"fmt"
	"net/http"
	"github.com/cosmos/cosmos-sdk/client/commands"
	"github.com/cosmos/cosmos-sdk/stack"
	"github.com/cosmos/cosmos-sdk/client/commands/query"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/irisnet/iris-explorer/modules/tools"
)

func RegisterAccount(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		RegisterQueryAccount,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

func queryAccount(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	address := args["address"]
	actor, err := commands.ParseActor(address)
	if err != nil {
		sdk.WriteError(w, err)
		return
	}

	actor = coin.ChainAddr(actor)
	key := stack.PrefixedKey(coin.NameCoin, actor.Bytes())
	account := new(coin.Account)
	prove := !viper.GetBool(commands.FlagTrustNode)
	height, err := tools.GetParsed(key, account, query.GetHeight(), prove)
	if client.IsNoDataErr(err) {
		err := fmt.Errorf("account bytes are empty for address: %q", address)
		sdk.WriteError(w, err)
		return
	} else if err != nil {
		sdk.WriteError(w, err)
		return
	}

	if err := query.FoutputProof(w, account, height); err != nil {
		sdk.WriteError(w, err)
	}
}

// mux.Router registrars

func RegisterQueryAccount(r *mux.Router) error {
	r.HandleFunc("/account/{address}", queryAccount).Methods("GET")
	return nil
}
