package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"chezmoi-tui/pkg/root"
)

// Version is the application version
var Version = "0.1.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `Print the version number of chezmoi-tui.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("chezmoi-tui version %s\\n", Version)
	},
}

func init() {
	root.RootCmd.AddCommand(versionCmd)
}