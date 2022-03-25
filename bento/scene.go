package bento

import "github.com/hajimehoshi/ebiten/v2"

// Scene represents a 'level', or a segment of animation that is rendered on the screen.
type Scene interface {
	// Update updates the state of the scene, if any.
	// If next is not nil, it will replace the current scene on stage.
	Update(stage *Stage) error
	// Render renders any sprites/animations on screen.
	Render(screen *ebiten.Image)
}
