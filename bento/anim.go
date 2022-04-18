package bento

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Animation is a effect rendered on a scene/sprite.
type Animation interface {
	// Update updates the logical state of the animation.
	Update() error
	// Draw renders the animation to the image.
	Draw(img *ebiten.Image)
	// Done checks if the animation has finished.
	Done() bool
}
