package chezmoi

import (
	"testing"
)

func TestNewChezmoi(t *testing.T) {
	client, err := New()
	if err != nil {
		t.Skipf("Skipping test: %v", err)
	}

	if client == nil {
		t.Error("Expected chezmoi client instance, got nil")
	}

	if client.binaryPath == "" {
		t.Error("Expected binary path, got empty string")
	}
}

func TestParseStatusOutput(t *testing.T) {
	client := &Chezmoi{}

	// Test with sample output
	// In chezmoi status: first char might be space, second is dest status, third is target status
	sampleOutput := ` M .bashrc
A  .gitconfig
   .vimrc`

	results := client.ParseStatusOutput(sampleOutput)

	if len(results) != 3 {
		t.Errorf("Expected 3 results, got %d", len(results))
	}

	// First line: ' M .bashrc' -> ["M", ".bashrc"]
	// This should result in dest_status=" ", target_status="M", filename=".bashrc"
	if results[0]["filename"] != ".bashrc" {
		t.Errorf("Expected .bashrc, got %s", results[0]["filename"])
	}

	// According to my logic: for " M .bashrc", it becomes [M, .bashrc]
	// Since len=2, destStatus=" ", targetStatus="M", filename=".bashrc"
	if results[0]["target_status"] != "M" {
		t.Errorf("Expected M as target_status, got %s", results[0]["target_status"])
	}

	if results[0]["dest_status"] != " " {
		t.Errorf("Expected space as dest_status, got '%s'", results[0]["dest_status"])
	}
}
