package controller

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/types"
)

func RegisterNodes(r *mux.Router) error {
	apis := []func(*mux.Router) error{
		RegisterQueryFaucet,
		RegisterApply,
	}

	for _, fn := range apis {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

func RegisterQueryFaucet(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryFaucet, "GET", func(request model.IrisReq) interface{} {
		res, err := lcd.Faucet(request.Request)
		if err != nil {
			panic(err)
		}
		return res
	})
	return nil
}

func RegisterApply(r *mux.Router) error {
	doApi(r, types.UrlRegisterApply, "POST", func(request model.IrisReq) interface{} {
		res, err := lcd.GetToken(request.Request)
		if err != nil {
			panic(errors.New("draw iris fail " + err.Error()))
		}
		return res
	})
	return nil
}
