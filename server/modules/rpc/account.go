package rpc

import (
	"github.com/gorilla/mux"
	"github.com/cosmos/cosmos-sdk"
	"github.com/cosmos/cosmos-sdk/modules/coin"
	"net/http"
	"github.com/cosmos/cosmos-sdk/client/commands"
	"github.com/cosmos/cosmos-sdk/stack"
	"github.com/cosmos/cosmos-sdk/client/commands/query"
	"github.com/irisnet/irisplorer.io/server/modules/tools"
	"log"
	"time"
	"github.com/irisnet/irisplorer.io/server/modules/store/m"
	"github.com/spf13/cast"
)

func RegisterAccount(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		RegisterQueryAccount,
		RegisterQueryAllAccount,
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
	account := QueryBalance(address, false)
	if err := query.FoutputProof(w, account, 0); err != nil {
		sdk.WriteError(w, err)
	}
}

func queryAccounts(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	page := args["page"]
	p := cast.ToInt(page)
	accounts := m.QueryAccountByPage(p)
	if err := query.FoutputProof(w, accounts, 0); err != nil {
		sdk.WriteError(w, err)
	}
}

func QueryBalance(address string, delay bool) *coin.Account {
	account := new(coin.Account)
	actor, err := commands.ParseActor(address)
	if err != nil {
		return account
	}

	actor = coin.ChainAddr(actor)
	key := stack.PrefixedKey(coin.NameCoin, actor.Bytes())
	if delay {
		time.Sleep(1 * time.Second)
	}
	_, err2 := tools.GetParsed(key, account, query.GetHeight(), false)
	if err2 != nil {
		log.Printf("account bytes are empty for address: %q", address)
	}
	return account
}

// mux.Router registrars

func RegisterQueryAccount(r *mux.Router) error {
	r.HandleFunc("/account/{address}", queryAccount).Methods("GET")
	return nil
}

func RegisterQueryAllAccount(r *mux.Router) error {
	r.HandleFunc("/accounts/{page}", queryAccounts).Methods("GET")
	return nil
}
