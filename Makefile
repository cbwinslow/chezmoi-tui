# Makefile for chezmoi-tui

# Variables
BINARY_NAME=chezmoi-tui
VERSION?=$(shell git describe --tags --always --dirty="-dev" 2>/dev/null || echo "v0.0.0")
BUILD_DIR=dist
DIST_DIR=dist
GOARCH?=amd64

# Build the project
build:
	go build -o ${BINARY_NAME} .

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

# Build for different architectures
build-linux-amd64:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${DIST_DIR}/${BINARY_NAME}-linux-amd64 .

build-linux-arm64:
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o ${DIST_DIR}/${BINARY_NAME}-linux-arm64 .

build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o ${DIST_DIR}/${BINARY_NAME}-darwin-amd64 .

build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -o ${DIST_DIR}/${BINARY_NAME}-darwin-arm64 .

build-windows-amd64:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o ${DIST_DIR}/${BINARY_NAME}-windows-amd64.exe .

# Build for all platforms
build-all: build-linux-amd64 build-linux-arm64 build-darwin-amd64 build-darwin-arm64 build-windows-amd64

# Create distributable packages (DEB, RPM, etc.)
package-deb:
	@echo "Creating Debian package..."
	# Create directory structure
	mkdir -p ${DIST_DIR}/pkg/deb/usr/local/bin
	mkdir -p ${DIST_DIR}/pkg/deb/DEBIAN
	# Copy binary
	cp ${DIST_DIR}/${BINARY_NAME}-linux-amd64 ${DIST_DIR}/pkg/deb/usr/local/bin/${BINARY_NAME}
	# Create control file
	echo "Package: ${BINARY_NAME}" > ${DIST_DIR}/pkg/deb/DEBIAN/control
	echo "Version: ${VERSION}" >> ${DIST_DIR}/pkg/deb/DEBIAN/control
	echo "Section: utils" >> ${DIST_DIR}/pkg/deb/DEBIAN/control
	echo "Priority: optional" >> ${DIST_DIR}/pkg/deb/DEBIAN/control
	echo "Architecture: amd64" >> ${DIST_DIR}/pkg/deb/DEBIAN/control
	echo "Depends: golang" >> ${DIST_DIR}/pkg/deb/DEBIAN/control
	echo "Maintainer: Your Name <your.email@example.com>" >> ${DIST_DIR}/pkg/deb/DEBIAN/control
	echo "Description: Enhanced TUI and CLI for chezmoi dotfile management" >> ${DIST_DIR}/pkg/deb/DEBIAN/control
	# Build package
	dpkg-deb --build ${DIST_DIR}/pkg/deb ${DIST_DIR}/${BINARY_NAME}_${VERSION}_amd64.deb

# Build and package for distribution
dist: build-all
	@echo "Creating release artifacts..."
	# Create checksums
	cd ${DIST_DIR} && sha256sum * > checksums.txt && cd ..
	@echo "Release artifacts created in ${DIST_DIR}/"

# Run the application
run: build
	./${BINARY_NAME}

# Docker targets
docker-build:
	docker build -t ${BINARY_NAME}:${VERSION} .

docker-push:
	docker tag ${BINARY_NAME}:${VERSION} ghcr.io/cbwinslow/${BINARY_NAME}:${VERSION}
	docker push ghcr.io/cbwinslow/${BINARY_NAME}:${VERSION}

docker-build-multi:
	docker buildx build --platform linux/amd64,linux/arm64 -t ${BINARY_NAME}:${VERSION} --push .

# Docker Compose
docker-compose-up:
	docker-compose up -d

docker-compose-down:
	docker-compose down

# Clean build artifacts
clean:
	rm -f ${BINARY_NAME}
	rm -rf ${DIST_DIR}/
	rm -f coverage.html
	rm -f coverage.out

# Generate mocks (if using go mock)
mock:
	# Add mock generation here if needed

# Show help
help:
	@echo "Available targets:"
	@echo "  build                 - Build the project"
	@echo "  install               - Install the project"
	@echo "  test                  - Run tests"
	@echo "  test-v                - Run tests with verbose output"  
	@echo "  test-coverage         - Run tests with coverage"
	@echo "  build-linux-amd64     - Build for Linux AMD64"
	@echo "  build-linux-arm64     - Build for Linux ARM64"
	@echo "  build-darwin-amd64    - Build for macOS AMD64"
	@echo "  build-darwin-arm64    - Build for macOS ARM64"
	@echo "  build-windows-amd64   - Build for Windows AMD64"
	@echo "  build-all             - Build for all platforms"
	@echo "  package-deb           - Create Debian package (amd64)"
	@echo "  dist                  - Create distributable packages"
	@echo "  run                   - Build and run the application"
	@echo "  docker-build          - Build Docker image"
	@echo "  docker-push           - Push Docker image to registry"
	@echo "  docker-build-multi    - Build multi-arch Docker images"
	@echo "  clean                 - Clean build artifacts"
	@echo "  help                  - Show this help"