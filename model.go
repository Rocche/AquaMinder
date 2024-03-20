package main

import (
	"aquaminder/notification"
	"aquaminder/tui"
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

// Model is a struct compliant with the bubbletea.Model interface
type Model struct {
	form                 *huh.Form
	notificationsStarted bool
}

// NewModel creates the Model and populates its form
func NewModel() Model {
	return Model{
		form: huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Key("duration").
					Options(huh.NewOptions(
						"5s",
						"1m",
						"5m",
						"10m",
						"15m",
						"20m",
						"25m",
						"30m")...).
					Title("Choose the time interval at which you want to drink"),
			),
		),
	}
}

// Init initializes the model
func (m Model) Init() tea.Cmd {
	return m.form.Init()
}

// Update takes action based on user input
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
	}

	// start the notifications process if the form is completed and the
	// notifications process is not yet started
	if m.form.State == huh.StateCompleted && !m.notificationsStarted {
		duration := m.form.GetString("duration")
		interval, err := time.ParseDuration(duration)
		if err != nil {
			panic(err)
		}
		go beginNotifications(interval)
		m.notificationsStarted = true
	}
	return m, cmd
}

// View builds the string to render
func (m Model) View() string {
	b := strings.Builder{}
	b.WriteString(tui.WelcomeMessage())
	if m.form.State == huh.StateCompleted {
		duration := m.form.GetString("duration")
		b.WriteString(fmt.Sprintf("You will recieve a notification every %s. Happy drinking!\n\n", duration))
	} else {
		b.WriteString(m.form.View())
	}
	b.WriteString(tui.ExitInstructions())

	return b.String()
}

// beginNotifications is responsible of notifying the user at regular
// time intervals given in input. It is called as a goroutine
func beginNotifications(duration time.Duration) {
	for {
		notification.Notify()
		select {
		case <-time.After(duration):
			// do nothing
		}
	}
}
