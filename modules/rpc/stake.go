package rpc

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/cosmos/cosmos-sdk"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/stack"
	"github.com/cosmos/cosmos-sdk/modules/coin"
	"github.com/cosmos/cosmos-sdk/client/commands"
	crypto "github.com/tendermint/go-crypto"

	"github.com/irisnet/iris-explorer/modules/stake"
	"github.com/tendermint/go-wire"
	"github.com/irisnet/iris-explorer/modules/tools"
)

func RegisterStake(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryCandidatesByAccount,
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


func queryCandidatesByAccount(w http.ResponseWriter, r *http.Request) {
	var pks []crypto.PubKey
	var bonds []stake.DelegatorBond
	key := stack.PrefixedKey(stake.Name(), stake.CandidatesPubKeysKey)
	height := int64(viper.GetInt(tools.FlagHeight))
	_, err := tools.GetParsed(key, &pks, height, false)
	if err != nil || len(pks) ==0 {
		sdk.WriteError(w, err)
		return
	}

	args := mux.Vars(r)
	delegatorAddr := args["address"]
	delegator, err := commands.ParseActor(delegatorAddr)
	delegator = coin.ChainAddr(delegator)
	if err != nil {
		sdk.WriteError(w, err)
		return
	}

	for _,pk := range  pks{
		var bond stake.DelegatorBond
		key := stack.PrefixedKey(stake.Name(), getDelegatorBondKey(delegator,pk))
		_, err := tools.GetParsed(key, &bond, height, false)
		if !client.IsNoDataErr(err) && err ==nil {
			bonds = append(bonds,bond)
		}
	}
	tools.FmtOutPutResult(w, bonds)
}

// GetDelegatorBondKey - get the key for delegator bond with candidate
func getDelegatorBondKey(delegator sdk.Actor, candidate crypto.PubKey) []byte {
	return append(getDelegatorBondKeyPrefix(delegator), candidate.Bytes()...)
}

// GetDelegatorBondKeyPrefix - get the prefix for a delegator for all candidates
func getDelegatorBondKeyPrefix(delegator sdk.Actor) []byte {
	return append(stake.DelegatorBondKeyPrefix, wire.BinaryBytes(&delegator)...)
}