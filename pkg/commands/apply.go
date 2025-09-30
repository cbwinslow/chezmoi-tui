package commands

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"chezmoi-tui/internal/chezmoi"
	"chezmoi-tui/pkg/root"
)

var applyCmd = &cobra.Command{
	Use:   "apply [targets...]",
	Short: "Update the destination directory to match the target state",
	Long:  `Update the destination directory to match the target state, applying any changes.`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := chezmoi.New()
		if err != nil {
			log.Fatalf("Failed to initialize chezmoi: %v", err)
		}

		output, err := c.Apply(args...)
		if err != nil {
			log.Fatalf("Failed to apply: %v", err)
		}

		fmt.Print(output)
	},
}

func init() {
	root.RootCmd.AddCommand(applyCmd)
}