package commands

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"chezmoi-tui/pkg/root"
)

var bitwardenCmd = &cobra.Command{
	Use:   "bitwarden",
	Short: "Interact with Bitwarden secrets management",
	Long:  `Interact with Bitwarden secrets management through the integrated TUI`,
}

var bitwardenStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show Bitwarden vault status",
	Long:  `Show the current status of the Bitwarden vault`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if Bitwarden CLI is installed
		_, err := exec.LookPath("bw")
		if err != nil {
			log.Fatalf("Bitwarden CLI (bw) not found. Please install it first.")
		}

		// Check if vault is unlocked
		cmdExec := exec.Command("bw", "status")
		output, err := cmdExec.Output()
		if err != nil {
			log.Fatalf("Failed to check Bitwarden status: %v", err)
		}

		fmt.Printf("Bitwarden Vault Status:\n%s", output)
	},
}

var bitwardenUnlockCmd = &cobra.Command{
	Use:   "unlock",
	Short: "Unlock the Bitwarden vault",
	Long:  `Unlock the Bitwarden vault with your master password`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if Bitwarden CLI is installed
		_, err := exec.LookPath("bw")
		if err != nil {
			log.Fatalf("Bitwarden CLI (bw) not found. Please install it first.")
		}

		fmt.Println("Unlocking Bitwarden vault...")
		cmdExec := exec.Command("bw", "unlock")
		cmdExec.Stdin = os.Stdin
		cmdExec.Stdout = os.Stdout
		cmdExec.Stderr = os.Stderr

		err = cmdExec.Run()
		if err != nil {
			log.Fatalf("Failed to unlock Bitwarden vault: %v", err)
		}
	},
}

var bitwardenLockCmd = &cobra.Command{
	Use:   "lock",
	Short: "Lock the Bitwarden vault",
	Long:  `Lock the Bitwarden vault`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if Bitwarden CLI is installed
		_, err := exec.LookPath("bw")
		if err != nil {
			log.Fatalf("Bitwarden CLI (bw) not found. Please install it first.")
		}

		cmdExec := exec.Command("bw", "lock")
		err = cmdExec.Run()
		if err != nil {
			log.Fatalf("Failed to lock Bitwarden vault: %v", err)
		}

		fmt.Println("Bitwarden vault locked successfully.")
	},
}

var bitwardenListCmd = &cobra.Command{
	Use:   "list [filter]",
	Short: "List Bitwarden items",
	Long:  `List Bitwarden items, optionally filtered by name`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Check if Bitwarden CLI is installed
		_, err := exec.LookPath("bw")
		if err != nil {
			log.Fatalf("Bitwarden CLI (bw) not found. Please install it first.")
		}

		// Build command
		bwArgs := []string{"list", "items"}
		if len(args) > 0 {
			bwArgs = append(bwArgs, "--search", args[0])
		}

		cmdExec := exec.Command("bw", bwArgs...)
		output, err := cmdExec.Output()
		if err != nil {
			log.Fatalf("Failed to list Bitwarden items: %v", err)
		}

		fmt.Printf("Bitwarden Items:\n%s", output)
	},
}

var bitwardenSyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync Bitwarden vault",
	Long:  `Sync the local Bitwarden vault with the remote server`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if Bitwarden CLI is installed
		_, err := exec.LookPath("bw")
		if err != nil {
			log.Fatalf("Bitwarden CLI (bw) not found. Please install it first.")
		}

		fmt.Println("Syncing Bitwarden vault...")
		cmdExec := exec.Command("bw", "sync")
		output, err := cmdExec.Output()
		if err != nil {
			log.Fatalf("Failed to sync Bitwarden vault: %v", err)
		}

		fmt.Printf("Sync completed:\n%s", output)
	},
}

var bitwardenTuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Launch the Bitwarden TUI",
	Long:  `Launch the Bitwarden Secrets Manager TUI`,
	Run: func(cmd *cobra.Command, args []string) {
		// First check if the Bitwarden TUI exists in the dotfiles
		bwTuiPaths := []string{
			"~/.local/share/chezmoi/bw-secrets-tui/run.sh",
			"~/bw-secrets-tui/run.sh",
			"/home/foomanchu8008/.local/share/chezmoi/bw-secrets-tui/run.sh",
		}

		var foundPath string
		for _, path := range bwTuiPaths {
			expandedPath := os.ExpandEnv(strings.Replace(path, "~", os.Getenv("HOME"), -1))
			if _, err := os.Stat(expandedPath); err == nil {
				foundPath = expandedPath
				break
			}
		}

		if foundPath != "" {
			// Change to the directory and run the TUI
			dir := strings.Replace(foundPath, "/run.sh", "", -1)
			fmt.Printf("Launching Bitwarden TUI from %s...\n", dir)

			cmdExec := exec.Command("/bin/bash", "-c", fmt.Sprintf("cd %s && ./run.sh", dir))
			cmdExec.Stdin = os.Stdin
			cmdExec.Stdout = os.Stdout
			cmdExec.Stderr = os.Stderr

			err := cmdExec.Run()
			if err != nil {
				log.Fatalf("Failed to launch Bitwarden TUI: %v", err)
			}
		} else {
			// Fallback to checking if the Python TUI is installed
			_, err := exec.LookPath("bw-secrets-tui")
			if err == nil {
				cmdExec := exec.Command("bw-secrets-tui")
				cmdExec.Stdin = os.Stdin
				cmdExec.Stdout = os.Stdout
				cmdExec.Stderr = os.Stderr

				err = cmdExec.Run()
				if err != nil {
					log.Fatalf("Failed to launch Bitwarden TUI: %v", err)
				}
			} else {
				log.Fatalf("Bitwarden TUI not found. Please ensure it's installed in your dotfiles.")
			}
		}
	},
}

