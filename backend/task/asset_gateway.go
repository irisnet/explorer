package task

import (
	"fmt"
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/service"
	"github.com/irisnet/explorer/backend/utils"
)

type UpdateAssetGateways struct{}

func (task UpdateAssetGateways) Name() string {
	return "update_asset_gateways"
}

func (task UpdateAssetGateways) Start() {
	utils.RunTimer(conf.Get().Server.CronTimeAssetGateways, utils.Sec, func() {
		if err := task.DoTask(); err != nil {
			logger.Error(fmt.Sprintf("%s fail", task.Name()), logger.String("err", err.Error()))
		} else {
			logger.Info(fmt.Sprintf("%s success", task.Name()))
		}
	})
}

func (task UpdateAssetGateways) DoTask() error {
	assetGateways, err := document.AssetGateways{}.GetAllAssetGateways()
	if err != nil {
		return err
	}

	assetService := service.Get(service.Asset).(*service.AssetsService)
	if err := assetService.UpdateAssetGateway(assetGateways); err != nil {
		return err
	}

	return nil
}
