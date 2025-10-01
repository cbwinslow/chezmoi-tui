# Quick Start Guide

Get up and running with Chezmoi TUI quickly with this guide.

## Installation

### Option 1: Pre-built Binary (Recommended)

```bash
# Download the latest release
wget https://github.com/cbwinslow/chezmoi-tui/releases/latest/download/chezmoi-tui-linux-amd64.tar.gz
tar -xzf chezmoi-tui-linux-amd64.tar.gz
sudo mv chezmoi-tui /usr/local/bin/

# Verify installation
chezmoi-tui --version
```

### Option 2: Build from Source

```bash
# Clone the repository
git clone https://github.com/cbwinslow/chezmoi-tui.git
cd chezmoi-tui

# Build the project
go build -o chezmoi-tui .

# Install globally (optional)
sudo cp chezmoi-tui /usr/local/bin/
```

## Initialize Your Dotfiles

### With Existing Repository

```bash
# Initialize with your dotfiles repository
chezmoi-tui init https://github.com/username/dotfiles.git

# Apply the configuration
chezmoi-tui apply
```

### Create New Repository

```bash
# Initialize and create new repository
chezmoi-tui init --apply

# Add your first file
chezmoi-tui add ~/.bashrc

# Commit and push to GitHub
chezmoi-tui git add .
chezmoi-tui git commit -m "Initial commit"
chezmoi-tui git push origin master
```

## Core Workflow

### Check Status

```bash
# View the status of your managed files
chezmoi-tui status
```

Output example:
```
 M .bashrc              # Modified file
 A .gitconfig           # Added file
   .vimrc                # Up-to-date file
 D .old-config          # Deleted file
```

### Add New Files

```bash
# Add a single file
chezmoi-tui add ~/.tmux.conf

# Add multiple files
chezmoi-tui add ~/.bashrc ~/.vimrc ~/.gitconfig

# Add with encryption (for sensitive files)
chezmoi-tui add --encrypt ~/.ssh/id_rsa
```

### Apply Changes

```bash
# Apply all changes
chezmoi-tui apply

# Apply with preview
chezmoi-tui apply --dry-run

# Apply specific files
chezmoi-tui apply ~/.bashrc
```

### View Differences

```bash
# Show all differences
chezmoi-tui diff

# Show differences for specific files
chezmoi-tui diff ~/.bashrc
```

## Using the TUI

### Launch the Interface

```bash
# Launch the Terminal User Interface
chezmoi-tui tui
```

### Navigate the TUI

1. **Main Menu**: Use arrow keys to navigate options
2. **Select**: Press `Enter` to select an option
3. **Go Back**: Press `h` or left arrow to return to previous screen
4. **Quit**: Press `q` or `Ctrl+C` to exit

### Main Menu Options

- **View Status**: See current status of all managed files
- **Add Files**: Add new files to management
- **Apply Changes**: Apply managed files to your system
- **Diff Changes**: Show differences between source and destination
- **Show Stats**: View detailed statistics about your dotfiles
- **Bitwarden Manager**: Manage Bitwarden secrets integration
- **Exit**: Quit the application

## Bitwarden Integration (Optional)

### Prerequisites

Install Bitwarden CLI:
```bash
# Ubuntu/Debian
sudo snap install bw

# macOS
brew install bitwarden-cli

# Or using npm
npm install -g @bitwarden/cli
```

### Unlock Your Vault

```bash
# Unlock the vault
chezmoi-tui bitwarden unlock

# Check status
chezmoi-tui bitwarden status
```

### Generate Templates

```bash
# Generate Chezmoi template from Bitwarden item
chezmoi-tui bitwarden template "api-keys"

# Apply the template
chezmoi-tui apply
```

### Export to Environment File

```bash
# Export secrets to .env file
chezmoi-tui bitwarden export .env.development
```

## Configuration

### Generate Default Config

```bash
# Generate default configuration file
chezmoi-tui config generate
```

Location: `~/.config/chezmoi-tui/config.yaml`

### Sample Configuration

```yaml
# ~/.config/chezmoi-tui/config.yaml
theme:
  primary_color: "#1793d1"
  secondary_color: "#0366d6"

tui:
  show_help: true
  refresh_interval: 5

cli:
  verbose: false
  color: true
```

## Daily Workflow Example

Here's a typical daily workflow:

```bash
# 1. Check for changes
chezmoi-tui status

# 2. Add any new configuration files
chezmoi-tui add ~/.new-config-file

# 3. Apply changes to system
chezmoi-tui apply

# 4. Commit and push to repository
chezmoi-tui git add .
chezmoi-tui git commit -m "Add new config file"
chezmoi-tui git push

# 5. Check statistics
chezmoi-tui stats
```

## Useful Commands

### Git Integration

```bash
# Initialize Git repository
chezmoi-tui git init

# Add all changes
chezmoi-tui git add .

# Commit changes
chezmoi-tui git commit -m "Update dotfiles"

# Push to remote
chezmoi-tui git push origin master
```

### Advanced Features

```bash
# Show detailed statistics
chezmoi-tui stats --details

# List managed files
chezmoi-tui managed

# List unmanaged files
chezmoi-tui unmanaged

# Verify integrity
chezmoi-tui verify
```

## Next Steps

After completing the quick start:

1. **[Read the Full Documentation](README.md)** - Dive deeper into all features
2. **[Explore CLI Commands](cli-commands.md)** - Learn about all available commands
3. **[Master the TUI](tui-user-guide.md)** - Become proficient with the interface
4. **[Set up Bitwarden Integration](bitwarden-integration.md)** - Securely manage secrets
5. **[Configure Advanced Features](advanced-features.md)** - Customize to your needs

## Getting Help

```bash
# Show help
chezmoi-tui --help

# Show help for specific command
chezmoi-tui <command> --help

# Enable debug output
chezmoi-tui --debug <command>
```

## Community Resources

- **[GitHub Repository](https://github.com/cbwinslow/chezmoi-tui)**
- **[Issue Tracker](https://github.com/cbwinslow/chezmoi-tui/issues)**
- **[Discussion Forum](https://github.com/cbwinslow/chezmoi-tui/discussions)**