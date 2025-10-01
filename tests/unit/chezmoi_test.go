package tests

import (
	"chezmoi-tui/internal/chezmoi"
	"testing"
)

// TestChezmoiWrapper tests the low-level chezmoi wrapper
func TestChezmoiWrapper(t *testing.T) {
	t.Run("Initialization", func(t *testing.T) {
		client, err := chezmoi.New()
		if err != nil {
			t.Skipf("Skipping test: %v", err) // Skip if chezmoi isn't available
		}

		if client == nil {
			t.Error("Expected chezmoi client instance, got nil")
		}

		if client.GetBinaryPath() == "" {
			t.Error("Expected binary path, got empty string")
		}
	})

	t.Run("ParseStatusOutput", func(t *testing.T) {
		client := &chezmoi.Chezmoi{}

		// Test with various status output formats
		testCases := []struct {
			name     string
			input    string
			expected int
		}{
			{
				name: "Normal status output",
				input: ` M .bashrc
A  .gitconfig
   .vimrc`,
				expected: 3,
			},
			{
				name:     "Empty output",
				input:    "",
				expected: 0,
			},
			{
				name:     "Single file",
				input:    `   .single`,
				expected: 1,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				results := client.ParseStatusOutput(tc.input)
				if len(results) != tc.expected {
					t.Errorf("Expected %d results, got %d", tc.expected, len(results))
				}
			})
		}
	})
}
