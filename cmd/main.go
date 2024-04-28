package main

import (
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nooooaaaaah/madMapper/config"
	"github.com/nooooaaaaah/madMapper/internal"
	"github.com/nooooaaaaah/madMapper/tui"
)

func main() {
	p := tea.NewProgram(tui.InitialModel())
	api.DiscoverMatterDevices("_matter._tcp", 10*time.Second)
	if _, err := p.Run(); err != nil {
		config.LogError("Error starting the program: %v", err)
		os.Exit(1)
	}
}
