package main

func (app *Application) drawText(x, y int, message string) {
	for _, val := range message {
		app.screen.SetContent(x, y, val, nil, app.defStyle)
		x++
	}
}

func (app *Application) drawBuffer() {
	for i := range app.buffer {
		app.drawText(0, i, string(app.buffer[i]))
	}
}

func (app *Application) render() {
	// clear screen for rerender
	app.screen.Clear()

	// how many lines in the textfile
	app.lines = len(app.buffer)
	// Draw all lines from buffer
	app.drawBuffer()
	// Update cursor position and show
	app.updateCursor()
}
