package task

import (
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
	timeInterval := conf.Get().Server.CronTimeAssetGateways
	taskName := task.Name()

	utils.RunTimer(timeInterval, utils.Sec, func() {
		if err := tcService.runTask(taskName, timeInterval, task.DoTask); err != nil {
			logger.Error(err.Error())
		}
	})
}

func (task UpdateAssetGateways) DoTask(fn func(string) chan bool) error {
	defer HeartQuit(fn(task.Name()))
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
