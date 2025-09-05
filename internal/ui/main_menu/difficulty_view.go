package main_menu

import (
	tea "github.com/charmbracelet/bubbletea"
)

type DifficultyModel struct {
	cursor  int
	choices []string
}

func InitialDifficultyMenu() tea.Model {
	return DifficultyModel{
		choices: []string{
			"Easy", "Medium", "Hard", "Back",
		},
	}
}

func (m DifficultyModel) Init() tea.Cmd {
	return nil
}

func (m DifficultyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			switch m.cursor {
			case 0:
				return InitialClassMenu(), nil
			case 1:
				return InitialClassMenu(), nil
			case 2:
				return InitialClassMenu(), nil
			case 3:
				return InitialMenu(), nil
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

func (m DifficultyModel) View() string {
	s := "----- RPG in Terminal -----\n\n"

	for i, choice := range m.choices {
		cursor := " "

		if m.cursor == i {
			cursor = "➜"
		}

		s += cursor + " " + choice + "\n"
	}

	return s
}
