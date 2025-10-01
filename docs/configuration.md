# Configuration Guide

This guide explains how to configure Chezmoi TUI to suit your preferences and workflow.

## Configuration File Location

Chezmoi TUI looks for configuration files in the following locations:

### Linux/macOS
```
~/.config/chezmoi-tui/config.yaml
~/.config/chezmoi-tui/config.json
~/.chezmoi-tui.yaml
~/.chezmoi-tui.json
```

### Windows
```
%APPDATA%\chezmoi-tui\config.yaml
%APPDATA%\chezmoi-tui\config.json
%USERPROFILE%\.chezmoi-tui.yaml
%USERPROFILE%\.chezmoi-tui.json
```

## Configuration File Formats

Chezmoi TUI supports both YAML and JSON configuration formats.

### YAML Format (Recommended)

```yaml
# ~/.config/chezmoi-tui/config.yaml
theme:
  primary_color: "#1793d1"
  secondary_color: "#0366d6"
  background_color: "#000000"
  text_color: "#ffffff"
  border_style: "rounded"

tui:
  show_help: true
  refresh_interval: 5
  confirm_actions: true
  auto_refresh: true

cli:
  verbose: false
  color: true
  pager: "less"

integration:
  chezmoi_binary_path: ""
  bitwarden_binary_path: ""
  timeout: 30

bitwarden:
  template_path: "~/.local/share/chezmoi/dot_secrets.tmpl"
  export_path: "~/.env"
```

### JSON Format

```json
{
  "theme": {
    "primary_color": "#1793d1",
    "secondary_color": "#0366d6",
    "background_color": "#000000",
    "text_color": "#ffffff",
    "border_style": "rounded"
  },
  "tui": {
    "show_help": true,
    "refresh_interval": 5,
    "confirm_actions": true,
    "auto_refresh": true
  },
  "cli": {
    "verbose": false,
    "color": true,
    "pager": "less"
  },
  "integration": {
    "chezmoi_binary_path": "",
    "bitwarden_binary_path": "",
    "timeout": 30
  },
  "bitwarden": {
    "template_path": "~/.local/share/chezmoi/dot_secrets.tmpl",
    "export_path": "~/.env"
  }
}
```

## Generating Configuration

### Create Default Configuration

Generate a default configuration file:

```bash
# Generate default YAML configuration
chezmoi-tui config generate

# Generate JSON configuration
chezmoi-tui config generate --format json

# Generate to custom location
chezmoi-tui config generate --output ~/.config/chezmoi-tui/my-config.yaml
```

### Show Current Configuration

View your current configuration:

```bash
# Show current configuration
chezmoi-tui config show

# Show in specific format
chezmoi-tui config show --format json

# Show configuration file path
chezmoi-tui config show --path
```

### Validate Configuration

Check if your configuration is valid:

```bash
# Validate configuration
chezmoi-tui config validate

# Validate specific file
chezmoi-tui config validate --file ~/.config/chezmoi-tui/custom-config.yaml
```

## Theme Configuration

### Color Customization

Customize the color scheme of the TUI:

```yaml
theme:
  # Primary accent color
  primary_color: "#1793d1"
  
  # Secondary accent color
  secondary_color: "#0366d6"
  
  # Background color
  background_color: "#000000"
  
  # Text color
  text_color: "#ffffff"
  
  # Success color (for positive feedback)
  success_color: "#28a745"
  
  # Warning color (for cautionary messages)
  warning_color: "#ffc107"
  
  # Danger/error color (for errors)
  danger_color: "#dc3545"
```

### Border Styles

Choose from different border styles:

```yaml
theme:
  border_style: "rounded"  # Options: rounded, square, double, thick, thin
```

Border style options:
- **rounded**: Rounded corners (default)
- **square**: Square corners
- **double**: Double line borders
- **thick**: Thick line borders
- **thin**: Thin line borders

### Typography

Customize text appearance:

```yaml
theme:
  # Font family (if supported by terminal)
  font_family: "monospace"
  
  # Base font size
  font_size: 12
  
  # Bold text enabled
  bold_text: true
```

## TUI Configuration

### Navigation and Behavior

```yaml
tui:
  # Show help information in TUI
  show_help: true
  
  # Auto-refresh interval in seconds (0 to disable)
  refresh_interval: 5
  
  # Confirm destructive actions
  confirm_actions: true
  
  # Auto-refresh when files change
  auto_refresh: true
  
  # Show line numbers in file views
  show_line_numbers: true
  
  # Wrap long lines in file views
  wrap_lines: false
  
  # Highlight syntax in file views
  syntax_highlighting: true
```

### Performance Settings

```yaml
tui:
  # Maximum number of files to display at once
  max_file_display: 100
  
  # Cache timeout in minutes
  cache_timeout: 30
  
  # Animation speed (0 = disabled, 1-10 = slow-fast)
  animation_speed: 3
```

## CLI Configuration

### Output and Display

```yaml
cli:
  # Enable verbose output by default
  verbose: false
  
  # Enable colored output
  color: true
  
  # Pager for long output
  pager: "less"
  
  # Maximum width for output
  max_width: 120
  
  # Show timestamps in output
  timestamps: true
  
  # Show emojis in output
  emojis: true
```

### Command Behavior

```yaml
cli:
  # Automatically confirm yes/no prompts
  auto_confirm: false
  
  # Continue on error instead of stopping
  continue_on_error: false
  
  # Number of retries for failed operations
  retry_attempts: 3
  
  # Delay between retries in seconds
  retry_delay: 1
```

