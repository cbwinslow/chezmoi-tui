#!/bin/bash
# release.sh - Script to build and package chezmoi-tui for release

set -e

# Default values
VERSION=${1:-$(git describe --tags --always --dirty="-dev" 2>/dev/null || echo "v0.0.0")}
BINARY_NAME="chezmoi-tui"
DIST_DIR="dist"
REPO_NAME="cbwinslow/chezmoi-tui"

echo "Building release for version: $VERSION"

# Create dist directory
mkdir -p $DIST_DIR

# Build for different platforms
echo "Building for Linux AMD64..."
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-X main.Version=$VERSION" -o $DIST_DIR/${BINARY_NAME}-linux-amd64 .

echo "Building for Linux ARM64..."
GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -ldflags="-X main.Version=$VERSION" -o $DIST_DIR/${BINARY_NAME}-linux-arm64 .

echo "Building for macOS AMD64..."
GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-X main.Version=$VERSION" -o $DIST_DIR/${BINARY_NAME}-darwin-amd64 .

echo "Building for macOS ARM64..."
GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -ldflags="-X main.Version=$VERSION" -o $DIST_DIR/${BINARY_NAME}-darwin-arm64 .

echo "Building for Windows AMD64..."
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-X main.Version=$VERSION" -o $DIST_DIR/${BINARY_NAME}-windows-amd64.exe .

# Create archives
cd $DIST_DIR

echo "Creating archives..."
tar -czf ${BINARY_NAME}-${VERSION}-linux-amd64.tar.gz ${BINARY_NAME}-linux-amd64
tar -czf ${BINARY_NAME}-${VERSION}-linux-arm64.tar.gz ${BINARY_NAME}-linux-arm64
tar -czf ${BINARY_NAME}-${VERSION}-darwin-amd64.tar.gz ${BINARY_NAME}-darwin-amd64
tar -czf ${BINARY_NAME}-${VERSION}-darwin-arm64.tar.gz ${BINARY_NAME}-darwin-arm64

# For Windows, create zip instead
zip -r ${BINARY_NAME}-${VERSION}-windows-amd64.zip ${BINARY_NAME}-windows-amd64.exe

# Create checksums
echo "Creating checksums..."
for file in *.tar.gz *.zip; do
    if [ -f "$file" ]; then
        sha256sum "$file" > "${file}.sha256"
    fi
done

# Create overall checksums file
echo "Creating overall checksums..."
sha256sum $(ls *.tar.gz *.zip 2>/dev/null) > checksums.txt

cd ..

echo "Release build completed! Files are in $DIST_DIR/"
ls -la $DIST_DIR/

echo ""
echo "To create a Debian package, run:"
echo "  make package-deb"