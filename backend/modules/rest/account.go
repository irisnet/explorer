package rest

import (
	"encoding/json"
	"github.com/irisnet/explorer/backend/service"
	"github.com/irisnet/explorer/backend/types"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/utils"
)

// mux.Router registrars
func RegisterAccount(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		RegisterQueryAccount,
		RegisterQueryAllAccount,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

func RegisterQueryAccount(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryAccount, func(writer http.ResponseWriter, request *http.Request) {
		args := mux.Vars(request)
		address := args["address"]

		result := service.GetAccount().Query(address)
		resultByte, _ := json.Marshal(result)
		writer.Write(resultByte)

	}).Methods("GET")
	return nil
}

func RegisterQueryAllAccount(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryAllAccount, func(writer http.ResponseWriter, request *http.Request) {
		page, size := utils.GetPage(request)

		result := service.GetAccount().QueryList(page, size)
		resultByte, _ := json.Marshal(result)
		writer.Write(resultByte)

	}).Methods("GET")
	return nil
}
