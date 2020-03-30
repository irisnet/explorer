package controller

import (
	"time"

	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/vo"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/lcd"
)

func RegisterTextSearch(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryText,
		registerQuerySysDate,
		registerQueryEnvConfig,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

//type Common struct {
//	*service.CommonService
//}
//
//var common = Common{
//	service.Get(service.Common).(*service.CommonService),
//}

// @Summary search text
// @Description search text
// @Tags common
// @Accept  json
// @Produce  json
// @Param   text   path   string  true    "text"
// @Success 200 {object} vo.ResultVo	"success"
// @Router /api/search/{text} [get]
func registerQueryText(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryText, "GET", func(request vo.IrisReq) interface{} {
		common.SetTid(request.TraceId)
		text := Var(request, "text")

		result := common.QueryText(text)
		return result
	})

	return nil
}

// @Summary sysdate
// @Description get sysdate
// @Tags common
// @Accept  json
// @Produce  json
// @Success 200 {object} int64	"success"
// @Router /api/sysdate [get]
func registerQuerySysDate(r *mux.Router) error {
	doApi(r, types.UrlRegisterQuerySysDate, "GET", func(request vo.IrisReq) interface{} {
		return time.Now().Unix()
	})

	return nil
}

// @Summary config
// @Description get config
// @Tags common
// @Accept  json
// @Produce  json
// @Success 200 {object} vo.EnvConfig	"success"
// @Router /api/config [get]
func registerQueryEnvConfig(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryConfig, "GET", func(request vo.IrisReq) interface{} {
		getconfig := func(dconfigs []document.Config) (ret []vo.ConfigVo) {
			for _, val := range dconfigs {
				item := vo.ConfigVo{
					NetworkName: val.NetworkName,
					Env:         val.Env,
					Host:        val.Host,
					ChainId:     val.ChainId,
					ShowFaucet:  val.ShowFaucet,
					UmengId:     val.UmengId,
				}
				if nodeinfo, err := lcd.NodeInfo(val.EnvLcd); err == nil {
					item.TendermintVersion = nodeinfo.Version
				}
				if version, err := lcd.NodeVersion(val.EnvLcd); err == nil {
					item.NodeVersion = version
				}
				ret = append(ret, item)
			}
			return ret
		}
		var envConf = vo.EnvConfig{
			CurEnv:  conf.Get().Server.CurEnv,
			ChainId: conf.Get().Hub.ChainId,
			Configs: getconfig(common.GetConfig()),
		}
		return envConf
	})

	return nil
}
