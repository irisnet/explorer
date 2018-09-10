package main

import (
	"fmt"
	"github.com/irisnet/explorer/backend/task"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/modules/rest"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/explorer/backend/version"
)

func AddRoutes(r *mux.Router) {
	routeRegistrars := []func(*mux.Router) error{
		rest.RegisterBlock,
		rest.RegisterTx,
		rest.RegisterAccount,
		rest.RegisterStake,
		version.RegisterQueryVersion,
		rest.RegisterChain,
		rest.RegisterProposal,
	}

	for _, routeRegistrar := range routeRegistrars {
		if err := routeRegistrar(r); err != nil {
			log.Fatal(err)
		}
	}
}

var FAUCET_URL string = utils.GetEnv("FAUCET_URL", "http://dev.faucet.irisplorer.io")
var CHAIN_ID string = utils.GetEnv("CHAIN_ID", "rainbow-dev")

func main() {
	go task.Start()
	router := mux.NewRouter()
	// latest
	AddRoutes(router)

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("../frontend/dist/"))))

	loggedRouter := handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}))(router))

	loggedRouter = handlers.LoggingHandler(os.Stdout, AddHeader(loggedRouter))
	loggedRouter = handlers.CompressHandler(loggedRouter)

	port := utils.GetEnv("PORT", "8080")
	addr := fmt.Sprintf(":%s", port)

	log.Printf("Serving on %q", addr)

	// loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	//return http.ListenAndServe(addr, router)
	//TODO 生产环境需要借助nginx配置
	http.ListenAndServe(addr, loggedRouter)
}

func AddHeader(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("FAUCET_URL", FAUCET_URL)
		w.Header().Add("CHAIN_ID", CHAIN_ID)
		h.ServeHTTP(w, r)
	})
}
