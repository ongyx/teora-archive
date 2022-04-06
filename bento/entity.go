package bento

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// RenderState represents the rendering state of a sprite/scene.
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

	sprite Sprite

	hidden, init bool

	transition  Transition
	renderState RenderState
}

// NewEntity constructs an entity from a sprite.
// Entities are hidden by default.
func NewEntity(sprite Sprite) *Entity {
	return &Entity{sprite: sprite, hidden: true}
}

// NewEntitySlice constructs a slice of entites from sprites.
func NewEntitySlice(sprites ...Sprite) []*Entity {
	entities := make([]*Entity, len(sprites))
	for i, s := range sprites {
		entities[i] = NewEntity(s)
	}

	return entities
}

// RenderState returns the rendering state of the sprite.
func (e *Entity) RenderState() RenderState {
	return e.renderState
}

// Hidden returns if the sprite is hidden or not.
func (e *Entity) Hidden() bool {
	return e.hidden
}

// Show shows the sprite if it's hidden, after rendering an enter transition.
// If t is nil, the sprite is immediately shown.
func (e *Entity) Show(t Transition) {
	if e.hidden {
		// NOTE: the sprite must be drawn during the enter transition.
		e.hidden = false

		if t != nil {
			e.transition = t
			e.renderState = Entering
		}
	}
}

// Hide hides the sprite if it's visible, after rendering an exit transition.
// If t is nil, the sprite is immediately hidden.
func (e *Entity) Hide(t Transition) {
	if !e.hidden {
		if t != nil {
			e.transition = t
			e.renderState = Exiting
		} else {
			e.hidden = true
		}
	}
}

func (e *Entity) update() error {
	if e.transition != nil {
		if err := e.transition.Update(); err != nil {
			return err
		}
	}

	if err := e.sprite.Update(e); err != nil {
		return err
	}

	return nil
}

func (e *Entity) draw(screen *ebiten.Image) {
	if !e.init {
		e.sprite.Init(e, screen.Bounds().Size())
		e.init = true
	}

	if !e.hidden {
		render := e.sprite.Render()

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(e.Position.X), float64(e.Position.Y))
		screen.DrawImage(render, op)

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
	}
}
