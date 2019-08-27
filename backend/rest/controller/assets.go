package controller

import (
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/service"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
)

var assets service.AssetsService

func RegisterAssets(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryNativeAsset,
		registerQueryGatewayAsset,
		registerAssetTokens,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

func registerQueryNativeAsset(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryNativeAsset, "GET", func(request model.IrisReq) interface{} {
		assets.SetTid(request.TraceId)
		page := int(utils.ParseIntWithDefault(QueryParam(request, "page"), 1))
		size := int(utils.ParseIntWithDefault(QueryParam(request, "size"), 5))
		total := QueryParam(request, "total")
		txtype := QueryParam(request, "tx_type")
		istotal := true
		if total == "false" {
			istotal = false
		}
		res, err := assets.GetNativeAsset(txtype, page, size, istotal)
		if err != nil {
			return model.NewResponse("-1", err.Error(), nil)
		}
		return model.NewResponse(types.CodeSuccess.Code, types.CodeSuccess.Msg, res)
	})
	return nil
}

func registerQueryGatewayAsset(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryGatewayAsset, "GET", func(request model.IrisReq) interface{} {
		assets.SetTid(request.TraceId)
		page := int(utils.ParseIntWithDefault(QueryParam(request, "page"), 1))
		size := int(utils.ParseIntWithDefault(QueryParam(request, "size"), 5))
		txtype := QueryParam(request, "tx_type")
		total := QueryParam(request, "total")
		istotal := true
		if total == "false" {
			istotal = false
		}
		res, err := assets.GetGatewayAsset(txtype, page, size, istotal)
		if err != nil {
			return model.NewResponse("-1", err.Error(), nil)
		}
		return model.NewResponse(types.CodeSuccess.Code, types.CodeSuccess.Msg, res)
	})
	return nil
}

func registerAssetTokens(r *mux.Router) error {
	doApi(r, types.UrlRegisterAssetTokens, "GET", func(request model.IrisReq) interface{} {
		assets.SetTid(request.TraceId)
		result := assets.QueryAssetToken()
		return result
	})
	return nil
}
