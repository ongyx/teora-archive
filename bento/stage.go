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

	transition  Transition
	renderState RenderState
}

// NewStage creates a stage with an inital scene.
// NOTE: The initial scene's enter animation is rendered!
func NewStage(initial Scene) *Stage {
	return &Stage{
		scene:       initial,
		transition:  initial.Enter(),
		renderState: Entering,
	}
}

// Change changes the scene to render in the next frame.
func (s *Stage) Change(newScene Scene) {
	oldScene := s.scene

	log.Printf("changing scene (%p) -> (%p)\n", oldScene, newScene)

	s.setTransition(oldScene.Exit(), Exiting)

	s.scene = newScene
}

// Update updates the current scene's state.
func (s *Stage) Update() error {
	if s.renderState != Exiting {
		if err := s.scene.Update(s); err != nil {
			return err
		}
	}

	if s.renderState != Normal {
		if err := s.transition.Update(); err != nil {
			return err
		}
	}

	return nil
}

// Draw renders the current scene to the screen.
func (s *Stage) Draw(screen *ebiten.Image) {
	if s.snapshot == nil {
		s.snapshot = ebiten.NewImage(screen.Size())
	}

	// render the scene only if we aren't exiting
	if s.renderState != Exiting {
		s.snapshot.Clear()
		s.scene.Draw(s.snapshot)
	}

	screen.DrawImage(s.snapshot, nil)

	if s.renderState != Normal {
		s.transition.Draw(screen)

		if s.transition.Done() {
			// transition finished, change rendering state
			switch s.renderState {
			case Entering:
				log.Println("enter transition finished")
				s.setTransition(nil, Normal)
			case Exiting:
				log.Println("exit transition finished")
				// render the enter transition of the new scene
				s.setTransition(s.scene.Enter(), Entering)
			}
		}
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

func (s *Stage) setTransition(t Transition, rs RenderState) {
	if t != nil {
		s.transition = t
		s.renderState = rs
	} else {
		s.transition = nil
		s.renderState = Normal
	}
}
