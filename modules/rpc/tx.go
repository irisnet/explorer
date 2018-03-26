package rpc


import (
	"github.com/gorilla/mux"
	"encoding/hex"
	"github.com/cosmos/cosmos-sdk"
	"net/http"
	"github.com/tendermint/tmlibs/common"
	"github.com/irisnet/iris-explorer/modules/tools"
	"github.com/irisnet/iris-explorer/modules/store"
)


func registerQueryTx(r *mux.Router) error {
	r.HandleFunc("/tx/{hash}", queryTx).Methods("GET")
	return nil
}

func registerQueryCoinTxByAccount(r *mux.Router) error {
	r.HandleFunc("/tx/coin/{address}", queryCoinTxByAccount).Methods("GET")
	return nil
}

func registerQueryStakeTxByAccount(r *mux.Router) error {
	r.HandleFunc("/tx/stake/{address}", queryStakeTxByAccount).Methods("GET")
	return nil
}

// queryTx is to query transaction by hash
func queryTx(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	hash := args["hash"]

	key, err := hex.DecodeString(common.StripHex(hash))
	if err != nil {
		sdk.WriteError(w, err)
		return
	}

	node := tools.GetNode()
	defer node.Release()

	res, err := node.Client.Tx(key, false)
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
	result := store.QueryCoinTxsByAccount(account)
	tools.FmtOutPutResult(w, result)
}

func queryStakeTxByAccount(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	account := args["address"]
	result := store.QueryStakeTxsByAccount(account)
	tools.FmtOutPutResult(w, result)
}

func RegisterTx(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryTx,
		registerQueryCoinTxByAccount,
		registerQueryStakeTxByAccount,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}