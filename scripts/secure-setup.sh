#!/bin/bash
# secure-setup.sh - Script to demonstrate proper Chezmoi secrets setup

set -e

echo "Setting up secure Chezmoi secrets management..."

# Create example structure for encrypted secrets
mkdir -p .chezmoiexternal.yaml.d

# Create a template for environment-specific configuration
cat > .chezmoitemplates/secrets.env.tmpl << 'EOF'
# Environment-specific secrets
# This file will be generated from chezmoi template

# GitHub
{{- if (and (exists ".github.token") (not (eq .github.token ""))) }}
export GITHUB_TOKEN={{ .github.token }}
{{- else }}
export GITHUB_TOKEN=""
{{- end }}

# Docker Registry
{{- if (and (exists ".docker.username") (not (eq .docker.username ""))) }}
export DOCKER_USERNAME={{ .docker.username }}
{{- else }}
export DOCKER_USERNAME=""
{{- end }}
{{- if (and (exists ".docker.password") (not (eq .docker.password ""))) }}
export DOCKER_PASSWORD={{ .docker.password }}
{{- else }}
export DOCKER_PASSWORD=""
{{- end }}

# AWS
{{- if (and (exists ".aws.accessKeyId") (not (eq .aws.accessKeyId ""))) }}
export AWS_ACCESS_KEY_ID={{ .aws.accessKeyId }}
{{- else }}
export AWS_ACCESS_KEY_ID=""
{{- end }}
{{- if (and (exists ".aws.secretAccessKey") (not (eq .aws.secretAccessKey ""))) }}
export AWS_SECRET_ACCESS_KEY={{ .aws.secretAccessKey }}
{{- else }}
export AWS_SECRET_ACCESS_KEY=""
{{- end }}
EOF

echo "Created secrets template for use with Chezmoi"

# Show how to initialize Chezmoi with secret data
cat > docs/ci-setup.md << 'EOF'
# CI/CD Setup with Chezmoi Secrets

## Local Development Setup

To set up Chezmoi with your secrets locally:

1. Install chezmoi:
   ```bash
   brew install chezmoi  # On macOS
   # or
   curl -fsSL https://get.chezmoi.io | sh  # On Linux
   ```

2. Initialize your chezmoi repository:
   ```bash
   chezmoi init --ssh git@github.com:username/dotfiles.git
   ```

3. Set your secret data:
   ```bash
   # Set GitHub token
   chezmoi data --set github.token=your_github_token_here
   
   # Set AWS credentials
   chezmoi data --set aws.accessKeyId=your_access_key_id
   chezmoi data --set aws.secretAccessKey=your_secret_access_key
   
   # Set Docker credentials
   chezmoi data --set docker.username=your_docker_username
   chezmoi data --set docker.password=your_docker_password
   ```

4. Add sensitive files with encryption:
   ```bash
   # Add encrypted credentials file
   chezmoi add --encrypt ~/.aws/credentials
   
   # Add encrypted SSH private key
   chezmoi add --encrypt ~/.ssh/id_rsa
   ```

## GitHub Actions Configuration

To use these secrets in GitHub Actions:

1. Store sensitive values as GitHub repository secrets
2. Use the following pattern in your workflow:

```yaml
- name: Setup chezmoi with secrets
  run: |
    # Use chezmoi to generate configuration with secrets
    # Secrets are provided via environment variables
    mkdir -p ~/.config/chezmoi
    echo "{}" > ~/.config/chezmoi/chezmoi.json
    chezmoi init --apply
```

## Security Best Practices

- Never commit plaintext secrets to the repository
- Use chezmoi's encryption for sensitive files
- Set different secrets for different environments
- Use GitHub Environments for production deployments
- Implement approval gates for sensitive operations
EOF

echo "Created CI/CD setup documentation"

# Make the script executable
chmod +x scripts/setup-secrets.sh

echo "Setup complete! The scripts and documentation for secure Chezmoi secrets management have been created."
echo ""
echo "Next steps:"
echo "1. Review the documentation in docs/secrets-management.md"
echo "2. Check the CI/CD workflow in .github/workflows/advanced-cicd.yml"
echo "3. Examine the secrets setup script in scripts/setup-secrets.sh"
echo "4. Follow the setup guide in docs/ci-setup.md"