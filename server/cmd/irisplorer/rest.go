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
	"github.com/irisnet/irisplorer.io/server/modules/rpc"
	"github.com/irisnet/irisplorer.io/server/modules/tools"
	"github.com/irisnet/irisplorer.io/server/modules/store"

	_ "github.com/cosmos/cosmos-sdk/modules/auth"
	_ "github.com/cosmos/cosmos-sdk/modules/base"
	_ "github.com/cosmos/cosmos-sdk/modules/coin"
	_ "github.com/cosmos/cosmos-sdk/modules/nonce"
	"github.com/irisnet/irisplorer.io/server/modules/sync"
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

	restServerCmd.PersistentFlags().Int64(tools.MaxConnectionNum, 500, "max connection amount of rpc client")
	restServerCmd.PersistentFlags().Int64(tools.InitConnectionNum, 100, "init connection amount of rpc client")
	restServerCmd.PersistentFlags().String(store.MgoUrl, "localhost:27017", "url of MongoDB")
	restServerCmd.PersistentFlags().String(tools.SyncCron, "@every 5s", "Cron Task")

	restServerCmd.PersistentFlags().IntP(tools.FlagPort, "p", 7998, "port to run the server on")
	restServerCmd.PersistentFlags().Bool(tools.WithOnlyRestServer, false, "start rest server without sync process (default false)")
}

func AddRoutes(r *mux.Router) {
	routeRegistrars := []func(*mux.Router) error{
		rpc.RegisterBlock,
		rpc.RegisterTx,
		rpc.RegisterAccount,
		rpc.RegisterStake,
	}

	for _, routeRegistrar := range routeRegistrars {
		if err := routeRegistrar(r); err != nil {
			log.Fatal(err)
		}
	}
}

func prepareSyncServer(cmd *cobra.Command, args []string) {
	if !viper.GetBool(tools.WithOnlyRestServer) {
		sync.Start()
	}else {
		sync.InitServer()
	}
}

func cmdRestServer(cmd *cobra.Command, args []string) error {

	prepareSyncServer(cmd, args)

	router := mux.NewRouter()
	// latest
	AddRoutes(router)

	addr := fmt.Sprintf(":%d", viper.GetInt(tools.FlagPort))

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
