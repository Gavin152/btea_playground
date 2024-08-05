package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
)

type model struct {
	items    []string
	cursor   int
	checked  map[int]bool
	viewport viewport.Model
}

var doneStyle lipgloss.Style = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#009688")).
	Strikethrough(true)

var openStyle lipgloss.Style = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#962244")).
	Strikethrough(false)

func initialModel() model {
	vp := viewport.New(80, 10)
	vp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#009688")).
		PaddingLeft(2)
	return model{
		items:    []string{"Zirbe", "Schopf", "Erdäpfel", "Krenn", "Senf"},
		cursor:   0,
		checked:  make(map[int]bool),
		viewport: vp,
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
			if m.cursor < len(m.items)-1 {
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
	uiString := "Was gibt's auf der Hütte?\n\n"

	for i, item := range m.items {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		style := openStyle
		if _, ok := m.checked[i]; ok {
			checked = "x"
			style = doneStyle
		}

		itemString := fmt.Sprintf("%s [%s] %s", cursor, checked, item)
		uiString += style.Render(itemString)
		uiString += "\n"
	}

	uiString += fmt.Sprintf("\nPress 'q' or 'crtl+x' to quit")
	m.viewport.SetContent(uiString)
	return m.viewport.View()
}

func main() {
	app := tea.NewProgram(initialModel())
	if _, err := app.Run(); err != nil {
		fmt.Println("Whoops, something went wrong here:\n", err)
		os.Exit(1)
	}
}
