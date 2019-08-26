package controller

import (
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/types"
)

func RegisterAsset(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerAssetTokens,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

func registerAssetTokens(r *mux.Router) error {
	doApi(r, types.UrlRegisterAssetTokens, "GET", func(request model.IrisReq) interface{} {
		asset.SetTid(request.TraceId)
		result := asset.QueryAssetToken()
		return result
	})
	return nil
}
