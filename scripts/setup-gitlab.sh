#!/bin/bash
# setup-gitlab.sh - Set up GitLab repository and push code

set -e

# Configuration
GITLAB_USERNAME=\"cbwinslow\"
REPO_NAME=\"chezmoi-tui\"
GITLAB_URL=\"https://gitlab.com\"
LOCAL_REPO_PATH=\"/home/foomanchu8008/chezmoi-tui\"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Print colored output
print_info() {
    echo -e \"${BLUE}[INFO]${NC} $1\"
}

print_success() {
    echo -e \"${GREEN}[SUCCESS]${NC} $1\"
}

print_warning() {
    echo -e \"${YELLOW}[WARNING]${NC} $1\"
}

print_error() {
    echo -e \"${RED}[ERROR]${NC} $1\" >&2
}

# Check if we're in the right directory
if [[ ! -d \"${LOCAL_REPO_PATH}/.git\" ]]; then
    print_error \"Not a git repository: ${LOCAL_REPO_PATH}\"
    exit 1
fi

cd \"${LOCAL_REPO_PATH}\"

# Check if GitLab CLI is available
if ! command -v glab &> /dev/null; then
    print_warning \"GitLab CLI (glab) not found. Will use API directly.\"
    
    # Check if we have GitLab personal access token
    if [[ -z \"${GITLAB_TOKEN}\" ]]; then
        print_warning \"GITLAB_TOKEN environment variable not set.\"
        print_info \"Please set GITLAB_TOKEN with a personal access token:\"
        echo \"export GITLAB_TOKEN=your_gitlab_personal_access_token\"
        print_info \"You can create a token at: https://gitlab.com/-/profile/personal_access_tokens\"
        print_info \"The token needs 'api' scope.\"
        exit 1
    fi
    
    # Create repository using GitLab API
    print_info \"Creating GitLab repository...\"
    
    # Use curl to create the repository
    RESPONSE=$(curl -s -w \"%{http_code}\" -X POST \
        -H \"PRIVATE-TOKEN: ${GITLAB_TOKEN}\" \
        -H \"Content-Type: application/json\" \
        -d \"{\\\"name\\\": \\\"${REPO_NAME}\\\", \\\"visibility\\\": \\\"public\\\"}\" \
        \"${GITLAB_URL}/api/v4/projects/\")
    
    HTTP_CODE=$(echo \"${RESPONSE}\" | tail -c 4)
    RESPONSE_BODY=$(echo \"${RESPONSE}\" | head -c -4)
    
    if [[ \"${HTTP_CODE}\" == \"201\" ]]; then
        print_success \"GitLab repository created successfully!\"
    elif [[ \"${HTTP_CODE}\" == \"400\" ]] && echo \"${RESPONSE_BODY}\" | grep -q \"already exists\"; then
        print_warning \"Repository already exists on GitLab.\"
    else
        print_error \"Failed to create GitLab repository. HTTP ${HTTP_CODE}\"
        echo \"Response: ${RESPONSE_BODY}\"
        exit 1
    fi
else
    # Use GitLab CLI
    print_info \"Using GitLab CLI to create repository...\"
    
    # Check if already logged in
    if ! glab auth status &> /dev/null; then
        print_info \"Logging into GitLab...\"
        glab auth login
    fi
    
    # Create repository
    if ! glab repo create \"${GITLAB_USERNAME}/${REPO_NAME}\" --public --clone=false &> /dev/null; then
        print_warning \"Repository might already exist or creation failed.\"
    else
        print_success \"GitLab repository created successfully with CLI!\"
    fi
fi

# Add GitLab remote if it doesn't exist
if ! git remote get-url gitlab &> /dev/null; then
    print_info \"Adding GitLab remote...\"
    git remote add gitlab \"https://gitlab.com/${GITLAB_USERNAME}/${REPO_NAME}.git\"
    print_success \"GitLab remote added!\"
else
    print_info \"GitLab remote already exists.\"
fi

# Set up credentials for HTTPS push (if using token)
if [[ -n \"${GITLAB_TOKEN}\" ]]; then
    print_info \"Setting up GitLab credentials...\"
    git config credential.helper store
    echo \"https://${GITLAB_USERNAME}:${GITLAB_TOKEN}@gitlab.com\" > ~/.git-credentials
    chmod 600 ~/.git-credentials
fi

# Push to GitLab
print_info \"Pushing to GitLab...\"
if git push gitlab master; then
    print_success \"Successfully pushed to GitLab!\"
else
    print_error \"Failed to push to GitLab.\"
    print_info \"You may need to set up authentication manually.\"
    print_info \"Try: git push https://oauth2:[YOUR_TOKEN]@gitlab.com/${GITLAB_USERNAME}/${REPO_NAME}.git master\"
fi

print_success \"GitLab setup completed!\"
print_info \"Repository URL: https://gitlab.com/${GITLAB_USERNAME}/${REPO_NAME}\"