package bento

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Sprite is an image with state.
type Sprite interface {
	Component

	// Render renders the sprite to an image, given the screen's size.
	Render(entity *Entity, size image.Point) *ebiten.Image
}
