package main

import (
	"github.com/irisnet/irisplorer.io/server/modules/store"
	"github.com/irisnet/irisplorer.io/server/modules/sync"
	"github.com/irisnet/irisplorer.io/server/modules/tools"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	cmn "github.com/tendermint/tmlibs/common"
	"log"
)

var (
	syncCmd = &cobra.Command{
		Use:  "sync",
		Long: `sync blockchain data from irisnet`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return StartWatch(cmd, args)
		},
	}
)

func prepareSyncCommands() {

	SyncPk := flag.NewFlagSet("", flag.ContinueOnError)
	SyncPk.Int64(tools.MaxConnectionNum, 100, "max amount of rest client")
	SyncPk.Int64(tools.InitConnectionNum, 50, "init amount of rest client")
	SyncPk.String(store.MgoUrl, "localhost:27017", "url of MongoDB")
	SyncPk.String(tools.SyncCron, "@every 5s", "Cron Task")
	syncCmd.Flags().AddFlagSet(SyncPk)
}

func StartWatch(cmd *cobra.Command, args []string) error {
	sync.Start()
	// Sleep forever and then...
	cmn.TrapSignal(func() {
		log.Printf("sync process exit")
	})

	return nil
}
