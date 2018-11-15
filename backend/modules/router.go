package modules

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/modules/rest"
	"log"
	"net/http"
	"net/http/pprof"
	"os"
	"time"
)

func registerApi(r *mux.Router) {
	routeRegistrars := []func(*mux.Router) error{
		rest.RegisterBlock,
		rest.RegisterTx,
		rest.RegisterAccount,
		rest.RegisterStake,
		rest.RegisterQueryVersion,
		rest.RegisterProposal,
		rest.RegisterNodes,
		rest.RegisterTextSearch,
		rest.RegisterPing,
		//reginsterPprof,
	}

	for _, routeRegistrar := range routeRegistrars {
		if err := routeRegistrar(r); err != nil {
			log.Fatal(err)
		}
	}
}

func reginsterPprof(r *mux.Router) error {
	r.HandleFunc("/debug/pprof", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)
	return nil
}

func NewAPIMux() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()
	registerApi(s)

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("../frontend/dist/"))))
	return r
}

func NewHandler(h http.Handler) http.Handler {
	handler := handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}))(h))

	handler = AddHeader(handler)
	handler = handlers.CompressHandler(handler)
	return handler
}

func AddHeader(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("FAUCET_URL", conf.Get().Server.HubFaucetUrl)
		w.Header().Add("CHAIN_ID", conf.Get().Server.ChainId)
		h.ServeHTTP(w, r)
	})
}

type Server struct {
	*http.Server
}

func NewServer() Server {
	router := NewAPIMux()
	handler := NewHandler(router)

	port := conf.Get().Server.ServerPort
	addr := fmt.Sprintf(":%d", port)

	log.Printf("Serving on %q", addr)

	server := &http.Server{
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		IdleTimeout:       10 * time.Second,
		Addr:              addr,
		Handler:           handler,
	}
	return Server{
		server,
	}
}
