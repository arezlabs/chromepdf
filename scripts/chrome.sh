#!/bin/bash

# Define the download URL for Chromium for Linux
CHROME_DOWNLOAD_URL="https://commondatastorage.googleapis.com/chromium-browser-snapshots/Linux_x64"

# Fetch the latest build version number
latest=$(curl -s $CHROME_DOWNLOAD_URL/LAST_CHANGE)

# Construct the download URL for the latest Linux Chromium build
download_url="$CHROME_DOWNLOAD_URL/$latest/chrome-linux.zip"

# Define the directory to store the Chromium binary
CHROME_DIR="../chromium/linux"

# Create the directory if it doesn't exist
mkdir -p $CHROME_DIR

# Download and unzip Chromium
echo "Downloading Chromium for Linux (Build $latest)..."
curl -L $download_url -o chrome-linux.zip

echo "Extracting Chromium..."
unzip chrome-linux.zip -d $CHROME_DIR

# Clean up unnecessary files
echo "Cleaning up unnecessary files..."
rm -rf $CHROME_DIR/chrome-linux/locales       # Remove locale files
rm -rf $CHROME_DIR/chrome-linux/swiftshader   # Remove SwiftShader (GPU-related)
rm -rf $CHROME_DIR/chrome-linux/resources     # Remove resources
rm -rf $CHROME_DIR/chrome-linux/README.md     # Remove README files
rm -rf $CHROME_DIR/chrome-linux/chrome_100_percent.pak  # Remove unnecessary package

# Cleanup the zip file
rm chrome-linux.zip

echo "Chromium has been downloaded, extracted, and cleaned up in $CHROME_DIR"