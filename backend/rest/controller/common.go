package controller

import (
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/service"
	"github.com/irisnet/explorer/backend/types"
	"time"
)

func RegisterTextSearch(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryText,
		registerQuerySysDate,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

type Common struct {
	*service.CommonService
}

var common = Common{
	service.Get(service.Common).(*service.CommonService),
}

func registerQueryText(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryText, "GET", func(request model.IrisReq) interface{} {
		common.SetTid(request.TraceId)
		text := Var(request, "text")

		result := common.QueryText(text)
		return result
	})

	return nil
}

func registerQuerySysDate(r *mux.Router) error {
	doApi(r, types.UrlRegisterQuerySysDate, "GET", func(request model.IrisReq) interface{} {
		return time.Now().Unix()
	})

	return nil
}
