package task

import (
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/service"
	"github.com/irisnet/explorer/backend/logger"
)

type UpdateAssetGateways struct{}

func (task UpdateAssetGateways) Name() string {
	return "update_asset_gateways"
}

func (task UpdateAssetGateways) Start() {
	utils.RunTimer(1, utils.Hour, func() {

		assetTokens, err := document.AssetGateways{}.GetAllAssetGateways()
		if err != nil {
			logger.Error("queryAssets failed", logger.String("taskName", task.Name()), logger.String("errmsg", err.Error()))
			return
		}

		assetService := service.Get(service.Asset).(*service.AssetsService)
		if err := assetService.UpdateAssetGateway(assetTokens); err != nil {
			logger.Error("UpdateAssetGateways task failed", logger.String("taskName", task.Name()), logger.String("errmsg", err.Error()))
			return
		}
		logger.Info("UpdateAssetGateways task is OK.")
	})
}
