package main

import (
	"os"
	"github.com/spf13/cobra"
	"github.com/tendermint/tmlibs/cli"
	"github.com/cosmos/cosmos-sdk/client/commands"
	"github.com/irisnet/irisplorer.io/server/version"
)

// entry point for this binary
var (
	ExplorerCmd = &cobra.Command{
		Use:   "explorer",
		Short: "explorer for IRIS Hub (a regional Cosmos Hub with a powerful iService infrastructure)",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

func prepareMainCmd() {
	commands.AddBasicFlags(ExplorerCmd)
}

func main() {
	// disable sorting
	cobra.EnableCommandSorting = false
	prepareMainCmd()
	prepareRestServerCommands()
	prepareSyncCommands()

	ExplorerCmd.AddCommand(
		commands.InitCmd,
		restServerCmd,
		syncCmd,
		version.VersionCmd,
	)

	// prepare and add flags
	executor := cli.PrepareMainCmd(ExplorerCmd, "EX", os.ExpandEnv("$HOME/.irisplorer.io"))
	executor.Execute()
}
