package controller

import (
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/service"
	"github.com/irisnet/explorer/backend/task"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/vo"
)

func RegisterCronTask(r *mux.Router) error {
	fns := []func(*mux.Router) error{
		registerCronTaskAssetGateways,
		registerCronTaskAssetTokens,
		registerCronTaskGovParams,
		registerCronTaskTxNumByDay,
		registerCronTaskValidators,
		registerCronTaskValidatorIcons,
		registerUpdateBlackValidators,
	}

	for _, fn := range fns {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

func registerCronTaskAssetGateways(r *mux.Router) error {
	doApi(r, types.UrlRegisterDoCronTaskAssetGateways, "PUT", func(request vo.IrisReq) interface{} {
		cronTask := task.UpdateAssetGateways{}
		if err := cronTask.DoTask(); err != nil {
			return err
		} else {
			return types.CodeSuccess
		}
	})

	return nil
}

func registerCronTaskAssetTokens(r *mux.Router) error {
	doApi(r, types.UrlRegisterDoCronTaskAssetTokens, "PUT", func(request vo.IrisReq) interface{} {
		cronTask := task.UpdateAssetTokens{}
		if err := cronTask.DoTask(); err != nil {
			return err
		} else {
			return types.CodeSuccess
		}
	})

	return nil
}

func registerCronTaskGovParams(r *mux.Router) error {
	doApi(r, types.UrlRegisterDoCronTaskGovParams, "PUT", func(request vo.IrisReq) interface{} {
		cronTask := task.UpdateGovParams{}
		if err := cronTask.DoTask(); err != nil {
			return err
		} else {
			return types.CodeSuccess
		}
	})

	return nil
}

func registerCronTaskTxNumByDay(r *mux.Router) error {
	doApi(r, types.UrlRegisterDoCronTaskTxNumByDay, "PUT", func(request vo.IrisReq) interface{} {
		cronTask := task.TxNumGroupByDayTask{}
		if err := cronTask.DoTask(); err != nil {
			return err
		} else {
			return types.CodeSuccess
		}
	})

	return nil
}

func registerCronTaskValidators(r *mux.Router) error {
	doApi(r, types.UrlRegisterDoCronTaskValidators, "PUT", func(request vo.IrisReq) interface{} {
		cronTask := task.UpdateValidator{}
		if err := cronTask.DoTask(); err != nil {
			return err
		} else {
			return types.CodeSuccess
		}
	})

	return nil
}

func registerCronTaskValidatorIcons(r *mux.Router) error {
	doApi(r, types.UrlRegisterDoCronTaskValidatorIcons, "PUT", func(request vo.IrisReq) interface{} {
		cronTask := task.UpdateValidatorIcons{}
		if err := cronTask.DoTask(); err != nil {
			return err
		} else {
			return types.CodeSuccess
		}
	})

	return nil
}

func registerUpdateBlackValidators(r *mux.Router) error {
	doApi(r, types.UrlRegisterTaskUpdateBlackValidators, "PUT", func(request vo.IrisReq) interface{} {
		baseService := service.BaseService{}
		baseService.ReloadBlackValidators()
		return types.CodeSuccess
	})

	return nil
}
