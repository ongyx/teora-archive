package bento

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// DebugOptions are options for debug mode.
// A font is required to render certain elements.
type DebugOptions struct {
	Font *Font
}

// Stage is a scene manager which implements the ebiten.Game interface.
// The current scene must never be nil.
// If debug is not nil, debug mode is enabled.
type Stage struct {
	Debug *DebugOptions

	scene Scene

	// snapshot holds the last rendered frame of the current scene.
	// This is used mainly for transitions.
	snapshot *ebiten.Image

	transition *Transition
}

// NewStage creates a stage with an inital scene.
// NOTE: The initial scene's enter animation is rendered!
func NewStage(initial Scene) *Stage {
	s := &Stage{scene: initial, transition: NewTransition()}
	s.transition.Show(initial.Enter())

	return s
}

// Change changes the scene to render in the next frame.
func (s *Stage) Change(newScene Scene) {
	oldScene := s.scene

	log.Printf("changing scene (%p) -> (%p)\n", oldScene, newScene)

	s.transition.Hide(oldScene.Exit())

	s.scene = newScene
}

// Update updates the current scene's state.
func (s *Stage) Update() error {
	if s.transition.RenderState() != Exiting {
		if err := s.scene.Update(s); err != nil {
			return err
		}
	}

	if err := s.transition.Update(); err != nil {
		return err
	}

	return nil
}

// Draw renders the current scene to the screen.
func (s *Stage) Draw(screen *ebiten.Image) {
	if s.snapshot == nil {
		s.snapshot = ebiten.NewImage(screen.Size())
	}

	// render the scene only if we aren't exiting
	if s.transition.RenderState() != Exiting {
		s.snapshot.Clear()
		s.scene.Draw(s.snapshot)
	}

	screen.DrawImage(s.snapshot, nil)

	s.transition.Draw(screen)

	if s.transition.RenderState() == Hidden {
		// finished old scene's exit transition
		// render the enter transition of the new scene
		s.transition.Show(s.scene.Enter())
	}

	if s.Debug != nil {
		// draw tps/fps at the top left of the screen
		s.Debug.Font.Write(
			fmt.Sprintf("tps: %0.2f", ebiten.CurrentTPS()),
			color.White,
			screen,
			image.Pt(0, 0),
			Default,
		)
	}
}

// Layout returns the screen's size.
func (s *Stage) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return DPIScale(outsideWidth), DPIScale(outsideHeight)
}
