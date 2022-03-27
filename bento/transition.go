package bento

import "github.com/hajimehoshi/ebiten/v2"

// Transition is a animation that is rendered in-between changing scenes.
type Transition interface {
	// Update updates the state of the transition.
	Update() error
	// Render renders the transition to the screen.
	Render(screen *ebiten.Image)
	// Done checks if the transition's animation has finished.
	Done() bool
}
