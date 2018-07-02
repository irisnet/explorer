package rest

import (
	"github.com/gorilla/mux"
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
}

// queryTx is to query transaction by hash
func queryTx(w http.ResponseWriter, r *http.Request) {
}

func queryAllCoinTxByPage(w http.ResponseWriter, r *http.Request) {
}

func queryCoinTxByAccount(w http.ResponseWriter, r *http.Request) {
}

func queryCoinPageTxByAccount(w http.ResponseWriter, r *http.Request) {
}

func queryAllStakeTxByPage(w http.ResponseWriter, r *http.Request) {
}

func queryStakeTxByAccount(w http.ResponseWriter, r *http.Request) {
}

func queryPageStakeTxByAccount(w http.ResponseWriter, r *http.Request) {
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
