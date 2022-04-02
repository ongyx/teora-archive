package bento

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Scene represents a 'level', or a segment of animation that is rendered on the screen.
type Scene interface {
	// Init initalizes the scene and is called only once.
	// This should be used to calculate points/sizes relative to the screen's size.
	Init(screen *ebiten.Image)
	// Update updates the state of the scene, if any.
	Update(stage *Stage) error
	// Draw renders any sprites/animations on screen.
	Draw(screen *ebiten.Image)
	// Enter returns the enter transition of the scene, if any.
	Enter() Transition
	// Exit returns the enter transition of the scene, if any.
	Exit() Transition
}
