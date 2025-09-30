# Chezmoi TUI - Enhanced Terminal User Interface for Chezmoi

An enhanced Terminal User Interface (TUI) and CLI for managing your dotfiles with chezmoi.

## Features

- Interactive terminal UI for managing dotfiles
- Enhanced CLI commands with better feedback
- Visual representation of dotfiles status
- Easy management of chezmoi configurations
- Integration with existing chezmoi functionality
- Real-time status updates
- Keyboard navigation and shortcuts
- Search and filter capabilities

## Installation

### Prerequisites
- Go 1.19 or later
- Chezmoi installed and in your PATH

### Building from source
```bash
git clone <repository-url>
cd chezmoi-tui
go build -o chezmoi-tui .
```

## Usage

### CLI Commands
- `chezmoi-tui status` - Show the status of targets
- `chezmoi-tui apply [targets...]` - Update destination directory to match target state
- `chezmoi-tui add [targets...]` - Add targets to the source state
- `chezmoi-tui diff [targets...]` - Show differences between target and destination states
- `chezmoi-tui tui` - Launch the Terminal User Interface
- `chezmoi-tui version` - Print version information

### TUI Interface
Launch the interactive terminal interface with `chezmoi-tui tui`:
- Use arrow keys or hjkl to navigate
- Press 'l'/'right' to view file status details
- Press 'h'/'left' to return to the main menu
- Press 'q' or Ctrl+C to quit

## Architecture

The system follows a modular architecture with separate components for CLI and TUI functionality, integrated with the existing chezmoi backend.

For detailed architecture, see [docs/architecture.md](docs/architecture.md).

## Development

For development guidelines, see [docs/development.md](docs/development.md).

For usage details, see [docs/usage.md](docs/usage.md).

## Contributing

We welcome contributions to the project! Please read our [CONTRIBUTING.md](CONTRIBUTING.md) guide for details on how to get started.

This project extends the functionality of the existing chezmoi tool to provide enhanced TUI and CLI experiences.

## Features

- Interactive Terminal User Interface (TUI) for managing dotfiles
- Enhanced Command Line Interface (CLI) with additional commands
- Seamless integration with existing chezmoi workflows
- Comprehensive test suite covering unit, integration, and performance tests
- CI/CD pipeline with automated testing and releases
- Docker support for containerized usage
- Cross-platform compatibility