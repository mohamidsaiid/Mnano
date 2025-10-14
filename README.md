# Nano-like Text Editor

A minimal, terminal-based text editor written in Go, inspired by GNU nano. Built for learning terminal manipulation, file I/O, and building interactive CLI applications.

![Version](https://img.shields.io/badge/version-1.0.0-blue)
![Go](https://img.shields.io/badge/go-%3E%3D1.19-00ADD8)
![License](https://img.shields.io/badge/license-MIT-green)

---

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)

---

## Features

### Core Editing

- ✅ **Full text editing** - Insert, delete, backspace with Unicode support
- ✅ **Cursor movement** - Arrow keys, Home, End with line wrapping
- ✅ **Multi-line support** - Create and delete lines with Enter/Backspace
- ✅ **Efficient buffer** - Uses `[][]rune` for proper Unicode handling

### File Operations

- ✅ **Load files** - Open existing files from command line
- ✅ **Save files** - Save with Ctrl+S, prompts for filename if needed
- ✅ **Modified indicator** - Shows `[Modified]` when file has unsaved changes
- ✅ **Quit protection** - Warns before quitting with unsaved changes

### User Interface

- ✅ **Status bar** - Shows filename, modified flag, cursor position
- ✅ **Help bar** - Quick reference for keyboard shortcuts
- ✅ **Status messages** - Real-time feedback for operations
- ✅ **Clean terminal UI** - Nano-inspired interface

### Platform Support

- ✅ **Cross-platform** - Works on Linux, macOS, and Windows
- ✅ **Terminal compatibility** - Works with most modern terminals

---

## Installation

### Prerequisites

- **Go 1.19 or higher** - [Download Go](https://go.dev/dl/)
- A terminal emulator (works best with modern terminals)

### Install using script

```bash
sudo sh -c "$(curl -sSL https://raw.githubusercontent.com/mohamidsaiid/Mnano/refs/heads/main/install.sh)"
```

or

### Install from Source

```bash
# Clone the repository
git clone https://github.com/yourusername/nano-editor.git
cd nano-editor

# Initialize Go module
go mod init nano-editor

# Install dependencies
go get github.com/gdamore/tcell/v2

# Build the editor
go build -o nano-editor

# Run it
./nano-editor
```

### Quick Start (Without Building)

```bash
# Run directly with Go
go run main.go file.go myfile.txt
```

---

## Usage

### Opening Files

```bash
# Open an existing file
mnano myfile.txt

# Open multiple files (opens first file)
mnano file1.txt

# Create a new file
mnano newfile.txt
```

### Basic Editing

1. **Type** to insert text
2. **Arrow keys** to move cursor
3. **Enter** to create new lines
4. **Backspace** to delete characters
5. **Ctrl+S** to save
6. **Ctrl+Q** to quit

### Saving Files

```bash
# If file was opened with a name:
Ctrl+S  # Saves immediately

# If new file without name:
Ctrl+S  # Prompts: "Save as: "
# Type filename and press Enter
```

### Example Workflow

```bash
# Open file
mnano hello.go

# Edit content
# ... type your code ...

# Save changes
# Press: Ctrl+S

# Quit
# Press: Ctrl+Q
```

---

### Recommended Reading

- [Build Your Own Text Editor](https://viewsourcecode.org/snaptoken/kilo/) - Excellent tutorial (in C)
- [tcell Documentation](https://github.com/gdamore/tcell)
- [How Nano Works](https://www.nano-editor.org/) - GNU nano source code

---

## Author

Built project to understand:

- Terminal manipulation
- Text editor architecture
- Interactive CLI applications

---

## Acknowledgments

- **GNU nano** - For inspiration and design principles
- **tcell library** - For excellent terminal handling
- **Go community** - For great documentation and examples

---

**Happy Editing!**

If you found this project helpful, please ⭐ star it on GitHub!
