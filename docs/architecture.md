# Chezmoi TUI - Architecture Overview

This document provides a comprehensive overview of the Chezmoi TUI project architecture.

## Project Structure

```
chezmoi-tui/
├── cmd/                    # CLI command definitions
│   └── main.go             # Application entry point
├── docs/                   # Documentation files
├── examples/               # Example usage files
├── internal/               # Internal packages (not importable by other projects)
│   ├── chezmoi/            # Low-level chezmoi wrapper
│   ├── integration/        # High-level integration layer
│   └── utils/              # Utility functions
├── pkg/                    # Shared packages (importable by other projects)
│   ├── commands/           # CLI command implementations
│   └── root/               # Root command definition
├── scripts/                # Utility scripts
├── tests/                  # Test files
│   ├── e2e/                # End-to-end tests
│   ├── performance/        # Performance tests
│   └── unit/               # Unit tests
├── ui/                     # Terminal User Interface components
├── .github/                # GitHub Actions workflows
│   └── workflows/          # CI/CD configuration
├── go.mod                  # Go module definition
├── go.sum                  # Go module checksums
├── Dockerfile              # Container configuration
├── Makefile                # Build and development tasks
└── README.md               # Project overview
```

## Core Architecture Layers

### 1. Presentation Layer
- **UI Package** (`ui/`): Bubble Tea based TUI
- **Commands Package** (`pkg/commands/`): CLI command definitions
- **Root Package** (`pkg/root/`): Main command configuration

### 2. Business Logic Layer
- **Integration Package** (`internal/integration/`): High-level operations
- **Chezmoi Package** (`internal/chezmoi/`): Low-level chezmoi wrapper

### 3. External Dependencies
- Original `chezmoi` binary
- Go standard library
- Third-party libraries (Cobra, Bubble Tea, etc.)

## Design Patterns

### Command Pattern (CLI)
- Uses Cobra framework for CLI command structure
- Each command is a separate package following the same pattern
- Commands use the integration layer for functionality
- Consistent error handling and return types

### Model-View-Update (TUI)
- Uses Bubble Tea framework for TUI
- Model represents the state of the UI
- Update function handles messages and state changes
- View function renders the UI

### Facade Pattern (Integration Layer)
- `internal/integration` provides a unified interface
- Hides complexity of the underlying chezmoi wrapper
- Provides enhanced functionality and error handling
- Abstracts away direct interactions with chezmoi

## Data Flow

### CLI Data Flow
1. User executes CLI command
2. Cobra routes to appropriate command handler
3. Command calls integration layer methods
4. Integration layer calls chezmoi wrapper
5. Chezmoi wrapper executes chezmoi binary
6. Results are processed and returned to user

### TUI Data Flow
1. TUI model receives user input
2. Model calls integration layer methods
3. Integration layer calls chezmoi wrapper
4. Results are processed and update the model
5. Updated model is rendered to the terminal

## Security Architecture

### Secrets Management
- Uses Chezmoi's built-in encryption for sensitive files
- Environment variables for runtime secrets
- GitHub Actions secrets for CI/CD
- No hardcoded credentials in source code

### Access Control
- All external command execution is validated
- Input validation for all user-provided values
- Proper error handling to prevent information disclosure

## Build and Deployment

### Local Development
- Standard Go build process
- Makefile for common tasks
- Docker support for containerization

### CI/CD Pipeline
- Automated testing on multiple platforms
- Security scanning
- Cross-platform builds
- Automated releases
- Docker image building and publishing

## Error Handling Strategy

### CLI Error Handling
- Proper exit codes for different failure types
- User-friendly error messages
- Detailed error logging for debugging
- Graceful degradation when possible

### TUI Error Handling
- Non-fatal errors displayed in UI
- Recovery from common errors
- Clear error messages to the user
- Preserved UI state when possible

## Testing Architecture

### Unit Tests
- Located in `internal/*/` directories alongside code
- Test individual functions and methods
- Mock external dependencies where appropriate
- Focus on core business logic

### Integration Tests
- Test interactions between components
- Validate integration layer functionality
- Test real interactions with chezmoi when available

### Performance Tests
- Located in `tests/performance/`
- Benchmark critical operations
- Monitor response times
- Track performance regressions

### End-to-End Tests
- Located in `tests/e2e/`
- Test complete user workflows
- Validate integration between all components
- Test CLI and TUI functionality

## Extensibility Points

### Adding CLI Commands
1. Create new file in `pkg/commands/`
2. Define command using Cobra
3. Implement using integration layer
4. Register with root command

### Adding TUI Features
1. Extend the Model struct
2. Update the Update function to handle new messages
3. Update the View function to render new elements
4. Add new keyboard bindings

### Adding Integration Features
1. Add new methods to integration layer
2. Implement using chezmoi wrapper
3. Add appropriate error handling
4. Update tests and documentation

## Deployment Architecture

### Containerized Deployment
- Multi-architecture Docker images
- Minimal base images for security
- Non-root execution
- Proper resource limits

### Binary Distribution
- Cross-platform builds
- Signed releases
- Checksum verification
- Multiple installation methods