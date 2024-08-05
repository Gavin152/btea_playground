package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

type model struct {
	items   []string
	cursor  int
	checked map[int]bool
}

func initialModel() model {
	return model{
		items:   []string{"Zirbe", "Schopf", "Erdäpfel", "Krenn", "Senf"},
		checked: make(map[int]bool),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor == len(m.items)-1 {
				m.cursor++
			}
		case "enter", "l":
			_, ok := m.checked[m.cursor]
			if ok {
				delete(m.checked, m.cursor)
			} else {
				m.checked[m.cursor] = true
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	uiString := "Was gibt's auf der Hütte?"
	return uiString
}

func main() {
	app := tea.NewProgram(initialModel())
	if _, err := app.Run(); err != nil {
		fmt.Println("Whoops, something went wrong here:\n", err)
		os.Exit(1)
	}
}
