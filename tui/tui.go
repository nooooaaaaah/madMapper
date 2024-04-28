package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	devices []string
	cursor  int
}

func InitialModel() *model {
	return &model{
		devices: []string{"Router", "Laptop", "Smartphone", "Printer"},
		cursor:  0,
	}
}

func (m *model) Init() tea.Cmd { return nil }

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.devices)-1 {
				m.cursor++
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m *model) View() string {
	s := "Network Devices:\n\n"
	for i, device := range m.devices {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor points at our selected device
		}
		s += fmt.Sprintf("%s %s\n", cursor, device)
	}
	s += "\nPress q to quit.\n"
	return s
}
