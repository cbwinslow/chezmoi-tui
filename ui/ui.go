package ui

import (
	"fmt"
	"strings"

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
	choices := []string{"View Status", "Add Files", "Apply Changes", "Diff Changes", "Edit Config", "Exit"}
	
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
	case "Edit Config":
		return "Edit configuration files"
	case "Exit":
		return "Quit the application"
	default:
		return "Select an option"
	}
}