package main

import (
	"os"
	"testing"
	"github.com/irisnet/iris-explorer/modules/tools"
	"time"
	"github.com/spf13/viper"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/gorilla/handlers"
)

func TestMain2(t *testing.T){
	viper.Set("node","tcp://47.104.155.125:46757")
	//viper.Set("node","tcp://116.62.62.39:11657")
	viper.Set(tools.InitConnectionNum,50)
	viper.Set(tools.MaxConnectionNum,100)
	viper.Set("cron","@every 3s")
	tools.Init()
	StartWatch(nil,nil)

	time.Sleep(1 * time.Hour)
}

func TestRunRestServer(t *testing.T){
	viper.Set("node","tcp://47.104.155.125:46757")
	viper.Set(tools.InitConnectionNum,5)
	viper.Set(tools.MaxConnectionNum,10)

	tools.Init()

	router := mux.NewRouter()
	// latest
	AddRoutes(router)
	http.ListenAndServe(":8998",
		handlers.LoggingHandler(os.Stdout, handlers.CORS(
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}))(router)))

	time.Sleep(1 * time.Hour)
}
