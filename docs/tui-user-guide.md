# TUI User Guide

This guide explains how to use the Terminal User Interface (TUI) of Chezmoi TUI for managing your dotfiles.

## Launching the TUI

To launch the TUI, use the `tui` command:

```bash
# Launch the TUI
chezmoi-tui tui

# Launch with specific dimensions
chezmoi-tui tui --width 120 --height 40

# Launch in fullscreen mode
chezmoi-tui tui --fullscreen
```

## Main Menu Navigation

Upon launching, you'll see the main menu with several options:

```
┌─ Chezmoi TUI - Enhanced dotfile management ─────────────────────┐
│                                                                  │
│  → View Status       Show status of all managed files           │
│    Add Files         Add new files to chezmoi management        │
│    Apply Changes     Apply managed files to your system         │
│    Diff Changes      Show differences between source/destination│
│    Show Stats        Show statistics about your dotfiles         │
│    Bitwarden Manager Manage Bitwarden secrets and integration  │
│    Exit              Quit the application                        │
│                                                                  │
│  Use arrow keys to navigate, Enter to select, q to quit          │
└──────────────────────────────────────────────────────────────────┘
```

### Navigation Controls

- **↑/↓ Arrow Keys**: Move selection up/down
- **Enter**: Select the highlighted option
- **h/j/k/l**: Vim-style navigation (h:left, j:down, k:up, l:right)
- **q/Ctrl+C**: Quit the application

## View Status Screen

The "View Status" screen shows the current status of all managed files:

```
┌─ Chezmoi File Status ────────────────────────────────────────────┐
│                                                                  │
│  → [M] .bashrc                                                   │
│    [A] .gitconfig                                                │
│    [ ] .vimrc                                                    │
│    [D] .old-config                                              │
│                                                                  │
│  4 files total | Use arrow keys to navigate, 'h' to go back      │
└──────────────────────────────────────────────────────────────────┘
```

### Status Symbols

- **M**: Modified file (changes exist in destination)
- **A**: Added file (newly added to management)
- **D**: Deleted file (removed from destination)
- **Space**: Up-to-date file (no changes)

### Navigation in File View

- **↑/↓ Arrow Keys**: Navigate between files
- **h/← Left Arrow**: Return to main menu
- **q/Ctrl+C**: Quit the application

## Add Files Workflow

To add new files to management:

1. Select "Add Files" from main menu
2. Enter file paths when prompted
3. Choose options (encryption, templating, etc.)
4. Files will be added to your chezmoi repository

### Supported Options

- **File Encryption**: Encrypt sensitive files
- **Templating**: Add as template for dynamic content
- **Auto-templating**: Automatically detect and template variables
- **Recursive**: Add directories recursively

## Apply Changes Workflow

To apply managed files to your system:

1. Select "Apply Changes" from main menu
2. Review changes that will be applied
3. Confirm to proceed with application
4. Changes will be applied to your destination directory

### Apply Options

- **Verbose**: Show detailed output during application
- **Dry-run**: Preview changes without applying them
- **Force**: Apply without confirmation prompts
- **Selective**: Apply only specific files

## Diff Changes Workflow

To view differences between source and destination:

1. Select "Diff Changes" from main menu
2. Choose specific files or view all differences
3. Differences will be displayed in unified diff format

### Diff Options

- **Unified Format**: Show in standard unified diff format
- **Side-by-side**: Show differences side-by-side
- **Color Output**: Use colors to highlight differences
- **Specific Files**: Show differences for specific files only

## Show Stats Screen

The "Show Stats" screen provides detailed statistics about your dotfiles:

```
┌─ Chezmoi Dotfiles Statistics ───────────────────────────────────┐
│ Last Updated: 2025-09-30 23:26:01                             │
├─────────────────────────────────────────────────────────────────┤
│ Total Managed Files:    861                                   │
│ Total Unmanaged Files:  393                                   │
├─────────────────────────────────────────────────────────────────┤
│ Up to Date:               2 (  0%)                           │
│ Modified:                41 (  4%)                           │
│ Added:                  745 ( 86%)                           │
│ Deleted:                  0 (  0%)                           │
├─────────────────────────────────────────────────────────────────┤
│ Actions: Use arrow keys to navigate, 'h' to go back, 'q' to quit │
└─────────────────────────────────────────────────────────────────┘
```

### Statistical Information

- **Managed Files**: Total number of files under management
- **Unmanaged Files**: Files in destination not under management
- **Status Distribution**: Breakdown of file statuses
- **Percentages**: Percentage distribution of statuses

## Bitwarden Manager

The "Bitwarden Manager" provides integration with Bitwarden secrets management:

```
┌─ Bitwarden Secrets Manager ─────────────────────────────────────┐
│                                                                 │
│  Manage your Bitwarden vault from within Chezmoi TUI           │
│                                                                 │
│  Features:                                                      │
│  • View and manage Bitwarden items                              │
│  • Unlock/Lock your vault                                       │
│  • Sync with remote server                                      │
│  • Integration with Chezmoi templates                            │
│  • Launch dedicated Bitwarden TUI                               │
│                                                                 │
│  CLI Commands:                                                   │
│  • chezmoi-tui bitwarden status  - Show vault status            │
│  • chezmoi-tui bitwarden unlock  - Unlock vault                 │
│  • chezmoi-tui bitwarden lock    - Lock vault                   │
│  • chezmoi-tui bitwarden list    - List items                  │
│  • chezmoi-tui bitwarden sync    - Sync with server             │
│  • chezmoi-tui bitwarden tui     - Launch Bitwarden TUI         │
│                                                                 │
│  Actions: Use arrow keys to navigate, 'h' to go back, 'q' to quit │
└─────────────────────────────────────────────────────────────────┘
```

