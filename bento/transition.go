package bento

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Transition struct {
	rs   RenderState
	anim Animation
}

// NewTransition creates a new transition.
func NewTransition() *Transition {
	return &Transition{rs: Hidden}
}

// Show sets the render state to Visible after rendering an enter animation.
// If enter is nil, the state is immediately changed.
func (t *Transition) Show(enter Animation) {
	log.Printf("(%p) entering (anim: %p)\n", t.anim, enter)

	if enter != nil {
		t.rs = Entering
		t.anim = enter
	} else {
		// NOTE: this clobbers the existing render state!
		t.rs = Visible
		t.anim = nil
	}
}

// Hide sets the render state to Hidden after rendering an exit transition.
// If exit is nil, the state is immediately changed.
func (t *Transition) Hide(exit Animation) {
	log.Printf("(%p) exiting (anim: %p)\n", t.anim, exit)

	if exit != nil {
		t.rs = Exiting
		t.anim = exit
	} else {
		t.rs = Hidden
		t.anim = nil
	}
}

// RenderState returns the rendering state of the transition.
func (t *Transition) RenderState() RenderState {
	return t.rs
}

// Update updates the transition's state.
func (t *Transition) Update() error {
	if a := t.transition(); a != nil {
		if err := a.Update(); err != nil {
			return err
		}
	}

	return nil
}

// Draw draws the transition onto the image.
func (t *Transition) Draw(image *ebiten.Image) {
	if a := t.transition(); a != nil {
		// draw transition over the image.
		a.Draw(image)

		if a.Done() {
			// transition finished, change rendering state
			switch t.rs {
			case Entering:
				log.Printf("(%p) entered\n", t.anim)
				t.rs = Visible
			case Exiting:
				log.Printf("(%p) exited\n", t.anim)
				t.rs = Hidden
			default:
				// this really shouldn't happen.
				panic("transition: inconsistent state")
			}

			t.anim = nil
		}
	}
}

func (t *Transition) transition() Animation {
	if t.rs == Entering || t.rs == Exiting {
		// sanity check
		if t.anim == nil {
			panic("transition: anim is nil")
		}

		return t.anim
	}

	return nil
}
