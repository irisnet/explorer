package controller

import (
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/vo"
	"github.com/irisnet/explorer/backend/utils"
	"gopkg.in/mgo.v2/bson"
)

func RegisterHtlc(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryHtlcTx,
		registerQueryHtlcInfo,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

// @Summary detail
// @Description get htlc info
// @Tags htlc
// @Accept  json
// @Produce  json
// @Param   hash_lock   path   string true    "hash_lock"
// @Success 200 {object} vo.HtlcInfo	"success"
// @Router /api/htlcs/{hash_lock} [get]
func registerQueryHtlcInfo(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryHtlc, "GET", func(request vo.IrisReq) interface{} {
		tx.SetTid(request.TraceId)
		hash_lock := Var(request, "hash_lock")
		//fmt.Println(hash_lock)
		var result vo.HtlcInfo
		result = htlc.QueryHtlcByHashLock(hash_lock)
		return result
	})
	return nil
}

// @Summary list
// @Description get htlc txs list
// @Tags htlc
// @Accept  json
// @Produce  json
// @Param   page    query   int true    "page num" Default(1)
// @Param   size   query   int true    "page size" Default(5)
// @Param   hash_lock   path   string true    "hash_lock"
// @Param   height   query   int64 false    "height"
// @Param   txType   query   string false    "txType"
// @Param   address   query   string false    "address"
// @Param   status   query   string false    "status" Enums(success,fail)
// @Param   total   query   bool false    "total" Enums(true,false)
// @Success 200 {object} vo.PageVo	"success"
// @Router /api/htlcs/{hash_lock}/txs [get]
func registerQueryHtlcTx(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryHtlcTxs, "GET", func(request vo.IrisReq) interface{} {
		tx.SetTid(request.TraceId)
		page := int(utils.ParseIntWithDefault(QueryParam(request, "page"), 1))
		size := int(utils.ParseIntWithDefault(QueryParam(request, "size"), 5))
		height := int(utils.ParseIntWithDefault(QueryParam(request, "height"), 0))
		total := QueryParam(request, "total")
		txType := QueryParam(request, "txType")
		address := QueryParam(request, "address")
		status := QueryParam(request, "status")
		hash_lock := Var(request, "hash_lock")
		//beginTime := int64(utils.ParseIntWithDefault(QueryParam(request, "beginTime"), 0))
		//endTime := int64(utils.ParseIntWithDefault(QueryParam(request, "endTime"), 0))
		//utc,_ := time.LoadLocation("UTC")
		istotal := false
		if total == "true" {
			istotal = true
		}

		query := bson.M{
			"msgs.msg.hash_lock": hash_lock,
		}
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
			query["$or"] = []bson.M{
				{"signers.addr_bech32": address},
				{"from": address},
				{"to": address},
			}

		}
		//if beginTime != 0 && endTime != 0 {
		//	query["time"] = bson.M{
		//		"$gte": time.Unix(beginTime, 0).In(utc),
		//		"$lt":  time.Unix(endTime, 0).In(utc),
		//	}
		//} else if beginTime != 0 {
		//	query["time"] = bson.M{
		//		"$gte": time.Unix(beginTime, 0).In(utc),
		//	}
		//} else if endTime != 0 {
		//	query["time"] = bson.M{
		//		"$lt": time.Unix(endTime, 0).In(utc),
		//	}
		//}
		var result vo.PageVo
		result = tx.QueryHtlcTx(query, page, size, istotal)
		return result
	})
	return nil
}
