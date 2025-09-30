# Makefile for chezmoi-tui

# Build the project
build:
	go build -o chezmoi-tui .

# Install the project
install:
	go install .

# Run tests
test:
	go test ./...

# Run tests with verbose output
test-v:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Build for multiple platforms
build-all: build-linux build-darwin build-windows

build-linux:
	GOOS=linux GOARCH=amd64 go build -o dist/chezmoi-tui-linux-amd64 .

build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o dist/chezmoi-tui-darwin-amd64 .

build-windows:
	GOOS=windows GOARCH=amd64 go build -o dist/chezmoi-tui-windows-amd64.exe .

# Clean build artifacts
clean:
	rm -f chezmoi-tui
	rm -rf dist/

# Generate mocks (if using go mock)
mock:
	# Add mock generation here if needed

# Run the application
run: build
	./chezmoi-tui

# Show help
help:
	@echo "Available targets:"
	@echo "  build        - Build the project"
	@echo "  install      - Install the project"
	@echo "  test         - Run tests"
	@echo "  test-v       - Run tests with verbose output"  
	@echo "  test-coverage - Run tests with coverage"
	@echo "  build-all    - Build for all platforms"
	@echo "  run          - Build and run the application"
	@echo "  clean        - Clean build artifacts"
	@echo "  help         - Show this help"