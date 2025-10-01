package root

import (
	"github.com/spf13/cobra"
)

// Version is the application version
var Version = "0.1.0"

// RootCmd is the root command for the application
var RootCmd = &cobra.Command{
	Use:     "chezmoi-tui",
	Short:   "Enhanced TUI and CLI for chezmoi dotfile management",
	Long:    `An enhanced Terminal User Interface and Command Line Interface for managing your dotfiles with chezmoi.`,
	Version: Version,
}
