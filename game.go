package teora

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

func scale(res int) int {
	return int(float64(res) * ebiten.DeviceScaleFactor())
}

// Game holds the main state of the game.
type Game struct {
	scroll *Scroll
	tick   int
}

// NewGame creates a game and initalises it's state.
func NewGame() *Game {
	return &Game{
		scroll: nil,
	}
}

// Update updates the game's state.
func (g *Game) Update() error {
	g.tick++
	return nil
}

// Draw draws to the screen based on the game's state.
func (g *Game) Draw(screen *ebiten.Image) {
	// draw tps/fps at the top left of the screen
	hack.Write(
		fmt.Sprintf("tps: %0.2f", ebiten.CurrentTPS()),
		screen,
		image.Pt(0, 0),
		AlignRight|AlignBottom,
	)

	// we can only init scroll here because we need to know the screen size
	if g.scroll == nil {
		cx, cy := screen.Size()
		g.scroll = NewScroll(teoran, "Hello World!", cx/2, cy/2, AlignCenter)
	}

	g.scroll.Render(screen, (g.tick%2) == 0)

	hack.Log("this is a log", screen)
}

// Layout returns the screen's size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return scale(outsideWidth), scale(outsideHeight)
}