## Integration Configuration

### Binary Paths

Specify custom paths for external binaries:

```yaml
integration:
  # Custom path to chezmoi binary
  chezmoi_binary_path: "/usr/local/bin/chezmoi"
  
  # Custom path to Bitwarden binary
  bitwarden_binary_path: "/usr/local/bin/bw"
  
  # Custom path to Git binary
  git_binary_path: "/usr/local/bin/git"
  
  # Timeout for external commands in seconds
  timeout: 30
```

### Advanced Integration Settings

```yaml
integration:
  # Enable/disable specific integrations
  enable_chezmoi: true
  enable_bitwarden: true
  enable_git: true
  
  # Parallel processing for faster operations
  parallel_processing: true
  
  # Number of parallel workers
  max_workers: 4
  
  # Buffer size for streaming output
  buffer_size: 4096
```

## Bitwarden Configuration

### Template and Export Settings

```yaml
bitwarden:
  # Default path for generated templates
  template_path: "~/.local/share/chezmoi/dot_secrets.tmpl"
  
  # Default path for exported environment files
  export_path: "~/.env"
  
  # Enable automatic template generation
  auto_generate_templates: true
  
  # Enable secure file permissions for exports
  secure_export_permissions: true
  
  # Default export format (env, json, yaml)
  export_format: "env"
```

### Vault Settings

```yaml
bitwarden:
  # Default sync interval in minutes
  sync_interval: 60
  
  # Auto-lock vault after inactivity (minutes)
  auto_lock_timeout: 30
  
  # Enable two-factor authentication reminder
  tfa_reminder: true
```

## Environment Variables

Configuration can also be overridden using environment variables:

```bash
# Set configuration file path
export CHEZMOI_TUI_CONFIG=~/.config/chezmoi-tui/custom-config.yaml

# Enable verbose output
export CHEZMOI_TUI_VERBOSE=true

# Enable debug mode
export CHEZMOI_TUI_DEBUG=true

# Set custom theme
export CHEZMOI_TUI_THEME=dark

# Override specific settings
export CHEZMOI_TUI_CLI_COLOR=false
export CHEZMOI_TUI_TUI_REFRESH_INTERVAL=10
```

Environment variables take precedence over configuration file settings.

## Advanced Configuration Examples

### Multi-Environment Setup

```yaml
# ~/.config/chezmoi-tui/config.yaml
environments:
  personal:
    theme:
      primary_color: "#1793d1"
    tui:
      show_help: true
    
  work:
    theme:
      primary_color: "#0366d6"
    tui:
      show_help: false
      confirm_actions: true

# Active environment
active_environment: "personal"
```

Switch environments:
```bash
# Set active environment
export CHEZMOI_TUI_ENVIRONMENT=work
```

### Conditional Configuration

```yaml
# Conditional settings based on OS
conditional:
  - when: "os == 'linux'"
    settings:
      integration:
        chezmoi_binary_path: "/usr/bin/chezmoi"
  
  - when: "os == 'darwin'"
    settings:
      integration:
        chezmoi_binary_path: "/usr/local/bin/chezmoi"
  
  - when: "os == 'windows'"
    settings:
      integration:
        chezmoi_binary_path: "C:\\Program Files\\chezmoi\\chezmoi.exe"
```

### Plugin Configuration

```yaml
plugins:
  - name: "backup-manager"
    enabled: true
    path: "~/.local/share/chezmoi-tui/plugins/backup-manager.so"
    config:
      backup_location: "~/backups"
      retention_days: 30
  
  - name: "sync-service"
    enabled: false
    path: "~/.local/share/chezmoi-tui/plugins/sync-service.so"
    config:
      sync_targets:
        - "dropbox"
        - "google-drive"
```

## Validation and Troubleshooting

### Configuration Validation

Validate your configuration:

```bash
# Validate configuration
chezmoi-tui config validate

# Validate with detailed output
chezmoi-tui config validate --verbose
```

### Common Issues

1. **Invalid YAML/JSON Syntax**: Use online validators or `yamllint`
2. **Unknown Configuration Keys**: Check spelling and consult documentation
3. **Permission Issues**: Ensure config file has appropriate permissions
4. **Path Issues**: Use absolute paths or proper tilde expansion

### Reset to Defaults

Reset configuration to defaults:

```bash
# Backup current configuration
cp ~/.config/chezmoi-tui/config.yaml ~/.config/chezmoi-tui/config.backup.yaml

# Generate fresh default configuration
chezmoi-tui config generate --force
```

## Best Practices

### Organization

1. **Keep configuration simple**: Start with minimal config and add as needed
2. **Use comments**: Document complex settings
3. **Version control**: Track configuration changes in Git
4. **Environment-specific**: Use different configs for different environments

### Security

1. **Secure permissions**: Set restrictive permissions on config files
2. **Avoid hardcoding secrets**: Use environment variables or external secret stores
3. **Regular reviews**: Periodically review configuration for outdated settings

### Performance

1. **Disable unused features**: Turn off integrations you don't use
2. **Optimize refresh intervals**: Balance between freshness and performance
3. **Use caching wisely**: Enable caching for frequently accessed data

## See Also

- [Quick Start Guide](quick-start.md)
- [CLI Commands Reference](cli-commands.md)
- [TUI User Guide](tui-user-guide.md)
- [Bitwarden Integration](bitwarden-integration.md)