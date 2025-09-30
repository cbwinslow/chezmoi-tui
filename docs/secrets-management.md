# Secrets Management with Chezmoi TUI

This document explains how to securely manage secrets using Chezmoi TUI and GitHub Actions.

## Overview

Proper secrets management is critical for secure software development. This project demonstrates best practices for:

1. Storing secrets securely using Chezmoi's encryption capabilities
2. Managing secrets across development, staging, and production environments
3. Using GitHub Actions with encrypted secrets
4. Safely handling sensitive configuration data

## Chezmoi Secrets Best Practices

### 1. Encrypting Sensitive Files

Use Chezmoi's built-in encryption for sensitive files:

```bash
# Add a file that should be encrypted
chezmoi add --encrypt path/to/secret/file

# Or add a template that includes sensitive data
chezmoi add --template --encrypt path/to/secret.tmpl
```

### 2. Template-Based Configuration

Store configuration in templates with variable substitution:

```yaml
# Example: .chezmoi.yaml.tmpl
data:
  github:
    token: {{ .github.token | quote }}
  aws:
    accessKeyId: {{ .aws.accessKeyId | quote }}
    secretAccessKey: {{ .aws.secretAccessKey | quote }}
```

### 3. Environment-Specific Secrets

Set different secrets for different environments:

```bash
# On your development machine, set development secrets
chezmoi data --set github.token=dev_token --set aws.accessKeyId=dev_key

# In CI/CD, use different secrets
chezmoi data --set github.token=${{ secrets.GITHUB_TOKEN }}
```

## GitHub Actions Integration

### 1. Using GitHub Secrets

Store secrets in GitHub repository settings:

1. Go to your repository on GitHub
2. Navigate to Settings > Secrets and variables > Actions
3. Add your secrets (e.g., GITHUB_TOKEN, AWS_ACCESS_KEY_ID, etc.)

### 2. Environment-Specific Variables

Use GitHub Environments for different deployment stages:

- Development: Minimal secrets required
- Staging: Staging-specific credentials
- Production: Production credentials with manual approval

### 3. Encrypted Checkouts

The workflow includes secure handling of encrypted files:

```yaml
- name: Decrypt secrets
  run: |
    # Use chezmoi to decrypt secrets during the workflow
    chezmoi init --apply  # Initialize with encrypted configuration
    source ~/.secrets.env  # Source decrypted secrets if needed
```

## Security Guidelines

### DO:
- Encrypt sensitive files with Chezmoi
- Use different secrets for each environment
- Implement approval gates for production deployments
- Rotate secrets regularly
- Use least-privilege access principles

### DON'T:
- Commit secrets in plain text to repositories
- Use the same secrets across all environments
- Store secrets in unencrypted configuration files
- Hardcode secrets in source code
- Share production secrets with all team members

## Example Setup Process

1. Initialize your chezmoi configuration:
   ```bash
   chezmoi init --ssh git@github.com:username/dotfiles.git
   ```

2. Encrypt sensitive files:
   ```bash
   chezmoi add --encrypt .config/gh-token
   chezmoi add --encrypt .aws/credentials
   ```

3. Create templates for environment configuration:
   ```bash
   chezmoi add --template --encrypt .envrc.tmpl
   ```

4. Apply configuration (this will prompt for decryption key):
   ```bash
   chezmoi apply
   ```

## CI/CD Pipeline Security

The included workflow demonstrates secure CI/CD practices:

- Secrets are only accessed in appropriate environments
- Production deployments require manual approval
- Multiple quality checks before deployment
- Automatic security scanning
- Proper access controls and environment isolation

## Troubleshooting

### Problem: "Secrets not available in CI/CD"
**Solution**: Ensure your encrypted files are properly added to chezmoi and that the CI/CD environment has the necessary decryption keys.

### Problem: "Permission denied accessing secrets"
**Solution**: Check that GitHub Actions permissions are properly configured and that secrets are defined in repository settings.

### Problem: "Build failing due to missing secrets"
**Solution**: Verify that non-sensitive defaults are provided and that the workflow handles missing secrets gracefully.