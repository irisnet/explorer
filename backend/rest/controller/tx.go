package controller

import (
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/explorer/backend/vo"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func RegisterTx(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryTx,
		registerQueryTxsByAccount,
		registerQueryTxsByDay,
		//new
		registerQueryTxListByType,
		registerQueryTxsCounter,
		registerQueryRecentTx,
		registerQueryTxList,
		registerQueryTxType,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

//type Tx struct {
//	*service.TxService
//}
//
//var tx = Tx{
//	service.Get(service.Tx).(*service.TxService),
//}

func registerQueryTxList(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryTxList, "GET", func(request vo.IrisReq) interface{} {
		tx.SetTid(request.TraceId)
		page := int(utils.ParseIntWithDefault(QueryParam(request, "page"), 1))
		size := int(utils.ParseIntWithDefault(QueryParam(request, "size"), 5))
		height := int(utils.ParseIntWithDefault(QueryParam(request, "height"), 0))
		total := QueryParam(request, "total")
		txType := QueryParam(request, "txType")
		status := QueryParam(request, "status")
		beginTime := int64(utils.ParseIntWithDefault(QueryParam(request, "beginTime"), 0))
		endTime := int64(utils.ParseIntWithDefault(QueryParam(request, "endTime"), 0))
		istotal := false
		if total == "true" {
			istotal = true
		}
		query := bson.M{}
		if height > 0 {
			query["height"] = height
		}
		if txType != "" {
			query["type"] = txType
		}

		if status != "" {
			query["status"] = status
		}
		if beginTime != 0 && endTime != 0 {
			query["time"] = bson.M{
				"$gte": time.Unix(beginTime, 0),
				"$lt":  time.Unix(endTime, 0),
			}
		} else if beginTime != 0 {
			query["time"] = bson.M{
				"$gte": time.Unix(beginTime, 0),
			}
		} else if endTime != 0 {
			query["time"] = bson.M{
				"$lt": time.Unix(endTime, 0),
			}
		}
		var result vo.PageVo
		result = tx.QueryBaseList(query, page, size, istotal)
		return result
	})
	return nil
}

func registerQueryTxListByType(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryTxListByType, "GET", func(request vo.IrisReq) interface{} {
		tx.SetTid(request.TraceId)
		query := bson.M{}
		total := GetString(request, "total")
		txTyp := QueryParam(request, "txType")
		status := QueryParam(request, "status")
		beginTime := int64(utils.ParseIntWithDefault(QueryParam(request, "beginTime"), 0))
		endTime := int64(utils.ParseIntWithDefault(QueryParam(request, "endTime"), 0))
		istotal := true
		if total == "false" {
			istotal = false
		}

		address := GetString(request, "address")

		if len(address) > 0 {
			query["$or"] = []bson.M{{"from": address}, {"to": address}, {"signers": bson.M{"$elemMatch": bson.M{"addr_bech32": address}}}}
		}

		height := GetInt(request, "height")
		if height > 0 {
			query["height"] = height
		}

		txType := Var(request, "type")
		page, size := GetPage(request)

		var result vo.PageVo
		if txTyp != "" {
			query["type"] = txTyp
		} else {
			switch types.TxTypeFromString(txType) {
			case types.Trans:
				query["type"] = bson.M{
					"$in": types.BankList,
				}
				return tx.QueryTxList(query, page, size, istotal)
			case types.Declaration:
				query["type"] = bson.M{
					"$in": types.DeclarationList,
				}
				return tx.QueryTxList(query, page, size, istotal)
			case types.Stake:
				query["type"] = bson.M{
					"$in": types.StakeList,
				}
				return tx.QueryTxList(query, page, size, istotal)
			case types.Gov:
				query["type"] = bson.M{
					"$in": types.GovernanceList,
				}
				break
			}
		}

		if status != "" {
			query["status"] = status
		}
		if beginTime != 0 && endTime != 0 {
			query["time"] = bson.M{
				"$gte": time.Unix(beginTime, 0),
				"$lt":  time.Unix(endTime, 0),
			}
		} else if beginTime != 0 {
			query["time"] = bson.M{
				"$gte": time.Unix(beginTime, 0),
			}
		} else if endTime != 0 {
			query["time"] = bson.M{
				"$lt": time.Unix(endTime, 0),
			}
		}

		result = tx.QueryTxList(query, page, size, istotal)
		return result
	})
	return nil
}

func registerQueryTx(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryTx, "GET", func(request vo.IrisReq) interface{} {
		tx.SetTid(request.TraceId)
		hash := Var(request, "hash")

		result := tx.Query(hash)
		return result
	})

	return nil
}

func registerQueryTxsCounter(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryTxsCounter, "GET", func(request vo.IrisReq) interface{} {
		tx.SetTid(request.TraceId)
		query := bson.M{}
		request.ParseForm()

		address := GetString(request, "address")
		if len(address) > 0 {
			query["$or"] = []bson.M{{"from": address}, {"to": address}, {"signers": bson.M{"$elemMatch": bson.M{"addr_bech32": address}}}}
		}

		height := GetInt(request, "height")
		if height > 0 {
			query["height"] = height
		}

		result := tx.CountByType(query)
		return result
	})

	return nil
}

func registerQueryTxsByAccount(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryTxsByAccount, "GET", func(request vo.IrisReq) interface{} {
		tx.SetTid(request.TraceId)
		address := Var(request, "address")
		page, size := GetPage(request)
		total := QueryParam(request, "total")
		istotal := true
		if total == "false" {
			istotal = false
		}
		result := tx.QueryByAcc(address, page, size, istotal)

		return result
	})

	return nil
}

func registerQueryTxsByDay(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryTxsByDay, "GET", func(request vo.IrisReq) interface{} {
		tx.SetTid(request.TraceId)
		result := tx.QueryTxNumGroupByDay()
		return result
	})
	return nil
}

func registerQueryRecentTx(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryRecentTx, "GET", func(request vo.IrisReq) interface{} {
		tx.SetTid(request.TraceId)
		result := tx.QueryRecentTx()
		return result
	})
	return nil
}

func registerQueryTxType(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryTxType, "GET", func(request vo.IrisReq) interface{} {
		tx.SetTid(request.TraceId)
		txType := QueryParam(request, "type")
		result := tx.QueryTxType(txType)
		return result
	})
	return nil
}
