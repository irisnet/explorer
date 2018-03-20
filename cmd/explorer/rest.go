package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/gorilla/handlers"
	"github.com/spf13/viper"
	"os"
	"github.com/irisnet/iris-explorer/modules/rpc"
	"github.com/irisnet/iris-explorer/modules/tools"
	"github.com/irisnet/iris-explorer/modules/sync"
)

const (
	FlagPort = "port"
	MgoUrl   = "mgo-url"
	ClientMaxCon   = "client-max-conn"
	WithSyncServer   = "with-sync-server"
)

var (
	restServerCmd = &cobra.Command{
		Use:  "rest-server",
		Long: `presents  a nice (not raw hex) interface to the gaia blockchain structure.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmdRestServer(cmd, args)
		},
	}
)

func prepareRestServerCommands() {
	restServerCmd.PersistentFlags().Int(ClientMaxCon, 10, "amount of rpc client")
	restServerCmd.PersistentFlags().String(MgoUrl, "localhost:27017", "url of MongoDB")
	restServerCmd.PersistentFlags().IntP(FlagPort, "p", 8998, "port to run the server on")
	restServerCmd.PersistentFlags().Bool(WithSyncServer,true,"start sync server")
}

func AddRoutes(r *mux.Router) {
	routeRegistrars := []func(*mux.Router) error{
		rpc.RegisterBlock,
	}

	for _, routeRegistrar := range routeRegistrars {
		if err := routeRegistrar(r); err != nil {
			log.Fatal(err)
		}
	}
}

func prepareSyncServer(){
	if viper.GetBool(WithSyncServer) {
		sync.StartWatch()
	}
}

func cmdRestServer(cmd *cobra.Command, args []string) error {
	//初始化连接池
	tools.Init()

	prepareSyncServer()

	router := mux.NewRouter()
	// latest
	AddRoutes(router)


	addr := fmt.Sprintf(":%d", viper.GetInt(FlagPort))

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
