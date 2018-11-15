package controller

import (
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
	RegisterApi(r, types.UrlRegisterQueryBlock, "GET", func(writer http.ResponseWriter, request *http.Request) {
		h := Var(request, "height")
		height, ok := utils.ParseInt(h)
		if !ok {
			panic(types.ErrorCodeInValidParam)
			return
		}

		result := service.GetBlock().Query(height)
		WriteResponse(writer, result)
	})
	return nil
}

func registerQueryBlocks(r *mux.Router) error {
	RegisterApi(r, types.UrlRegisterQueryBlocks, "GET", func(writer http.ResponseWriter, request *http.Request) {
		page, size := GetPage(request)
		result := service.GetBlock().QueryList(page, size)
		WriteResponse(writer, result)
	})

	return nil
}

func registerQueryBlocksPrecommits(r *mux.Router) error {

	RegisterApi(r, types.UrlRegisterQueryBlocksPrecommits, "GET", func(writer http.ResponseWriter, request *http.Request) {
		address := Var(request, "address")
		page, size := GetPage(request)

		result := service.GetBlock().QueryPrecommits(address, page, size)
		WriteResponse(writer, result)
	})

	return nil
}
