package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"chezmoi-tui/internal/integration"
)

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = itemStyle.Copy().Foreground(lipgloss.Color("#1793d1"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
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

// item implements list.Item
type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

// Model represents the state of the TUI
type Model struct {
	// Integration layer
	integration *integration.ChezmoiIntegration
	
	// Main menu state
	choice    int
	choices   []string
	cursor    int
	quitting  bool
	
	// Status list view
	statusList list.Model
	
	// File status view
	fileCursor int
	fileStatus []FileStatus
	showFiles  bool
	help       help.Model
	viewport   viewport.Model
}

// RunTUI starts the terminal user interface
func RunTUI() error {
	// Initialize integration layer
	integ, err := integration.New()
	if err != nil {
		return fmt.Errorf("failed to initialize chezmoi integration: %w", err)
	}
	
	model := initialModel(integ)
	p := tea.NewProgram(&model, tea.WithAltScreen())
	_, err = p.Run()
	return err
}

// initialModel returns the initial state of the UI
func initialModel(integ *integration.ChezmoiIntegration) Model {
	choices := []string{"View Status", "Add Files", "Apply Changes", "Diff Changes", "Show Stats", "Bitwarden Manager", "Exit"}
	
	// Create items for the list
	var items []list.Item
	for _, choice := range choices {
		items = append(items, item{title: choice, desc: getDescription(choice)})
	}

	// Create a custom delegate for our file list
	delegate := list.NewDefaultDelegate()
	delegate.Styles.NormalTitle = itemStyle
	delegate.Styles.SelectedTitle = selectedItemStyle
	
	// Create the status list
	statusList := list.New(items, delegate, 0, 0)
	statusList.Title = "Chezmoi TUI - Enhanced dotfile management"
	statusList.SetShowStatusBar(false)
	statusList.SetFilteringEnabled(false)
	statusList.Styles.Title = titleStyle
	
	// Set key bindings
	statusList.KeyMap.Quit = key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	)
	
	return Model{
		choices:     choices,
		integration: integ,
		fileStatus:  []FileStatus{},
		statusList:  statusList,
		help:        help.New(),
		viewport:    viewport.New(78, 20), // width and height
	}
}

// Init is the initial command for the TUI
func (m *Model) Init() tea.Cmd {
	return nil
}

// Update handles messages and updates the model
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := titleStyle.GetFrameSize()
		m.statusList.SetSize(msg.Width-h, msg.Height-v)
		m.viewport.Width = msg.Width
		m.viewport.Height = msg.Height - 15 // Leave space for header and footer

	case tea.KeyMsg:
		// Don't forward quit or back commands to the list when showing files
		if m.showFiles && (msg.String() == "h" || msg.String() == "left") {
			m.showFiles = false
			return m, nil
		}

		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			if !m.showFiles {
				m.choice = m.statusList.Index()
				selectedItem := m.statusList.SelectedItem()
				if selectedItem != nil {
					item := selectedItem.(item)
					if item.title == "Exit" {
						m.quitting = true
						return m, tea.Quit
					} else if item.title == "View Status" {
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
					} else if item.title == "Show Stats" {
						// Show statistics about the dotfiles
						statsContent, err := generateStatsContent(m.integration)
						if err != nil {
							statsContent = fmt.Sprintf("Error loading stats: %v", err)
						}
						m.viewport.SetContent(statsContent)
						m.showFiles = true // Reuse the file view for stats display
					} else if item.title == "Bitwarden Manager" {
						// Show Bitwarden manager information
						bwContent := generateBitwardenContent()
						m.viewport.SetContent(bwContent)
						m.showFiles = true
					}
				}
			}
		case "l", "right":
			if !m.showFiles {
				// Check if "View Status" is selected
				selectedItem := m.statusList.SelectedItem()
				if selectedItem != nil {
					item := selectedItem.(item)
					if item.title == "View Status" {
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
				}
			}
		}
	}

	// Update the status list unless we're showing files
	if !m.showFiles {
		m.statusList, cmd = m.statusList.Update(msg)
		cmds = append(cmds, cmd)
	}

	// Update the viewport if needed
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
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
func (m *Model) View() string {
	if m.quitting {
		return quitTextStyle.Render("Bye!")
	}

	if m.showFiles {
		// File status view
		if len(m.fileStatus) == 0 {
			return "No files to display. Press 'h' to go back.\n"
		}

		// Create content for the viewport
		var content strings.Builder
		content.WriteString("Chezmoi File Status\n\n")

		for i, file := range m.fileStatus {
			cursor := " "
			if m.fileCursor == i {
				cursor = "→"
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
			
			content.WriteString(fmt.Sprintf("%s [%s] %s\n", cursor, statusSymbol, file.Name))
		}

		content.WriteString(fmt.Sprintf("\n%d files total | Use arrow keys to navigate, 'h' to go back, 'q' to quit\n", len(m.fileStatus)))

		m.viewport.SetContent(content.String())
		return m.viewport.View()
	} else {
		// Main menu
		return m.statusList.View()
	}
}

func generateStatsContent(integ *integration.ChezmoiIntegration) (string, error) {
	// Get status information
	statusData, err := integ.GetStatus()
	if err != nil {
		return "", fmt.Errorf("could not get status data: %w", err)
	}

	// Get all managed files
	managedOutput, err := integ.GetManagedFiles()
	if err != nil {
		return "", fmt.Errorf("could not get managed files: %w", err)
	}

	// Get unmanaged files
	unmanagedOutput, err := integ.GetUnmanagedFiles()
	if err != nil {
		// This is okay, sometimes there are no unmanaged files
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

	var content strings.Builder
	content.WriteString("┌─ Chezmoi Dotfiles Statistics ──────────────────────────────────┐\n")
	content.WriteString(fmt.Sprintf("│ Last Updated: %-47s │\n", time.Now().Format("2006-01-02 15:04:05")))
	content.WriteString("├─────────────────────────────────────────────────────────────────┤\n")
	content.WriteString(fmt.Sprintf("│ Total Managed Files:    %3d                                   │\n", len(validManagedFiles)))
	content.WriteString(fmt.Sprintf("│ Total Unmanaged Files:  %3d                                   │\n", len(validManagedFiles)))
	content.WriteString("├─────────────────────────────────────────────────────────────────┤\n")
	content.WriteString(fmt.Sprintf("│ Up to Date:             %3d (%3d%%)                           │\n", 
		upToDateCount, calculatePercentage(upToDateCount, len(validManagedFiles))))
	content.WriteString(fmt.Sprintf("│ Modified:               %3d (%3d%%)                           │\n", 
		modifiedCount, calculatePercentage(modifiedCount, len(validManagedFiles))))
	content.WriteString(fmt.Sprintf("│ Added:                  %3d (%3d%%)                           │\n", 
		addedCount, calculatePercentage(addedCount, len(validManagedFiles))))
	content.WriteString(fmt.Sprintf("│ Deleted:                %3d (%3d%%)                           │\n", 
		deletedCount, calculatePercentage(deletedCount, len(validManagedFiles))))
	content.WriteString("├─────────────────────────────────────────────────────────────────┤\n")
	content.WriteString("│ Actions: Use arrow keys to navigate, 'h' to go back, 'q' to quit │\n")
	content.WriteString("└─────────────────────────────────────────────────────────────────┘\n")

	return content.String(), nil
}

func calculatePercentage(part, total int) int {
	if total <= 0 {
		return 0
	}
	return int(float64(part) / float64(total) * 100)
}

func generateBitwardenContent() string {
	return "┌─ Bitwarden Secrets Manager ─────────────────────────────────────┐\n" +
		"│                                                                 │\n" +
		"│  Manage your Bitwarden vault from within Chezmoi TUI           │\n" +
		"│                                                                 │\n" +
		"│  Features:                                                      │\n" +
		"│  • View and manage Bitwarden items                              │\n" +
		"│  • Unlock/Lock your vault                                       │\n" +
		"│  • Sync with remote server                                      │\n" +
		"│  • Integration with Chezmoi templates                            │\n" +
		"│  • Launch dedicated Bitwarden TUI                               │\n" +
		"│                                                                 │\n" +
		"│  CLI Commands:                                                   │\n" +
		"│  • chezmoi-tui bitwarden status  - Show vault status            │\n" +
		"│  • chezmoi-tui bitwarden unlock  - Unlock vault                 │\n" +
		"│  • chezmoi-tui bitwarden lock    - Lock vault                   │\n" +
		"│  • chezmoi-tui bitwarden list    - List items                  │\n" +
		"│  • chezmoi-tui bitwarden sync    - Sync with server             │\n" +
		"│  • chezmoi-tui bitwarden tui     - Launch Bitwarden TUI         │\n" +
		"│                                                                 │\n" +
		"│  Actions: Use arrow keys to navigate, 'h' to go back, 'q' to quit │\n" +
		"└─────────────────────────────────────────────────────────────────┘\n"
}

func getDescription(choice string) string {
	switch choice {
	case "View Status":
		return "Show status of all managed files"
	case "Add Files":
		return "Add new files to chezmoi management"
	case "Apply Changes":
		return "Apply managed files to your system"
	case "Diff Changes":
		return "Show differences between source and destination"
	case "Show Stats":
		return "Show statistics about your dotfiles"
	case "Bitwarden Manager":
		return "Manage Bitwarden secrets and integration"
	case "Exit":
		return "Quit the application"
	default:
		return "Select an option"
	}
}