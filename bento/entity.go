package bento

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// RenderState represents the rendering state of a sprite/scene.
//go:generate stringer -type=RenderState
type RenderState int

const (
	// Normal means a sprite is rendering normally (default state).
	Normal RenderState = iota
	// Entering means a enter transition is rendering over a sprite.
	Entering
	// Exiting means a exit transition is rendering over a sprite.
	Exiting
)

// Entity is a sprite with rendering state.
// While a sprite renders to an image, an entity handles drawing the rendered image to the screen.
type Entity struct {
	// The position of the sprite on screen.
	Position image.Point
	Sprite   Sprite

	hidden, guard bool

	transition  Transition
	renderState RenderState
}

// NewEntity constructs an entity from a sprite.
// Entities are hidden by default.
func NewEntity(sprite Sprite) *Entity {
	return &Entity{Sprite: sprite, hidden: true}
}

// RenderState returns the rendering state of the sprite.
func (e *Entity) RenderState() RenderState {
	return e.renderState
}

// Hidden returns if the sprite is hidden or not.
func (e *Entity) Hidden() bool {
	return e.hidden
}

// Show shows the sprite, after rendering an enter transition.
// If t is nil, the sprite is immediately shown.
func (e *Entity) Show(t Transition) {
	// NOTE: the sprite must be drawn during the enter transition.
	e.hidden = false

	if t != nil {
		e.transition = t
		e.renderState = Entering
	}
}

// Hide hides the sprite, after rendering an exit transition.
// If t is nil, the sprite is immediately hidden.
func (e *Entity) Hide(t Transition) {
	if t != nil {
		e.transition = t
		e.renderState = Exiting
	} else {
		e.hidden = true
	}
}

// Update updates the sprite's state.
func (e *Entity) Update() error {
	if e.transition != nil {
		if err := e.transition.Update(); err != nil {
			return err
		}
	}

	if err := e.Sprite.Update(); err != nil {
		return err
	}

	return nil
}

// Draw draws the sprite's render onto the screen.
func (e *Entity) Draw(screen *ebiten.Image) {
	render := e.Sprite.Render(e, screen.Bounds().Size())

	if e.transition != nil {
		// draw transition over the render.
		e.transition.Draw(render)

		if e.transition.Done() {
			// transition finished, change rendering state
			switch e.renderState {
			case Entering:
				e.renderState = Normal
			case Exiting:
				e.hidden = true
				e.renderState = Entering
			}

			e.transition = nil
		}
	}

	if !e.hidden {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(e.Position.X), float64(e.Position.Y))
		screen.DrawImage(render, op)
	}
}
