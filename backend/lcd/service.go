package lcd

import (
	"time"
	"github.com/irisnet/explorer/backend/logger"
)

type ServiceBinding struct {
	ServiceName  string    `json:"service_name"`
	Provider     string    `json:"provider"`
	Deposit      []Coin    `json:"deposit"`
	Pricing      string    `json:"pricing"`
	Qos          uint64    `json:"qos"`
	Available    bool      `json:"available"`
	DisabledTime time.Time `json:"disabled_time"`
	Owner        string    `json:"owner"`
}

// RequestContext defines a context which holds request-related data
type RequestContext struct {
	ServiceName        string   `json:"service_name"`
	Providers          []string `json:"providers"`
	Consumer           string   `json:"consumer"`
	Input              string   `json:"input"`
	ServiceFeeCap      []Coin   `json:"service_fee_cap"`
	ModuleName         string   `json:"module_name"`
	Timeout            int64    `json:"timeout"`
	SuperMode          bool     `json:"super_mode"`
	Repeated           bool     `json:"repeated"`
	RepeatedFrequency  uint64   `json:"repeated_frequency"`
	RepeatedTotal      int64    `json:"repeated_total"`
	BatchCounter       uint64   `json:"batch_counter"`
	BatchRequestCount  uint32   `json:"batch_request_count"`
	BatchResponseCount uint32   `json:"batch_response_count"`
	ResponseThreshold  uint32   `json:"response_threshold"`
	BatchState         int32    `json:"batch_state"`
	State              int32    `json:"state"`
}

func GetServiceBindings(servicename, provider string) (resp []ServiceBinding) {
	bindingData, err := client.Service().QueryBindings(servicename)
	if err != nil {
		logger.Error("Query Bindings failed",
			logger.String("servicename", servicename),
			logger.String("err", err.Error()))
		return
	}

	for _, val := range bindingData {
		if provider == "" {
			resp = append(resp, ServiceBinding{
				ServiceName:  val.ServiceName,
				Provider:     val.Provider,
				Deposit:      LoadCoins(val.Deposit),
				Pricing:      val.Pricing,
				Available:    val.Available,
				DisabledTime: val.DisabledTime,
				Owner:        val.Owner,
				Qos:          val.Qos,
			})
		} else {
			if provider == val.Provider {
				resp = append(resp, ServiceBinding{
					ServiceName:  val.ServiceName,
					Provider:     val.Provider,
					Deposit:      LoadCoins(val.Deposit),
					Pricing:      val.Pricing,
					Available:    val.Available,
					DisabledTime: val.DisabledTime,
					Owner:        val.Owner,
					Qos:          val.Qos,
				})
			}
		}

	}
	return
}

func GetServiceContexts(reqcontextid string) (resp RequestContext) {
	reqContext, err := client.Service().QueryRequestContext(reqcontextid)
	if err != nil {
		logger.Error("Query Request Contexts failed",
			logger.String("reqcontextid", reqcontextid),
			logger.String("err", err.Error()))
		return
	}
	ConverAddrs := func() (ret []string) {
		for _, val := range reqContext.Providers {
			ret = append(ret, val.String())
		}
		return
	}
	resp = RequestContext{
		ServiceName:        reqContext.ServiceName,
		Providers:          ConverAddrs(),
		Consumer:           reqContext.Consumer.String(),
		Input:              reqContext.Input,
		ServiceFeeCap:      LoadCoins(reqContext.ServiceFeeCap),
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
	return
}
