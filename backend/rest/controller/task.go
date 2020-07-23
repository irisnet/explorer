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
		//registerCronTaskAssetGateways,
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

//// @Summary update asset_gateways
//// @Description update asset gateways
//// @Tags task
//// @Accept  json
//// @Produce  json
//// @Success 200 {object} types.BizCode  "success"
//// @Router /api/task/asset_gateways [put]
//func registerCronTaskAssetGateways(r *mux.Router) error {
//	doApi(r, types.UrlRegisterDoCronTaskAssetGateways, "PUT", func(request vo.IrisReq) interface{} {
//		cronTask := task.UpdateAssetGateways{}
//		if err := cronTask.DoTask(task.HeartBeat); err != nil {
//			return err
//		} else {
//			return types.CodeSuccess
//		}
//	})
//
//	return nil
//}

// @Summary update asset_tokens
// @Description update asset tokens
// @Tags task
// @Accept  json
// @Produce  json
// @Success 200 {object} types.BizCode  "success"
// @Router /api/task/asset_tokens [put]
func registerCronTaskAssetTokens(r *mux.Router) error {
	doApi(r, types.UrlRegisterDoCronTaskAssetTokens, "PUT", func(request vo.IrisReq) interface{} {
		cronTask := task.UpdateAssetTokens{}
		if err := cronTask.DoTask(task.HeartBeat); err != nil {
			return err
		} else {
			return types.CodeSuccess
		}
	})

	return nil
}

// @Summary update gov_params
// @Description update gov params
// @Tags task
// @Accept  json
// @Produce  json
// @Success 200 {object} types.BizCode  "success"
// @Router /api/task/gov_params [put]
func registerCronTaskGovParams(r *mux.Router) error {
	doApi(r, types.UrlRegisterDoCronTaskGovParams, "PUT", func(request vo.IrisReq) interface{} {
		cronTask := task.UpdateGovParams{}
		if err := cronTask.DoTask(task.HeartBeat); err != nil {
			return err
		} else {
			return types.CodeSuccess
		}
	})

	return nil
}

// @Summary update tx_num_by_day
// @Description update tx_num_by_day
// @Tags task
// @Accept  json
// @Produce  json
// @Success 200 {object} types.BizCode  "success"
// @Router /api/task/tx_num_by_day [put]
func registerCronTaskTxNumByDay(r *mux.Router) error {
	doApi(r, types.UrlRegisterDoCronTaskTxNumByDay, "PUT", func(request vo.IrisReq) interface{} {
		cronTask := task.TxNumGroupByDayTask{}
		if err := cronTask.DoTask(task.HeartBeat); err != nil {
			return err
		} else {
			return types.CodeSuccess
		}
	})

	return nil
}

// @Summary update validators
// @Description update validators
// @Tags task
// @Accept  json
// @Produce  json
// @Success 200 {object} types.BizCode  "success"
// @Router /api/task/validators [put]
func registerCronTaskValidators(r *mux.Router) error {
	doApi(r, types.UrlRegisterDoCronTaskValidators, "PUT", func(request vo.IrisReq) interface{} {
		cronTask := task.UpdateValidator{}
		if err := cronTask.DoTask(task.HeartBeat); err != nil {
			return err
		} else {
			return types.CodeSuccess
		}
	})

	return nil
}

// @Summary update validator icons
// @Description update validator icons
// @Tags task
// @Accept  json
// @Produce  json
// @Success 200 {object} types.BizCode  "success"
// @Router /api/task/validator_icons [put]
func registerCronTaskValidatorIcons(r *mux.Router) error {
	doApi(r, types.UrlRegisterDoCronTaskValidatorIcons, "PUT", func(request vo.IrisReq) interface{} {
		cronTask := task.UpdateValidatorIcons{}
		if err := cronTask.DoTask(task.HeartBeat); err != nil {
			return err
		} else {
			return types.CodeSuccess
		}
	})

	return nil
}

// @Summary update black validators
// @Description update black validators
// @Tags task
// @Accept  json
// @Produce  json
// @Success 200 {object} types.BizCode  "success"
// @Router /api/task/black_validators [put]
func registerUpdateBlackValidators(r *mux.Router) error {
	doApi(r, types.UrlRegisterTaskUpdateBlackValidators, "PUT", func(request vo.IrisReq) interface{} {
		baseService := service.BaseService{}
		baseService.ReloadBlackValidators()
		return types.CodeSuccess
	})

	return nil
}
