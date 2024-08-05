package main

import (
	tea "github.com/charmbracelet/bubbletea"
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
