package main

import (
	"fmt"
	"github.com/irisnet/explorer/backend/task"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/modules"
	"github.com/irisnet/explorer/backend/utils"
	_ "net/http/pprof"
)

var faucetUrl = utils.GetEnv("FAUCET_URL", "http://dev.faucet.irisplorer.io")
var chainId = utils.GetEnv("CHAIN_ID", "rainbow-dev")

func main() {
	go task.Start()
	router := mux.NewRouter()
	r := router.PathPrefix("/api").Subrouter()

	modules.Register(r)

	loggedRouter := handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}))(r))

	loggedRouter = handlers.LoggingHandler(os.Stdout, AddHeader(loggedRouter))
	loggedRouter = handlers.CompressHandler(loggedRouter)

	port := utils.GetEnv("PORT", "8080")
	addr := fmt.Sprintf(":%s", port)

	log.Printf("Serving on %q", addr)

	http.Handle("/", loggedRouter)
	http.ListenAndServe(addr, nil)
}

func AddHeader(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("FAUCET_URL", faucetUrl)
		w.Header().Add("CHAIN_ID", chainId)
		h.ServeHTTP(w, r)
	})
}
