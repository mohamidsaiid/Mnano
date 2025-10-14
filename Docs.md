# Documentation

## Table of Contents

- [Keyboard Shortcuts](#keyboard-shortcuts)
- [Architecture](#architecture)
- [How It Works](#how-it-works)
- [Development](#development)
- [Roadmap](#roadmap)

- ***

## Keyboard Shortcuts

### Navigation

| Key             | Action                |
| --------------- | --------------------- |
| `↑` `↓` `←` `→` | Move cursor           |
| `Home`          | Jump to start of line |
| `End`           | Jump to end of line   |

### Editing

| Key         | Action                         |
| ----------- | ------------------------------ |
| `Backspace` | Delete character before cursor |
| `Delete`    | Delete character at cursor     |
| `Enter`     | Create new line                |
| `Tab`       | Insert tab character           |

### File Operations

| Key      | Action    |
| -------- | --------- |
| `Ctrl+S` | Save file |

### Other

| Key      | Action                     |
| -------- | -------------------------- |
| `Ctrl+C` | Cancel/Quit (with warning) |
| `Esc`    | Cancel current operation   |

---

## Architecture

### Project Structure

```
nano-editor/
├── main.go          # Main editor logic, event loop, UI rendering
├── file.go          # File I/O operations (load/save)
├── go.mod           # Go module dependencies
├── go.sum           # Dependency checksums
└── README.md        # This file
```

### Key Components

#### **Application Struct**

```go
type Application struct {
    cursorX, cursorY int        // Cursor position
    buffer           [][]rune   // Text content (2D array of runes)
    screen           tcell.Screen
    defStyle         tcell.Style
    filename         string     // Current filename
    isModified       bool       // Unsaved changes flag
    statusMsg        string     // Status message to display
}
```

#### **Text Buffer**

- Stored as `[][]rune` - array of lines, each line is array of runes
- Uses `rune` type for proper Unicode support (handles emoji, Arabic, Chinese, etc.)
- Dynamic resizing as lines are added/removed

#### **Event Loop**

```go
for {
    ev := screen.PollEvent()  // Wait for keyboard/resize events
    switch ev := ev.(type) {
        case *tcell.EventKey:
            // Handle keyboard input
        case *tcell.EventResize:
            // Handle terminal resize
    }
}
```

---

## How It Works

### Text Buffer Management

The editor uses a 2D array structure for efficient text storage:

```go
buffer = [][]rune{
    {'H', 'e', 'l', 'l', 'o'},  // Line 0
    {'W', 'o', 'r', 'l', 'd'},  // Line 1
}
```

**Key Operations:**

1. **Insert Character**

   ```go
   // Split line at cursor, insert char, join back
   line = line[:cursorX] + char + line[cursorX:]
   ```

2. **New Line (Enter)**

   ```go
   // Split current line at cursor
   line1 = line[:cursorX]
   line2 = line[cursorX:]
   // Insert line2 as new line
   ```

3. **Delete (Backspace)**
   ```go
   // If in middle of line: remove char before cursor
   // If at start of line: merge with previous line
   ```

### File Operations

**Loading:**

```go
file, _ := os.Open(filename)
scanner := bufio.NewScanner(file)
for scanner.Scan() {
    line := scanner.Text()
    buffer = append(buffer, []rune(line))
}
```

**Saving:**

```go
file, _ := os.Create(filename)
writer := bufio.NewWriter(file)
for _, line := range buffer {
    writer.WriteString(string(line) + "\n")
}
writer.Flush()
```

### Screen Rendering

**Layout:**

```
Row 0:     Status bar (reversed colors)
Row 1-N:   Text buffer
Row N+1:   Status message (if any)
Row N+2:   Help bar (reversed colors)
```

**Rendering Pipeline:**

```go
1. Clear screen
2. Draw status bar
3. Draw text buffer
4. Draw help bar
5. Position cursor
6. Show screen
```

---

## Development

### Running Tests

```bash
# Create a test file
echo "Test line 1" > test.txt
echo "Test line 2" >> test.txt

# Open and test
go run main.go file.go test.txt
```

### Debugging

Add debug output (remember to remove before release):

```go
// In main.go, temporarily add:
log.Printf("CursorX: %d, CursorY: %d", app.cursorX, app.cursorY)
```

### Building for Release

```bash
# Build for current platform
go build -o nano-editor

# Build for Linux
GOOS=linux GOARCH=amd64 go build -o nano-editor-linux

# Build for macOS
GOOS=darwin GOARCH=amd64 go build -o nano-editor-mac

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o nano-editor.exe
```

---

## Roadmap

### Completed ✅

- [x] Basic text editing (insert, delete, cursor movement)
- [x] File loading from command line
- [x] File saving with Ctrl+S
- [x] Modified indicator
- [x] Quit protection
- [x] Status bar with position
- [x] Help bar with shortcuts

### Planned Features

#### **Phase 1: Essential Features**

- [ ] **Search** (Ctrl+W) - Find text in buffer
- [ ] **Cut/Paste** (Ctrl+K, Ctrl+U) - Cut and paste lines
- [ ] **Line numbers** - Display line numbers on the left
- [ ] **Undo/Redo** (Ctrl+Z, Ctrl+Y) - Undo/redo changes

#### **Phase 2: Enhanced Editing**

- [ ] **Find and Replace** - Search and replace text
- [ ] **Go to line** (Ctrl+G) - Jump to specific line number
- [ ] **Select text** - Visual selection with Shift+Arrows
- [ ] **Copy** (Ctrl+C) - Copy without cutting

#### **Phase 3: Advanced Features**

- [ ] **Syntax highlighting** - Color coding for Go, Python, etc.
- [ ] **Multiple buffers** - Edit multiple files at once
- [ ] **Split view** - View two files side by side
- [ ] **Auto-indent** - Automatic indentation
- [ ] **Tab completion** - Autocomplete file paths

#### **Phase 4: Polish**

- [ ] **Configuration file** - ~/.nanorc support
- [ ] **Themes** - Customizable color schemes
- [ ] **Mouse support** - Click to position cursor
- [ ] **Performance** - Optimize for large files (>10MB)

---

## Testing Checklist

Before releasing, test these scenarios:

- [ ] Open existing file
- [ ] Create new file
- [ ] Edit text (insert, delete, backspace)
- [ ] Navigate with all arrow keys
- [ ] Create multiple lines
- [ ] Delete lines (backspace at line start)
- [ ] Save existing file
- [ ] Save new file (with filename prompt)
- [ ] Try to quit with unsaved changes
- [ ] Quit with no changes
- [ ] Open non-existent file
- [ ] Test with Unicode characters (emoji, Arabic, etc.)
- [ ] Test with very long lines
- [ ] Test with empty file
- [ ] Resize terminal window

---
