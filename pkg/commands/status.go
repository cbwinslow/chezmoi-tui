package commands

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"chezmoi-tui/internal/chezmoi"
	"chezmoi-tui/pkg/root"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the status of targets",
	Long:  `Show the status of targets in a format similar to git status.`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := chezmoi.New()
		if err != nil {
			log.Fatalf("Failed to initialize chezmoi: %v", err)
		}

		output, err := c.Status()
		if err != nil {
			log.Fatalf("Failed to get status: %v", err)
		}

		fmt.Print(output)
	},
}

func init() {
	root.RootCmd.AddCommand(statusCmd)
}
