package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// start the bubbletea program
	_, err := tea.NewProgram(NewModel()).Run()
	if err != nil {
		panic(err)
	}
}
