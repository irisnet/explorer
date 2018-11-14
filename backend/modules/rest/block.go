package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/service"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"net/http"
	"strconv"
)

// mux.Router registrars
func RegisterBlock(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryBlock,
		registerQueryBlocks,
		registerQueryBlocksPrecommits,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

func registerQueryBlock(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryBlock, func(writer http.ResponseWriter, request *http.Request) {
		args := mux.Vars(request)

		var heigth int64
		if iHeight, err := strconv.Atoi(args["height"]); err == nil {
			heigth = int64(iHeight)
		}

		result := service.GetBlock().Query(heigth)
		resultByte, _ := json.Marshal(result)
		writer.Write(resultByte)

	}).Methods("GET")
	return nil
}

func registerQueryBlocks(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryBlocks, func(writer http.ResponseWriter, request *http.Request) {
		page, size := utils.GetPage(request)
		result := service.GetBlock().QueryList(page, size)
		resultByte, _ := json.Marshal(result)
		writer.Write(resultByte)

	}).Methods("GET")
	return nil
}

func registerQueryBlocksPrecommits(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryBlocksPrecommits, func(writer http.ResponseWriter, request *http.Request) {
		args := mux.Vars(request)
		address := args["address"]
		page, size := utils.GetPage(request)

		result := service.GetBlock().QueryPrecommits(address, page, size)
		resultByte, _ := json.Marshal(result)
		writer.Write(resultByte)
	}).Methods("GET")
	return nil
}
