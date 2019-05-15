//go:generate statik -src=./swagger-ui
package lite

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/conf"
	_ "github.com/irisnet/explorer/backend/lcd/lite/statik"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/rakyll/statik/fs"
	"net/http"
	_ "net/http/pprof"
)

func RegisterSwaggerUI(r *mux.Router) {
	if conf.Get().Server.CurEnv == conf.EnvironmentDevelop || conf.Get().Server.CurEnv == conf.EnvironmentLocal || conf.Get().Server.CurEnv == conf.EnvironmentQa {
		statikFS, err := fs.New()
		if err != nil {
			panic(err)
		}

		staticServer := http.FileServer(statikFS)
		r.PathPrefix("/swagger-ui/").Handler(http.StripPrefix("/swagger-ui/", staticServer))
		logger.Info(fmt.Sprintf("enalbe swagger ui in %s environment.", conf.Get().Server.CurEnv))

	} else {
		logger.Info(fmt.Sprintf("disable swagger ui in %s environment.", conf.Get().Server.CurEnv))
	}
}
