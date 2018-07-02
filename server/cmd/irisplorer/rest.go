package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"github.com/irisnet/explorer/server/modules/rest"
)

var (
	restServerCmd = &cobra.Command{
		Use:  "rest-server",
		Long: `presents  a nice (not raw hex) interface to the irisnet blockchain structure.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmdRestServer(cmd, args)
		},
	}
)

func prepareRestServerCommands() {
	restServerCmd.PersistentFlags().IntP("port", "p", 8888, "port to run the server on")
}

func AddRoutes(r *mux.Router) {
	routeRegistrars := []func(*mux.Router) error{
		rest.RegisterBlock,
		rest.RegisterTx,
		rest.RegisterAccount,
		rest.RegisterStake,
	}

	for _, routeRegistrar := range routeRegistrars {
		if err := routeRegistrar(r); err != nil {
			log.Fatal(err)
		}
	}
}


func cmdRestServer(cmd *cobra.Command, args []string) error {

	router := mux.NewRouter()
	// latest
	AddRoutes(router)

	addr := fmt.Sprintf(":%d", viper.GetInt("port"))

	log.Printf("Serving on %q", addr)

	// loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	//return http.ListenAndServe(addr, router)
	//TODO 生产环境需要借助nginx配置
	return http.ListenAndServe(addr,
		handlers.LoggingHandler(os.Stdout, handlers.CORS(
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}))(router)))
}
