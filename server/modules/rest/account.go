package rest

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
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

	accByte,_ := json.Marshal(address)
	w.Write(accByte)
}

func queryAccounts(w http.ResponseWriter, r *http.Request) {
}


// mux.Router registrars

func RegisterQueryAccount(r *mux.Router) error {
	r.HandleFunc("/account/{address}", queryAccount).Methods("GET")
	return nil
}

func RegisterQueryAllAccount(r *mux.Router) error {
	r.HandleFunc("/accounts/{page}/{size}", queryAccounts).Methods("GET")
	return nil
}
