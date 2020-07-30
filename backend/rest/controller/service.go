package controller

import (
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/vo"
)

func RegisterServices(r *mux.Router) error {
	apis := []func(*mux.Router) error{
		RegisterQueryServiceBindings,
		RegisterQueryServiceRequests,
	}

	for _, fn := range apis {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

// @Summary bindings
// @Description get service by servicename
// @Tags service
// @Accept  json
// @Produce  json
// @Param   servicename    path   string true    "servicename"
// @Success 200 {object} vo.ServiceBindingRespond   "success"
// @Router /api/service/bindings/{servicename} [get]
func RegisterQueryServiceBindings(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryServiceBindings, "GET", func(request vo.IrisReq) interface{} {
		tx.SetTid(request.TraceId)
		servicename := Var(request, "servicename")
		return iservice.QueryServiceBindings(servicename)
	})
	return nil
}

// @Summary requests
// @Description get service requests by contextid
// @Tags service
// @Accept  json
// @Produce  json
// @Param   contextid    path   string true    "contextid"
// @Success 200 {object} vo.ServiceRequestRespond   "success"
// @Router /api/service/requests/{contextid} [get]
func RegisterQueryServiceRequests(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryServiceRequest, "GET", func(request vo.IrisReq) interface{} {
		tx.SetTid(request.TraceId)
		contextid := Var(request, "contextid")
		return iservice.QueryServiceRequests(contextid)
	})
	return nil
}
