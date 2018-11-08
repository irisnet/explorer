package modules

import (
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/modules/rest"
	"github.com/irisnet/explorer/backend/version"
	"log"
)

func Register(r *mux.Router) {
	routeRegistrars := []func(*mux.Router) error{
		rest.RegisterBlock,
		rest.RegisterTx,
		rest.RegisterAccount,
		rest.RegisterStake,
		version.RegisterQueryVersion,
		rest.RegisterChain,
		rest.RegisterProposal,
		rest.RegisterQueryIp,
		rest.RegisterNodes,
		rest.RegisterTextSearch,
	}

	for _, routeRegistrar := range routeRegistrars {
		if err := routeRegistrar(r); err != nil {
			log.Fatal(err)
		}
	}
}
