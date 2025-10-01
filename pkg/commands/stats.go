package commands

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"chezmoi-tui/internal/integration"
	"chezmoi-tui/pkg/root"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Show statistics about your dotfiles",
	Long:  `Show statistics and analytics about your dotfiles management`,
	Run: func(cmd *cobra.Command, args []string) {
		// Create integration instance
		integ, err := integration.New()
		if err != nil {
			log.Fatalf("Failed to initialize integration: %v", err)
		}

		// Get status information
		statusData, err := integ.GetStatus()
		if err != nil {
			log.Printf("Could not get status data: %v", err)
		}

		// Get all managed files
		managedOutput, err := integ.GetManagedFiles()
		if err != nil {
			log.Printf("Could not get managed files: %v", err)
		}

		// Get unmanaged files
		unmanagedOutput, err := integ.GetUnmanagedFiles()
		if err != nil {
			log.Printf("Could not get unmanaged files: %v", err)
		}

		// Calculate stats
		var modifiedCount, addedCount, deletedCount, upToDateCount int
		for _, entry := range statusData {
			destStatus := entry["dest_status"]
			targetStatus := entry["target_status"]

			if strings.Contains(destStatus, "M") || strings.Contains(targetStatus, "M") {
				modifiedCount++
			} else if strings.Contains(destStatus, "A") || strings.Contains(targetStatus, "A") {
				addedCount++
			} else if strings.Contains(destStatus, "D") || strings.Contains(targetStatus, "D") {
				deletedCount++
			} else {
				upToDateCount++
			}
		}

		managedFiles := strings.Split(strings.TrimSpace(managedOutput), "\n")
		var validManagedFiles []string
		for _, file := range managedFiles {
			if strings.TrimSpace(file) != "" {
				validManagedFiles = append(validManagedFiles, file)
			}
		}

		unmanagedFiles := strings.Split(strings.TrimSpace(unmanagedOutput), "\n")
		var validUnmanagedFiles []string
		for _, file := range unmanagedFiles {
			if strings.TrimSpace(file) != "" {
				validUnmanagedFiles = append(validUnmanagedFiles, file)
			}
		}

		// Display statistics
		fmt.Println("┌─ Chezmoi Dotfiles Statistics ──────────────────────────────────┐")
		fmt.Printf("│ Last Updated: %-47s │\n", time.Now().Format("2006-01-02 15:04:05"))
		fmt.Println("├─────────────────────────────────────────────────────────────────┤")
		fmt.Printf("│ Total Managed Files:    %3d                                   │\n", len(validManagedFiles))
		fmt.Printf("│ Total Unmanaged Files:  %3d                                   │\n", len(validUnmanagedFiles))
		fmt.Println("├─────────────────────────────────────────────────────────────────┤")
		fmt.Printf("│ Up to Date:             %3d (%3d%%)                           │\n",
			upToDateCount, calculatePercentage(upToDateCount, len(validManagedFiles)))
		fmt.Printf("│ Modified:               %3d (%3d%%)                           │\n",
			modifiedCount, calculatePercentage(modifiedCount, len(validManagedFiles)))
		fmt.Printf("│ Added:                  %3d (%3d%%)                           │\n",
			addedCount, calculatePercentage(addedCount, len(validManagedFiles)))
		fmt.Printf("│ Deleted:                %3d (%3d%%)                           │\n",
			deletedCount, calculatePercentage(deletedCount, len(validManagedFiles)))
		fmt.Println("└─────────────────────────────────────────────────────────────────┘")

		// Show additional details if requested
		details, _ := cmd.Flags().GetBool("details")
		if details {
			fmt.Println("\nDetailed Breakdown:")
			fmt.Printf("Managed files (%d): %v\n", len(validManagedFiles), validManagedFiles)
			// Only show unmanaged if there are any
			if len(validUnmanagedFiles) > 0 {
				fmt.Printf("Unmanaged files (%d): %v\n", len(validUnmanagedFiles), validUnmanagedFiles)
			}
		}
	},
}

func calculatePercentage(part, total int) int {
	if total <= 0 {
		return 0
	}
	return int(float64(part) / float64(total) * 100)
}

func init() {
	// Add flags to the stats command
	statsCmd.Flags().BoolP("details", "d", false, "Show detailed breakdown")

	// Add the stats command to the root
	root.RootCmd.AddCommand(statsCmd)
}
