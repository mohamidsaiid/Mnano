package main

import (
	"bufio"
	"os"
)

func (app *Application) save() error {
	file, err := os.Create(app.filename)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(file)
	for _, line := range app.buffer {
		writer.WriteString(string(line))
		writer.WriteString("\n")
	}
	writer.Flush()
	app.isModified = false
	return nil
}

func loadFile(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return [][]rune{{}}, nil 
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	buffer := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		buffer = append(buffer, []rune(line))
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}
	return buffer, nil
}
