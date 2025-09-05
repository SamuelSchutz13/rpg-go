package main_menu

import (
	"strconv"

	"github.com/SamuelSchutz13/rpg-go/internal/game/model"
	tea "github.com/charmbracelet/bubbletea"
)

type ClassModel struct {
	cursor  int
	choices []model.Class
}

func InitialClassMenu() tea.Model {
	return ClassModel{
		choices: []model.Class{
			model.Class{
				Name: "Warrior",
				Hp:   200,
				Atk:  15,
				Def:  10,
			},
			model.Class{
				Name: "Wizard",
				Hp:   100,
				Atk:  25,
				Def:  5,
			},
			model.Class{
				Name: "Archer",
				Hp:   150,
				Atk:  20,
				Def:  3,
			},
		},
	}
}

func (m ClassModel) Init() tea.Cmd {
	return nil
}

func (m ClassModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
				return m, nil
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

func (m ClassModel) View() string {
	s := "----- RPG in Terminal -----\n\n"

	for i, choice := range m.choices {
		cursor := " "

		if m.cursor == i {
			cursor = "➜"
		}

		s += cursor + " " + choice.Name + " = " + " - Atk: " + strconv.Itoa(choice.Atk) + " - Def: " + strconv.Itoa(choice.Def) + " - Hp: " + strconv.Itoa(choice.Hp) + "\n"
	}

	return s
}
