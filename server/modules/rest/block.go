package rest

import (
	"net/http"
	"github.com/gorilla/mux"
)

func queryBlock(w http.ResponseWriter, r *http.Request) {

}

func queryValidators(w http.ResponseWriter, r *http.Request) {

}

func queryRecentBlock(w http.ResponseWriter, r *http.Request) {
}

func queryBlocks(w http.ResponseWriter, r *http.Request) {
}

// mux.Router registrars

func registerQueryBlock(r *mux.Router) error {
	r.HandleFunc("/block/{height}", queryBlock).Methods("GET")
	return nil
}

func registerQueryValidators(r *mux.Router) error {
	r.HandleFunc("/validators/{height}", queryValidators).Methods("GET")
	return nil
}

func registerQueryRecentBlock(r *mux.Router) error {
	r.HandleFunc("/blocks/recent", queryRecentBlock).Methods("GET")
	return nil
}

func registerQueryBlocks(r *mux.Router) error {
	r.HandleFunc("/blocks/{page}/{size}", queryBlocks).Methods("GET")
	return nil
}

func RegisterBlock(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryBlock,
		registerQueryValidators,
		registerQueryRecentBlock,
		registerQueryBlocks,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}
