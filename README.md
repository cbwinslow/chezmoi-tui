# Chezmoi TUI - Enhanced Terminal User Interface for Chezmoi

An enhanced Terminal User Interface (TUI) and CLI for managing your dotfiles with chezmoi.

[![Go Report Card](https://goreportcard.com/badge/github.com/cbwinslow/chezmoi-tui)](https://goreportcard.com/report/github.com/cbwinslow/chezmoi-tui)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.21-blue)](https://golang.org/dl/)

## Table of Contents
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Architecture](#architecture)
- [CI/CD & Security](#cicd--security)
- [Development](#development)
- [Contributing](#contributing)
- [GitLab Mirror](#gitlab-mirror)
- [Project Status](#project-status)

## Features

- **Interactive Terminal UI**: Modern TUI built with Bubble Tea for visual dotfile management
- **Enhanced CLI**: Extended command set with improved feedback and functionality
- **Seamless Integration**: Full compatibility with existing chezmoi workflows
- **Security First**: Proper secrets management using Chezmoi encryption
- **Cross-Platform**: Builds for Linux, macOS, and Windows (AMD64 & ARM64)
- **Comprehensive Testing**: Unit, integration, performance, and end-to-end tests
- **Docker Support**: Multi-architecture container images available
- **Advanced CI/CD**: Automated testing, security scanning, and releases

## Installation

### Prerequisites
- Go 1.21 or later
- Chezmoi installed and in your PATH

### Building from source
```bash
git clone https://github.com/cbwinslow/chezmoi-tui.git
cd chezmoi-tui
go build -o chezmoi-tui .
```

### Using Go install
```bash
go install github.com/cbwinslow/chezmoi-tui@latest
```

### Docker
```bash
# Pull the latest image
docker pull ghcr.io/cbwinslow/chezmoi-tui:latest

# Run in container
docker run -it --rm -v ~/.config/chezmoi:/home/nobody/.config/chezmoi ghcr.io/cbwinslow/chezmoi-tui:latest
```

## Usage

### CLI Commands
- `chezmoi-tui version` - Print the application version
- `chezmoi-tui status` - Show the status of targets
- `chezmoi-tui apply [targets...]` - Update destination directory to match target state
- `chezmoi-tui add [targets...]` - Add targets to the source state
- `chezmoi-tui diff [targets...]` - Show differences between target and destination states
- `chezmoi-tui init [repo]` - Initialize the source directory with a remote repository
- `chezmoi-tui tui` - Launch the Terminal User Interface

### TUI Interface
Launch the interactive terminal interface with `chezmoi-tui tui`:
- Use arrow keys or hjkl to navigate
- Press 'l'/'right' to view file status details
- Press 'h'/'left' to return to the main menu
- Press 'q' or Ctrl+C to quit

## Architecture

The system follows a modular architecture with separate components for CLI and TUI functionality, integrated with the existing chezmoi backend.

- **Command Layer**: CLI commands using Cobra framework
- **Integration Layer**: High-level operations and error handling
- **Wrapper Layer**: Low-level interaction with chezmoi binary
- **Presentation Layer**: TUI using Bubble Tea framework

For detailed architecture, see [docs/architecture.md](docs/architecture.md).

## CI/CD & Security

### GitHub Actions Workflows
- **CI Pipeline**: Automated testing across multiple platforms and Go versions
- **Security Scanning**: Trivy-based vulnerability scanning
- **Release Pipeline**: Automated binary generation for multiple platforms
- **Docker Builds**: Multi-architecture container image building

### Secrets Management
- Proper encryption using Chezmoi's built-in capabilities
- Environment-based configuration
- Secure credential handling in CI/CD
- No hardcoded secrets in the codebase

For more details, see [docs/github-actions.md](docs/github-actions.md) and [docs/secrets-management.md](docs/secrets-management.md).

## Development

### Getting Started
```bash
# Clone the repository
git clone https://github.com/cbwinslow/chezmoi-tui.git
cd chezmoi-tui

# Install dependencies
go mod download

# Build the project
go build -o chezmoi-tui .

# Run tests
go test ./...
```

### Using Makefile
```bash
# Build the project
make build

# Run tests
make test

# Run tests with coverage
make test-coverage

# Build for all platforms
make build-all
```

For development guidelines, see [docs/development.md](docs/development.md).

## Contributing

We welcome contributions to the project! Please read our [CONTRIBUTING.md](CONTRIBUTING.md) guide for details on how to get started.

For improvement suggestions, see [tasks.md](tasks.md).

## GitLab Mirror

This project is primarily hosted on GitHub. If you need a GitLab mirror, see [docs/gitlab-setup.md](docs/gitlab-setup.md) for setup instructions.

## Project Status

**Active Development**: This project is actively maintained with regular updates and improvements.

### Completed Features
- ✅ Interactive TUI for dotfile management
- ✅ Enhanced CLI with additional commands
- ✅ Full chezmoi integration and compatibility
- ✅ Comprehensive test suite
- ✅ Advanced CI/CD with security scanning
- ✅ Multi-platform support (Docker, binaries)
- ✅ Proper secrets management
- ✅ Complete documentation set

### Planned Enhancements
For upcoming features and improvements, check [tasks.md](tasks.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built as an enhancement to the excellent [chezmoi](https://github.com/twpayne/chezmoi) project
- TUI built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) framework
- CLI commands powered by [Cobra](https://github.com/spf13/cobra) framework