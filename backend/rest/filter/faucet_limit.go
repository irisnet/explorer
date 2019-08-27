package filter

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/vo"
	"time"
)

var rateLimitMap = make(map[string]int, 0)

func init() {
	go func() {
		for {
			now := time.Now()
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			select {
			case <-t.C:
				logger.Warn("clear count")
				rateLimitMap = make(map[string]int, 0)
			}
		}
	}()
}

type FaucetLimitPreFilter struct{}

func (FaucetLimitPreFilter) Match() []Condition {
	return []Condition{
		{
			path:   types.UrlRegisterApply,
			method: "POST",
		},
	}
}

func (FaucetLimitPreFilter) Type() Type {
	return Pre
}

func (FaucetLimitPreFilter) Do(request *vo.IrisReq, data interface{}) (interface{}, types.BizCode) {
	traceId := logger.String("traceId", request.TraceId)
	logger.Info("FaucetLimitPreFilter", traceId)

	args := mux.Vars(request.Request)
	var addr = args["address"]
	cnt := rateLimitMap[addr]
	if cnt >= conf.Get().Server.MaxDrawCnt {
		logger.Warn("exceed draw coin limit", traceId, logger.String("addr", addr), logger.Int("count", cnt))
		return nil, types.CodeRateLimit
	}
	return nil, types.CodeSuccess
}

type FaucetLimitPostFilter struct{}

func (FaucetLimitPostFilter) Match() []Condition {
	return []Condition{
		{
			path:   types.UrlRegisterApply,
			method: "POST",
		},
	}
}

func (FaucetLimitPostFilter) Type() Type {
	return Post
}

func (FaucetLimitPostFilter) Do(request *vo.IrisReq, data interface{}) (interface{}, types.BizCode) {
	traceId := logger.String("traceId", request.TraceId)
	logger.Info("FaucetLimitPostFilter", traceId)
	d := data.([]byte)
	var res fErr
	if err := json.Unmarshal(d, &res); err == nil {
		args := mux.Vars(request.Request)
		var addr = args["address"]
		cnt := rateLimitMap[addr]
		rateLimitMap[addr] = cnt + 1
		logger.Info("already draw count", logger.String("addr", addr), logger.Int("count", cnt+1))
	}

	return nil, types.CodeSuccess
}

type fErr struct {
	Code string `json:"err_code,omitempty"`
	Msg  string `json:"err_msg,omitempty"`
}
