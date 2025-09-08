package main_menu

import (
	tea "github.com/charmbracelet/bubbletea"
)

type MenuModel struct {
	cursor  int
	choices []string
}

func InitialMenu() tea.Model {
	return MenuModel{
		choices: []string{
			"New Game", "Load Game", "Quit",
		},
	}
}

func (m MenuModel) Init() tea.Cmd {
	return nil
}

func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			switch m.cursor {
			case 0:
				return InitialDifficultyMenu(), nil
			case 1:
				return InitialLoadViewMenu(), nil
			case 2:
				return m, tea.Quit
			}
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		}
	}

	return m, nil
}

func (m MenuModel) View() string {
	s := "----- RPG in Terminal -----\n\n"

	for i, choice := range m.choices {
		cursor := " "

		if m.cursor == i {
			cursor = ">"
		}

		s += cursor + " " + choice + "\n"
	}

	return s
}
