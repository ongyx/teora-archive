package bento

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Transition is a animation that is rendered on top of a sprite/scene.
type Transition interface {
	// Update updates the state of the transition.
	Update() error
	// Draw renders the transition to the image.
	Draw(img *ebiten.Image)
	// Done checks if the transition's animation has finished.
	Done() bool
}
