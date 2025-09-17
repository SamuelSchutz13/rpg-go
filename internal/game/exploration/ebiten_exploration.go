package exploration

import (
	"image/color"

	"github.com/SamuelSchutz13/rpg-go/internal/game/model"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type GameActualScreen int

const (
	ScreenMenu GameActualScreen = iota
	ScreenExploration
	ScreenBattle
)

type Game struct {
	GameState model.GameState
	Screen    GameActualScreen
	PlayerImg *ebiten.Image
}

func NewEbitenGameExploration(gs model.GameState, screen GameActualScreen) *Game {
	player := ebiten.NewImage(16, 16)
	player.Fill(color.RGBA{255, 0, 0, 255})

	return &Game{
		GameState: gs,
		Screen:    ScreenMenu,
		PlayerImg: player,
	}
}

func (g *Game) Update() error {
	switch g.Screen {
	case ScreenMenu:
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			g.Screen = ScreenExploration
		}
	case ScreenExploration:
		if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
			g.GameState.PosX -= 2
		}
		if ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
			g.GameState.PosX += 2
		}
		if ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
			g.GameState.PosY -= 2
		}
		if ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
			g.GameState.PosY += 2
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.Screen {
	case ScreenMenu:
		ebitenutil.DebugPrint(screen, "xxx")
	case ScreenExploration:
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(g.GameState.PosX), float64(g.GameState.PosY))
		screen.DrawImage(g.PlayerImg, opts)
		ebitenutil.DebugPrint(screen, "explore")
	case ScreenBattle:
		ebitenutil.DebugPrint(screen, "battle")
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
