package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/modules"
	"github.com/irisnet/explorer/backend/task"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	go task.Start()
	router := mux.NewRouter().
		PathPrefix("/api").Subrouter()

	modules.Register(router)

	loggedRouter := handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}))(router))

	//loggedRouter = handlers.LoggingHandler(os.Stdout, )
	loggedRouter = AddHeader(loggedRouter)
	loggedRouter = handlers.CompressHandler(loggedRouter)
	port := conf.Get().Server.ServerPort
	addr := fmt.Sprintf(":%d", port)

	log.Printf("Serving on %q", addr)

	server := &http.Server{
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		IdleTimeout:       10 * time.Second,
		Addr:              addr,
		Handler:           loggedRouter,
	}
	server.ListenAndServe()
}

func AddHeader(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("FAUCET_URL", conf.Get().Server.HubFaucetUrl)
		w.Header().Add("CHAIN_ID", conf.Get().Server.ChainId)
		h.ServeHTTP(w, r)
	})
}
