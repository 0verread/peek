#!/bin/bash

set -e

# Configuration
BINARY_NAME="peek"
INSTALL_DIR="/usr/local/bin"
GITHUB_REPO="0verread/peek"
VERSION="latest"
MAIN_PATH="cmd/peek/main.go"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m'

# Print step information
info() {
    echo -e "${BLUE}==>${NC} $1"
}

# Print success messages
success() {
    echo -e "${GREEN}==>${NC} $1"
}

# Print error messages
error() {
    echo -e "${RED}==>${NC} $1"
}

# Check for required tools
check_requirements() {
    info "Checking system requirements..."
    
    if ! command -v go >/dev/null 2>&1; then
        error "Go is not installed. Please install Go first."
        exit 1
    fi

    if ! command -v git >/dev/null 2>&1; then
        error "Git is not installed. Please install Git first."
        exit 1
    fi
}

# Parse command line arguments
parse_args() {
    while getopts "v:d:" opt; do
        case $opt in
            v) VERSION="$OPTARG";;
            d) INSTALL_DIR="$OPTARG";;
            \?) error "Invalid option -$OPTARG"; exit 1;;
        esac
    done
}

# Detect OS and architecture
detect_platform() {
    info "Detecting platform..."
    
    OS=$(uname -s | tr '[:upper:]' '[:lower:]')
    ARCH=$(uname -m)
    
    # Convert architecture names
    case $ARCH in
        x86_64) ARCH="amd64";;
        aarch64) ARCH="arm64";;
        armv7l) ARCH="arm";;
    esac
    
    success "Detected: $OS/$ARCH"
}

# Create temporary directory
setup_temp_dir() {
    TEMP_DIR=$(mktemp -d)
    trap 'rm -rf "$TEMP_DIR"' EXIT
    
    info "Created temporary directory: $TEMP_DIR"
}

# Download and build the project
build_project() {
    info "Building $BINARY_NAME..."
    
    cd "$TEMP_DIR"
    
    if [ "$VERSION" = "latest" ]; then
        git clone "https://github.com/$GITHUB_REPO" .
    else
        git clone -b "$VERSION" "https://github.com/$GITHUB_REPO" .
    fi
    
    # Check if the main.go exists in the expected location
    if [ -f "$MAIN_PATH" ]; then
        info "Building from $MAIN_PATH"
        go build -o "$BINARY_NAME" "./$MAIN_PATH"
    else
        # Try to find main.go recursively
        MAIN_FILES=$(find . -name "main.go")
        if [ -z "$MAIN_FILES" ]; then
            error "No main.go found in the repository"
            error "Repository structure:"
            ls -R
            exit 1
        fi
        
        # Use the first main.go found
        MAIN_FILE=$(echo "$MAIN_FILES" | head -n 1)
        info "Building from found main.go at: $MAIN_FILE"
        go build -o "$BINARY_NAME" "$(dirname "$MAIN_FILE")"
    fi
    
    if [ ! -f "$BINARY_NAME" ]; then
        error "Build failed: Binary was not created"
        exit 1
    fi
    
    success "Build completed successfully"
}

# Install the binary
install_binary() {
    info "Installing to $INSTALL_DIR..."
    
    # Create install directory if it doesn't exist
    mkdir -p "$INSTALL_DIR"
    
    # Check if we have write permissions
    if [ ! -w "$INSTALL_DIR" ]; then
        error "No write permission to $INSTALL_DIR. Please run with sudo."
        exit 1
    fi
    
    # Move binary to install directory
    mv "$TEMP_DIR/$BINARY_NAME" "$INSTALL_DIR/"
    chmod +x "$INSTALL_DIR/$BINARY_NAME"
    
    success "Installation completed!"
    success "You can now run '$BINARY_NAME' from your terminal."
}

# Check if the binary is properly installed
verify_installation() {
    if command -v "$BINARY_NAME" >/dev/null 2>&1; then
        success "Verification successful: $BINARY_NAME is installed and accessible"
        "$BINARY_NAME" --version 2>/dev/null || true
    else
        error "Verification failed: $BINARY_NAME is not accessible in PATH"
        error "Please check your installation"
        exit 1
    fi
}

# Main installation process
main() {
    echo "Installing $BINARY_NAME..."
    
    parse_args "$@"
    check_requirements
    detect_platform
    setup_temp_dir
    build_project
    install_binary
    verify_installation
    
    echo
    success "Installation completed successfully!"
    echo "Run '$BINARY_NAME --help' to get started"
}

# Run main function
main "$@"