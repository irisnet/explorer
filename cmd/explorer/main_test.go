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
	tools.Init()
	sync.StartWatch()

	time.Sleep(1 * time.Hour)
}

//func TestLoadTx(t *testing.T){
//
//	bin := []byte{22,3,1,5,112,97,110,103,117,0,0,0,0,0,0,0,0,105,0,0,0,31,1,1,0,1,4,115,105,103,115,1,20,246,7,46,65,200,94,227,39,181,172,33,215,214,241,93,28,78,86,242,64,32,1,1,0,1,4,115,105,103,115,1,20,246,7,46,65,200,94,227,39,181,172,33,215,214,241,93,28,78,86,242,64,1,1,1,4,105,114,105,115,0,0,0,0,0,0,0}
//	txb, _ := sdk.LoadTx(bin)
//	log.Printf(" finish %s",txb)
//	//sdk.TxMapper.RegisterImplementation(base.ChainTx{}, base.TypeChainTx, base.ByteChainTx)
//}
