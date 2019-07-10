package service

import "github.com/irisnet/explorer/backend/model"

type TokenStatsService struct {
	BaseService
}

func (service *TokenStatsService) QueryTokenStats() (result model.TokenStatsVo,err error) {

	return
}

func (service *TokenStatsService) QueryTokensAccountTotal() (result map[string]string,err error) {

	return
}
