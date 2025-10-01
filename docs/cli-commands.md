# CLI Commands Reference

This reference documents all available CLI commands in Chezmoi TUI.

## Global Options

All Chezmoi TUI commands support these global options:

```bash
-h, --help      Show help for the command
-v, --version   Show version information
--config        Specify configuration file path
--verbose       Enable verbose output
--debug         Enable debug logging
```

## Core Commands

### `version`

Display the application version.

```bash
# Show version
chezmoi-tui version

# Output example:
# chezmoi-tui version 0.3.0
```

### `help`

Display help information for commands.

```bash
# Show general help
chezmoi-tui help

# Show help for specific command
chezmoi-tui help <command>

# Show help for subcommand
chezmoi-tui <command> --help
```

### `completion`

Generate shell completion scripts.

```bash
# Generate bash completion
chezmoi-tui completion bash

# Generate zsh completion
chezmoi-tui completion zsh

# Generate fish completion
chezmoi-tui completion fish

# Generate PowerShell completion
chezmoi-tui completion powershell
```

## Dotfile Management Commands

### `status`

Show the status of managed files.

```bash
# Show status of all files
chezmoi-tui status

# Show status with verbose output
chezmoi-tui status --verbose

# Show status in JSON format
chezmoi-tui status --json
```

**Output Format:**
```
 M .bashrc              # Modified file
 A .gitconfig           # Added file
   .vimrc                # Up-to-date file
 D .old-config          # Deleted file
```

### `apply`

Apply managed files to the destination directory.

```bash
# Apply all changes
chezmoi-tui apply

# Apply specific files
chezmoi-tui apply ~/.bashrc ~/.vimrc

# Apply with verbose output
chezmoi-tui apply --verbose

# Apply without confirmation prompts
chezmoi-tui apply --force

# Preview changes without applying
chezmoi-tui apply --dry-run
```

### `add`

Add files to the source state.

```bash
# Add a single file
chezmoi-tui add ~/.bashrc

# Add multiple files
chezmoi-tui add ~/.bashrc ~/.vimrc ~/.gitconfig

# Add with encryption
chezmoi-tui add --encrypt ~/.ssh/id_rsa

# Add as template
chezmoi-tui add --template ~/.bashrc

# Add with automatic templating
chezmoi-tui add --autotemplate ~/.bashrc
```

### `diff`

Show differences between target and destination states.

```bash
# Show differences for all files
chezmoi-tui diff

# Show differences for specific files
chezmoi-tui diff ~/.bashrc

# Show differences in unified format
chezmoi-tui diff --unified

# Show differences with color
chezmoi-tui diff --color
```

### `init`

Initialize the source directory.

```bash
# Initialize with default settings
chezmoi-tui init

# Initialize with remote repository
chezmoi-tui init https://github.com/username/dotfiles.git

# Initialize and apply immediately
chezmoi-tui init --apply

# Initialize with specific branch
chezmoi-tui init --branch develop
```

## Configuration Commands

### `config`

Manage configuration for chezmoi-tui.

```bash
# Generate default configuration
chezmoi-tui config generate

# Show current configuration
chezmoi-tui config show

# Validate configuration
chezmoi-tui config validate

# Edit configuration
chezmoi-tui config edit
```

**Configuration File Location:**
- Linux/macOS: `~/.config/chezmoi-tui/config.yaml`
- Windows: `%APPDATA%\chezmoi-tui\config.yaml`

**Sample Configuration:**
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

integration:
  chezmoi_binary_path: ""
  timeout: 30
```

### `stats`

Show statistics about your dotfiles.

```bash
# Show basic statistics
chezmoi-tui stats

# Show detailed statistics
chezmoi-tui stats --details

# Show statistics in JSON format
chezmoi-tui stats --json
```

## TUI Command

### `tui`

Launch the Terminal User Interface.

```bash
# Launch the TUI
chezmoi-tui tui

# Launch with specific dimensions
chezmoi-tui tui --width 120 --height 40

# Launch in fullscreen mode
chezmoi-tui tui --fullscreen
```

## Bitwarden Commands

### `bitwarden`

Interact with Bitwarden secrets management.

```bash
# Show general Bitwarden help
chezmoi-tui bitwarden --help

# Show vault status
chezmoi-tui bitwarden status

# Unlock vault
chezmoi-tui bitwarden unlock

# Lock vault
chezmoi-tui bitwarden lock

