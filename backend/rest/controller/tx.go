package controller

import (
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/explorer/backend/vo"
	"gopkg.in/mgo.v2/bson"
	"time"
	"github.com/irisnet/explorer/backend/conf"
)

func RegisterTx(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryTx,
		registerQueryTxsByAccount,
		registerQueryTxsByDay,
		//new
		registerQueryTxListByType,
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

// @Summary list
// @Description get txs list
// @Tags txs
// @Accept  json
// @Produce  json
// @Param   page    query   int true    "page num" Default(1)
// @Param   size   query   int true    "page size" Default(5)
// @Param   height   query   int64 false    "height"
// @Param   txType   query   string false    "txType"
// @Param   status   query   string false    "status" Enums(success,fail)
// @Param   address   query   string false    "address"
// @Param   beginTime   query  int64 false    "beginTime"
// @Param   endTime   query   int64 false    "endTime"
// @Param   total   query   bool false    "total" Enums(true,false)
// @Success 200 {object} vo.PageVo	"success"
// @Router /api/txs [get]
func registerQueryTxList(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryTxList, "GET", func(request vo.IrisReq) interface{} {
		tx.SetTid(request.TraceId)
		page := int(utils.ParseIntWithDefault(QueryParam(request, "page"), 1))
		size := int(utils.ParseIntWithDefault(QueryParam(request, "size"), 5))
		height := int(utils.ParseIntWithDefault(QueryParam(request, "height"), 0))
		total := QueryParam(request, "total")
		txType := QueryParam(request, "txType")
		status := QueryParam(request, "status")
		address := QueryParam(request, "address")
		beginTime := int64(utils.ParseIntWithDefault(QueryParam(request, "beginTime"), 0))
		endTime := int64(utils.ParseIntWithDefault(QueryParam(request, "endTime"), 0))
		utc,_ := time.LoadLocation("UTC")
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
		if address != "" {
			if txType != "" {
				switch txType {
				case types.TxTypeTransfer:
					query["$or"] = []bson.M{
						{"from": address},
						{"to": address},
					}
				case types.TxTypeCreateHTLC, types.TxTypeClaimHTLC, types.TxTypeRefundHTLC:
					query["$or"] = []bson.M{
						{"signers.addr_bech32": address},
						{"from": address},
						{"to": address},
					}
				default:
					query["signers.addr_bech32"] = address
				}

			} else {
				query["$or"] = []bson.M{
					{"signers.addr_bech32": address},
					{"to": address},
				}
			}

		}
		if beginTime != 0 && endTime != 0 {
			query["time"] = bson.M{
				"$gte": time.Unix(beginTime, 0).In(utc),
				"$lt":  time.Unix(endTime, 0).In(utc),
			}
		} else if beginTime != 0 {
			query["time"] = bson.M{
				"$gte": time.Unix(beginTime, 0).In(utc),
			}
		} else if endTime != 0 {
			query["time"] = bson.M{
				"$lt": time.Unix(endTime, 0).In(utc),
			}
		}
		var result vo.PageVo
		result = tx.QueryBaseList(query, page, size, istotal)
		return result
	})
	return nil
}

