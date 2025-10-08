package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

type Application struct {
	cursorX, cursorY int
	lines            int
	buffer           [][]rune
	screen           tcell.Screen
	defStyle         tcell.Style
	statusBarStyle   tcell.Style
	helpMenuStyle    tcell.Style
	statusMsg        string
	filename         string
	isModified       bool
}

func newApplication() (*Application, error) {
	filename := ""
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	if err := screen.Init(); err != nil {
		return nil, err
	}

	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	screen.SetStyle(defStyle)
	screen.Clear()

	buffer, err := loadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Application{
		buffer:   buffer,
		screen:   screen,
		defStyle: defStyle,
		filename: filename,
	}, nil
}

func main() {
	app, err := newApplication()
	if err != nil {
		log.Fatalf("%+v", err)
		return
	}

	message := "Nano-like Editor | Type to begin | ESC or Ctrl+C to quit"
	app.drawText(0, 0, message, app.defStyle)
	app.screen.Show()

	for {
		ev := app.screen.PollEvent()

		app.screen.ShowCursor(app.cursorX+1, app.cursorY)

		switch ev := ev.(type) {
		case *tcell.EventResize:
			app.screen.Sync()
			app.render()
		case *tcell.EventKey:
			// Handle keyboard events
			app.handleKeyEvents(ev)
		}

	}
}
