package controller

import (
	"strconv"

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
		registerQueryValidatorSet,
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
		block.SetTid(request.TraceId)

		height, err := strconv.ParseInt(Var(request, "height"), 10, 0)
		if err != nil || height < 1 {
			panic(types.CodeInValidParam)
		}
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

func registerQueryValidatorSet(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryValidatorSet, "GET", func(request model.IrisReq) interface{} {
		block.SetTid(request.TraceId)
		page := int(utils.ParseIntWithDefault(QueryParam(request, "page"), DefaultPageNum))
		size := int(utils.ParseIntWithDefault(QueryParam(request, "size"), DefaultPageSize))
		height := utils.ParseIntWithDefault(QueryParam(request, "height"), DefaultBlockHeight)
		if height < 1 {
			panic(types.CodeInValidParam)
		}
		result := block.GetValidatorSet(height, page, size)
		return result
	})
	return nil
}
