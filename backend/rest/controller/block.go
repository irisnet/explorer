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
	doApi(r, types.UrlRegisterQueryBlock, "GET", func(request *http.Request) interface{} {
		h := Var(request, "height")
		height, ok := utils.ParseInt(h)
		if !ok {
			panic(types.ErrorCodeInValidParam)
			return nil
		}

		result := service.GetBlock().Query(height)
		return result
	})
	return nil
}

func registerQueryBlocks(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryBlocks, "GET", func(request *http.Request) interface{} {
		page, size := GetPage(request)
		result := service.GetBlock().QueryList(page, size)
		return result
	})

	return nil
}

func registerQueryBlocksPrecommits(r *mux.Router) error {

	doApi(r, types.UrlRegisterQueryBlocksPrecommits, "GET", func(request *http.Request) interface{} {
		address := Var(request, "address")
		page, size := GetPage(request)

		result := service.GetBlock().QueryPrecommits(address, page, size)
		return result
	})

	return nil
}
