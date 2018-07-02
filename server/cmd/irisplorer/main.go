package main

import (
	"github.com/irisnet/explorer/server/version"
	"github.com/spf13/cobra"
	"github.com/tendermint/tmlibs/cli"
	"os"
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


func main() {
	// disable sorting
	cobra.EnableCommandSorting = false
	prepareRestServerCommands()

	ExplorerCmd.AddCommand(
		restServerCmd,
		version.VersionCmd,
	)

	// prepare and add flags
	executor := cli.PrepareMainCmd(ExplorerCmd, "EX", os.ExpandEnv("$HOME/.explorer"))
	executor.Execute()
}
