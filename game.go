package teora

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

func scale(res int) int {
	return int(float64(res) * ebiten.DeviceScaleFactor())
}

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
	// draw tps at the top left of the screen
	hack.Draw(
		fmt.Sprintf("tps: %0.2f", ebiten.CurrentTPS()),
		screen,
		0, 0,
		AlignRight|AlignBottom,
	)

	teoran.DrawCenter("Hello World!", screen)
}

// Layout returns the screen's size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return scale(outsideWidth), scale(outsideHeight)
}
