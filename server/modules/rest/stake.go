package rest

import (
	"github.com/gorilla/mux"
	"net/http"

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
}
