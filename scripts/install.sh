#!/bin/bash
# install.sh - Install Chezmoi TUI by creating a symlink

set -e

# Default values
INSTALL_DIR="${HOME}/.local/bin"
BINARY_NAME="chezmoi-tui"
SOURCE_DIR="$(pwd)"

# Colors for output
RED="\033[0;31m"
GREEN="\033[0;32m"
YELLOW="\033[1;33m"
BLUE="\033[0;34m"
NC="\033[0m" # No Color

# Print usage information
usage() {
    echo "Usage: $0 [OPTIONS]"
    echo ""
    echo "Install Chezmoi TUI by creating a symlink."
    echo ""
    echo "OPTIONS:"
    echo "  -d, --dir DIR        Installation directory (default: ~/.local/bin)"
    echo "  -n, --name NAME      Binary name (default: chezmoi-tui)"
    echo "  -f, --force          Force overwrite existing symlink"
    echo "  -h, --help           Show this help message"
    echo ""
    echo "EXAMPLES:"
    echo "  $0                    # Install to ~/.local/bin/chezmoi-tui"
    echo "  $0 -d /usr/local/bin  # Install to /usr/local/bin/chezmoi-tui"
    echo "  $0 -n cmt             # Install as cmt binary"
}

# Print colored output
print_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1" >&2
}

# Parse command line arguments
FORCE=""
while [[ $# -gt 0 ]]; do
    case $1 in
        -d|--dir)
            INSTALL_DIR="$2"
            shift 2
            ;;
        -n|--name)
            BINARY_NAME="$2"
            shift 2
            ;;
        -f|--force)
            FORCE="true"
            shift
            ;;
        -h|--help)
            usage
            exit 0
            ;;
        *)
            print_error "Unknown option: $1"
            usage
            exit 1
            ;;
    esac
done

# Check if source binary exists
BINARY_PATH="${SOURCE_DIR}/${BINARY_NAME}"
if [[ ! -f "${BINARY_PATH}" ]]; then
    # Try with .exe extension for Windows
    BINARY_PATH="${SOURCE_DIR}/${BINARY_NAME}.exe"
    if [[ ! -f "${BINARY_PATH}" ]]; then
        print_error "Binary not found at ${SOURCE_DIR}/${BINARY_NAME} or ${SOURCE_DIR}/${BINARY_NAME}.exe"
        print_error "Please build the binary first with go build -o ${BINARY_NAME} ."
        exit 1
    fi
fi

# Create installation directory if it does not exist
if [[ ! -d "${INSTALL_DIR}" ]]; then
    print_info "Creating installation directory: ${INSTALL_DIR}"
    mkdir -p "${INSTALL_DIR}"
fi

# Check if installation directory is in PATH
if [[ ":$PATH:" != *":${INSTALL_DIR}:"* ]]; then
    print_warning "Installation directory ${INSTALL_DIR} is not in your PATH"
    echo "Add it to your PATH by adding this line to your shell configuration (~/.bashrc, ~/.zshrc, etc.):"
    echo "export PATH=\"\$PATH:${INSTALL_DIR}\""
fi

# Check if destination already exists
DEST_PATH="${INSTALL_DIR}/${BINARY_NAME}"
if [[ -e "${DEST_PATH}" ]] || [[ -L "${DEST_PATH}" ]]; then
    if [[ "${FORCE}" == "true" ]]; then
        print_info "Removing existing file/symlink: ${DEST_PATH}"
        rm -f "${DEST_PATH}"
    else
        print_error "Destination already exists: ${DEST_PATH}"
        print_error "Use --force to overwrite or remove it manually"
        exit 1
    fi
fi

# Create symlink
print_info "Creating symlink: ${DEST_PATH} -> ${BINARY_PATH}"
ln -s "${BINARY_PATH}" "${DEST_PATH}"

# Set executable permissions
chmod +x "${DEST_PATH}"

# Verify installation
print_info "Verifying installation..."
if command -v "${BINARY_NAME}" >/dev/null 2>&1; then
    VERSION_OUTPUT=$("${BINARY_NAME}" --version 2>&1)
    print_success "Installation successful!"
    echo "Installed version: ${VERSION_OUTPUT}"
    echo "You can now use ${BINARY_NAME} from anywhere in your system."
else
    print_warning "Installation completed, but verification failed."
    echo "You may need to restart your shell or add ${INSTALL_DIR} to your PATH."
fi

print_success "Chezmoi TUI installed successfully to ${DEST_PATH}"
