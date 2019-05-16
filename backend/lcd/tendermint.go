package lcd

import (
	"encoding/json"
	"fmt"

	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/utils"
)

func NodeInfo() (result NodeInfoVo, err error) {
	url := fmt.Sprintf(UrlNodeInfo, conf.Get().Hub.LcdUrl)
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

func Genesis() (result GenesisVo, err error) {
	url := fmt.Sprintf(UrlGenesis, conf.Get().Hub.NodeUrl)
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

func Block(height int64) (result BlockVo) {
	url := fmt.Sprintf(UrlBlock, conf.Get().Hub.LcdUrl, height)
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("Block error", logger.Int64("height", height))
		return result
	}

	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("Block error", logger.String("err", err.Error()))
		return result
	}
	return result
}

func BlockLatest() (result BlockVo) {
	url := fmt.Sprintf(UrlBlockLatest, conf.Get().Hub.LcdUrl)
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("BlockLatest error", logger.String("err", err.Error()))
		return result
	}

	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("BlockLatest error", logger.String("err", err.Error()))
		return result
	}
	return result
}

func ValidatorSet(height int64) (result ValidatorSetVo) {
	url := fmt.Sprintf(UrlValidatorSet, conf.Get().Hub.LcdUrl, height)
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("BlockLatest error", logger.String("err", err.Error()))
		return result
	}

	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("BlockLatest error", logger.String("err", err.Error()))
		return result
	}
	return result
}

func LatestValidatorSet() (result ValidatorSetVo) {
	url := fmt.Sprintf(UrlValidatorSetLatest, conf.Get().Hub.LcdUrl)
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("BlockLatest error", logger.String("err", err.Error()))
		return result
	}

	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("BlockLatest error", logger.String("err", err.Error()))
		return result
	}
	return result
}

func BlockResult(height int64) (result BlockResultVo) {

	url := fmt.Sprintf(UrlBlocksResult, conf.Get().Hub.LcdUrl, height)
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("BlockResult error", logger.String("err", err.Error()))
		return result
	}

	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("BlockResult error", logger.String("err", err.Error()))
		return result
	}
	return result

}
