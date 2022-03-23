package bento

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Game is a game with one or more scenes.
// It implements the ebiten.Game interface.
// TODO: implement scene switching
type Game struct {
	Scenes  map[string]Scene
	current string
}

// NewGame creates a new game.
func NewGame() *Game {
	return &Game{
		Scenes: make(map[string]Scene),
	}
}

// Scene returns the current scene.
func (g *Game) Scene() Scene {
	return g.Scenes[g.current]
}

// SetScene sets the scene to render next.
// This must be called at least once.
func (g *Game) SetScene(name string) {
	g.current = name
}

// Update updates the current scene's state.
func (g *Game) Update() error {
	return g.Scene().Update()
}

// Draw renders the current scene to the screen.
func (g *Game) Draw(screen *ebiten.Image) {
	g.Scene().Render(screen)
}

// Layout returns the screen's size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return scale(outsideWidth), scale(outsideHeight)
}

func scale(res int) int {
	return int(float64(res) * ebiten.DeviceScaleFactor())
}
