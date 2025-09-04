package main

import (
	"github.com/SamuelSchutz13/rpg-go/internal/ui/main_menu"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(main_menu.InitialMenu(), tea.WithAltScreen())
	_, err := p.Run()

	if err != nil {
		panic(err)
	}
}
