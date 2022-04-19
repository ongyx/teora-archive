package bento

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Entity is a sprite with rendering state.
// While a sprite renders to an image, an entity handles drawing the rendered image to the screen.
// Op is the options used to draw the entity to the screen.
type Entity struct {
	Sprite
	*Transition

	Op *ebiten.DrawImageOptions
}

// NewEntity constructs an entity from a sprite.
// Entities are hidden by default.
func NewEntity(sprite Sprite) *Entity {
	return &Entity{Sprite: sprite, Transition: NewTransition()}
}

// NewEntities constructs a slice of entities from several sprites.
func NewEntities(sprites ...Sprite) []*Entity {
	// alloc slice with the exact size
	es := make([]*Entity, len(sprites))
	for i, s := range sprites {
		es[i] = NewEntity(s)
	}

	return es
}

// Update updates the sprite's state.
func (e *Entity) Update() error {
	if err := e.Sprite.Update(); err != nil {
		return err
	}

	if err := e.Transition.Update(); err != nil {
		return err
	}

	return nil
}

// Draw draws the sprite's render onto the screen.
func (e *Entity) Draw(screen *ebiten.Image) {
	render := e.Sprite.Render(e, screen.Bounds().Size())

	e.Transition.Draw(render)

	if e.Transition.RenderState() != Hidden {
		screen.DrawImage(render, e.Op)
	}
}
