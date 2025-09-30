#!/bin/bash
# setup-secrets.sh - Script to demonstrate how to securely handle secrets with chezmoi-tui

set -e

echo "Setting up secure secrets management with chezmoi-tui..."

# Create a template file for secrets that can be encrypted by chezmoi
cat > .config/chezmoi/chezmoi-secrets.tmpl << 'EOF'
# This file is a chezmoi template
# Secrets are stored encrypted in chezmoi and decrypted during application

# GitHub Tokens
{{- if (and (exists ".github.token") (not (eq .github.token ""))) }}
GITHUB_TOKEN="{{ .github.token }}"
{{- end }}

# Docker Registry Credentials
{{- if (and (exists ".docker.username") (not (eq .docker.username ""))) }}
DOCKER_USERNAME="{{ .docker.username }}"
{{- end }}
{{- if (and (exists ".docker.password") (not (eq .docker.password ""))) }}
DOCKER_PASSWORD="{{ .docker.password }}"
{{- end }}

# AWS Credentials
{{- if (and (exists ".aws.accessKeyId") (not (eq .aws.accessKeyId ""))) }}
AWS_ACCESS_KEY_ID="{{ .aws.accessKeyId }}"
{{- end }}
{{- if (and (exists ".aws.secretAccessKey") (not (eq .aws.secretAccessKey ""))) }}
AWS_SECRET_ACCESS_KEY="{{ .aws.secretAccessKey }}"
{{- end }}

# SSH Keys
{{- if (and (exists ".ssh.privateKey") (not (eq .ssh.privateKey ""))) }}
SSH_PRIVATE_KEY="{{ .ssh.privateKey }}"
{{- end }}
EOF

echo "Created secrets template at .config/chezmoi/chezmoi-secrets.tmpl"
echo "This file should be managed by chezmoi for secure secret handling"

# Show how to encrypt a file with chezmoi
echo ""
echo "To encrypt secrets with chezmoi, use:"
echo "  chezmoi add --encrypt path/to/secret/file"
echo ""
echo "To decrypt secrets for use in workflows, use:"
echo "  chezmoi cat path/to/encrypted/secret"
echo ""

# Create an example GitHub Actions workflow that shows how to use these secrets
mkdir -p .github/workflows

cat > .github/workflows/secrets-example.yml << 'EOF'
name: Secrets Management Example

on:
  push:
    branches: [main, master]
  pull_request:
    branches: [main, master]

env:
  # Default environment variables (overridden by secrets where needed)
  DEFAULT_REGISTRY: ghcr.io

jobs:
  # Job to demonstrate secure secret handling
  secret-setup:
    runs-on: ubuntu-latest
    steps:
    - name: Checkout code
      uses: actions/checkout@v4
      with:
        fetch-depth: 0

    - name: Setup Go
      uses: actions/setup-go@v5
      with:
        go-version: '1.23'

    - name: Cache Go modules
      uses: actions/cache@v4
      with:
        path: |
          ~/.go/pkg/mod
          ~/go/pkg/mod
        key: ${{ runner.os }}-go-${{ hashFiles('**/go.mod') }}
        restore-keys: |
          ${{ runner.os }}-go-

    - name: Install chezmoi
      run: |
        curl -fsSL https://get.chezmoi.io | sh

    - name: Decrypt secrets
      run: |
        # Use gpg to decrypt secrets from chezmoi if they exist
        # In a real scenario, you'd use chezmoi properly to manage encrypted files
        echo "Decrypted secrets would be available here for use in the workflow"
        
        # Example: If you had a decrypted secret file, you'd source it like:
        # source .secrets.env || true  # The || true prevents failure if the file doesn't exist

    - name: Build application
      run: |
        go build -o chezmoi-tui .

    - name: Run tests
      run: |
        go test ./... -v

  # Job to demonstrate Docker build and push with secrets
  docker-build-and-push:
    runs-on: ubuntu-latest
    needs: secret-setup
    if: github.ref == 'refs/heads/master'
    steps:
    - name: Checkout code
      uses: actions/checkout@v4

    - name: Set up Docker Buildx
      uses: docker/setup-buildx-action@v3

    - name: Login to GitHub Container Registry
      uses: docker/login-action@v3
      if: env.GITHUB_REGISTRY_PASSWORD != null
      with:
        registry: ghcr.io
        username: ${{ github.actor }}
        password: ${{ secrets.GITHUB_TOKEN }}

    - name: Build and push Docker image
      uses: docker/build-push-action@v5
      with:
        context: .
        push: true
        tags: |
          ghcr.io/${{ github.repository }}:latest
          ghcr.io/${{ github.repository }}:${{ github.sha }}
        labels: |
          org.opencontainers.image.revision=${{ github.sha }}
          org.opencontainers.image.version=${{ github.ref_name }}

  # Job to demonstrate deployment with secrets
  deploy:
    runs-on: ubuntu-latest
    needs: [secret-setup, docker-build-and-push]
    if: github.ref == 'refs/heads/master'
    steps:
    - name: Deploy to production
      run: |
        echo "Deployment would happen here using secrets"
        # In real usage, you'd use the decrypted secrets for deployment
        # Example: Deploy to Kubernetes, AWS, etc.
EOF

echo "Created example GitHub Actions workflow at .github/workflows/secrets-example.yml"
echo "This shows how to properly handle secrets in GitHub Actions without exposing them"