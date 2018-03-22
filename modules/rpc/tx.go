package rpc


import (
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"encoding/hex"
	"github.com/cosmos/cosmos-sdk"
	"net/http"
	"github.com/cosmos/cosmos-sdk/client/commands"
	"github.com/tendermint/tmlibs/common"
	"github.com/irisnet/iris-explorer/modules/tools"
	"github.com/irisnet/iris-explorer/modules/store"
	"fmt"
)


func RegisterQueryTx(r *mux.Router) error {
	r.HandleFunc("/tx/{hash}", queryTx).Methods("GET")
	return nil
}

func RegisterQueryCoinTxByAccount(r *mux.Router) error {
	r.HandleFunc("/tx/coin/{address}", queryCoinTxByAccount).Methods("GET")
	return nil
}

func RegisterQueryStakeTxByAccount(r *mux.Router) error {
	r.HandleFunc("/tx/stake/{address}", queryStakeTxByAccount).Methods("GET")
	return nil
}

// queryTx is to query transaction by hash
func queryTx(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	hash := args["hash"]

	prove := !viper.GetBool(commands.FlagTrustNode)
	key, err := hex.DecodeString(common.StripHex(hash))
	if err != nil {
		sdk.WriteError(w, err)
		return
	}

	node := tools.GetNode()
	defer node.Release()

	res, err := node.Client.Tx(key, prove)
	if err != nil {
		sdk.WriteError(w, err)
		return
	}

	wrap, err := tools.BuildTxResp(res.Height, res.Proof.Data, false, hash)
	if err != nil {
		sdk.WriteError(w, err)
		return
	}

	tools.FmtOutPutResult(w, wrap)
}

func queryCoinTxByAccount(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	account := args["address"]
	actor, err := commands.ParseActor(account)
	if err != nil {
		sdk.WriteError(w, err)
		return
	}

	findSender := fmt.Sprintf("coin.sender='%s'", actor)
	findReceiver := fmt.Sprintf("coin.receiver='%s'", actor)

	wrap, err := tools.SearchTx(w, findSender, findReceiver)
	if err != nil {
		sdk.WriteError(w, err)
	}

	tools.FmtOutPutResult(w, wrap)
}

func queryStakeTxByAccount(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	account := args["address"]
	result := store.Mgo.QueryStakeTxsByAccount(account)
	tools.FmtOutPutResult(w, result)
}

func RegisterTx(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		RegisterQueryTx,
		RegisterQueryCoinTxByAccount,
		RegisterQueryStakeTxByAccount,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}