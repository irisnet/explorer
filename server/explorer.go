package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/irisnet/explorer/server/modules/rest"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/irisnet/explorer/server/utils"
	"github.com/irisnet/explorer/server/version"
)

func AddRoutes(r *mux.Router) {
	routeRegistrars := []func(*mux.Router) error{
		rest.RegisterBlock,
		rest.RegisterTx,
		rest.RegisterAccount,
		rest.RegisterStake,
		version.RegisterQueryVersion,
	}

	for _, routeRegistrar := range routeRegistrars {
		if err := routeRegistrar(r); err != nil {
			log.Fatal(err)
		}
	}
}

func main()  {
	router := mux.NewRouter()
	// latest
	AddRoutes(router)

	router.PathPrefix("/").Handler(http.StripPrefix("/",http.FileServer(http.Dir("../frontend/dist/"))))

	port := utils.GetEnv("PORT","8080")
	addr := fmt.Sprintf(":%s", port)


	log.Printf("Serving on %q", addr)

	// loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	//return http.ListenAndServe(addr, router)
	//TODO 生产环境需要借助nginx配置
	http.ListenAndServe(addr,
		handlers.LoggingHandler(os.Stdout, handlers.CORS(
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}))(router)))
}
