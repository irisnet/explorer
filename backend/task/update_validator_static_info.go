package task

import (
	"fmt"
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/service"
	"github.com/irisnet/explorer/backend/utils"
)

type ValidatorStaticInfo struct {
}

func (task ValidatorStaticInfo) Name() string {
	return "update_validator_static_data"
}

func (task ValidatorStaticInfo) Start() {
	utils.RunTimer(conf.Get().Server.CronTimeValidatorStaticInfo, utils.Sec, func() {
		if err := task.DoTask(); err != nil {
			logger.Error(fmt.Sprintf("%s fail", task.Name()), logger.String("err", err.Error()))
		} else {
			logger.Info(fmt.Sprintf("%s success", task.Name()))
		}
	})
}

func (task ValidatorStaticInfo) DoTask() error {
	validatorService := service.ValidatorService{}
	return validatorService.UpdateValidatorStaticInfo()
}
