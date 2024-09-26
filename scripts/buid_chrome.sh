#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# Directory where the chromium source will be downloaded and built
CHROMIUM_DIR=../chromium

# Number of CPU cores to use during the build process
CORES=$(nproc)

# Function to check if a command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Step 1: Install required dependencies
install_dependencies() {
    echo "Installing dependencies..."
    sudo apt-get update
    sudo apt-get install -y git python3 curl lsb-release sudo

    # For Ubuntu/Debian-based distributions, install build tools as recommended by Chromium
    sudo apt-get install -y clang g++-multilib libc6-dev-i386 libgconf2-dev \
        libnss3-dev libxss-dev libxtst-dev
}

# Step 2: Download depot_tools and Chromium source code
download_chromium_source() {
    echo "Downloading Chromium source..."

    if [ ! -d "$CHROMIUM_DIR" ]; then
        mkdir -p "$CHROMIUM_DIR"
    fi

    cd "$CHROMIUM_DIR"

    # Install depot_tools if not already installed
    if [ ! -d "$CHROMIUM_DIR/depot_tools" ]; then
        echo "Installing depot_tools..."
        git clone https://chromium.googlesource.com/chromium/tools/depot_tools.git
        export PATH="$CHROMIUM_DIR/depot_tools:$PATH"
    fi

    # Fetch the Chromium source code
    if [ ! -d "$CHROMIUM_DIR/src" ]; then
        echo "Fetching Chromium..."
        fetch --nohooks chromium
    fi

    cd "$CHROMIUM_DIR/src"

    # Install the build dependencies (automated for Ubuntu/Debian)
    echo "Installing build dependencies..."
    ./build/install-build-deps.sh --no-prompt --no-arm --no-chromeos-fonts

    # Sync and update dependencies
    echo "Running gclient sync..."
    gclient sync
}

# Step 3: Configure the build for headless Chromium
configure_headless_build() {
    echo "Configuring headless Chromium build..."

    cd "$CHROMIUM_DIR/src"

    # Generate build files with GN, enabling headless mode and disabling unnecessary features
    gn gen out/Headless --args='
        is_debug=false
        headless=true
        use_ozone=false
        enable_nacl=false
        use_allocator="none"
    '
}

# Step 4: Build Chromium (headless mode)
build_chromium() {
    echo "Building Chromium (this may take a while)..."

    # Start the build using Ninja
    ninja -C out/Headless chrome -j$CORES
}

# Step 5: Check the binary size and completion
check_build() {
    echo "Build completed! Checking binary size..."

    # Check the size of the compiled Chromium binary
    BINARY_PATH="$CHROMIUM_DIR/src/out/Headless/chrome"

    if [ -f "$BINARY_PATH" ]; then
        echo "Chromium binary size:"
        du -h "$BINARY_PATH"
        echo "Binary path: $BINARY_PATH"
    else
        echo "Error: Chromium binary not found!"
        exit 1
    fi
}

# Step 6: Provide usage instructions
usage_instructions() {
    echo "To use the headless Chromium binary for PDF generation, run:"
    echo "$CHROMIUM_DIR/src/out/Headless/chrome --headless --disable-gpu --print-to-pdf=output.pdf https://www.example.com"
}

# Main script logic
main() {
    if ! command_exists "git"; then
        echo "Git is not installed. Installing dependencies first."
        install_dependencies
    fi

    download_chromium_source
    configure_headless_build
    build_chromium
    check_build
    usage_instructions
}

# Execute the main function
main