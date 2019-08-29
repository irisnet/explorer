package task

import (
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
	utils.RunTimer(1, utils.Hour, func() {

		assetTokens, err := document.Asset{}.GetAllAssets()
		if err != nil {
			logger.Error("queryAssets failed", logger.String("taskName", task.Name()), logger.String("errmsg", err.Error()))
			return
		}

		assetService := service.Get(service.Asset).(*service.AssetsService)
		if err := assetService.UpdateAssetTokens(assetTokens); err != nil {
			logger.Error("UpdateAssetTokens task failed", logger.String("taskName", task.Name()), logger.String("errmsg", err.Error()))
			return
		}
		logger.Info("UpdateAssetTokens task is OK.")
	})
}
