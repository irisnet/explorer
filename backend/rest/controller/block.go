package controller

import (
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/service"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
)

// mux.Router registrars
func RegisterBlock(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryBlock,
		registerQueryBlocks,
		registerQueryRecentBlocks,
		registerQueryBlocksPrecommits,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

type Block struct {
	*service.BlockService
}

var block = Block{
	service.Get(service.Block).(*service.BlockService),
}

func registerQueryBlock(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryBlock, "GET", func(request model.IrisReq) interface{} {
		h := Var(request, "height")
		height, ok := utils.ParseInt(h)
		if !ok {
			panic(types.CodeInValidParam)
			return nil
		}
		block.SetTid(request.TraceId)
		result := block.Query(height)
		return result
	})
	return nil
}

func registerQueryBlocks(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryBlocks, "GET", func(request model.IrisReq) interface{} {
		block.SetTid(request.TraceId)
		page := int(utils.ParseIntWithDefault(QueryParam(request, "page"), 1))
		size := int(utils.ParseIntWithDefault(QueryParam(request, "size"), 100))
		result := block.QueryList(page, size)
		return result
	})

	return nil
}

func registerQueryRecentBlocks(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryRecentBlocks, "GET", func(request model.IrisReq) interface{} {
		return block.QueryRecent()
	})

	return nil
}

func registerQueryBlocksPrecommits(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryBlocksPrecommits, "GET", func(request model.IrisReq) interface{} {
		block.SetTid(request.TraceId)

		address := Var(request, "address")
		page, size := GetPage(request)

		result := block.QueryPrecommits(address, page, size)
		return result
	})

	return nil
}
