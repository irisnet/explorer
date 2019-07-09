package controller

import (
	"strconv"

	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
)

// mux.Router registrars
func RegisterBlock(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryBlocks,
		registerQueryRecentBlocks,
		registerQueryValidatorSet,
		registerQueryBlockInfoByBlock,
		registerQueryBlockLatestHeight,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

//type Block struct {
//	*service.BlockService
//}
//
//var block = Block{
//	service.Get(service.Block).(*service.BlockService),
//}

func registerQueryBlockLatestHeight(r *mux.Router) error {

	doApi(r, types.UrlRegisterQueryBlockLatestHeight, "GET", func(request model.IrisReq) interface{} {
		var block = lcd.BlockLatest()
		var height, ok = utils.ParseInt(block.BlockMeta.Header.Height)
		if !ok {
			panic(types.CodeNotFound)
		}
		return height
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

func registerQueryValidatorSet(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryBlockValidatorSet, "GET", func(request model.IrisReq) interface{} {
		block.SetTid(request.TraceId)
		height, err := strconv.ParseInt(Var(request, "height"), 10, 0)
		if err != nil || height < 1 {
			panic(types.CodeInValidParam)
		}

		page := int(utils.ParseIntWithDefault(QueryParam(request, "page"), DefaultPageNum))
		size := int(utils.ParseIntWithDefault(QueryParam(request, "size"), DefaultPageSize))
		result := block.GetValidatorSet(height, page, size)

		return result
	})
	return nil
}

func registerQueryBlockInfoByBlock(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryBlockInfo, "GET", func(request model.IrisReq) interface{} {
		tx.SetTid(request.TraceId)
		height, err := strconv.ParseInt(Var(request, "height"), 10, 0)
		if err != nil || height < 1 {
			panic(types.CodeInValidParam)
		}
		return block.QueryBlockInfo(height)
	})
	return nil
}
