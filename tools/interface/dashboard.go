package main

import (
	"fmt"
	"os"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct{}

func (m model) Init() tea.Cmd { return nil }
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if _, ok := msg.(tea.KeyMsg); ok {
		return m, tea.Quit
	}
	return m, nil
}
func (m model) View() string {
	return "PhoenixOS Dashboard (Press q to quit)"
}

func main() {
	p := tea.NewProgram(model{})
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
