package bento

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Stage is a scene manager which implements the ebiten.Game interface.
// The current scene must never be nil.
type Stage struct {
	Current Scene
}

// NewStage creates a stage with an inital scene.
func NewStage(inital Scene) *Stage {
	return &Stage{
		Current: inital,
	}
}

// Update updates the current scene's state.
func (s *Stage) Update() error {
	next, err := s.Current.Update()
	if next != nil {
		s.Current = next
	}

	return err
}

// Draw renders the current scene to the screen.
func (s *Stage) Draw(screen *ebiten.Image) {
	s.Current.Render(screen)
}

// Layout returns the screen's size.
func (s *Stage) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return scale(outsideWidth), scale(outsideHeight)
}

func scale(res int) int {
	return int(float64(res) * ebiten.DeviceScaleFactor())
}
