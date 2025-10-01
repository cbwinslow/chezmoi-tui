package commands

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

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

	// Add the bitwarden command to the root
	root.RootCmd.AddCommand(bitwardenCmd)
}
