package lcd

import (
	"fmt"
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/explorer/backend/logger"
	"encoding/json"
)

type LcdHtlc struct {
	Type  string   `json:"type"`
	Value HtlcData `json:"value"`
}

type HtlcData struct {
	Sender               string  `json:"sender"`
	To                   string  `json:"to"`
	ReceiverOnOtherChain string  `json:"receiver_on_other_chain"`
	Amount               []*Coin `json:"amount"`
	Secret               []byte  `json:"secret"`
	Timestamp            int64   `json:"timestamp,string"`
	ExpireHeight         int64   `json:"expire_height,string"`
	State                string  `json:"state"`
}

func HtlcInfo(hashlock string) (result LcdHtlc, err error) {
	url := fmt.Sprintf(UrlHtlcInfo, conf.Get().Hub.LcdUrl, hashlock)
	resBytes, err := utils.Get(url)
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("get account error", logger.String("err", err.Error()))
		return result, err
	}
	return result, nil
}
