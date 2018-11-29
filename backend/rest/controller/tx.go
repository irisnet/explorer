package controller

import (
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/service"
	"github.com/irisnet/explorer/backend/types"
	"gopkg.in/mgo.v2/bson"
	"net/http"
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
	doApi(r, types.UrlRegisterQueryTxList, "GET", func(request *http.Request) interface{} {
		query := bson.M{}

		address := GetString(request, "address")
		if len(address) > 0 {
			query["$or"] = []bson.M{{"from": address}, {"to": address}}
		}

		height := GetInt(request, "height")
		if height > 0 {
			query["height"] = height
		}

		txType := Var(request, "type")
		page, size := GetPage(request)

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
		return result
	})
	return nil
}

func registerQueryTx(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryTx, "GET", func(request *http.Request) interface{} {
		hash := Var(request, "hash")

		result := service.GetTx().Query(hash)
		return result
	})

	return nil
}

func registerQueryTxs(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryTxs, "GET", func(request *http.Request) interface{} {
		query := bson.M{}
		var typeArr []string
		typeArr = append(typeArr, types.TypeTransfer)
		typeArr = append(typeArr, types.DeclarationList...)
		typeArr = append(typeArr, types.StakeList...)
		typeArr = append(typeArr, types.GovernanceList...)
		query["type"] = bson.M{
			"$in": typeArr,
		}
		page, pageSize := GetPage(request)
		result := service.GetTx().QueryLatest(query, page, pageSize)

		return result
	})

	return nil
}

func registerQueryTxsCounter(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryTxsCounter, "GET", func(request *http.Request) interface{} {
		query := bson.M{}
		request.ParseForm()

		address := GetString(request, "address")
		if len(address) > 0 {
			query["$or"] = []bson.M{{"from": address}, {"to": address}}
		}

		height := GetInt(request, "height")
		if height > 0 {
			query["height"] = height
		}

		result := service.GetTx().CountByType(query)
		return result
	})

	return nil
}

func registerQueryTxsByAccount(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryTxsByAccount, "GET", func(request *http.Request) interface{} {
		address := Var(request, "address")
		page, size := GetPage(request)
		result := service.GetTx().QueryByAcc(address, page, size)

		return result
	})

	return nil
}

func registerQueryTxsByDay(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryTxsByDay, "GET", func(request *http.Request) interface{} {
		result := service.GetTx().CountByDay()
		return result
	})
	return nil
}
