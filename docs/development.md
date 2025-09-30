# Development Guide

## Project Structure

```
chezmoi-tui/
├── cmd/                 # CLI command definitions
├── docs/               # Documentation files
├── examples/           # Example usage files
├── internal/           # Internal packages (not importable by other projects)
│   ├── chezmoi/        # Low-level chezmoi wrapper
│   ├── integration/    # High-level integration layer
│   └── utils/          # Utility functions
├── pkg/                # Shared packages (importable by other projects)
│   └── root/           # Root command definition
├── scripts/            # Utility scripts
├── tests/              # Test files
├── ui/                 # Terminal User Interface components
├── go.mod             # Go module definition
├── go.sum             # Go module checksums
├── main.go            # Application entry point
└── README.md          # Project overview
```

## Setting up Development Environment

1. Install Go 1.19 or later
2. Ensure `chezmoi` is installed and available in your PATH
3. Clone the repository:
   ```bash
   git clone <repository-url>
   cd chezmoi-tui
   ```

## Building

```bash
go build -o chezmoi-tui .
```

## Testing

Run all tests:
```bash
go test ./...
```

Run tests for a specific package:
```bash
go test ./internal/chezmoi
```

Run tests with verbose output:
```bash
go test -v ./...
```

## Adding New Features

### CLI Commands
1. Create a new file in `pkg/commands/` (e.g., `newcmd.go`)
2. Define your command using Cobra
3. Add the command to the root command in the init() function
4. Ensure the command package is imported in `main.go`

### TUI Components
1. Add new functionality to the `ui/` package
2. Update the Model struct as needed
3. Implement Update and View functions following Bubble Tea patterns
4. Add any new dependencies to go.mod if needed

### Integration Layer
1. Add new methods to the integration package in `internal/integration/`
2. These methods should wrap and enhance the low-level chezmoi operations
3. Ensure proper error handling and data transformation

## Code Style

- Follow Go standard formatting (use `go fmt`)
- Write clear, descriptive comments
- Add tests for new functionality
- Follow the existing project architecture patterns
- Use meaningful variable and function names

## Dependency Management

Add new dependencies with:
```bash
go get <package-name>
```

Tidy dependencies after adding/removing:
```bash
go mod tidy
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Ensure all tests pass
6. Submit a pull request