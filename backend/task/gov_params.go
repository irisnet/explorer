package task

import (
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
	utils.RunTimer(60, utils.Sec, func() {
		curModuleKv, err := lcd.GetAllGovModuleParam()
		if err != nil {
			logger.Error("UpdateGovParams task failed", logger.String("taskName", task.Name()), logger.String("errmsg", err.Error()))
			return
		}

		err = document.GovParams{}.UpdateCurrentModuleParamValue(curModuleKv)
		if err != nil {
			logger.Error("UpdateGovParams task failed", logger.String("taskName", task.Name()), logger.String("errmsg", err.Error()))
		}
	})
}
