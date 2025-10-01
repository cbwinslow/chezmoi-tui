# Chezmoi TUI - Terminal User Interface

This document details the Terminal User Interface component of the Chezmoi TUI project.

## Overview

The TUI (Terminal User Interface) provides an interactive way to manage your dotfiles with a visual interface in the terminal. Built using the Bubble Tea framework, it offers a modern, responsive interface for common chezmoi operations.

## Features

### Main Menu
- **View Status**: See the current status of all managed files
- **Add Files**: Add new files to chezmoi management
- **Apply Changes**: Apply managed files to your system
- **Diff Changes**: View differences between source and destination
- **Edit Config**: Edit configuration files
- **Exit**: Quit the TUI

### File Status View
- Visual representation of file states (Modified, Added, Unmanaged, etc.)
- Navigation with keyboard controls
- Color-coded status indicators
- Real-time status updates

## Navigation

The TUI supports multiple navigation methods:

### Keyboard Controls
- **Arrow Keys** / **hjkl**: Navigate menu and file lists
- **Enter**: Select menu items or confirm actions
- **Right Arrow** / **l**: Enter file status view from main menu
- **Left Arrow** / **h**: Return to main menu from file view
- **q** / **Ctrl+C**: Quit the application

### File Status Navigation
- Use arrow keys to move cursor in file list
- See status symbols: M (Modified), A (Added), D (Deleted), etc.
- Navigate between main menu and file view using arrow keys

## Architecture

The TUI follows the Bubble Tea architecture pattern:

```
Model: Represents the state of the UI
- Main menu state
- File status state
- Navigation state

Update: Handles messages and updates the model
- Keyboard input handling
- Navigation logic
- Chezmoi command execution

View: Renders the UI to the terminal
- Menu rendering
- File status display
- Status information
```

## Integration with Chezmoi

The TUI integrates with the core chezmoi functionality through the internal integration layer:
- Status checking using `chezmoi status`
- File management using `chezmoi add`/`apply`/`diff`
- Configuration management
- Error handling and feedback

## Future Enhancements

Potential improvements for the TUI component:
- More detailed file operations (per-file apply, diff, etc.)
- Search and filter capabilities
- Batch operations
- Configuration editing capabilities
- Plugin system for custom commands