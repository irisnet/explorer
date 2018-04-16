package rest

import (
	"encoding/hex"
	"github.com/cosmos/cosmos-sdk"
	"github.com/gorilla/mux"
	"github.com/irisnet/irisplorer.io/server/modules/store"
	"github.com/irisnet/irisplorer.io/server/modules/store/m"
	"github.com/irisnet/irisplorer.io/server/modules/tools"
	"github.com/spf13/cast"
	"github.com/tendermint/tmlibs/common"
	"net/http"
)

func registerQueryTx(r *mux.Router) error {
	r.HandleFunc("/tx/{hash}", queryTx).Methods("GET")
	return nil
}

func registerQueryAllCoinTxByPage(r *mux.Router) error {
	r.HandleFunc("/txs/coin/{page}/{size}", queryAllCoinTxByPage).Methods("GET")
	return nil
}

func registerQueryCoinTxByAccount(r *mux.Router) error {
	r.HandleFunc("/tx/coin/{address}", queryCoinTxByAccount).Methods("GET")
	return nil
}

func registerQueryPageCoinTxByAccount(r *mux.Router) error {
	r.HandleFunc("/tx/coin/{address}/{page}/{size}", queryCoinPageTxByAccount).Methods("GET")
	return nil
}

func registerQueryAllStakeTxByPage(r *mux.Router) error {
	r.HandleFunc("/txs/stake/{page}/{size}", queryAllStakeTxByPage).Methods("GET")
	return nil
}

func registerQueryStakeTxByAccount(r *mux.Router) error {
	r.HandleFunc("/tx/stake/{address}", queryStakeTxByAccount).Methods("GET")
	return nil
}

func registerQueryPageStakeTxByAccount(r *mux.Router) error {
	r.HandleFunc("/tx/stake/{address}/{page}/{size}", queryPageStakeTxByAccount).Methods("GET")
	return nil
}

func registerQueryTxs(r *mux.Router) error {
	r.HandleFunc("/txs/{txType}", queryTxs).Methods("GET")
	return nil
}

func queryTxs(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	if tools.ValidateReq(w, r) != nil {
		return
	}

	//过滤条件
	query, collectionNm := buildQuery(r)

	//分页条件
	page := r.Form["page"][0]
	size := r.Form["size"][0]

	p := cast.ToInt(page)
	s := cast.ToInt(size)

	//排序条件
	var sort string
	if orderBy := r.Form["orderBy"]; len(orderBy) > 0 {
		sort = orderBy[0]
	}

	result, _ := store.SelectByPage(collectionNm, query, sort, p, s)
	tools.FmtOutPutResult(w, result)
}

func buildQuery(r *http.Request) (query *store.Query, collectionNm string) {
	//过滤条件
	address := r.Form["address"]
	if len(address) > 0 {
		value := address[0]
		vector := r.Form["type"]

		args := mux.Vars(r)
		txType := args["txType"]

		if txType == "coin" {
			if len(vector) == 0 || len(vector[0]) == 0{
				query = store.NewQuery().Or(store.M{{"from": value}, {"to": value}})
			} else {
				query = store.CreateQuery(vector[0], value)
			}
			collectionNm = m.DocsNmCoinTx
		} else {
			if len(vector) == 0 || len(vector[0]) == 0{
				query = store.CreateQuery("from", value)
			} else {
				query = store.NewQuery().And(store.M{{"from": value}, {"type": vector[0]}})
			}
			collectionNm = m.DocsNmStakeTx
		}
	}
	return query, collectionNm
}

// queryTx is to query transaction by hash
func queryTx(w http.ResponseWriter, r *http.Request) {
	if tools.ValidateReq(w, r) != nil {
		return
	}

	args := mux.Vars(r)
	hash := args["hash"]

	key, err := hex.DecodeString(common.StripHex(hash))
	if err != nil {
		sdk.WriteError(w, err)
		return
	}

	node := tools.GetNode()
	defer node.Release()

	res, err := node.Client.Tx(key, true)
	if err != nil {
		sdk.WriteError(w, err)
		return
	}

	wrap, err := tools.BuildTxResp(res.Height, res.Proof.Data, true, hash)
	if err != nil {
		sdk.WriteError(w, err)
		return
	}

	tools.FmtOutPutResult(w, wrap)
}

func queryAllCoinTxByPage(w http.ResponseWriter, r *http.Request) {
	if tools.ValidateReq(w, r) != nil {
		return
	}
	args := mux.Vars(r)
	page := args["page"]
	size := args["size"]
	p := cast.ToInt(page)
	s := cast.ToInt(size)
	result := m.QueryAllPageCoinTxs(p, s)
	tools.FmtOutPutResult(w, result)
}

func queryCoinTxByAccount(w http.ResponseWriter, r *http.Request) {
	if tools.ValidateReq(w, r) != nil {
		return
	}
	args := mux.Vars(r)
	account := args["address"]
	result := m.QueryCoinTxsByAccount(account)
	tools.FmtOutPutResult(w, result)
}

func queryCoinPageTxByAccount(w http.ResponseWriter, r *http.Request) {
	if tools.ValidateReq(w, r) != nil {
		return
	}
	args := mux.Vars(r)
	account := args["address"]
	page := args["page"]
	size := args["size"]
	p := cast.ToInt(page)
	s := cast.ToInt(size)
	result := m.QueryPageCoinTxsByAccount(account, p, s)
	tools.FmtOutPutResult(w, result)
}

func queryAllStakeTxByPage(w http.ResponseWriter, r *http.Request) {
	if tools.ValidateReq(w, r) != nil {
		return
	}
	args := mux.Vars(r)
	page := args["page"]
	size := args["size"]
	p := cast.ToInt(page)
	s := cast.ToInt(size)
	result := m.QueryPageStakeTxs(p, s)
	tools.FmtOutPutResult(w, result)
}

func queryStakeTxByAccount(w http.ResponseWriter, r *http.Request) {
	if tools.ValidateReq(w, r) != nil {
		return
	}
	args := mux.Vars(r)
	account := args["address"]

	result := m.QueryStakeTxsByAccount(account)
	tools.FmtOutPutResult(w, result)
}

func queryPageStakeTxByAccount(w http.ResponseWriter, r *http.Request) {
	if tools.ValidateReq(w, r) != nil {
		return
	}
	args := mux.Vars(r)
	account := args["address"]
	page := args["page"]
	size := args["size"]
	p := cast.ToInt(page)
	s := cast.ToInt(size)
	result := m.QueryPageStakeTxsByAccount(account, p, s)
	tools.FmtOutPutResult(w, result)
}

func RegisterTx(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryTx,
		registerQueryTxs,
		registerQueryCoinTxByAccount,
		registerQueryPageCoinTxByAccount,
		registerQueryAllCoinTxByPage,
		registerQueryStakeTxByAccount,
		registerQueryPageStakeTxByAccount,
		registerQueryAllStakeTxByPage,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}
