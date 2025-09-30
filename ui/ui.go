package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbletea"
	"chezmoi-tui/internal/integration"
)

type StatusType int

const (
	StatusUpToDate StatusType = iota
	StatusModified
	StatusUnmanaged
	StatusIgnored
)

type FileStatus struct {
	Name         string
	Type         StatusType
	DestStatus   string
	TargetStatus string
}

// Model represents the state of the TUI
type Model struct {
	// Integration layer
	integration *integration.ChezmoiIntegration
	
	// UI state
	choice    int
	choices   []string
	cursor    int
	quitting  bool
	
	// File status view
	fileCursor int
	fileStatus []FileStatus
	showFiles  bool
}

// RunTUI starts the terminal user interface
func RunTUI() error {
	// Initialize integration layer
	integ, err := integration.New()
	if err != nil {
		return fmt.Errorf("failed to initialize chezmoi integration: %w", err)
	}
	
	p := tea.NewProgram(initialModel(integ), tea.WithAltScreen())
	_, err = p.Run()
	return err
}

// initialModel returns the initial state of the UI
func initialModel(integ *integration.ChezmoiIntegration) Model {
	return Model{
		choices:     []string{"View Status", "Add Files", "Apply Changes", "Diff Changes", "Edit Config", "Exit"},
		integration: integ,
		fileStatus:  []FileStatus{},
	}
}

// Init is the initial command for the TUI
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles messages and updates the model
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit

		case "up", "k":
			if m.showFiles {
				if m.fileCursor > 0 {
					m.fileCursor--
				}
			} else {
				if m.cursor > 0 {
					m.cursor--
				}
			}

		case "down", "j":
			if m.showFiles {
				if m.fileCursor < len(m.fileStatus)-1 {
					m.fileCursor++
				}
			} else {
				if m.cursor < len(m.choices)-1 {
					m.cursor++
				}
			}

		case "left", "h":
			if m.showFiles {
				m.showFiles = false
			}

		case "right", "l":
			if !m.showFiles && m.choices[m.cursor] == "View Status" {
				// Load file status using the integration layer
				statusData, err := m.integration.GetStatus()
				if err != nil {
					// Handle error - for now just show an error message
					m.fileStatus = []FileStatus{
						{Name: fmt.Sprintf("Error loading status: %v", err), Type: StatusIgnored},
					}
				} else {
					// Convert the status data to our internal format
					m.fileStatus = make([]FileStatus, len(statusData))
					for i, entry := range statusData {
						destStatus := entry["dest_status"]
						targetStatus := entry["target_status"]
						filename := entry["filename"]
						
						statusType := getStatusType(destStatus, targetStatus)
						
						m.fileStatus[i] = FileStatus{
							Name:         filename,
							Type:         statusType,
							DestStatus:   destStatus,
							TargetStatus: targetStatus,
						}
					}
				}
				m.showFiles = true
			}

		case "enter":
			if m.showFiles {
				// Handle file-specific action - just ignore for now
			} else {
				m.choice = m.cursor
				switch m.choices[m.cursor] {
				case "View Status":
					// Already handled with right arrow
				case "Exit":
					m.quitting = true
					return m, tea.Quit
				default:
					// Other actions can be implemented as needed
				}
			}
		}
	}

	return m, nil
}

// getStatusType determines the status type based on chezmoi status codes
func getStatusType(destStatus, targetStatus string) StatusType {
	// This is a simplification - in real world you'd have more complex logic
	if strings.Contains(destStatus, "M") || strings.Contains(targetStatus, "M") {
		return StatusModified
	}
	if strings.Contains(destStatus, "A") || strings.Contains(targetStatus, "A") {
		return StatusUnmanaged
	}
	if strings.Contains(destStatus, "D") || strings.Contains(targetStatus, "D") {
		return StatusUnmanaged // deleted
	}
	return StatusUpToDate
}

// View renders the UI
func (m Model) View() string {
	if m.quitting {
		return "Bye!\n"
	}

	if m.showFiles {
		// File status view
		s := "Chezmoi File Status\n\n"

		for i, file := range m.fileStatus {
			cursor := " "
			if m.fileCursor == i {
				cursor = ">"
			}
			
			statusSymbol := " "
			switch file.Type {
			case StatusModified:
				statusSymbol = "M"
			case StatusUnmanaged:
				statusSymbol = "A"
			case StatusIgnored:
				statusSymbol = "I"
			default:
				statusSymbol = " "
			}
			
			s += fmt.Sprintf("%s [%s] %s\n", cursor, statusSymbol, file.Name)
		}

		s += fmt.Sprintf("\n%d files total | Use arrow keys to navigate, left/right to switch views, q to quit\n", len(m.fileStatus))
		return s
	} else {
		// Main menu
		s := "Chezmoi TUI - Enhanced dotfile management\n\n"

		for i, choice := range m.choices {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			s += fmt.Sprintf("%s %s\n", cursor, choice)
		}

		s += "\nUse arrow keys to navigate, enter to select, q to quit\n"
		return s
	}
}