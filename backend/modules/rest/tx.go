package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/service"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
)

func RegisterTx(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryTx,
		registerQueryTxs,
		registerQueryTxsByAccount,
		registerQueryTxsByDay,
		//new
		registerQueryTxList,
		registerQueryTxsCounter,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

func registerQueryTxList(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryTxList, func(writer http.ResponseWriter, request *http.Request) {
		args := mux.Vars(request)
		query := bson.M{}
		request.ParseForm()
		if len(request.Form["address"]) > 0 {
			address := request.Form["address"][0]
			query["$or"] = []bson.M{{"from": address}, {"to": address}}
		}

		if len(request.Form["height"]) > 0 {
			height, _ := strconv.Atoi(request.Form["height"][0])
			query["height"] = height
		}

		txType := args["type"]
		page, size := utils.GetPage(request)

		var result model.Page
		switch types.TxTypeFromString(txType) {
		case types.Trans:
			query["type"] = types.TypeTransfer
			break
		case types.Declaration:
			query["type"] = bson.M{
				"$in": types.DeclarationList,
			}
			break
		case types.Stake:
			query["type"] = bson.M{
				"$in": types.StakeList,
			}
			break
		case types.Gov:
			query["type"] = bson.M{
				"$in": types.GovernanceList,
			}
			break
		}
		result = service.GetTx().QueryList(query, page, size)
		resp, err := json.Marshal(result)
		if err == nil {
			writer.Write(resp)
		}

	}).Methods("GET")
	return nil
}

func registerQueryTx(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryTx, func(writer http.ResponseWriter, request *http.Request) {
		args := mux.Vars(request)
		hash := args["hash"]

		tx, err := service.GetTx().Query(hash)
		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		}
		resultByte, _ := json.Marshal(tx)
		writer.Write(resultByte)
	}).Methods("GET")
	return nil
}

func registerQueryTxs(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryTxs, func(writer http.ResponseWriter, request *http.Request) {
		query := bson.M{}
		var typeArr []string
		typeArr = append(typeArr, types.TypeTransfer)
		typeArr = append(typeArr, types.DeclarationList...)
		typeArr = append(typeArr, types.StakeList...)
		typeArr = append(typeArr, types.GovernanceList...)
		query["type"] = bson.M{
			"$in": typeArr,
		}
		page, pageSize := utils.GetPage(request)
		result := service.GetTx().QueryLatest(query, page, pageSize)

		res, _ := json.Marshal(result)
		writer.Write(res)
	}).Methods("GET")
	return nil
}

func registerQueryTxsCounter(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryTxsCounter, func(writer http.ResponseWriter, request *http.Request) {
		query := bson.M{}
		request.ParseForm()
		if len(request.Form["address"]) > 0 {
			address := request.Form["address"][0]
			query["$or"] = []bson.M{{"from": address}, {"to": address}}
		}

		if len(request.Form["height"]) > 0 {
			height, _ := strconv.Atoi(request.Form["height"][0])
			query["height"] = height
		}

		result := service.GetTx().CountByType(query)
		res, _ := json.Marshal(result)
		writer.Write(res)
	}).Methods("GET")
	return nil
}

func registerQueryTxsByAccount(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryTxsByAccount, func(writer http.ResponseWriter, request *http.Request) {
		args := mux.Vars(request)
		address := args["address"]
		page, size := utils.GetPage(request)
		result := service.GetTx().QueryByAcc(address, page, size)

		res, _ := json.Marshal(result)
		writer.Write(res)
	}).Methods("GET")
	return nil
}

func registerQueryTxsByDay(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryTxsByDay, func(writer http.ResponseWriter, request *http.Request) {
		result := service.GetTx().CountByDay()
		res, _ := json.Marshal(result)
		writer.Write(res)
	}).Methods("GET")
	return nil
}
