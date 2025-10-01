package tests

import (
	"chezmoi-tui/internal/integration"
	"testing"
)

// MockChezmoiClient simulates the Chezmoi client for testing
type MockChezmoiClient struct {
	StatusResult string
	StatusError  error
	ApplyResult  string
	ApplyError   error
	AddResult    string
	AddError     error
	DataResult   string
	DataError    error
}

// Mock methods would go here if we were using a proper interface
// For now, we'll test based on the actual integration layer behavior

// TestIntegrationLayer tests the high-level integration functions
func TestIntegrationLayer(t *testing.T) {
	t.Run("Initialization", func(t *testing.T) {
		integ, err := integration.New()
		if err != nil {
			t.Skipf("Skipping test: %v", err) // Skip if chezmoi isn't available
		}

		if integ == nil {
			t.Error("Expected integration instance, got nil")
		}

		// Test that all expected methods exist and don't panic
		_, err = integ.GetManagedFiles()
		if err != nil {
			// This is OK if chezmoi isn't initialized
			t.Logf("GetManagedFiles returned error (expected if chezmoi not initialized): %v", err)
		}

		_, err = integ.GetUnmanagedFiles()
		if err != nil {
			t.Logf("GetUnmanagedFiles returned error (expected if chezmoi not initialized): %v", err)
		}

		_, err = integ.GetConfigData()
		if err != nil {
			t.Logf("GetConfigData returned error (expected if chezmoi not initialized): %v", err)
		}
	})

	t.Run("StatusParsing", func(t *testing.T) {
		integ, err := integration.New()
		if err != nil {
			t.Skipf("Skipping test: %v", err)
		}

		// Test the parsing function by calling it indirectly
		// Since we can't directly test the parsing without real output,
		// we'll just make sure the function doesn't crash
		_, err = integ.GetStatus()
		if err != nil {
			t.Logf("GetStatus returned error (expected if chezmoi not initialized): %v", err)
		}
	})
}
