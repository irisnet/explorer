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
		registerQueryBaseDenom,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

// @Summary tokenstats
// @Description tokenstats
// @Tags tokenstats
// @Accept  json
// @Produce  json
// @Success 200 {object} vo.Response  "success"
// @Router /api/tokenstats [get]
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

// @Summary account_total
// @Description get tokenstats account_total
// @Tags tokenstats
// @Accept  json
// @Produce  json
// @Success 200 {object} vo.Response  "success"
// @Router /api/tokenstats/account_total [get]
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

// @Summary basedenom
// @Description get base denom
// @Tags tokenstats
// @Accept  json
// @Produce  json
// @Success 200 {object} vo.Response  "success"
// @Router /api/basedenom [get]
func registerQueryBaseDenom(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryBaseDenom, "GET", func(request vo.IrisReq) interface{} {
		tokenstats.SetTid(request.TraceId)
		result, err := tokenstats.QueryBaseDenom()
		if err != nil {
			return vo.NewResponse("-1", err.Error(), nil)
		}
		return vo.NewResponse(types.CodeSuccess.Code, types.CodeSuccess.Msg, result)
	})

	return nil
}
