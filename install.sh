#!/bin/bash

set -e

# C for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}  Nano Editor Installer${NC}"
echo -e "${GREEN}========================================${NC}"
echo ""

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo -e "${RED}Error: Go is not installed.${NC}"
    echo "Please install Go from https://golang.org/dl/"
    exit 1
fi

echo -e "${GREEN}✓ Go is installed ($(go version))${NC}"

# Check if git is installed
if ! command -v git &> /dev/null; then
    echo -e "${RED}Error: Git is not installed.${NC}"
    echo "Please install Git first."
    exit 1
fi

echo -e "${GREEN}✓ Git is installed${NC}"
echo ""

# Set variables
REPO_URL="https://github.com/mohamidsaiid/mnano.git"
INSTALL_DIR="/usr/bin/"
TEMP_DIR=$(mktemp -d)

echo -e "${YELLOW}Cloning repository...${NC}"
git clone "$REPO_URL" "$TEMP_DIR/mnano"
cd "$TEMP_DIR/mnano"

echo -e "${YELLOW}Installing dependencies...${NC}"
go mod tidy

echo -e "${YELLOW}Building nano-editor...${NC}"
go build -o mnano

# Create install directory if it doesn't exist
mkdir -p "$INSTALL_DIR"

echo -e "${YELLOW}Installing to $INSTALL_DIR...${NC}"
mv mnano "$INSTALL_DIR/"

# Check if install directory is in PATH
if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
    echo ""
    echo -e "${YELLOW}Warning: $INSTALL_DIR is not in your PATH${NC}"
    echo "Add this line to your ~/.bashrc or ~/.zshrc:"
    echo ""
    echo "    export PATH=\"\$PATH:$INSTALL_DIR\""
    echo ""
fi

# Cleanup
cd ~
rm -rf "$TEMP_DIR"

echo ""
echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}  Installation Complete!${NC}"
echo -e "${GREEN}========================================${NC}"
echo ""
echo "Run the editor with: mnano"
echo ""
