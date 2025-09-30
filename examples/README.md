# Examples

## Basic Usage

### Check status of your dotfiles
```bash
chezmoi-tui status
```

### Add a new file to management
```bash 
chezmoi-tui add ~/.gitconfig
```

### Apply changes to your system
```bash
chezmoi-tui apply
```

### Launch the TUI
```bash
chezmoi-tui tui
```

## Advanced Usage

### Apply only specific files
```bash
chezmoi-tui apply ~/.bashrc ~/.gitconfig
```

### View differences for specific files
```bash
chezmoi-tui diff ~/.vimrc
```

### Add multiple files at once
```bash
chezmoi-tui add ~/.bashrc ~/.gitconfig ~/.vimrc
```