# Contributing to Chezmoi TUI

Thank you for your interest in contributing to Chezmoi TUI! This document outlines the processes and guidelines for contributing.

## Code of Conduct

Please follow our Code of Conduct in all interactions. We are committed to providing a welcoming and inclusive environment for everyone.

## How to Contribute

### Reporting Bugs
- Use the GitHub issue tracker to report bugs
- Verify the bug exists in the latest version
- Explain the problem clearly with steps to reproduce
- Include system information (OS, Go version, etc.)

### Suggesting Features
- Use the GitHub issue tracker for feature requests
- Explain the feature clearly and why it would be useful
- Check existing issues to avoid duplicates

### Pull Requests
1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Add tests for new functionality
5. Update documentation as needed
6. Ensure all tests pass (`go test ./...`)
7. Commit your changes (`git commit -m 'Add amazing feature'`)
8. Push to the branch (`git push origin feature/amazing-feature`)
9. Open a pull request

## Development Setup

1. Install Go 1.19 or later
2. Fork and clone the repository
3. Install dependencies: `go mod download`
4. Build the project: `go build -o chezmoi-tui .`

## Testing

- Write tests for new functionality
- Run all tests before submitting: `go test ./...`
- Add benchmarks for performance-critical code
- Ensure test coverage remains high

## Code Style

- Follow Go standard formatting (`go fmt`)
- Write clear, descriptive comments
- Add GoDoc comments for exported functions/types
- Keep functions focused and small when possible
- Use meaningful variable and function names

## Architecture Overview

The project follows a modular architecture:

- `cmd/` - CLI command definitions
- `internal/chezmoi/` - Low-level wrapper around chezmoi functionality
- `internal/integration/` - High-level integration layer
- `ui/` - Terminal User Interface components
- `pkg/root/` - Root command definition
- `tests/` - Test suites

## Commit Messages

Use clear, descriptive commit messages in the format:
```
Add feature name
Fix bug in component
Update documentation for feature
```

## Questions?

If you have questions, feel free to open an issue for discussion.

Thank you for contributing!