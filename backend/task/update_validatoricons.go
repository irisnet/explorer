package task

import (
	"fmt"
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/service"
	"github.com/irisnet/explorer/backend/utils"
)

type UpdateValidatorIcons struct{}

func (task UpdateValidatorIcons) Name() string {
	return "update_validator_icons"
}

func (task UpdateValidatorIcons) Start() {
	utils.RunTimer(conf.Get().Server.CronTimeValidatorIcons, utils.Sec, func() {
		if err := task.DoTask(); err != nil {
			logger.Error(fmt.Sprintf("%s fail", task.Name()), logger.String("err", err.Error()))
		} else {
			logger.Info(fmt.Sprintf("%s success", task.Name()))
		}
	})

}

func (task UpdateValidatorIcons) DoTask() error {
	err := new(service.ValidatorService).UpdateValidatorIcons()
	if err != nil {
		return err
	}

	return nil
}
