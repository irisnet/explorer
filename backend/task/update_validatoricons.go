package task

import (
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/service"
)

type UpdateValidatorIcons struct{}

func (task UpdateValidatorIcons) Name() string {
	return "update_validator_icons"
}

func (task UpdateValidatorIcons) Start() {
	taskName := task.Name()
	timeInterval := conf.Get().Server.CronTimeValidatorIcons

	//utils.RunTimer(timeInterval, utils.Sec, func() {
	if err := tcService.runTask(taskName, timeInterval, task.DoTask); err != nil {
		logger.Error(err.Error())
	}
	//})
}

func (task UpdateValidatorIcons) DoTask(fn func(string) chan bool) error {
	defer HeartQuit(fn(task.Name()))
	err := new(service.ValidatorService).UpdateValidatorIcons()
	if err != nil {
		return err
	}

	return nil
}
