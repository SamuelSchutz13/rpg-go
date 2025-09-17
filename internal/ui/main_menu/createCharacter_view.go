package main_menu

import (
	"fmt"

	"github.com/SamuelSchutz13/rpg-go/internal/game"
	"github.com/SamuelSchutz13/rpg-go/internal/game/exploration"
	"github.com/SamuelSchutz13/rpg-go/internal/game/model"
	"github.com/SamuelSchutz13/rpg-go/internal/pkg"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/hajimehoshi/ebiten/v2"
)

type CreateCharacterMenu struct {
	textInput textinput.Model
	class     model.Class
	err       error
}

type (
	errMsg error
)

func InitialCreateCharacterMenu(class model.Class) tea.Model {
	ti := textinput.New()
	ti.Placeholder = "Enter character name"
	ti.Focus()
	ti.CharLimit = 100
	ti.Width = 20

	return CreateCharacterMenu{
		textInput: ti,
		class:     class,
		err:       nil,
	}
}

func (m CreateCharacterMenu) Init() tea.Cmd {
	return textinput.Blink
}

func (m CreateCharacterMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			fileForSave, err := pkg.CreateRandomFile("saves", "save")

			if err != nil {
				panic(err)
			}

			character := model.Character{
				Name:  m.textInput.Value(),
				Class: m.class,
				Level: model.Level_1,
				Hp:    m.class.Hp,
				Xp:    0,
			}

			gameState := model.GameState{
				Character: character,
				PosX:      0,
				PosY:      0,
				Filename:  fileForSave,
			}

			game.Save(gameState, "")

			ebiten.SetWindowSize(640, 480)
			ebiten.SetWindowTitle("RPG in Go - Exploration Mode")

			if err := ebiten.RunGame(exploration.NewEbitenGameExploration(gameState, exploration.ScreenMenu)); err != nil {
				panic(err)
			}

			return m, nil
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m CreateCharacterMenu) View() string {
	return fmt.Sprintf(
		"Create a new character\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to return to choose class)",
	) + "\n"
}
