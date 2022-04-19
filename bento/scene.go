package bento

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Scene is a special kind of entity that draws directly to a screen instead of rendering to a image.
type Scene interface {
	// Update updates the state of the scene, if any.
	Update(stage *Stage) error
	// Draw renders the scene on screen.
	Draw(screen *ebiten.Image)
	// Enter returns the enter transition of the scene, if any.
	Enter() Animation
	// Exit returns the enter transition of the scene, if any.
	Exit() Animation
	// Entities returns a slice of entities to render on the scene.
	Entities() []*Entity
}
