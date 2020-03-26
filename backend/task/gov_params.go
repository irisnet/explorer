package task

import (
	"fmt"
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/utils"
)

type UpdateGovParams struct{}

func (task UpdateGovParams) Name() string {
	return "update_gov_params"
}
func (task UpdateGovParams) Start() {
	taskName := task.Name()
	timeInterval := conf.Get().Server.CronTimeGovParams
	utils.RunTimer(timeInterval, utils.Sec, func() {
		if notBeExec, err := tcService.assetTaskShouldNotBeExecuted(taskName, timeInterval); err != nil {
			logger.Error("assetTaskShouldNotBeExecuted fail", logger.String("taskName", taskName),
				logger.String("err", err.Error()))
		} else {
			if !notBeExec {
				if err := tcService.lockTask(taskName); err != nil {
					logger.Error("lockTask fail", logger.String("taskName", taskName),
						logger.String("err", err.Error()))
				} else {
					if err := task.DoTask(); err != nil {
						logger.Error(fmt.Sprintf("%s fail", task.Name()), logger.String("err", err.Error()))
					} else {
						logger.Info(fmt.Sprintf("%s success", task.Name()))
					}

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

func (task UpdateGovParams) DoTask() error {
	curModuleKv, err := lcd.GetAllGovModuleParam()
	if err != nil {
		return err
	}

	err = document.GovParams{}.UpdateCurrentModuleParamValue(curModuleKv)
	if err != nil {
		return err
	}

	return nil
}
