# Nano-like Text Editor

A minimal, terminal-based text editor written in Go, inspired by GNU nano. Built for learning terminal manipulation, file I/O, and building interactive CLI applications.

![Version](https://img.shields.io/badge/version-1.0.0-blue)
![Go](https://img.shields.io/badge/go-%3E%3D1.19-00ADD8)
![License](https://img.shields.io/badge/license-MIT-green)

---

## 📋 Table of Contents

- [Features](#features)
- [Screenshots](#screenshots)
- [Installation](#installation)
- [Usage](#usage)
- [Keyboard Shortcuts](#keyboard-shortcuts)
- [Architecture](#architecture)
- [How It Works](#how-it-works)
- [Development](#development)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)

---

## ✨ Features

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

## 📸 Screenshots

```
┌────────────────────────────────────────────┐
│ myfile.txt [Modified] | Ln 5, Col 12       │  ← Status Bar
├────────────────────────────────────────────┤
│ package main                               │
│                                            │
│ import "fmt"                               │
│                                            │
│ func main() {█                             │  ← Your Text + Cursor
│     fmt.Println("Hello, World!")           │
│ }                                          │
│                                            │
│ Saved to myfile.txt                        │  ← Status Message
└────────────────────────────────────────────┘
│ ^S Save  ^Q Quit  ^C Cancel  Arrows Move  │  ← Help Bar
└────────────────────────────────────────────┘
```

---

## 🚀 Installation

### Prerequisites
- **Go 1.19 or higher** - [Download Go](https://go.dev/dl/)
- A terminal emulator (works best with modern terminals)

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

## 📖 Usage

### Opening Files

```bash
# Open an existing file
./nano-editor myfile.txt

# Open multiple files (opens first file)
./nano-editor file1.txt

# Create a new file
./nano-editor newfile.txt
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
./nano-editor hello.go

# Edit content
# ... type your code ...

# Save changes
# Press: Ctrl+S

# Quit
# Press: Ctrl+Q
```

---

## 🤝 Contributing

Contributions are welcome! Here's how to help:

### Reporting Bugs
1. Check existing issues first
2. Include steps to reproduce
3. Mention your OS and Go version

### Suggesting Features
1. Open an issue with `[Feature Request]` prefix
2. Explain the use case
3. Propose implementation approach if possible

### Submitting Pull Requests
1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Code Style
- Follow standard Go formatting (`gofmt`)
- Add comments for complex logic
- Keep functions small and focused
- Write tests for new features

---

## 📚 Learning Resources

This project was built to learn:
- **Terminal programming** with tcell
- **Text editor internals** (buffers, cursors, rendering)
- **File I/O** in Go
- **Event-driven programming**
- **Go best practices**

### Recommended Reading
- [Build Your Own Text Editor](https://viewsourcecode.org/snaptoken/kilo/) - Excellent tutorial (in C)
- [tcell Documentation](https://github.com/gdamore/tcell)
- [How Nano Works](https://www.nano-editor.org/) - GNU nano source code

---

## 🐛 Known Issues

- Large files (>100MB) may be slow to load
- Some special Unicode characters may not render correctly
- Windows terminal may have different backspace key codes

---

## 📄 License

MIT License

Copyright (c) 2025

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

---

## 👨‍💻 Author

Built with ❤️ as a learning project to understand:
- Terminal manipulation
- Text editor architecture  
- Go programming
- Interactive CLI applications

---

## 🙏 Acknowledgments

- **GNU nano** - For inspiration and design principles
- **tcell library** - For excellent terminal handling
- **Go community** - For great documentation and examples

---

## 📞 Support

- **Issues**: [GitHub Issues](https://github.com/mohamidsaiid/mnano/issues)
- **Discussions**: [GitHub Discussions](https://github.com/mohamidsaiid/mnano/discussions)

---

**Happy Editing!** 🚀

If you found this project helpful, please ⭐ star it on GitHub!
