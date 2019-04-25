package rest

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/rest/controller"
	"github.com/irisnet/explorer/backend/rest/filter"
)

type ApiServer struct {
	*http.Server
}

func NewApiServer() *ApiServer {
	router := NewAPIMux()
	handler := handlers.CompressHandler(router)
	handler = AddHeader(handler)

	port := conf.Get().Server.ServerPort
	addr := fmt.Sprintf(":%d", port)

	logger.Info("server will start", logger.String("addr", addr))

	server := &http.Server{
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      60 * time.Second,
		ReadHeaderTimeout: 60 * time.Second,
		IdleTimeout:       10 * time.Second,
		Addr:              addr,
		Handler:           handler,
	}

	return &ApiServer{
		server,
	}
}

func (s *ApiServer) Start() error {
	go func() {
		logger.Info("pprof will Serve on 6061")
		if err := http.ListenAndServe("localhost:6061", nil); err != nil {
			logger.Fatal("ListenAndServe Failed: ", logger.Any("err", err))
		}
	}()
	return s.ListenAndServe()
}

func registerApi(r *mux.Router) {
	routeRegistrars := []func(*mux.Router) error{
		controller.RegisterBlock,
		controller.RegisterTx,
		controller.RegisterAccount,
		controller.RegisterStake,
		controller.RegisterQueryVersion,
		controller.RegisterProposal,
		controller.RegisterNodes,
		controller.RegisterTextSearch,
		controller.RegisterPing,
		controller.RegisterHome,
	}

	for _, routeRegistrar := range routeRegistrars {
		if err := routeRegistrar(r); err != nil {
			log.Fatal(err)
		}
	}
}

func registerFilters() {
	filter.Register(filter.LogPreFilter{})
	filter.Register(filter.LogPostFilter{})
	filter.Register(filter.FaucetLimitPreFilter{})
	filter.Register(filter.FaucetLimitPostFilter{})
}

func NewAPIMux() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()
	registerApi(s)
	registerFilters()

	if conf.Get().Server.CurEnv == conf.EnvironmentDevelop || conf.Get().Server.CurEnv == conf.EnvironmentLocal || conf.Get().Server.CurEnv == conf.EnvironmentQa {
		logger.Info(fmt.Sprintf("enalbe swagger ui in %s environment.", conf.Get().Server.CurEnv))
		r.PathPrefix("/swagger-ui/").Handler(http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./swagger-ui"))))
	} else {
		logger.Info(fmt.Sprintf("disable swagger ui in %s environment.", conf.Get().Server.CurEnv))
	}

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("../frontend/dist/"))))
	return r
}

func AddHeader(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
}
