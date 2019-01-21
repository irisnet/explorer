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

	logger.Info("LogPreFilter", traceId, apiPath, agent, formValue, urlValue)
	return nil, types.CodeSuccess
}

func (LogPreFilter) Paths() []string {
	return []string{GlobalFilterPath}
}

func (LogPreFilter) Type() Type {
	return Pre
}

// display user's request information,optional
type LogPostFilter struct {
}

func (LogPostFilter) Do(request *model.IrisReq, data interface{}) (interface{}, types.BizCode) {
	traceId := logger.Int64("traceId", request.TraceId)
	coastSecond := time.Now().Unix() - request.Start.Unix()
	coast := logger.Int64("coast", coastSecond)
	defer func() {
		if r := recover(); r != nil {
			logger.Error("LogPostFilter failed", traceId)
		}
	}()
	logger.Info("LogPostFilter", traceId, coast)
	if coastSecond >= 3 {
		logger.Warn("LogPostFilter api coast most time", traceId)
	}
	return nil, types.CodeSuccess
}

func (LogPostFilter) Paths() []string {
	return []string{GlobalFilterPath}
}

func (LogPostFilter) Type() Type {
	return Pre
}
