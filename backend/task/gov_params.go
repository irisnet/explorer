package task

import (
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/utils"
	"encoding/json"
)

type UpdateGovParams struct{}

func (task UpdateGovParams) Name() string {
	return "update_gov_params"
}
func (task UpdateGovParams) Start() {
	taskName := task.Name()
	timeInterval := conf.Get().Server.CronTimeGovParams

	utils.RunTimer(timeInterval, utils.Sec, func() {
		if err := tcService.runTask(taskName, timeInterval, task.DoTask); err != nil {
			logger.Error(err.Error())
		}
	})
}

func (task UpdateGovParams) DoTask(fn func(string) chan bool) error {
	stop := fn(task.Name())
	defer HeartQuit(stop)
	//curModuleKv := lcd.GetAllGovModuleParam()

	params, err := lcd.GetGovModuleParam("")
	if err != nil {
		return err
	}
	curModuleKv := make(map[string]interface{}, len(params))

	for _, one := range params {
		data := map[string]interface{}{}
		if err := json.Unmarshal([]byte(one.Value), &data); err == nil {
			for k, v := range data {
				curModuleKv[k] = v
			}
		}

	}
	if err := (document.GovParams{}).UpdateCurrentModuleParamValue(curModuleKv); err != nil {
		return err
	}

	return nil
}
