package controller

import (
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/vo"
)

func RegisterQueryVersion(r *mux.Router) error {

	doApi(r, types.UrlRegisterQueryApiVersion, "GET", func(request vo.IrisReq) interface{} {
		return map[string]string{"version": conf.Get().Server.ApiVersion}
	})

	return nil
}
