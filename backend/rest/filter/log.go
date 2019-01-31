package filter

import (
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/types"
	"time"
)

// display user's request information,optional
type LogPreFilter struct {
}

func (LogPreFilter) Do(request *model.IrisReq, data interface{}) (interface{}, types.BizCode) {
	start := time.Now()
	request.TraceId = string(start.UnixNano())
	request.Start = start

	traceId := logger.String("traceId", request.TraceId)
	apiPath := logger.String("path", request.RequestURI)
	formValue := logger.Any("form", request.Form)
	urlValue := logger.Any("url", mux.Vars(request.Request))
	agent := logger.String("agent", request.UserAgent())
	defer func() {
		if r := recover(); r != nil {
			logger.Error("LogPostFilter failed", traceId)
		}
	}()

	logger.Info("LogPreFilter", traceId, apiPath, agent, formValue, urlValue)
	return nil, types.CodeSuccess
}

func (LogPreFilter) Match() []Condition {
	return []Condition{globalCondition}
}

func (LogPreFilter) Type() Type {
	return Pre
}

// display user's request information,optional
type LogPostFilter struct {
}

func (LogPostFilter) Do(request *model.IrisReq, data interface{}) (interface{}, types.BizCode) {
	traceId := logger.String("traceId", request.TraceId)
	coastSecond := time.Now().Unix() - request.Start.Unix()
	coast := logger.Int64("coast", coastSecond)
	apiPath := logger.String("path", request.RequestURI)
	defer func() {
		if r := recover(); r != nil {
			logger.Error("LogPostFilter failed", traceId)
		}
	}()
	logger.Info("LogPostFilter", apiPath, traceId, coast)
	if coastSecond >= 3 {
		logger.Warn("LogPostFilter api coast most time", apiPath, traceId, coast)
	}
	return nil, types.CodeSuccess
}

func (LogPostFilter) Match() []Condition {
	return []Condition{globalCondition}
}

func (LogPostFilter) Type() Type {
	return Pre
}
