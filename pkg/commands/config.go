package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"chezmoi-tui/pkg/root"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration for chezmoi-tui",
	Long:  `Manage configuration for chezmoi-tui`,
}

var generateConfigCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a default configuration file",
	Long:  `Generate a default configuration file for chezmoi-tui`,
	Run: func(cmd *cobra.Command, args []string) {
		// Define the default config template
		configTemplate := `# Chezmoi TUI Configuration
# This file configures the enhanced TUI and CLI for chezmoi

# Theme settings
theme:
  primary_color: "#1793d1"
  secondary_color: "#0366d6"
  success_color: "#28a745"
  warning_color: "#ffc107"
  danger_color: "#dc3545"

# TUI settings
tui:
  show_help: true
  refresh_interval: 5 # seconds
  max_file_display: 100

# CLI settings
cli:
  verbose: false
  color: true
  pager: "less"

# Integration settings
integration:
  chezmoi_binary_path: ""
  timeout: 30 # seconds
`

		// Check if config file already exists
		configPath := os.Getenv("HOME") + "/.config/chezmoi-tui/config.yaml"
		if _, err := os.Stat(configPath); err == nil {
			force, _ := cmd.Flags().GetBool("force")
			if !force {
				log.Fatalf("Config file already exists at %s. Use --force to overwrite.", configPath)
			}
		}

		// Create directory if it doesn't exist
		dir := os.Getenv("HOME") + "/.config/chezmoi-tui"
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatalf("Failed to create config directory: %v", err)
		}

		// Write the config file
		file, err := os.Create(configPath)
		if err != nil {
			log.Fatalf("Failed to create config file: %v", err)
		}
		defer file.Close()

		_, err = file.WriteString(configTemplate)
		if err != nil {
			log.Fatalf("Failed to write config file: %v", err)
		}

		fmt.Printf("Configuration file generated at: %s\n", configPath)
	},
}

func init() {
	// Add flags to the generate command
	generateConfigCmd.Flags().BoolP("force", "f", false, "Force overwrite existing config file")

	// Add subcommands to config command
	configCmd.AddCommand(generateConfigCmd)

	// Add the config command to the root
	root.RootCmd.AddCommand(configCmd)
}
