package vo

import (
	"time"
)

type ServiceBindingRespond []ServiceBinding

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

type ServiceRequestRespond struct {
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
	State              string   `json:"state"`
}