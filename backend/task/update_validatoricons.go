package task

import (
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/explorer/backend/service"
	"github.com/irisnet/explorer/backend/logger"
)

type UpdateValidatorIcons struct{}

func (task UpdateValidatorIcons) Name() string {
	return "update_validator_icons"
}
func (task UpdateValidatorIcons) Start() {
	utils.RunTimer(1, utils.Day, func() {
		err := new(service.ValidatorService).UpdateValidatorIcons()
		if err != nil {
			logger.Error("UpdateValidatorIcons task failed", logger.String("taskName", task.Name()), logger.String("errmsg", err.Error()))
			return
		}
		logger.Info("UpdateValidatorIcons task is OK.")
	})

}
