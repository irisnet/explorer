package controller

import (
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/vo"
)

func RegisterTokenStats(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryTokenStats,
		registerQueryTokensAccountTotal,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

func registerQueryTokenStats(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryTokenStats, "GET", func(request vo.IrisReq) interface{} {
		tokenstats.SetTid(request.TraceId)
		result, err := tokenstats.QueryTokenStats()
		if err != nil {
			return vo.NewResponse("-1", err.Error(), nil)
		}
		return vo.NewResponse(types.CodeSuccess.Code, types.CodeSuccess.Msg, result)
	})

	return nil
}

func registerQueryTokensAccountTotal(r *mux.Router) error {
	doApi(r, types.UrlRegisterTokensAccountTotal, "GET", func(request vo.IrisReq) interface{} {
		tokenstats.SetTid(request.TraceId)
		result, err := tokenstats.QueryTokensAccountTotal()
		if err != nil {
			return vo.NewResponse("-1", err.Error(), nil)
		}
		return vo.NewResponse(types.CodeSuccess.Code, types.CodeSuccess.Msg, result)
	})

	return nil
}
