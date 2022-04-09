package bento

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Scene represents a 'level', or a segment of animation that is rendered on the screen.
type Scene interface {
	// Update updates the state of the scene, if any.
	Update(stage *Stage) error
	// Draw renders the scene on screen.
	Draw(screen *ebiten.Image)
	// Enter returns the enter transition of the scene, if any.
	Enter() Transition
	// Exit returns the enter transition of the scene, if any.
	Exit() Transition
}
