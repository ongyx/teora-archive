package bento

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Animation is a effect rendered on a scene/sprite.
type Animation interface {
	Component

	// Draw renders the animation to the image.
	Draw(img *ebiten.Image)
	// Done checks if the animation has finished.
	Done() bool
}
