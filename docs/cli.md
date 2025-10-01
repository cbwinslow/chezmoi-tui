# Chezmoi TUI - Command Line Interface

This document details the Command Line Interface component of the Chezmoi TUI project.

## Overview

The CLI (Command Line Interface) provides enhanced functionality for chezmoi with additional commands, improved feedback, and better integration capabilities. It maintains compatibility with the original chezmoi CLI while adding new features.

## Commands

### Core Commands

- `chezmoi-tui version` - Display the application version
- `chezmoi-tui help` - Display help information
- `chezmoi-tui completion` - Generate shell completion scripts

### Chezmoi Integration Commands

- `chezmoi-tui status` - Show the status of targets (similar to `git status`)
- `chezmoi-tui apply [targets...]` - Apply targets to destination directory
- `chezmoi-tui add [targets...]` - Add targets to the source state
- `chezmoi-tui diff [targets...]` - Show differences between target and destination states
- `chezmoi-tui init [repo]` - Initialize the source directory with a remote repository
- `chezmoi-tui tui` - Launch the Terminal User Interface

### Command Options

Most commands support common options:
- `-h, --help`: Show help for the command
- `-v, --verbose`: Enable verbose output
- `--config`: Specify config file location
- `--source`: Specify source directory
- `--destination`: Specify destination directory

## Architecture

The CLI is built using the Cobra framework with a modular architecture:

```
pkg/root/ - Main command definition and root configuration
pkg/commands/ - Individual command implementations
internal/integration/ - High-level integration layer
internal/chezmoi/ - Low-level chezmoi wrapper
```

### Command Structure

Each command follows a consistent pattern:
- Defined in `pkg/commands/command-name.go`
- Uses the integration layer for functionality
- Follows Cobra command conventions
- Implements proper error handling

## Integration Layer

The CLI uses an internal integration layer that provides:
- Unified interface to chezmoi functionality
- Enhanced error handling
- Data transformation and parsing
- Consistent return types across commands

## Usage Examples

### Basic Operations
```bash
# Check the status of all managed files
chezmoi-tui status

# Apply all changes
chezmoi-tui apply

# Add a specific file
chezmoi-tui add ~/.bashrc

# Show differences for a specific file
chezmoi-tui diff ~/.vimrc

# Launch the TUI
chezmoi-tui tui
```

### Advanced Operations
```bash
# Apply with verbose output
chezmoi-tui apply -v

# Initialize from a remote repository
chezmoi-tui init user/repo --apply

# Get specific file status
chezmoi-tui status ~/.bashrc
```

## Extensibility

The CLI architecture allows for easy addition of new commands:
1. Create a new file in `pkg/commands/`
2. Define the command using Cobra
3. Implement using the integration layer
4. Add to the root command in the init function

## Future Enhancements

Potential improvements for the CLI component:
- Additional command aliases
- Batch operation support
- Plugin system for custom commands
- Enhanced output formatting (JSON, YAML)
- Performance improvements