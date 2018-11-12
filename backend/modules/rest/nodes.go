package rest

import (
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"net/http"
)

func RegisterNodes(r *mux.Router) error {
	apis := []func(*mux.Router) error{
		RegisterQueryNodes,
	}

	for _, fn := range apis {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

func RegisterQueryNodes(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryNodes, func(w http.ResponseWriter, request *http.Request) {
		bz := utils.GetNodes()
		w.Write(bz)
	}).Methods("GET")
	return nil
}
