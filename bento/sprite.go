package bento

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Sprite is a image with state.
type Sprite interface {
	// Init initalizes the sprite and is called only once.
	// This should be used to calculate points/sizes relative to the screen's size.
	Init(screen *ebiten.Image)
	// Update updates the state of the sprite, if any.
	Update() error
	// Render renders the sprite to an image.
	Render() *ebiten.Image
}
