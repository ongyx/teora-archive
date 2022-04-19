package teora

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ongyx/teora/bento"
	"github.com/ongyx/teora/bento/anim"
)

var (
	StartScene bento.Scene
)

func init() {
	StartScene = &Start{
		subtitle: bento.NewScroll(Hack, "press [enter] or [space] to start"),
	}
}

type Start struct {
	subtitle *bento.Scroll
}

func (s *Start) Update(stage *bento.Stage) error {
	s.subtitle.Update()
	return nil
}

func (s *Start) Draw(screen *ebiten.Image) {
	teoran.WriteCenter("teora", color.White, screen)

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
