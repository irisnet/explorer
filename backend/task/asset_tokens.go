package task

import (
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/service"
	"github.com/irisnet/explorer/backend/utils"
)

type UpdateAssetTokens struct{}

func (task UpdateAssetTokens) Name() string {
	return "update_asset_tokens"
}

func (task UpdateAssetTokens) Start() {
	taskName := task.Name()
	timeInterval := conf.Get().Server.CronTimeAssetTokens

	utils.RunTimer(timeInterval, utils.Sec, func() {
		if err := tcService.runTask(taskName, timeInterval, task.DoTask); err != nil {
			logger.Error(err.Error())
		}
	})
}

func (task UpdateAssetTokens) DoTask(fn func(string) chan bool) error {
	stop := fn(task.Name())
	defer HeartQuit(stop)

	var assetService service.AssetsService
	assetService.UpdateAssetTokens()

	return nil
}
