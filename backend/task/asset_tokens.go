package task

import (
	"fmt"
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/service"
	"github.com/irisnet/explorer/backend/utils"
)

type UpdateAssetTokens struct{}

func (task UpdateAssetTokens) Name() string {
	return "update_asset_tokens"
}

func (task UpdateAssetTokens) Start() {
	utils.RunTimer(conf.Get().Server.CronTimeAssetTokens, utils.Sec, func() {
		if err := task.DoTask(); err != nil {
			logger.Error(fmt.Sprintf("%s fail", task.Name()), logger.String("err", err.Error()))
		} else {
			logger.Info(fmt.Sprintf("%s success", task.Name()))
		}
	})
}

func (task UpdateAssetTokens) DoTask() error {
	assetTokens, err := document.Asset{}.GetAllAssets()
	if err != nil {
		return err
	}

	assetService := service.Get(service.Asset).(*service.AssetsService)
	if err := assetService.UpdateAssetTokens(assetTokens); err != nil {
		return err
	}

	return nil
}
