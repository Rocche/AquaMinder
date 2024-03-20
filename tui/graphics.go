package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Banner is... well... the banner?
const Banner = `                                           
 _____             _____ _       _         
|  _  |___ _ _ ___|     |_|___ _| |___ ___ 
|     | . | | | .'| | | | |   | . | -_|  _|
|__|__|_  |___|__,|_|_|_|_|_|_|___|___|_|  
        |_|                                `

// WelcomeMessage returns the initial message of the program
func WelcomeMessage() string {
	b := strings.Builder{}
	b.WriteString(Banner)
	b.WriteString("\n\n")
	style := lipgloss.NewStyle().
		Bold(true).
		Background(lipgloss.Color("33"))
	b.WriteString("Welcome to ")
	b.WriteString(style.Render("AquaMinder"))
	b.WriteString(" and congratulations for beginning your quest of staying hydrated!\n\n")
	return b.String()
}

// ExitInstructions returns the phrase telling how to exit the program
func ExitInstructions() string {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("244"))
	return style.Render("(You can quit the program in any moment by pressing 'q', 'ESC' or 'CTRL+c')")
}
