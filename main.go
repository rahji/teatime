package main

import (
	"log"
	"time"

	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	statusMsg string
)

type model struct {
	timer timer.Model
}

const timeout = time.Second * 3

func InitialModel() model {
	m := model{
		timer: timer.NewWithInterval(0, time.Millisecond*50),
	}
	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case timer.TickMsg:
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyCtrlG:
			m.timer = timer.NewWithInterval(timeout, time.Millisecond*50)
			return m, m.timer.Init()
		}
	}
	return m, nil
}

func (m model) View() string {

	if m.timer.Timedout() {
		return "Press Ctrl+G"
	}

	return "PRESSED!"
}

func main() {

	p := tea.NewProgram(InitialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

}
