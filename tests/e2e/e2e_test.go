package e2e

import (
	"os"
	"os/exec"
	"testing"
	"time"
)

const testBinary = "../test-binary"

// setupTestEnv builds the binary for testing
func setupTestEnv(t *testing.T) {
	// Build the binary in the project root directory
	// Run from project root where go.mod is located
	cmd := exec.Command("go", "build", "-o", "../../test-binary", ".")
	cmd.Dir = "../../../" // Go to project root
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to build binary for testing: %v", err)
	}
}

// tearDownTestEnv removes the test binary
func tearDownTestEnv(t *testing.T) {
	// Remove the binary from project root
	err := os.Remove("../../../test-binary")
	if err != nil {
		t.Logf("Warning: failed to remove test binary: %v", err)
	}
}

// TestE2EWorkflow tests a complete user workflow
func TestE2EWorkflow(t *testing.T) {
	setupTestEnv(t)
	defer tearDownTestEnv(t)

	t.Run("CompleteWorkflow", func(t *testing.T) {
		// Test that the binary exists and can show help
		t.Run("BinaryExistsAndRuns", func(t *testing.T) {
			_, err := os.Stat(testBinary)
			if os.IsNotExist(err) {
				t.Fatal("Test binary does not exist")
			}

			cmd := exec.Command(testBinary, "--help")
			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("Help command failed: %v, output: %s", err, output)
			}

			if len(output) == 0 {
				t.Error("Help command returned empty output")
			}
		})

		// Test version command
		t.Run("VersionCommandWorks", func(t *testing.T) {
			cmd := exec.Command(testBinary, "version")
			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("Version command failed: %v, output: %s", err, output)
			}

			expected := "chezmoi-tui version"
			if !stringContains(string(output), expected) {
				t.Errorf("Expected version output to contain '%s', got: %s", expected, output)
			}
		})

		// Test status command (may fail if chezmoi isn't set up, but shouldn't crash)
		t.Run("StatusCommandExists", func(t *testing.T) {
			// Use a timeout to prevent hanging if TUI tries to start
			cmd := exec.Command(testBinary, "status")
			done := make(chan error, 1)

			go func() {
				_, err := cmd.CombinedOutput()
				done <- err
			}()

			select {
			case <-time.After(5 * time.Second):
				t.Log("Status command completed or timed out (expected behavior)")
				if cmd.Process != nil {
					cmd.Process.Kill()
				}
			case err := <-done:
				if err != nil {
					// This is OK - chezmoi might not be initialized
					t.Logf("Status command returned with error (expected if chezmoi not set up): %v", err)
				} else {
					t.Log("Status command completed successfully")
				}
			}
		})
	})
}

// TestCLIFeatures tests individual CLI features
func TestCLIFeatures(t *testing.T) {
	setupTestEnv(t)
	defer tearDownTestEnv(t)

	t.Run("AllCommandsAvailable", func(t *testing.T) {
		// Get help output to check for commands
		cmd := exec.Command(testBinary, "--help")
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Logf("Help command failed: %v, output: %s", err, output)
		}

		expectedCommands := []string{"add", "apply", "status", "tui", "version"}
		for _, cmd := range expectedCommands {
			if !stringContains(string(output), cmd) {
				t.Errorf("Expected help output to contain command '%s'", cmd)
			}
		}
	})
}

// TestTUIStarts tests that the TUI command at least starts without crashing
func TestTUIStarts(t *testing.T) {
	setupTestEnv(t)
	defer tearDownTestEnv(t)

	t.Run("TUICommandExists", func(t *testing.T) {
		// We can't fully test the TUI without a terminal,
		// but we can check that the command is recognized
		cmd := exec.Command(testBinary, "tui", "--help")
		output, err := cmd.CombinedOutput()

		// The command should be recognized even if it can't run in CI
		if err != nil {
			outputStr := string(output)
			if stringContains(outputStr, "unknown command") {
				t.Error("TUI command not recognized")
			}
		}
	})
}

// Helper function to check if a string contains a substring
func stringContains(s, substr string) bool {
	return len(s) >= len(substr) &&
		(s == substr ||
			len(substr) == 0 ||
			(len(s) > len(substr) &&
				(s[:len(substr)] == substr ||
					stringContains(s[1:], substr))))
}
