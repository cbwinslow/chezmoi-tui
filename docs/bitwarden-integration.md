# Bitwarden Integration Guide

This guide explains how to integrate Chezmoi TUI with Bitwarden for secure secrets management.

## Overview

Chezmoi TUI provides comprehensive integration with Bitwarden, allowing you to manage your secrets alongside your dotfiles. This integration offers:

- Seamless CLI commands for Bitwarden operations
- TUI interface for managing secrets
- Chezmoi template generation from Bitwarden items
- Secure environment file exports
- Integration with existing Bitwarden TUI installations

## Prerequisites

Before using Bitwarden integration, ensure you have:

1. **Bitwarden CLI (`bw`)** installed and in your PATH
2. **Bitwarden account** with secrets to manage
3. **Chezmoi TUI** properly installed
4. **Bitwarden TUI** (optional, for advanced features)

## Authentication

### Unlocking Your Vault

Before performing Bitwarden operations, you need to unlock your vault:

```bash
# Unlock the vault
chezmoi-tui bitwarden unlock

# Check vault status
chezmoi-tui bitwarden status
```

### Session Management

Chezmoi TUI automatically manages your Bitwarden session:

```bash
# Lock the vault (logout)
chezmoi-tui bitwarden lock

# Sync with remote server
chezmoi-tui bitwarden sync
```

## CLI Commands

### Basic Operations

```bash
# View vault status
chezmoi-tui bitwarden status

# Unlock vault
chezmoi-tui bitwarden unlock

# Lock vault
chezmoi-tui bitwarden lock

# Sync with server
chezmoi-tui bitwarden sync

# List items
chezmoi-tui bitwarden list [filter]

# Launch Bitwarden TUI
chezmoi-tui bitwarden tui
```

### Advanced Operations

```bash
# Generate Chezmoi template from Bitwarden item
chezmoi-tui bitwarden template <item-id>

# Export secrets to environment file
chezmoi-tui bitwarden export [.env-file]
```

## Chezmoi Template Generation

### Creating Templates

Generate Chezmoi templates directly from Bitwarden items:

```bash
# Generate template for a specific item
chezmoi-tui bitwarden template "api-keys"

# The template will be created at ~/.local/share/chezmoi/dot_secrets.tmpl
```

### Template Structure

Generated templates follow this structure:

```gotemplate
# Bitwarden Secrets Template
# Generated from item ID: api-keys
# This file is auto-generated - do not edit manually

# Example template using Bitwarden integration
{{- if (bitwarden "api-keys") }}
export EXAMPLE_SECRET="{{ (bitwarden "api-keys").password }}"
{{- end }}

# You can add more secrets here as needed
```

### Applying Templates

After generating templates, apply them with Chezmoi:

```bash
# Apply the generated templates
chezmoi apply
```

## Environment File Export

Export Bitwarden secrets to environment files for development:

```bash
# Export to default .env file
chezmoi-tui bitwarden export

# Export to custom file
chezmoi-tui bitwarden export development.env

# Export with specific filtering
chezmoi-tui bitwarden export --filter "development"
```

## TUI Integration

### Accessing Bitwarden Manager

Navigate to the Bitwarden Manager in the Chezmoi TUI:

1. Launch the TUI: `chezmoi-tui tui`
2. Navigate to "Bitwarden Manager" using arrow keys
3. Press Enter to access the manager

### Bitwarden Manager Features

The Bitwarden Manager provides:

- **Vault Status**: Current lock/unlock status
- **Item Management**: View and manage Bitwarden items
- **Template Generation**: Create Chezmoi templates
- **Export Functionality**: Export secrets to environment files
- **TUI Launch**: Direct access to Bitwarden TUI

## Security Considerations

### Best Practices

1. **Never commit secrets** to version control
2. **Use strong master passwords** for your Bitwarden vault
3. **Enable two-factor authentication** on your Bitwarden account
4. **Regularly rotate secrets** and update templates
5. **Review access logs** for unauthorized access

### Secure File Permissions

Ensure exported files have secure permissions:

```bash
# Set secure permissions on exported files
chmod 600 .env
```

### Template Security

When creating templates:

1. **Use Chezmoi's built-in templating** for sensitive values
2. **Avoid hardcoding secrets** in templates
3. **Review generated templates** before committing
4. **Use appropriate file permissions** for template files

## Integration with Existing Setups

### Dotfiles Integration

Chezmoi TUI automatically detects existing Bitwarden TUI installations in your dotfiles:

```bash
# Check for existing installation
ls ~/.local/share/chezmoi/bw-secrets-tui/
```

### Migration from Manual Setup

If you're migrating from a manual Bitwarden setup:

1. **Backup existing configurations**
2. **Install Chezmoi TUI**
3. **Import existing templates**
4. **Update workflow documentation**

## Advanced Configuration

### Custom Template Paths

Configure custom paths for generated templates:

```yaml
# In ~/.config/chezmoi-tui/config.yaml
bitwarden:
  template_path: "~/.local/share/chezmoi/private_dot_secrets.tmpl"
  export_path: "~/.env"
```

### Automation Scripts

Create automation scripts for regular operations:

```bash
#!/bin/bash
# daily-sync.sh - Daily Bitwarden sync script

# Unlock vault
chezmoi-tui bitwarden unlock

# Sync with server
chezmoi-tui bitwarden sync

# Generate/update templates
chezmoi-tui bitwarden template "api-keys"
chezmoi-tui bitwarden template "database-credentials"

# Apply changes
chezmoi apply

# Lock vault
chezmoi-tui bitwarden lock
```

## Troubleshooting

### Common Issues

1. **Vault locked**: Ensure vault is unlocked before operations
2. **Item not found**: Verify item ID/name exists in vault
3. **Permission denied**: Check file permissions and Bitwarden session
4. **Network issues**: Verify internet connectivity and Bitwarden server status

### Debugging

Enable verbose output for debugging:

```bash
# Enable debug mode
export DEBUG=true
chezmoi-tui bitwarden status
```

### Getting Help

For additional help:

```bash
# View command help
chezmoi-tui bitwarden --help
chezmoi-tui bitwarden <command> --help

# Check documentation
man chezmoi-tui-bitwarden
```

## Next Steps

After setting up Bitwarden integration:

1. [Explore Advanced Features](advanced-features.md)
2. [Learn About Security Best Practices](security.md)
3. [Set up Automated Workflows](automation.md)
4. [Review Configuration Options](configuration.md)

## Community Resources

Join the community for additional support:

- [GitHub Issues](https://github.com/cbwinslow/chezmoi-tui/issues)
- [Discussion Forum](https://github.com/cbwinslow/chezmoi-tui/discussions)
- [Community Slack](#) (if available)