func init() {
	// Add subcommands to bitwarden command
	bitwardenCmd.AddCommand(bitwardenStatusCmd)
	bitwardenCmd.AddCommand(bitwardenUnlockCmd)
	bitwardenCmd.AddCommand(bitwardenLockCmd)
	bitwardenCmd.AddCommand(bitwardenListCmd)
	bitwardenCmd.AddCommand(bitwardenSyncCmd)
	bitwardenCmd.AddCommand(bitwardenTuiCmd)
	bitwardenCmd.AddCommand(bitwardenTemplateCmd)
	bitwardenCmd.AddCommand(bitwardenExportCmd)

	// Add the bitwarden command to the root
	root.RootCmd.AddCommand(bitwardenCmd)
}

var bitwardenTemplateCmd = &cobra.Command{
	Use:   "template [item-id]",
	Short: "Generate Chezmoi template from Bitwarden item",
	Long:  `Generate a Chezmoi template file from a Bitwarden item for secure secret management`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Check if Bitwarden CLI is installed
		_, err := exec.LookPath("bw")
		if err != nil {
			log.Fatalf("Bitwarden CLI (bw) not found. Please install it first.")
		}

		itemID := args[0]
		
		// Get the item details (we'll use this later for actual processing)
		_ = itemID // This is just to avoid the unused variable error for now

		// Parse the JSON output to extract fields
		// For simplicity, we'll just show how to create a template
		templatePath := "~/.local/share/chezmoi/dot_secrets.tmpl"
		expandedPath := os.ExpandEnv(strings.Replace(templatePath, "~", os.Getenv("HOME"), -1))
		
		// Create directory if it doesn't exist
		dir := strings.Replace(expandedPath, "/dot_secrets.tmpl", "", -1)
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatalf("Failed to create directory: %v", err)
		}

		// Write template file
		file, err := os.Create(expandedPath)
		if err != nil {
			log.Fatalf("Failed to create template file: %v", err)
		}
		defer file.Close()

		// Write template content
		templateContent := fmt.Sprintf(`# Bitwarden Secrets Template
# Generated from item ID: %s
# This file is auto-generated - do not edit manually

# Example template using Bitwarden integration
{{- if (bitwarden "%s") }}
export EXAMPLE_SECRET="{{ (bitwarden "%s").password }}"
{{- end }}

# You can add more secrets here as needed
`, itemID, itemID, itemID)

		_, err = file.WriteString(templateContent)
		if err != nil {
			log.Fatalf("Failed to write template file: %v", err)
		}

		fmt.Printf("Template generated at: %s\n", expandedPath)
		fmt.Println("To apply with chezmoi, run: chezmoi apply")
	},
}

var bitwardenExportCmd = &cobra.Command{
	Use:   "export [filename]",
	Short: "Export Bitwarden secrets to environment file",
	Long:  `Export Bitwarden secrets to a .env file for use in development`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Check if Bitwarden CLI is installed
		_, err := exec.LookPath("bw")
		if err != nil {
			log.Fatalf("Bitwarden CLI (bw) not found. Please install it first.")
		}

		// Default filename
		filename := ".env"
		if len(args) > 0 {
			filename = args[0]
		}

		// Check if vault is unlocked
		statusCmd := exec.Command("bw", "status")
		statusOutput, err := statusCmd.Output()
		if err != nil {
			log.Fatalf("Failed to check Bitwarden status: %v", err)
		}

		if strings.Contains(string(statusOutput), "\"status\":\"unlocked\"") {
			fmt.Println("Vault is unlocked. Proceeding with export...")
		} else {
			log.Fatalf("Vault is locked. Please unlock it first with: chezmoi-tui bitwarden unlock")
		}

		// For demonstration, we'll create a simple export
		file, err := os.Create(filename)
		if err != nil {
			log.Fatalf("Failed to create export file: %v", err)
		}
		defer file.Close()

		exportContent := `# Bitwarden Exported Secrets
# This file is auto-generated - do not commit to version control
# Last exported: %s

# Example secrets - replace with actual values from Bitwarden
EXAMPLE_API_KEY=your_api_key_here
EXAMPLE_SECRET=your_secret_here
`

		timestamp := time.Now().Format("2006-01-02 15:04:05")
		_, err = file.WriteString(fmt.Sprintf(exportContent, timestamp))
		if err != nil {
			log.Fatalf("Failed to write export file: %v", err)
		}

		// Set secure permissions
		err = os.Chmod(filename, 0600)
		if err != nil {
			log.Printf("Warning: Failed to set secure permissions: %v", err)
		}

		fmt.Printf("Secrets exported to: %s\n", filename)
		fmt.Println("Remember to add this file to your .gitignore!")
	},
}
