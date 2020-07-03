package controller

import (
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/vo"
	"strings"
	"github.com/irisnet/explorer/backend/utils"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

func RegisterCaculate(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerCaculateMonthData,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

// @Summary list
// @Description get caculate validators or delegators monthdata
// @Tags caculate
// @Accept  json
// @Produce  json
// @Param   type   query   string  true    "role type" Enums(delegator,validator)
// @Param   page    query   int true    "page num" Default(1)
// @Param   size   query   int true    "page size" Default(5)
// @Param   caculateTime   query  string false    "caculateTime (eg 2006.01.02)"
// @Param   total   query   bool false    "total" Enums(true,false)
// @Success 200 {object} vo.ExStaticMonthDataRespond	"success"
// @Router /api/caculate/monthdata [get]
func registerCaculateMonthData(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryCaculateMonthData, "GET", func(request vo.IrisReq) interface{} {
		bondedtokens.SetTid(request.TraceId)
		page := int(utils.ParseIntWithDefault(QueryParam(request, "page"), 1))
		size := int(utils.ParseIntWithDefault(QueryParam(request, "size"), 100))
		vtype := strings.ToLower(QueryParam(request, "type"))
		total := QueryParam(request, "total")
		istotal := false
		if total == "true" {
			istotal = true
		}
		fmt.Println(page, size)

		caculateTime := QueryParam(request, "caculateTime")
		//beginTime := QueryParam(request, "beginTime")
		//endTime := QueryParam(request, "endTime")
		//cst := time.FixedZone("CST", 8*3600)

		//starttime, _ := time.ParseInLocation(types.TimeLayout, beginTime, cst)
		//endtime, _ := time.ParseInLocation(types.TimeLayout, endTime, cst)

		cond := bson.M{
		}
		if caculateTime != "" {
			cond["date"] = caculateTime
		}
		//if !starttime.IsZero() && !endtime.IsZero() {
		//	cond["date"] = bson.M{
		//		"$gte": starttime,
		//		"$lt":  endTime,
		//	}
		//} else if !starttime.IsZero() {
		//	cond["date"] = bson.M{
		//		"$gte": starttime,
		//	}
		//} else if !endtime.IsZero() {
		//	cond["date"] = bson.M{
		//		"$lt": endtime,
		//	}
		//}

		resp := vo.ExStaticMonthDataRespond{
			Type: vtype,
		}
		var err error
		switch vtype {
		case "delegator":
			resp.Data, resp.Total, err = caculate.GetDelegatorCaculateMonth(cond, page, size, istotal)
		case "validator":
			resp.Data, resp.Total, err = caculate.GetValidatorCaculateMonth(cond, page, size, istotal)
		}
		if err != nil {
			return vo.NewResponse("-1", err.Error(), nil)
		}
		return vo.NewResponse(types.CodeSuccess.Code, types.CodeSuccess.Msg, resp)
	})

	return nil
}
