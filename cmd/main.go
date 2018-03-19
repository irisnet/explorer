package main

import (
	"os"
	"github.com/spf13/cobra"
	"github.com/tendermint/tmlibs/cli"
	"github.com/cosmos/cosmos-sdk/client/commands"
	"github.com/irisnet/iris-explorer/version"
)

// entry point for this binary
var (
	explorer = &cobra.Command{
		Use:   "iris-explorer",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

func main() {
	// disable sorting
	cobra.EnableCommandSorting = false

	prepareRestServerCommands()

	explorer.AddCommand(
		commands.InitCmd,
		restServerCmd,
		//syncCmd,
		version.VersionCmd,
	)

	// prepare and add flags
	executor := cli.PrepareMainCmd(explorer, "EX", os.ExpandEnv("$HOME/.iris-explorer"))
	executor.Execute()
}
