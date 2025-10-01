# Chezmoi TUI - GitLab Mirror Setup

This document explains how to set up a GitLab mirror of this GitHub repository if needed.

## GitLab Mirror Setup Instructions

### Prerequisites
1. GitLab account with repository creation permissions
2. GitLab CLI installed (glab):
   ```bash
   # On Ubuntu/Debian
   sudo apt install glab
   
   # On macOS
   brew install glab
   
   # On Windows (with Chocolatey)
   choco install glab
   ```

### Initial Setup

1. Authenticate with GitLab:
   ```bash
   glab auth login
   ```

2. Create a new repository on GitLab:
   ```bash
   glab repo create cbwinslow/chezmoi-tui --public
   ```

3. Add GitLab as a remote:
   ```bash
   git remote add gitlab https://gitlab.com/cbwinslow/chezmoi-tui.git
   ```

4. Push all branches and tags to GitLab:
   ```bash
   git push gitlab --all
   git push gitlab --tags
   ```

### CI/CD Configuration for GitLab

If using GitLab CI/CD, create `.gitlab-ci.yml` with similar functionality to the GitHub Actions workflows:

```yaml
stages:
  - test
  - build
  - release

variables:
  DOCKER_DRIVER: overlay2
  DOCKER_TLS_CERTDIR: "/certs"
  CGO_ENABLED: "0"

.test_template: &test_definition
  stage: test
  image: golang:1.23
  script:
    - go mod download
    - go build -v ./...
    - go test -v ./...
  artifacts:
    reports:
      coverage: coverage.txt
  coverage: '/total:.*\s(\d+%)$/'

unit_test:
  <<: *test_definition
  coverage: '/total:.*\s(\d+\.\d+%)$/'
  script:
    - go test -v -coverprofile=coverage.txt -covermode=atomic ./...

security_scan:
  stage: test
  image: aquasec/trivy:latest
  script:
    - trivy fs --security-checks vuln,config --exit-code 1 .

build_linux:
  stage: build
  image: golang:1.23
  script:
    - mkdir -p dist
    - GOOS=linux GOARCH=amd64 go build -o dist/chezmoi-tui-linux-amd64 .
    - GOOS=linux GOARCH=arm64 go build -o dist/chezmoi-tui-linux-arm64 .
  artifacts:
    paths:
      - dist/
    expire_in: 1 week

build_macos:
  stage: build
  image: golang:1.23
  script:
    - mkdir -p dist
    - GOOS=darwin GOARCH=amd64 go build -o dist/chezmoi-tui-darwin-amd64 .
    - GOOS=darwin GOARCH=arm64 go build -o dist/chezmoi-tui-darwin-arm64 .
  artifacts:
    paths:
      - dist/
    expire_in: 1 week

build_windows:
  stage: build
  image: golang:1.23
  script:
    - mkdir -p dist
    - GOOS=windows GOARCH=amd64 go build -o dist/chezmoi-tui-windows-amd64.exe .
  artifacts:
    paths:
      - dist/
    expire_in: 1 week

docker_build:
  stage: build
  image: docker:20.10.16
  services:
    - docker:20.10.16-dind
  before_script:
    - docker login -u $CI_REGISTRY_USER -p $CI_REGISTRY_PASSWORD $CI_REGISTRY
  script:
    - docker build -t $CI_REGISTRY_IMAGE:$CI_COMMIT_SHA .
    - docker tag $CI_REGISTRY_IMAGE:$CI_COMMIT_SHA $CI_REGISTRY_IMAGE:latest
    - docker push $CI_REGISTRY_IMAGE:$CI_COMMIT_SHA
    - docker push $CI_REGISTRY_IMAGE:latest
  only:
    - master

release:
  stage: release
  image: registry.gitlab.com/gitlab-org/release-cli:latest
  before_script:
    - apt-get update && apt-get install -y curl
  script:
    - echo "Creating release for $CI_COMMIT_TAG"
  release:
    tag_name: $CI_COMMIT_TAG
    description: "Release $CI_COMMIT_TAG"
  only:
    - tags
```

### Synchronization

To keep GitHub and GitLab in sync, you can set up automatic mirroring:

1. In GitLab repository settings, enable pull mirroring from GitHub
2. Or use a GitHub Action to push changes to GitLab:
   ```yaml
   name: Sync to GitLab
   on:
     push:
       branches: [master]
       tags: ['*']
   jobs:
     sync:
       runs-on: ubuntu-latest
       steps:
         - uses: actions/checkout@v4
           with:
             fetch-depth: 0
         - name: Add GitLab remote
           run: |
             git remote set-url --add --push origin https://gitlab.com/username/repo.git
         - name: Push to GitLab
           run: |
             git push https://gitlab.com/username/repo.git ${GITHUB_REF}
   ```

### Feature Parity

The GitLab version should maintain feature parity with the GitHub version:
- Same CI/CD functionality
- Same security scanning
- Same multi-platform builds
- Same documentation
- Same testing coverage
```

This README provides instructions for setting up a GitLab mirror if needed, though the current GitHub repository is complete with all functionality.