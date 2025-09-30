package commands

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"chezmoi-tui/pkg/root"
	"chezmoi-tui/internal/integration"
)

var initCmd = &cobra.Command{
	Use:   "init [repo]",
	Short: "Setup the source directory and update the destination directory to match the target state",
	Long:  `Setup the source directory, generate the config file, and optionally update the destination directory to match the target state.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Create integration instance
		integ, err := integration.New()
		if err != nil {
			log.Fatalf("Failed to initialize integration: %v", err)
		}

		var repo string
		if len(args) > 0 {
			repo = args[0]
		}

		// Get flags
		apply, _ := cmd.Flags().GetBool("apply")
		purge, _ := cmd.Flags().GetBool("purge")
		
		var output string
		if repo != "" {
			output, err = integ.InitializeRepo(repo, apply)
		} else {
			// Just init without a repo
			output, err = integ.InitializeRepo("", apply)
		}
		
		if err != nil {
			log.Fatalf("Failed to initialize: %v", err)
		}

		if output != "" {
			fmt.Print(output)
		}
		
		if purge {
			// In a real implementation, this would purge the config, source, and cache directories
			fmt.Println("Purge functionality would remove config, source, and cache directories")
		}
	},
}

func init() {
	// Add flags
	initCmd.Flags().BoolP("apply", "a", false, "Update destination directory")
	initCmd.Flags().BoolP("purge", "p", false, "Purge config and source directories after running")
	
	// Add the command to the root
	root.RootCmd.AddCommand(initCmd)
}