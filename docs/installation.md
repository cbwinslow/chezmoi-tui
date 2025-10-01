# Installation Guide

This guide provides step-by-step instructions for installing Chezmoi TUI on various platforms.

## Prerequisites

Before installing Chezmoi TUI, ensure you have the following prerequisites:

1. **Go 1.21 or later** - Required for building from source
2. **Chezmoi** - The underlying dotfile management tool
3. **Bitwarden CLI (optional)** - For Bitwarden integration features

### Installing Prerequisites

#### Go Installation

**Ubuntu/Debian:**
```bash
sudo apt update
sudo apt install golang
```

**macOS:**
```bash
brew install go
```

**Windows:**
Download from [golang.org](https://golang.org/dl/)

#### Chezmoi Installation

**Using Package Managers:**
```bash
# Ubuntu/Debian
sudo apt install chezmoi

# macOS
brew install chezmoi

# Windows (using Chocolatey)
choco install chezmoi
```

**Using Script:**
```bash
curl -sfL https://get.chezmoi.io | sh
```

#### Bitwarden CLI Installation (Optional)

**Using Package Managers:**
```bash
# Ubuntu/Debian
sudo snap install bw

# macOS
brew install bitwarden-cli

# Windows (using Chocolatey)
choco install bitwarden-cli
```

**Using NPM:**
```bash
npm install -g @bitwarden/cli
```

## Installation Methods

### Method 1: Pre-built Binaries (Recommended)

Download the latest release from the [GitHub Releases page](https://github.com/cbwinslow/chezmoi-tui/releases).

**Linux:**
```bash
wget https://github.com/cbwinslow/chezmoi-tui/releases/latest/download/chezmoi-tui-linux-amd64.tar.gz
tar -xzf chezmoi-tui-linux-amd64.tar.gz
sudo mv chezmoi-tui /usr/local/bin/
```

**macOS:**
```bash
wget https://github.com/cbwinslow/chezmoi-tui/releases/latest/download/chezmoi-tui-darwin-amd64.tar.gz
tar -xzf chezmoi-tui-darwin-amd64.tar.gz
sudo mv chezmoi-tui /usr/local/bin/
```

**Windows:**
Download the Windows executable and place it in your PATH.

### Method 2: Building from Source

**Clone the Repository:**
```bash
git clone https://github.com/cbwinslow/chezmoi-tui.git
cd chezmoi-tui
```

**Build the Project:**
```bash
go build -o chezmoi-tui .
```

**Install Globally:**
```bash
sudo cp chezmoi-tui /usr/local/bin/
```

### Method 3: Using Go Install

```bash
go install github.com/cbwinslow/chezmoi-tui@latest
```

## Verification

After installation, verify that Chezmoi TUI is working correctly:

```bash
chezmoi-tui --version
chezmoi-tui --help
```

You should see version information and help text.

## Post-Installation Setup

### Basic Configuration

Initialize Chezmoi TUI with your existing dotfiles repository:

```bash
chezmoi-tui init <your-dotfiles-repo-url>
```

Or create a new dotfiles repository:

```bash
chezmoi-tui init --apply
```

### First-time Configuration

Generate a default configuration file:

```bash
chezmoi-tui config generate
```

## Platform-Specific Instructions

### Linux

On Linux systems, ensure the binary has execute permissions:

```bash
chmod +x /usr/local/bin/chezmoi-tui
```

### macOS

On macOS, you may need to grant permission to run the application:

```bash
xattr -d com.apple.quarantine /usr/local/bin/chezmoi-tui
```

### Windows

On Windows, ensure the binary is added to your PATH environment variable.

## Troubleshooting

### Common Issues

1. **Command not found**: Ensure the binary is in your PATH
2. **Permission denied**: Check file permissions with `ls -l`
3. **Version mismatch**: Ensure all dependencies are compatible

### Getting Help

If you encounter issues during installation:

1. Check the [Troubleshooting Guide](troubleshooting.md)
2. Review the [FAQ](faq.md)
3. Open an issue on [GitHub](https://github.com/cbwinslow/chezmoi-tui/issues)

## Next Steps

After successful installation:

1. [Configure Chezmoi TUI](configuration.md)
2. [Explore CLI Commands](cli-commands.md)
3. [Try the TUI](tui-user-guide.md)
4. [Set up Bitwarden Integration](bitwarden-integration.md)

## Updating

To update to the latest version:

```bash
# If installed via go install
go install github.com/cbwinslow/chezmoi-tui@latest

# If installed from binaries
# Download and replace the binary as described in Method 1
```