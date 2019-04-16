package lcd

import (
	"encoding/json"
	"fmt"
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/utils"
)

func Validator(address string) (result ValidatorVo, err error) {
	url := fmt.Sprintf(UrlValidator, conf.Get().Hub.LcdUrl, address)
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

func Validators(page, size int) (result []ValidatorVo) {
	url := fmt.Sprintf(UrlValidators, conf.Get().Hub.LcdUrl, page, size)
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("get Validators error", logger.String("err", err.Error()))
		return result
	}

	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("Unmarshal Validators error", logger.String("err", err.Error()))
		return result
	}
	return result
}

func QueryWithdrawAddr(address string) (result string) {
	url := fmt.Sprintf(UrlWithdrawAddress, conf.Get().Hub.LcdUrl, address)
	resBytes, err := utils.Get(url)
	if err != nil {
		return result
	}
	result = string(resBytes)
	return
}

func DelegationByValidator(address string) (result []DelegationVo) {
	url := fmt.Sprintf(UrlDelegationByVal, conf.Get().Hub.LcdUrl, address)
	resBytes, err := utils.Get(url)
	if err != nil {
		return result
	}
	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("Unmarshal Delegation error", logger.String("err", err.Error()))
		return result
	}
	return
}

func StakePool() (result StakePoolVo) {
	url := fmt.Sprintf(UrlStakePool, conf.Get().Hub.LcdUrl)
	resBytes, err := utils.Get(url)
	if err != nil {
		return result
	}
	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("Unmarshal StakePool error", logger.String("err", err.Error()))
		return result
	}
	return
}

func SignInfo(consensusPubkey string) (result SignInfoVo) {
	url := fmt.Sprintf(UrlSignInfo, conf.Get().Hub.LcdUrl, consensusPubkey)
	resBytes, err := utils.Get(url)
	if err != nil {
		return result
	}
	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("Unmarshal SignInfoVo error", logger.String("err", err.Error()))
		return result
	}
	return
}
