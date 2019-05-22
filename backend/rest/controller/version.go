package controller

import (
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/types"
)

func RegisterQueryVersion(r *mux.Router) error {

	doApi(r, types.UrlRegisterQueryApiVersion, "GET", func(request model.IrisReq) interface{} {
		return map[string]string{"version": conf.Get().Server.ApiVersion}
	})

	return nil
}
