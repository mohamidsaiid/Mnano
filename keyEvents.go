package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

func (app *Application) handleUp() {
	if app.cursorY > 0 {
		app.cursorY--
		if app.cursorX > len(app.buffer[app.cursorY]) {
			app.cursorX = len(app.buffer[app.cursorY])
		}
		app.updateCursor()
	}
}

func (app *Application) handleDown() {
	if app.cursorY < app.lines-1 {
		app.cursorY++
		if app.cursorX > len(app.buffer[app.cursorY])-1 {
			app.cursorX = len(app.buffer[app.cursorY])
		}
		app.updateCursor()
	}
}

func (app *Application) handleRight() {
	if app.cursorX < len(app.buffer[app.cursorY]) {
		app.cursorX++
	} else if app.cursorY < app.lines-1 {
		app.cursorY++
		app.cursorX = 0
	}
	app.updateCursor()
}

func (app *Application) handleLeft() {
	if app.cursorX > 0 {
		app.cursorX--
	} else if app.cursorY > 0 {
		app.cursorY--
		app.cursorX = len(app.buffer[app.cursorY])
	}
	app.updateCursor()
}

func (app *Application) handleInsert(r rune) {
	//insertion would be based on the cursor posiotion
	currentLine := app.buffer[app.cursorY]
	newLine := make([]rune, 0, len(currentLine)+1)
	newLine = append(newLine, currentLine[:app.cursorX]...)
	newLine = append(newLine, r)
	newLine = append(newLine, currentLine[app.cursorX:]...)
	app.buffer[app.cursorY] = newLine
	app.cursorX++

	app.isModified = true
	app.render()
}

func (app *Application) handleEnter() {
	currentLine := app.buffer[app.cursorY]

	beforeCursor := append([]rune{}, currentLine[:app.cursorX]...)
	afterCursor := append([]rune{}, currentLine[app.cursorX:]...)

	app.buffer[app.cursorY] = beforeCursor
	app.buffer = append(
		app.buffer[:app.cursorY+1],
		append([][]rune{afterCursor}, app.buffer[app.cursorY+1:]...)...,
	)

	app.cursorY++
	app.cursorX = 0
	app.isModified = true
	app.render()
}

func (app *Application) handleBackspace() {
	currentLine := app.buffer[app.cursorY]

	if app.cursorX > 0 && app.cursorX <= len(app.buffer[app.cursorY]) {
		// backspace in the middle of a line
		newLine := make([]rune, 0, len(currentLine))

		newLine = append(newLine, currentLine[:app.cursorX-1]...)
		newLine = append(newLine, currentLine[app.cursorX:]...)

		app.buffer[app.cursorY] = newLine
		app.cursorX--
	} else if app.cursorX == 0 && app.cursorY > 0 {
		// backspace at the start of the line to handle merging with the previousLine
		previousLine := app.buffer[app.cursorY-1]
		app.cursorX = len(previousLine)

		app.buffer[app.cursorY-1] = append(previousLine, currentLine...)

		newBuffer := app.buffer[:app.cursorY]
		newBuffer = append(newBuffer, app.buffer[app.cursorY+1:]...)

		app.buffer = newBuffer
		app.cursorY--
	}

	app.isModified = true
	app.render()
}

func (app *Application) handleDelete() {
	currentLine := app.buffer[app.cursorY]

	if app.cursorX > 0 && app.cursorX < len(currentLine) {
		newLine := make([]rune, 0, len(currentLine))

		newLine = append(newLine, currentLine[:app.cursorX]...)
		newLine = append(newLine, currentLine[app.cursorX+1:]...)

		app.buffer[app.cursorY] = append([]rune{}, newLine...)
		app.cursorX--
	} else if app.cursorX == 0 && app.cursorX < len(currentLine) {
		newLine := make([]rune, 0, len(currentLine))

		newLine = append(newLine, currentLine[:app.cursorX]...)
		newLine = append(newLine, currentLine[app.cursorX+1:]...)

		app.buffer[app.cursorY] = append([]rune{}, newLine...)
	} else if app.cursorX == len(currentLine) && app.cursorY < app.lines-1 {
		nextLine := app.buffer[app.cursorY+1]

		// Append next line to current line
		app.buffer[app.cursorY] = append(currentLine, nextLine...)

		// Remove next line from buffer
		app.buffer = append(app.buffer[:app.cursorY+1], app.buffer[app.cursorY+2:]...)
	}

	app.isModified = true
	app.render()
}

func (app *Application) updateCursor() {
	app.screen.ShowCursor(app.cursorX, app.cursorY)
	app.screen.Show()
}

func (app *Application) handleQuit() {
	if app.isModified {
		app.drawText(0, 100, "if you need to save the file press CTRL+s")
		return
	}
	app.screen.Fini()
	os.Exit(0)
}

func (app *Application) handleKeyEvents(ev *tcell.EventKey) {
	switch ev.Key() {
	case tcell.KeyUp:
		app.handleUp()

	case tcell.KeyDown:
		app.handleDown()
	case tcell.KeyRight:
		app.handleRight()
	case tcell.KeyLeft:
		app.handleLeft()
	case tcell.KeyRune:
		// adding the new character to the buffer
		// Regular character typed
		char := ev.Rune()

		// Filter out control characters (optional)
		if char >= 32 || char == '\t' {
			app.handleInsert(char)
		}
	case tcell.KeyBackspace, tcell.KeyBackspace2:
		app.handleBackspace()
	case tcell.KeyDelete:
		app.handleDelete()
	case tcell.KeyEnter:
		app.handleEnter()
	case tcell.KeyHome:
		// Jump to start of line
		app.cursorX = 0
		app.updateCursor()

	case tcell.KeyEnd:
		// Jump to end of line
		app.cursorX = len(app.buffer[app.cursorY])
		app.updateCursor()

	case tcell.KeyCtrlS:
		err := app.save()
		if err != nil {
			log.Fatalf("error while saving file: %v", err)
		}

	case tcell.KeyEscape, tcell.KeyCtrlC, tcell.KeyCtrlQ:
		// Multiple ways to quit
		app.handleQuit()

	}
}
