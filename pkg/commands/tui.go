package commands

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"chezmoi-tui/pkg/root"
	"chezmoi-tui/ui"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Launch the Terminal User Interface",
	Long:  `Launch the enhanced Terminal User Interface for managing dotfiles.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Launching Chezmoi TUI...")
		// TUI logic will be implemented here
		err := ui.RunTUI()
		if err != nil {
			log.Fatalf("Failed to run TUI: %v", err)
		}
	},
}

func init() {
	root.RootCmd.AddCommand(tuiCmd)
}