package filter

import (
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
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

func (FaucetLimitPreFilter) Paths() []string {
	return []string{types.UrlRegisterApply}
}

func (FaucetLimitPreFilter) Type() Type {
	return Pre
}

func (FaucetLimitPreFilter) Do(request *model.IrisReq, data interface{}) (interface{}, types.BizCode) {
	traceId := logger.Int64("traceId", request.TraceId)
	logger.Warn("FaucetLimitPreFilter", traceId)

	var addr = utils.GetIpAddr(request.Request)
	cnt, ok := rateLimitMap[addr]
	if ok {
		if cnt >= conf.Get().Server.MaxDrawCnt {
			logger.Warn("exceed draw coin limit", traceId, logger.String("addr", addr))
			return nil, types.CodeRateLimit
		}
	}
	rateLimitMap[addr] = cnt + 1
	return nil, types.CodeSuccess
}
