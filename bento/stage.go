package bento

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// Stage is a scene manager which implements the ebiten.Game interface.
// The current scene must never be nil.
type Stage struct {
	Current Scene

	debug bool
	font  *Font
}

// NewStage creates a stage with an inital scene.
func NewStage(initial Scene) *Stage {
	return &Stage{
		Current: initial,
	}
}

// Debug enables/disables debug mode.
// Among other things, this shows a TPS counter on the top left corner of the screen.
// A font is required to render certain elements.
func (s *Stage) Debug(enabled bool, font *Font) {
	s.debug = enabled
	s.font = font
}

// Update updates the current scene's state.
func (s *Stage) Update() error {
	return s.Current.Update(s)
}

// Draw renders the current scene to the screen.
func (s *Stage) Draw(screen *ebiten.Image) {
	s.Current.Render(screen)

	if s.debug {
		// draw tps/fps at the top left of the screen
		s.font.Write(
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
	return dpiscale(outsideWidth), dpiscale(outsideHeight)
}

func dpiscale(res int) int {
	return int(float64(res) * ebiten.DeviceScaleFactor())
}
