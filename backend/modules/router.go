package modules

import (
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/modules/rest"
	"log"
	"net/http/pprof"
)

func Register(r *mux.Router) {
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
		//ReginsterPprof,
	}

	for _, routeRegistrar := range routeRegistrars {
		if err := routeRegistrar(r); err != nil {
			log.Fatal(err)
		}
	}
}

func ReginsterPprof(r *mux.Router) error {
	r.HandleFunc("/debug/pprof", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)
	return nil
}
