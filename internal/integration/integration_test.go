package integration

import (
	"testing"
)

func TestNewIntegration(t *testing.T) {
	// This test can only run if chezmoi is installed and available
	// In a real project, you would want to use mocks instead
	integ, err := New()
	if err != nil {
		t.Skipf("Skipping test: %v", err)
	}
	
	if integ == nil {
		t.Error("Expected integration instance, got nil")
	}
	
	if integ.client == nil {
		t.Error("Expected client instance, got nil")
	}
}

func TestGetConfigData(t *testing.T) {
	integ, err := New()
	if err != nil {
		t.Skipf("Skipping test: %v", err)
	}
	
	data, err := integ.GetConfigData()
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	
	if data == "" {
		t.Log("Config data is empty, which may be normal if chezmoi is not initialized")
	}
}