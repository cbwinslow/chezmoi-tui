# Chezmoi TUI Usage Guide

## Overview

Chezmoi TUI is an enhanced Terminal User Interface and CLI for managing your dotfiles with chezmoi. It provides both a command-line interface with improved functionality and an interactive terminal user interface.

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

## CLI Usage

### Available Commands

- `chezmoi-tui status` - Show the status of targets (similar to git status)
- `chezmoi-tui apply [targets...]` - Update destination directory to match target state
- `chezmoi-tui add [targets...]` - Add targets to the source state
- `chezmoi-tui diff [targets...]` - Show differences between target and destination states
- `chezmoi-tui tui` - Launch the Terminal User Interface
- `chezmoi-tui version` - Print version information

### Examples

```bash
# View status of all managed files
chezmoi-tui status

# Apply changes to all files
chezmoi-tui apply

# Add a specific file to management
chezmoi-tui add ~/.bashrc

# Show differences for a specific file
chezmoi-tui diff ~/.bashrc
```

## TUI Usage

The Terminal User Interface provides an interactive way to manage your dotfiles:

### Main Menu
- Use arrow keys (or hjkl) to navigate the menu options
- Press Enter to select an option
- Press 'q' or Ctrl+C to quit

### File Status View
- Press 'l' or 'right arrow' when on "View Status" to see detailed file status
- Navigate files with arrow keys (or hjkl)
- Press 'h' or 'left arrow' to return to main menu
- View status symbols: M (Modified), A (Added/Unmanaged), I (Ignored)

## Architecture

The application is structured as follows:

```
chezmoi-tui/
├── cmd/                 # CLI command definitions
├── internal/            # Internal packages
│   ├── chezmoi/        # Wrapper around chezmoi functionality
│   ├── integration/    # High-level operations integration
│   └── utils/          # Utility functions
├── ui/                  # TUI components
├── docs/               # Documentation
└── pkg/                # Shared packages
    └── root/           # Root command definition
```

## Development

### Running Tests

```bash
go test ./...
```

### Building

```bash
go build -o chezmoi-tui .
```

## Features

- Enhanced CLI with improved feedback and error handling
- Interactive TUI for visual status management
- Integration with existing chezmoi functionality
- Real-time status updates
- Keyboard navigation
- Compatible with existing chezmoi workflows