package main_menu

import (
	"fmt"
	"os"

	"github.com/SamuelSchutz13/rpg-go/internal/game"
	tea "github.com/charmbracelet/bubbletea"
)

type LoadViewModel struct {
	cursor  int
	choices []string
}

func InitialLoadViewMenu() tea.Model {
	files, err := os.ReadDir("saves")

	if err != nil {
		panic(err)
	}

	var filesArray []string

	for _, v := range files {
		filesArray = append(filesArray, v.Name())
	}

	return LoadViewModel{
		choices: filesArray,
	}
}

func (m LoadViewModel) Init() tea.Cmd {
	return nil
}

func (m LoadViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			gameState := game.LoadSave(m.choices[m.cursor])
			fmt.Println(gameState)
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

func (m LoadViewModel) View() string {
	s := "----- Load Game -----\n\n"

	for i, choice := range m.choices {
		cursor := " "

		if m.cursor == i {
			cursor = ">"
		}

		s += cursor + " " + choice + "\n"
	}

	return s
}
