package filter

import (
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/types"
	"time"
)

type LogPreFilter struct {
}

type LogPostFilter struct {
}

// display user's request information,optional
func (LogPreFilter) Do(request *model.IrisReq) (bool, interface{}, types.BizCode) {
	start := time.Now()
	request.TraceId = start.UnixNano()
	request.Start = start

	traceId := logger.Int64("traceId", request.TraceId)
	apiPath := logger.String("path", request.RequestURI)
	formValue := logger.Any("form", request.Form)
	urlValue := logger.Any("url", mux.Vars(request.Request))
	agent := logger.String("agent", request.UserAgent())
	defer func() {
		if r := recover(); r != nil {
			logger.Error("LogPostFilter failed", traceId)
		}
	}()

	logger.Debug("LogPreFilter", traceId, apiPath, agent, formValue, urlValue)
	return true, nil, types.CodeSysmaintenance
}

// display user's request information,optional
func (LogPostFilter) Do(request *model.IrisReq) (bool, interface{}, types.BizCode) {
	traceId := logger.Int64("traceId", request.TraceId)
	coastSecond := time.Now().Unix() - request.Start.Unix()
	coast := logger.Int64("coast", coastSecond)
	defer func() {
		if r := recover(); r != nil {
			logger.Error("LogPostFilter failed", traceId)
		}
	}()
	logger.Debug("LogPostFilter", traceId, coast)
	if coastSecond >= 3 {
		logger.Warn("LogPostFilter api coast most time", traceId)
	}
	return true, nil, types.CodeSysmaintenance
}