# List items
chezmoi-tui bitwarden list [filter]

# Sync with server
chezmoi-tui bitwarden sync

# Launch Bitwarden TUI
chezmoi-tui bitwarden tui

# Generate Chezmoi template
chezmoi-tui bitwarden template <item-id>

# Export to environment file
chezmoi-tui bitwarden export [filename]
```

## Advanced Commands

### `verify`

Verify that the destination state matches the target state.

```bash
# Verify all files
chezmoi-tui verify

# Verify specific files
chezmoi-tui verify ~/.bashrc

# Verify with verbose output
chezmoi-tui verify --verbose
```

### `managed`

List managed entries in the destination directory.

```bash
# List all managed files
chezmoi-tui managed

# List with full paths
chezmoi-tui managed --full-path

# List in JSON format
chezmoi-tui managed --json
```

### `unmanaged`

List unmanaged files in the destination directory.

```bash
# List all unmanaged files
chezmoi-tui unmanaged

# Exclude specific patterns
chezmoi-tui unmanaged --exclude "*.tmp"

# Include only specific file types
chezmoi-tui unmanaged --include "*.conf"
```

## Environment Variables

Chezmoi TUI respects these environment variables:

```bash
# Configuration file path
export CHEZMOI_TUI_CONFIG=~/.config/chezmoi-tui/custom-config.yaml

# Enable verbose output
export CHEZMOI_TUI_VERBOSE=true

# Enable debug mode
export CHEZMOI_TUI_DEBUG=true

# Set custom theme
export CHEZMOI_TUI_THEME=dark
```

## Exit Codes

Chezmoi TUI uses these exit codes:

- `0`: Success
- `1`: General error
- `2`: Misuse of shell builtins
- `126`: Command invoked cannot execute
- `127`: Command not found
- `128+n`: Fatal error signal "n"

## Command Aliases

Some commands have shorter aliases:

```bash
# Short aliases
chezmoi-tui st     # alias for status
chezmoi-tui ap     # alias for apply
chezmoi-tui ad     # alias for add
chezmoi-tui df     # alias for diff
chezmoi-tui cfg    # alias for config
chezmoi-tui stat   # alias for stats
```

## Usage Examples

### Basic Workflow

```bash
# Initialize with your dotfiles repository
chezmoi-tui init https://github.com/username/dotfiles.git

# Check the status of your files
chezmoi-tui status

# Add a new configuration file
chezmoi-tui add ~/.tmux.conf

# Apply changes
chezmoi-tui apply

# View statistics
chezmoi-tui stats
```

### Secret Management

```bash
# Unlock Bitwarden vault
chezmoi-tui bitwarden unlock

# Generate template from Bitwarden item
chezmoi-tui bitwarden template "api-keys"

# Apply templates
chezmoi-tui apply

# Export secrets for development
chezmoi-tui bitwarden export .env.development

# Lock vault when done
chezmoi-tui bitwarden lock
```

### TUI Usage

```bash
# Launch the interactive interface
chezmoi-tui tui

# Navigate with arrow keys
# Press Enter to select options
# Press 'q' or Ctrl+C to quit
```

## Shell Integration

### Auto-completion

Enable shell auto-completion:

```bash
# For bash (add to ~/.bashrc)
source <(chezmoi-tui completion bash)

# For zsh (add to ~/.zshrc)
source <(chezmoi-tui completion zsh)

# For fish (add to ~/.config/fish/completions/chezmoi-tui.fish)
chezmoi-tui completion fish > ~/.config/fish/completions/chezmoi-tui.fish
```

### Aliases

Create convenient aliases:

```bash
# Add to your shell configuration
alias ct='chezmoi-tui'
alias cts='chezmoi-tui status'
alias cta='chezmoi-tui apply'
alias cti='chezmoi-tui init'
```

## Troubleshooting

### Common Issues

1. **Command not found**: Ensure binary is in your PATH
2. **Permission denied**: Check file permissions
3. **Configuration errors**: Validate with `chezmoi-tui config validate`
4. **Network issues**: Check internet connectivity

### Getting Help

```bash
# Show detailed help
chezmoi-tui --help

# Show help for specific command
chezmoi-tui <command> --help

# Enable debug output
chezmoi-tui --debug <command>
```

## See Also

- [Installation Guide](installation.md)
- [TUI User Guide](tui-user-guide.md)
- [Bitwarden Integration](bitwarden-integration.md)
- [Configuration Management](configuration.md)