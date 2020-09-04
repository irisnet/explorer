package service

import (
	"github.com/irisnet/explorer/backend/vo"
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/utils"
)

type IServiceService struct {
	BaseService
}

func (service *IServiceService) QueryServiceBindings(servicename, provider string) vo.ServiceBindingRespond {
	bindings := lcd.GetServiceBindings(servicename, provider)
	resp := []vo.ServiceBinding{}
	if len(bindings) == 0 {
		return resp
	}
	for _, val := range bindings {
		resp = append(resp, vo.ServiceBinding{
			ServiceName:  val.ServiceName,
			Provider:     val.Provider,
			Deposit:      converCoin(val.Deposit),
			Pricing:      val.Pricing,
			Available:    val.Available,
			DisabledTime: val.DisabledTime,
			Owner:        val.Owner,
			Qos:          val.Qos,
		})
	}
	return resp
}

func (service *IServiceService) QueryServiceRequests(contextid string) vo.ServiceRequestRespond {
	reqContext := lcd.GetServiceContexts(contextid)

	return vo.ServiceRequestRespond{
		ServiceName:        reqContext.ServiceName,
		Providers:          reqContext.Providers,
		Consumer:           reqContext.Consumer,
		Input:              reqContext.Input,
		ServiceFeeCap:      converCoin(reqContext.ServiceFeeCap),
		State:              reqContext.State,
		SuperMode:          reqContext.SuperMode,
		ModuleName:         reqContext.ModuleName,
		Timeout:            reqContext.Timeout,
		ResponseThreshold:  reqContext.ResponseThreshold,
		RepeatedFrequency:  reqContext.RepeatedFrequency,
		RepeatedTotal:      reqContext.RepeatedTotal,
		Repeated:           reqContext.Repeated,
		BatchCounter:       reqContext.BatchCounter,
		BatchRequestCount:  reqContext.BatchRequestCount,
		BatchResponseCount: reqContext.BatchResponseCount,
		BatchState:         reqContext.BatchState,
	}
}

func converCoin(coins []lcd.Coin) (ret []vo.Coin) {
	for _, val := range coins {
		amount, _ := utils.ParseStringToFloat(val.Amount)
		ret = append(ret, vo.Coin{Amount: amount, Denom: val.Denom})
	}
	return
}
