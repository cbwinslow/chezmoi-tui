package commands

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"chezmoi-tui/internal/chezmoi"
	"chezmoi-tui/pkg/root"
)

var addCmd = &cobra.Command{
	Use:   "add [targets...]",
	Short: "Add targets to the source state",
	Long:  `Add targets to the source state. If any target is already in the source state, then its source state is replaced with its current state in the destination directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := chezmoi.New()
		if err != nil {
			log.Fatalf("Failed to initialize chezmoi: %v", err)
		}

		output, err := c.Add(args...)
		if err != nil {
			log.Fatalf("Failed to add: %v", err)
		}

		fmt.Print(output)
	},
}

func init() {
	root.RootCmd.AddCommand(addCmd)
}
