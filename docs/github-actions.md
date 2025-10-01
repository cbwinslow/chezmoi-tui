# Chezmoi TUI - GitHub Actions Workflows

This document details the GitHub Actions workflows implemented in this project.

## Workflows Overview

The project includes several GitHub Actions workflows designed for continuous integration, security scanning, and deployment:

1. **CI Workflow** (`.github/workflows/ci.yml`) - Automated testing across multiple platforms and Go versions
2. **Release Workflow** (`.github/workflows/release.yml`) - Automated release generation for multiple platforms
3. **Advanced CI/CD** (`.github/workflows/advanced-cicd.yml`) - Comprehensive pipeline with security scanning and deployment

## CI Workflow Features

- Tests across multiple Go versions (1.21.x, 1.22.x, 1.23.x)
- Tests on multiple operating systems (Ubuntu, macOS, Windows)
- Code coverage reporting to Codecov
- Linting with golangci-lint

## Release Workflow Features

- Automatic binary generation for multiple platforms
- Cross-platform builds (Linux, macOS, Windows)
- Automatic checksum generation
- GitHub release creation with assets

## Advanced CI/CD Workflow Features

- Security vulnerability scanning with Trivy
- Matrix builds for multiple platforms
- Docker build and push with multi-arch support
- Environment-based deployments (staging, production)
- Manual approval gates for production
- Code quality checks

## Secrets Management

The workflows demonstrate secure handling of secrets using:
- GitHub repository secrets
- Encrypted files managed by Chezmoi
- Environment-specific configurations
- Secure credential handling patterns