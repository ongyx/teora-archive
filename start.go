package teora

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ongyx/teora/assets"
	"github.com/ongyx/bento"
	"github.com/ongyx/bento/anim"
)

type Start struct {
	subtitle *bento.Scroll
}

func NewStart() bento.Scene {
	return &Start{
		subtitle: bento.NewScroll(assets.Hack, "press [enter] or [space] to start"),
	}
}

func (s *Start) Update(stage *bento.Stage) error {
	s.subtitle.Update()

	if bento.Keypress(confirmKeys) {
		stage.Change(NewOpenWorld())
	}

	return nil
}

func (s *Start) Draw(screen *ebiten.Image) {
	assets.Teoran.WriteCenter("teora", color.White, screen)

	b := screen.Bounds()

	s.subtitle.Draw(
		color.White,
		bento.Center.Align(
			image.Point{
				X: bento.Center.Point(b).X,
				Y: int(float64(b.Dy()) * 0.75),
			},
			s.subtitle.Size(),
		),
		screen,
	)
}

func (s *Start) Enter() bento.Animation {
	return anim.NewFade(true, color.Black, 0.5)
}

func (s *Start) Exit() bento.Animation {
	return anim.NewFade(false, color.Black, 0.5)
}
