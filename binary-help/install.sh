#!/bin/bash

# Check if script is run with sudo
if [ "$EUID" -ne 0 ]; then
  echo "Please run this script with sudo to install to /usr/local/bin"
  echo "Usage: sudo ./install.sh"
  exit 1
fi

# Make the binary executable
chmod +x saika

# Copy to system path
cp saika /usr/local/bin/

echo "Saika has been installed to /usr/local/bin/saika"
echo "You can now run it from anywhere by typing 'saika'"
echo "Installation complete!"