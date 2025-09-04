package main_menu

import (
	tea "github.com/charmbracelet/bubbletea"
)

type classModel struct {
	cursor  int
	choices []string
}

func InitialClassMenu() tea.Model {
	return classModel{
		choices: []string{
			"Warrior", "Wizard", "Archer", "Back",
		},
	}
}

func (m classModel) Init() tea.Cmd {
	return nil
}

func (m classModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			switch m.cursor {
			case 0:
				return m, nil
			case 1:
				return m, nil
			case 2:
				return m, nil
			case 3:
				return InitialDifficultyMenu(), nil
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

func (m classModel) View() string {
	s := "----- RPG in Terminal -----n\n"

	for i, choice := range m.choices {
		cursor := " "

		if m.cursor == i {
			cursor = "->"
		}

		s += cursor + " " + choice + "\n"
	}

	return s
}
