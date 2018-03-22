package main

import (
	"testing"
	"github.com/irisnet/iris-explorer/modules/sync"
	"github.com/irisnet/iris-explorer/modules/tools"
	"time"
	"github.com/spf13/viper"
)

func TestMain2(t *testing.T){
	viper.Set("client-max-conn",5)
	viper.Set("node","tcp://47.104.155.125:46757")
	viper.Set("with-sync",false)
	tools.Init()
	sync.StartWatch()

	time.Sleep(1 * time.Hour)
}
