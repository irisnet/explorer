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
	utils.RunTimer(conf.Get().Server.CronTimeGovParams, utils.Sec, func() {
		if err := task.DoTask(); err != nil {
			logger.Error(fmt.Sprintf("%s fail", task.Name()), logger.String("err", err.Error()))
		} else {
			logger.Info(fmt.Sprintf("%s success", task.Name()))
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
