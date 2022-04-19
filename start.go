package teora

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ongyx/teora/assets"
	"github.com/ongyx/teora/bento"
	"github.com/ongyx/teora/bento/anim"
)

var (
	fgmap = [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 26, 27, 28, 29, 30, 31, 0, 0, 0, 0},
		{0, 0, 0, 0, 51, 52, 53, 54, 55, 56, 0, 0, 0, 0},
		{0, 0, 0, 0, 76, 77, 78, 79, 80, 81, 0, 0, 0, 0},
		{0, 0, 0, 0, 101, 102, 103, 104, 105, 106, 0, 0, 0, 0},
		{0, 0, 0, 0, 126, 127, 128, 129, 130, 131, 0, 0, 0, 0},
		{0, 0, 0, 0, 303, 303, 245, 242, 303, 303, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0},
	}
	bgmap = [][]int{
		{243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243},
		{243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243},
		{243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243},
		{218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 244, 243},
		{243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243},
		{243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243},
		{243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243},
		{243, 244, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243},
		{243, 243, 243, 243, 243, 243, 243, 243, 219, 243, 243, 243, 219, 243},
		{243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243},
		{243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243},
		{243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243},
		{243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243},
		{218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 244, 243, 243, 243},
		{243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243},
	}

	StartScene bento.Scene
)

func init() {
	StartScene = &Start{
		subtitle: bento.NewScroll(assets.Hack, "press [enter] or [space] to start"),
	}
}

type Start struct {
	subtitle *bento.Scroll
	bg, fg   *ebiten.Image
	op       *ebiten.DrawImageOptions
}

func (s *Start) Update(stage *bento.Stage) error {
	s.subtitle.Update()
	return nil
}

func (s *Start) Draw(screen *ebiten.Image) {
	if s.op == nil {
		s.bg = assets.Demo.Render(bgmap)
		s.fg = assets.Demo.Render(fgmap)

		// scale to x axis (proportionally).
		// TODO: maybe move algorithm to bento?
		sw, sh := screen.Size()
		w, h := s.bg.Size()
		scale := float64(sh) / float64(h)

		s.op = &ebiten.DrawImageOptions{}
		s.op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
		s.op.GeoM.Scale(scale, scale)
		s.op.GeoM.Translate(float64(sw)/2, float64(sh)/2)
	}

	screen.DrawImage(s.bg, s.op)
	screen.DrawImage(s.fg, s.op)

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

func (s *Start) Entities() []*bento.Entity {
	return nil
}