### Bitwarden Manager Features

1. **Vault Status**: Check current lock/unlock status
2. **Authentication**: Unlock/lock your vault
3. **Item Management**: View and manage Bitwarden items
4. **Template Generation**: Create Chezmoi templates from Bitwarden items
5. **Secret Export**: Export secrets to environment files
6. **TUI Launch**: Direct access to dedicated Bitwarden TUI

## Customization

### Theme Configuration

Customize the appearance of the TUI through configuration:

```yaml
# ~/.config/chezmoi-tui/config.yaml
theme:
  primary_color: "#1793d1"
  secondary_color: "#0366d6"
  background_color: "#000000"
  text_color: "#ffffff"
  border_style: "rounded"
```

### Behavior Configuration

Adjust TUI behavior to suit your preferences:

```yaml
# ~/.config/chezmoi-tui/config.yaml
tui:
  show_help: true
  refresh_interval: 5
  confirm_actions: true
  auto_refresh: true
```

## Keyboard Shortcuts

### Global Shortcuts

- **q/Ctrl+C**: Quit the application
- **h/← Left Arrow**: Go back/return to previous screen
- **l/→ Right Arrow**: Enter/select current item
- **?**: Show help information

### Navigation Shortcuts

- **↑/k**: Move cursor up
- **↓/j**: Move cursor down
- **←/h**: Go back
- **→/l**: Enter/select

### Action Shortcuts

- **Enter**: Confirm selection
- **Esc**: Cancel/close dialog
- **Space**: Toggle selection (checkboxes)

## Help System

### Accessing Help

The TUI includes a comprehensive help system:

1. **Contextual Help**: Press `?` for context-specific help
2. **Main Menu Help**: Available on all screens
3. **Command Help**: Detailed help for each command

### Help Content

- **Navigation Guide**: Instructions for moving through the interface
- **Command Descriptions**: Explanation of each available command
- **Keyboard Shortcuts**: List of all available shortcuts
- **Troubleshooting Tips**: Common solutions to issues

## Error Handling

### Error Messages

When errors occur, they are displayed clearly:

```
┌─ Error ─────────────────────────────────────────────────────────┐
│                                                                 │
│  Failed to apply changes to ~/.bashrc: Permission denied       │
│                                                                 │
│  [OK]                                                           │
└─────────────────────────────────────────────────────────────────┘
```

### Error Resolution

Common error resolutions:

1. **Permission Errors**: Check file permissions and run with appropriate privileges
2. **Network Errors**: Verify internet connectivity and Bitwarden server status
3. **File Conflicts**: Resolve conflicts manually or use merge tools
4. **Vault Locked**: Unlock Bitwarden vault before accessing secrets

## Performance Considerations

### Large Repositories

For repositories with many files:

- **Pagination**: Files are paginated for better performance
- **Filtering**: Use filters to narrow down file lists
- **Caching**: Results are cached to improve responsiveness

### Slow Operations

Slow operations show progress indicators:

```
┌─ Applying Changes ──────────────────────────────────────────────┐
│                                                                 │
│  [•••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••] │
│  Processing ~/.config/application/settings.json (42/127)        │
│                                                                 │
│  Press Ctrl+C to cancel                                         │
└─────────────────────────────────────────────────────────────────┘
```

## Integration with External Tools

### Editor Integration

Launch your preferred editor from within the TUI:

```yaml
# ~/.config/chezmoi-tui/config.yaml
editor:
  command: "vim"
  args: []
```

### Git Integration

Perform Git operations without leaving the TUI:

- **Commit**: Commit changes with descriptive messages
- **Push**: Push changes to remote repository
- **Pull**: Pull latest changes from remote
- **Status**: View Git repository status

### Shell Integration

Execute shell commands directly from the TUI:

- **Quick Commands**: Run predefined shell commands
- **Custom Scripts**: Execute custom automation scripts
- **Environment Management**: Manage different environments

## Advanced Features

### Workspaces

Manage multiple dotfile configurations:

```bash
# Switch between workspaces
chezmoi-tui tui --workspace personal
chezmoi-tui tui --workspace work
```

### Profiles

Use different profiles for various contexts:

```yaml
# ~/.config/chezmoi-tui/config.yaml
profiles:
  personal:
    theme: "dark"
    include_unmanaged: false
  work:
    theme: "light"
    include_unmanaged: true
```

### Plugins

Extend functionality with plugins:

```yaml
# ~/.config/chezmoi-tui/config.yaml
plugins:
  - name: "backup-manager"
    path: "~/.local/share/chezmoi-tui/plugins/backup-manager.so"
  - name: "sync-service"
    path: "~/.local/share/chezmoi-tui/plugins/sync-service.so"
```

## Troubleshooting

### Common Issues

1. **Display Problems**: Adjust terminal size or font settings
2. **Navigation Issues**: Use standard arrow keys or hjkl navigation
3. **Performance Problems**: Enable pagination or filtering for large repositories
4. **Integration Failures**: Verify external tool installations and PATH settings

### Getting Help

For additional assistance:

1. **Built-in Help**: Use `?` for context-sensitive help
2. **Documentation**: Refer to online documentation
3. **Community**: Join the user community for peer support
4. **Issues**: Report bugs or request features on GitHub

## See Also

- [CLI Commands Reference](cli-commands.md)
- [Bitwarden Integration](bitwarden-integration.md)
- [Configuration Management](configuration.md)
- [Advanced Features](advanced-features.md)