// @Summary list by type
// @Description get txs list by type
// @Tags txs
// @Accept  json
// @Produce  json
// @Param   page    path   int true    "page num" Default(1)
// @Param   size   path   int true    "page size" Default(5)
// @Param   type   path   string true    "type"
// @Param   height   query   int64 false    "height"
// @Param   txType   query   string false    "txType"
// @Param   status   query   string false    "status" Enums(success,fail)
// @Param   address   query   string false    "address"
// @Param   beginTime   query  int64 false    "beginTime"
// @Param   endTime   query   int64 false    "endTime"
// @Param   total   query   bool false    "total" Enums(true,false)
// @Success 200 {object} vo.PageVo	"success"
// @Router /api/txs/{type}/{page}/{size} [get]
func registerQueryTxListByType(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryTxListByType, "GET", func(request vo.IrisReq) interface{} {
		tx.SetTid(request.TraceId)
		query := bson.M{}
		total := GetString(request, "total")
		txType := QueryParam(request, "txType")
		status := QueryParam(request, "status")
		beginTime := int64(utils.ParseIntWithDefault(QueryParam(request, "beginTime"), 0))
		endTime := int64(utils.ParseIntWithDefault(QueryParam(request, "endTime"), 0))
		utc,_ := time.LoadLocation("UTC")
		istotal := true
		if total == "false" {
			istotal = false
		}

		address := GetString(request, "address")

		if len(address) > 0 {
			condition := []bson.M{{"to": address}}
			prefix, _, _ := utils.DecodeAndConvert(address)
			if prefix == conf.Get().Hub.Prefix.ValAddr {
				condition = append(condition, bson.M{"from": address})
			}else{
				condition = append(condition, bson.M{"signers.addr_bech32": address})
			}
			query["$or"] = condition
		}

		height := GetInt(request, "height")
		if height > 0 {
			query["height"] = height
		}

		txTypeGroup := Var(request, "type")
		page, size := GetPage(request)

		var result vo.PageVo
		if status != "" {
			query["status"] = status
		}
		if beginTime != 0 && endTime != 0 {
			query["time"] = bson.M{
				"$gte": time.Unix(beginTime, 0).In(utc),
				"$lt":  time.Unix(endTime, 0).In(utc),
			}
		} else if beginTime != 0 {
			query["time"] = bson.M{
				"$gte": time.Unix(beginTime, 0).In(utc),
			}
		} else if endTime != 0 {
			query["time"] = bson.M{
				"$lt": time.Unix(endTime, 0).In(utc),
			}
		}

		if txType != "" {
			query["type"] = txType
		} else {
			switch types.TxTypeFromString(txTypeGroup) {
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

		result = tx.QueryTxList(query, page, size, istotal)
		return result
	})
	return nil
}

// @Summary tx detail
// @Description get txs detail
// @Tags txs
// @Accept  json
// @Produce  json
// @Param   hash    path   string true    "txhash"
// @Success 200 {object} vo.StakeTx	"success"
// @Router /api/tx/{hash} [get]
func registerQueryTx(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryTx, "GET", func(request vo.IrisReq) interface{} {
		tx.SetTid(request.TraceId)
		hash := Var(request, "hash")

		result := tx.Query(hash)
		return result
	})

	return nil
}


// @Summary txsByAddress
// @Description txsByAddress
// @Tags txs
// @Accept  json
// @Produce  json
// @Param   address  path   string true    "address"
// @Param   page   path   int64 true    "pagenum"
// @Param   size   path   int64 true    "pagesize"
// @Param   total   query   bool false    "total" Enums(true,false)
// @Success 200 {object} vo.PageVo	"success"
// @Router /api/txsByAddress/{address}/{page}/{size} [get]
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

// @Summary txsByDay
// @Description get txs ByDay
// @Tags txs
// @Accept  json
// @Produce  json
// @Success 200 {object} vo.TxNumGroupByDayVoRespond	"success"
// @Router /api/txsByDay [get]
func registerQueryTxsByDay(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryTxsByDay, "GET", func(request vo.IrisReq) interface{} {
		tx.SetTid(request.TraceId)
		result := tx.QueryTxNumGroupByDay()
		return result
	})
	return nil
}

// @Summary txs recent
// @Description get txs recent
// @Tags txs
// @Accept  json
// @Produce  json
// @Success 200 {object} vo.RecentTxRespond	"success"
// @Router /api/txs/recent [get]
func registerQueryRecentTx(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryRecentTx, "GET", func(request vo.IrisReq) interface{} {
		tx.SetTid(request.TraceId)
		result := tx.QueryRecentTx()
		return result
	})
	return nil
}

// @Summary tx_types detail
// @Description get tx_types detail
// @Tags txs
// @Accept  json
// @Produce  json
// @Param   type   path   string true    "type"
// @Success 200 {object} vo.QueryTxTypeRespond	"success"
// @Router /api/tx_types/{type} [get]
func registerQueryTxType(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryTxType, "GET", func(request vo.IrisReq) interface{} {
		tx.SetTid(request.TraceId)
		txType := Var(request, "type")
		result := tx.QueryTxType(txType)
		return result
	})
	return nil
}
