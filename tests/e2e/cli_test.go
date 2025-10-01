package e2e

import (
	"os"
	"os/exec"
	"testing"
)

// TestCLICommands tests the basic CLI functionality
func TestCLICommands(t *testing.T) {
	// Build the binary first
	cmd := exec.Command("go", "build", "-o", "test-binary", ".")
	cmd.Dir = "../" // Build from the parent directory
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer os.Remove("./test-binary") // Clean up

	// Test version command
	t.Run("VersionCommand", func(t *testing.T) {
		cmd := exec.Command("./test-binary", "version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Version command failed: %v, output: %s", err, output)
		}

		expected := "chezmoi-tui version"
		if !contains(string(output), expected) {
			t.Errorf("Expected version output to contain '%s', got: %s", expected, output)
		}
	})

	// Test help command
	t.Run("HelpCommand", func(t *testing.T) {
		cmd := exec.Command("./test-binary", "--help")
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Help command failed: %v, output: %s", err, output)
		}

		expectedCommands := []string{"add", "apply", "status", "tui", "version"}
		for _, cmd := range expectedCommands {
			if !contains(string(output), cmd) {
				t.Errorf("Expected help output to contain command '%s'", cmd)
			}
		}
	})
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) &&
		(s == substr ||
			len(substr) == 0 ||
			(len(s) > len(substr) &&
				(s[:len(substr)] == substr ||
					contains(s[1:], substr))))
}
