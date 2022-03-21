package teora

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ongyx/teora/data"
)

var teoran = NewFont(data.TeoranStandard, color.White)

// Game holds the main state of the game.
type Game struct {
}

// NewGame creates a game and initalises it's state.
func NewGame() *Game {
	return &Game{}
}

// Update updates the game's state.
func (g *Game) Update() error {
	return nil
}

// Draw draws to the screen based on the game's state.
func (g *Game) Draw(screen *ebiten.Image) {
	teoran.DrawCenter("Hello World!", screen, AlignCenter)
}

// Layout returns the screen's size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
