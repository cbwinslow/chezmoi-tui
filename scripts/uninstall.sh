#!/bin/bash
# uninstall.sh - Uninstall Chezmoi TUI by removing the symlink

set -e

# Default values
INSTALL_DIR="${HOME}/.local/bin"
BINARY_NAME="chezmoi-tui"

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
    echo "Uninstall Chezmoi TUI by removing the symlink."
    echo ""
    echo "OPTIONS:"
    echo "  -d, --dir DIR        Installation directory (default: ~/.local/bin)"
    echo "  -n, --name NAME      Binary name (default: chezmoi-tui)"
    echo "  -h, --help           Show this help message"
    echo ""
    echo "EXAMPLES:"
    echo "  $0                    # Uninstall from ~/.local/bin/chezmoi-tui"
    echo "  $0 -d /usr/local/bin  # Uninstall from /usr/local/bin/chezmoi-tui"
    echo "  $0 -n cmt             # Uninstall cmt binary"
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

# Check if destination exists
DEST_PATH="${INSTALL_DIR}/${BINARY_NAME}"
if [[ ! -e "${DEST_PATH}" ]] && [[ ! -L "${DEST_PATH}" ]]; then
    print_warning "Destination does not exist: ${DEST_PATH}"
    print_warning "Nothing to uninstall."
    exit 0
fi

# Check if it is a symlink
if [[ ! -L "${DEST_PATH}" ]]; then
    print_warning "Destination is not a symlink: ${DEST_PATH}"
    print_warning "This might be a regular file. Are you sure you want to remove it?"
    read -p "Remove file? (y/N): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        print_info "Aborting uninstall."
        exit 0
    fi
fi

# Remove the symlink/file
print_info "Removing: ${DEST_PATH}"
rm -f "${DEST_PATH}"

# Verify removal
if [[ ! -e "${DEST_PATH}" ]] && [[ ! -L "${DEST_PATH}" ]]; then
    print_success "Successfully uninstalled ${BINARY_NAME} from ${INSTALL_DIR}"
else
    print_error "Failed to remove ${DEST_PATH}"
    exit 1
fi

print_success "Chezmoi TUI uninstalled successfully"
