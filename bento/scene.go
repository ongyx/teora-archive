package bento

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Scene represents a 'level', or a segment of animation that is rendered on the screen.
type Scene interface {
	// Init initalizes the scene with the screen size and is called only once.
	Init(size image.Point)
	// Update updates the state of the scene, if any.
	Update(stage *Stage) error
	// Draw renders the scene on screen.
	Draw(screen *ebiten.Image)
	// Entities returns the entities on this scene.
	// Entities are always updated and drawn before the scene.
	Entities() []*Entity
	// Enter returns the enter transition of the scene, if any.
	Enter() Transition
	// Exit returns the enter transition of the scene, if any.
	Exit() Transition
}
