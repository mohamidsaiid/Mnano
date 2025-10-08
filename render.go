package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
)

func (app *Application) drawText(x, y int, message string, style tcell.Style) {
	for _, val := range message {
		app.screen.SetContent(x, y, val, nil, app.defStyle)
		x++
	}
}

func (app *Application) drawBuffer() {
	for i := range app.buffer {
		app.drawText(0, i, string(app.buffer[i]), app.defStyle)
	}
}

// handleSave saves the file
func (app *Application) handleSave() {
	// If no filename, prompt for one
	if app.filename == "" {
		newFilename := app.promptForInput("Save as: ")
		if newFilename == "" {
			app.statusMsg = "Save cancelled"
			return
		}
		app.filename = newFilename
	}

	// Save the file
	err := app.save()
	if err != nil {
		app.statusMsg = fmt.Sprintf("Error: %v", err)
	} else {
		app.statusMsg = fmt.Sprintf("Saved to %s", app.filename)
		app.isModified = false
	}
}

// promptForInput shows a prompt and gets user input
func (app *Application) promptForInput(prompt string) string {
	_, height := app.screen.Size()
	input := ""
	promptRow := height - 2

	for {
		// Render normal screen
		app.render()

		// Draw prompt over the status message area
		app.screen.Clear()
		app.renderStatusBar()
		app.renderHelpBar()
		app.renderBuffer()

		// Draw prompt
		promptStyle := app.defStyle.Foreground(tcell.ColorYellow)
		app.drawTextAt(0, promptRow, prompt+input, promptStyle)
		app.screen.ShowCursor(len(prompt)+len(input), promptRow)
		app.screen.Show()

		// Wait for input
		ev := app.screen.PollEvent()
		if ev, ok := ev.(*tcell.EventKey); ok {
			switch ev.Key() {
			case tcell.KeyEnter:
				return input
			case tcell.KeyEscape, tcell.KeyCtrlC:
				return ""
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				if len(input) > 0 {
					input = input[:len(input)-1]
				}
			case tcell.KeyRune:
				input += string(ev.Rune())
			}
		}
	}
}

// handleQuit attempts to quit, warns if modified
func (app *Application) handleQuit() {
	if app.isModified {
		app.statusMsg = "File modified! Press Ctrl+S to save, or Ctrl+Q again to quit"
		app.render()

		// Wait for confirmation
		ev := app.screen.PollEvent()
		if ev, ok := ev.(*tcell.EventKey); ok {
			if ev.Key() == tcell.KeyCtrlQ {
				app.quit()
			}
		}
	} else {
		app.quit()
	}
}

// render redraws the entire screen
func (app *Application) render() {
	app.screen.Clear()
	app.renderStatusBar()
	app.renderBuffer()
	app.renderHelpBar()
	app.updateCursor()
}

// renderStatusBar draws the top status bar
func (app *Application) renderStatusBar() {
	width, _ := app.screen.Size()

	// Build status text
	modFlag := ""
	if app.isModified {
		modFlag = " [Modified]"
	}

	filenameDisplay := app.filename
	if filenameDisplay == "" {
		filenameDisplay = "[No Name]"
	}

	status := fmt.Sprintf(" %s%s | Ln %d, Col %d ",
		filenameDisplay, modFlag, app.cursorY+1, app.cursorX+1)

	// Pad to full width
	for len(status) < width {
		status += " "
	}

	// Draw with reversed colors
	reverseStyle := app.defStyle.Reverse(true)
	app.drawTextAt(0, 0, status, reverseStyle)
}

// renderHelpBar draws the bottom help bar
func (app *Application) renderHelpBar() {
	width, height := app.screen.Size()
	helpRow := height - 1

	help := " ^S Save  ^C Quit  ^Q Cancel  Arrows Move "
	for len(help) < width {
		help += " "
	}

	reverseStyle := app.defStyle.Reverse(true)
	app.drawTextAt(0, helpRow, help, reverseStyle)
}

// renderBuffer draws the text buffer
func (app *Application) renderBuffer() {
	_, height := app.screen.Size()
	startRow := 1         // Leave space for status bar
	maxRows := height - 2 // Leave space for status and help bars

	for i := 0; i < len(app.buffer) && i < maxRows; i++ {
		line := app.buffer[i]
		for col, char := range line {
			app.screen.SetContent(col, startRow+i, char, nil, app.defStyle)
		}
	}

	// Draw status message if present
	if app.statusMsg != "" {
		msgRow := height - 2
		msgStyle := app.defStyle.Foreground(tcell.ColorYellow)
		app.drawTextAt(0, msgRow, app.statusMsg, msgStyle)
	}
}

// drawTextAt draws text at specific position with style
func (app *Application) drawTextAt(x, y int, text string, style tcell.Style) {
	col := x
	for _, char := range text {
		app.screen.SetContent(col, y, char, nil, style)
		col++
	}
}

// quit cleanly exits the application
func (app *Application) quit() {
	app.screen.Fini()
	os.Exit(0)
}
