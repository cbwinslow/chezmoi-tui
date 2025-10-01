package chezmoi

import (
	"fmt"
	"os/exec"
	"strings"
)

// Chezmoi wraps the chezmoi command-line tool
type Chezmoi struct {
	binaryPath string
}

// New creates a new Chezmoi wrapper
func New() (*Chezmoi, error) {
	binaryPath, err := exec.LookPath("chezmoi")
	if err != nil {
		return nil, fmt.Errorf("chezmoi binary not found in PATH: %w", err)
	}

	return &Chezmoi{
		binaryPath: binaryPath,
	}, nil
}

// Run executes a chezmoi command with the given arguments
func (c *Chezmoi) Run(args ...string) (string, error) {
	cmd := exec.Command(c.binaryPath, args...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("chezmoi %v failed: %w (output: %s)", args, err, string(output))
	}

	return string(output), nil
}

// Status runs the chezmoi status command
func (c *Chezmoi) Status() (string, error) {
	return c.Run("status")
}

// Apply runs the chezmoi apply command
func (c *Chezmoi) Apply(targets ...string) (string, error) {
	args := []string{"apply"}
	args = append(args, targets...)
	return c.Run(args...)
}

// Add runs the chezmoi add command
func (c *Chezmoi) Add(targets ...string) (string, error) {
	args := []string{"add"}
	args = append(args, targets...)
	return c.Run(args...)
}

// Diff runs the chezmoi diff command
func (c *Chezmoi) Diff(targets ...string) (string, error) {
	args := []string{"diff"}
	args = append(args, targets...)
	return c.Run(args...)
}

// Init runs the chezmoi init command
func (c *Chezmoi) Init(args ...string) (string, error) {
	initWithArgs := []string{"init"}
	initWithArgs = append(initWithArgs, args...)
	return c.Run(initWithArgs...)
}

// GetBinaryPath returns the path to the chezmoi binary
func (c *Chezmoi) GetBinaryPath() string {
	return c.binaryPath
}

// Managed runs the chezmoi managed command to list managed entries
func (c *Chezmoi) Managed() (string, error) {
	return c.Run("managed")
}

// Unmanaged runs the chezmoi unmanaged command to list unmanaged files
func (c *Chezmoi) Unmanaged() (string, error) {
	return c.Run("unmanaged")
}

// Ignored runs the chezmoi ignored command to list ignored targets
func (c *Chezmoi) Ignored() (string, error) {
	return c.Run("ignored")
}

// Doctor runs the chezmoi doctor command to check for potential problems
func (c *Chezmoi) Doctor() (string, error) {
	return c.Run("doctor")
}

// Data runs the chezmoi data command to print template data
func (c *Chezmoi) Data() (string, error) {
	return c.Run("data")
}

// ParseStatusOutput parses the output of the status command into structured data
func (c *Chezmoi) ParseStatusOutput(output string) []map[string]string {
	var result []map[string]string

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		// Split by whitespace but preserve file paths with spaces
		parts := strings.Fields(line)
		if len(parts) >= 1 { // At least 2 parts: status and filename
			// Extract the status column and filename
			// In chezmoi status, there could be 1 column (status only) or 2+ columns
			var destStatus, targetStatus, filename string

			if len(parts) == 2 {
				// Format: status filename
				destStatus = " "
				targetStatus = parts[0]
				filename = parts[1]
			} else if len(parts) >= 3 {
				// Format: dest_status target_status filename(s)
				destStatus = parts[0]
				targetStatus = parts[1]
				// Join remaining parts to reconstruct full filename (which might have spaces)
				filename = strings.Join(parts[2:], " ")
			} else {
				// If only one part, it's likely just the filename with no status changes
				destStatus = " "
				targetStatus = " "
				filename = parts[0]
			}

			entry := map[string]string{
				"dest_status":   destStatus,
				"target_status": targetStatus,
				"filename":      filename,
			}
			result = append(result, entry)
		}
	}

	return result
}
