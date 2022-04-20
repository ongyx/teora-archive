package bento

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// Stage is a scene manager which implements the ebiten.Game interface.
// The current scene must never be nil.
// If debug is not nil, debug mode is enabled.
type Stage struct {
	debug   *Debug
	scene   Scene
	entites []*Entity

	transition *Transition

	// snapshot holds the last rendered frame of the current scene.
	// This is used mainly for transitions.
	snapshot *ebiten.Image
}

// NewStage creates a stage with an inital scene.
// NOTE: The initial scene's enter animation is rendered!
func NewStage(initial Scene, debug *Debug) *Stage {
	s := &Stage{
		debug:      debug,
		transition: NewTransition(),
	}
	s.Change(initial)

	return s
}

// Change changes the scene to render in the next frame.
func (s *Stage) Change(newScene Scene) {
	oldScene := s.scene

	log.Printf("(%p) changing scene to %p\n", oldScene, newScene)

	if oldScene != nil {
		s.transition.Hide(oldScene.Exit())
	}

	s.scene = newScene
	s.entites = newScene.Entities()
}

// Update updates the current scene's state.
func (s *Stage) Update() error {
	if s.transition.RenderState() != Exiting {
		if err := s.updateEntities(); err != nil {
			return err
		}

		if err := s.scene.Update(s); err != nil {
			return err
		}
	}

	if s.transition.RenderState() == Hidden {
		// finished old scene's exit transition
		// render the enter transition of the new scene
		s.transition.Show(s.scene.Enter())
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
		//log.Printf("(%p) drawing to snapshot with %v state\n", s.scene, s.transition.RenderState())
		s.snapshot.Clear()
		s.drawEntities(s.snapshot)
		s.scene.Draw(s.snapshot)
	}

	screen.DrawImage(s.snapshot, nil)

	s.transition.Draw(screen)

	if s.debug != nil {
		// draw tps/fps at the top left of the screen
		s.debug.Font.Write(
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

func (s *Stage) updateEntities() error {
	for _, e := range s.entites {
		if err := e.Update(); err != nil {
			return err
		}
	}

	return nil
}

func (s *Stage) drawEntities(screen *ebiten.Image) {
	for _, e := range s.entites {
		e.Draw(screen)
	}
}
