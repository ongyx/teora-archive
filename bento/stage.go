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

	enter, exit Transition
	state       RenderState
}

// NewStage creates a stage with an inital scene.
// NOTE: The initial scene's enter animation is rendered!
func NewStage(initial Scene) *Stage {
	return &Stage{
		scene: initial,
		enter: initial.Enter(),
		state: Entering,
	}
}

// Change changes the scene to render in the next frame.
func (s *Stage) Change(newScene Scene) {
	// set the exit transition to the old scene's, and the enter transition to the new scene's.
	oldScene := s.scene

	s.exit = oldScene.Exit()
	s.enter = newScene.Enter()

	log.Printf("changing scene (%p) -> (%p)\n", oldScene, newScene)

	s.scene = newScene
	s.state = Exiting
}

// Update updates the current scene's state.
func (s *Stage) Update() error {
	if t := s.transition(); t != nil {
		if err := t.Update(); err != nil {
			return err
		}
	}

	if s.state != Exiting {
		if err := s.scene.Update(s); err != nil {
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
	if s.state != Exiting {
		s.snapshot.Clear()
		s.scene.Draw(s.snapshot)
	}

	screen.DrawImage(s.snapshot, nil)

	if t := s.transition(); t != nil {
		t.Draw(screen)

		if t.Done() {
			// transition finished, change rendering state
			switch s.state {
			case Entering:
				log.Println("enter transition finished")
				s.state = Normal
				s.enter = nil
			case Exiting:
				log.Println("exit transition finished")
				// render the enter transition of the new scene
				s.state = Entering
				s.exit = nil
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

func (s *Stage) transition() Transition {
	var t Transition

	switch s.state {
	case Entering:
		t = s.enter
	case Exiting:
		t = s.exit
	}

	return t
}
