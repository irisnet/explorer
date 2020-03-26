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
	taskName := task.Name()
	timeInterval := conf.Get().Server.CronTimeAssetTokens
	utils.RunTimer(timeInterval, utils.Sec, func() {
		if notBeExec, err := tcService.assetTaskShouldNotBeExecuted(taskName, timeInterval); err != nil {
			logger.Error("assetTaskShouldNotBeExecuted fail", logger.String("taskName", taskName),
				logger.String("err", err.Error()))
		} else {
			if !notBeExec {
				// lock task
				if err := tcService.lockTask(taskName); err != nil {
					logger.Error("lockTask fail", logger.String("taskName", taskName),
						logger.String("err", err.Error()))
				} else {
					// do task
					if err := task.DoTask(); err != nil {
						logger.Error(fmt.Sprintf("%s fail", task.Name()), logger.String("err", err.Error()))
					} else {
						logger.Info(fmt.Sprintf("%s success", task.Name()))
					}

					// unlock task
					if err := tcService.unlockTask(taskName); err != nil {
						logger.Error("unlockTask fail", logger.String("taskName", taskName),
							logger.String("err", err.Error()))
					}
				}
			} else {
				logger.Debug(fmt.Sprintf("%s shouldn't be executed on this instance", task.Name()))
			}
		}
	})
}

func (task UpdateAssetTokens) DoTask() error {
	assetTokens, err := document.AssetToken{}.GetAllAssets()
	if err != nil {
		return err
	}

	assetService := service.Get(service.Asset).(*service.AssetsService)
	if err := assetService.UpdateAssetTokens(assetTokens); err != nil {
		return err
	}

	return nil
}
