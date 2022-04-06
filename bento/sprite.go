package bento

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Sprite is an image with state.
type Sprite interface {
	// Init initalizes the sprite/entity state with the screen size and is called only once.
	Init(entity *Entity, size image.Point)
	// Update updates the state of the sprite.
	Update(entity *Entity) error
	// Render renders the sprite to an image.
	Render() *ebiten.Image
}
