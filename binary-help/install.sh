#!/bin/bash

# Check if script is run with sudo
if [ "$EUID" -ne 0 ]; then
  echo "Please run this script with sudo to install to /usr/local/bin"
  echo "Usage: sudo ./install.sh"
  exit 1
fi

# Get the directory where the script is located
SCRIPT_DIR=$(dirname "$(readlink -f "$0")")
cd "$SCRIPT_DIR"

# Make the binary executable
chmod +x saika

# Copy to system path
cp saika /usr/local/bin/

echo "Saika has been installed to /usr/local/bin/saika"
echo "You can now run it from anywhere by typing 'saika'"

# Ask if user wants to clean up
read -p "Do you want to remove the downloaded archive and extracted files after installation? (y/n): " CLEANUP

if [ "$CLEANUP" = "y" ] || [ "$CLEANUP" = "Y" ]; then
  # Get parent directory
  PARENT_DIR=$(dirname "$SCRIPT_DIR")
  
  # Go up one level and remove the directory
  cd ..
  rm -rf "$(basename "$SCRIPT_DIR")"
  
  # Find and remove the archive file
  ARCHIVE_FILE=$(find "$PARENT_DIR" -name "saika-*.tar.gz" | head -1)
  if [ -n "$ARCHIVE_FILE" ]; then
    rm "$ARCHIVE_FILE"
    echo "Cleaned up downloaded files."
  fi
  
  echo "Installation complete and cleanup done."
else
  echo "Installation complete. Downloaded files were kept."
fi