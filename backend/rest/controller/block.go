package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/service"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"net/http"
)

// mux.Router registrars
func RegisterBlock(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryBlock,
		registerQueryBlocks,
		registerQueryBlocksPrecommits,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

func registerQueryBlock(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryBlock, func(writer http.ResponseWriter, request *http.Request) {

		h := Var(request, "height")
		height, ok := utils.ParseInt(h)
		if !ok {
			err := types.ErrorCodeInValidParam
			err.Msg = fmt.Sprintf("%s must be int type,but value is %s", "height", h)
			WriteResonse(writer, err)
			return
		}

		result := service.GetBlock().Query(height)
		WriteResonse(writer, result)

	}).Methods("GET")
	return nil
}

func registerQueryBlocks(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryBlocks, func(writer http.ResponseWriter, request *http.Request) {
		page, size := GetPage(request)
		result := service.GetBlock().QueryList(page, size)
		WriteResonse(writer, result)

	}).Methods("GET")
	return nil
}

func registerQueryBlocksPrecommits(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryBlocksPrecommits, func(writer http.ResponseWriter, request *http.Request) {
		address := Var(request, "address")
		page, size := GetPage(request)

		result := service.GetBlock().QueryPrecommits(address, page, size)
		WriteResonse(writer, result)
	}).Methods("GET")
	return nil
}
