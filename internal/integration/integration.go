package integration

import (
	"fmt"
	"chezmoi-tui/internal/chezmoi"
)

// ChezmoiIntegration provides high-level operations for interacting with chezmoi
type ChezmoiIntegration struct {
	client *chezmoi.Chezmoi
}

// New creates a new integration layer
func New() (*ChezmoiIntegration, error) {
	client, err := chezmoi.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create chezmoi client: %w", err)
	}
	
	return &ChezmoiIntegration{
		client: client,
	}, nil
}

// GetStatus returns the current status of all managed files
func (ci *ChezmoiIntegration) GetStatus() ([]map[string]string, error) {
	statusOutput, err := ci.client.Status()
	if err != nil {
		return nil, fmt.Errorf("failed to get status: %w", err)
	}
	
	return ci.client.ParseStatusOutput(statusOutput), nil
}

// ApplyFiles applies the specified target files
func (ci *ChezmoiIntegration) ApplyFiles(targets ...string) (string, error) {
	return ci.client.Apply(targets...)
}

// AddFiles adds the specified files to the source state
func (ci *ChezmoiIntegration) AddFiles(targets ...string) (string, error) {
	return ci.client.Add(targets...)
}

// GetManagedFiles returns a list of all managed files
func (ci *ChezmoiIntegration) GetManagedFiles() (string, error) {
	return ci.client.Managed()
}

// GetUnmanagedFiles returns a list of all unmanaged files
func (ci *ChezmoiIntegration) GetUnmanagedFiles() (string, error) {
	return ci.client.Unmanaged()
}

// GetIgnoredFiles returns a list of all ignored files
func (ci *ChezmoiIntegration) GetIgnoredFiles() (string, error) {
	return ci.client.Ignored()
}

// GetConfigData returns the template data
func (ci *ChezmoiIntegration) GetConfigData() (string, error) {
	return ci.client.Data()
}

// RunDoctor checks the system for potential problems
func (ci *ChezmoiIntegration) RunDoctor() (string, error) {
	return ci.client.Doctor()
}

// DiffFiles shows the differences for the specified files
func (ci *ChezmoiIntegration) DiffFiles(targets ...string) (string, error) {
	return ci.client.Diff(targets...)
}

// InitializeRepo initializes the source directory with a remote repository
func (ci *ChezmoiIntegration) InitializeRepo(repo string, apply bool) (string, error) {
	args := []string{repo}
	if apply {
		args = append(args, "--apply")
	}
	return ci.client.Init(args...)
}