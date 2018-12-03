package rest

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/rest/controller"
	"github.com/irisnet/irishub-sync/module/logger"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

type ApiServer struct {
	*http.Server
}

func NewApiServer() *ApiServer {
	router := NewAPIMux()
	handler := newHandler(router)

	port := conf.Get().Server.ServerPort
	addr := fmt.Sprintf(":%d", port)

	logger.Info("server will start", logger.String("addr", addr))

	server := &http.Server{
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
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
		logger.Info("pprof will Serve on 6060")
		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
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
	}

	for _, routeRegistrar := range routeRegistrars {
		if err := routeRegistrar(r); err != nil {
			log.Fatal(err)
		}
	}
}

func NewAPIMux() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()
	registerApi(s)

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("../frontend/dist/"))))
	return r
}

func newHandler(h http.Handler) http.Handler {
	//handler := handlers.LoggingHandler(os.Stdout, handlers.CORS(
	//	handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
	//	handlers.AllowedOrigins([]string{"*"}),
	//	handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}))(h))

	handler := addHeader(h)
	handler = handlers.CompressHandler(handler)
	return handler
}

func addHeader(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("FAUCET_URL", conf.Get().Server.HubFaucetUrl)
		w.Header().Add("CHAIN_ID", conf.Get().Server.ChainId)
		h.ServeHTTP(w, r)
	})
}
