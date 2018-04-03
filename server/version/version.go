package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

const Version = "0.1.0"

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("v%s\n", Version)
	},
}
