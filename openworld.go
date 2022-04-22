package teora

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ongyx/teora/assets"
	"github.com/ongyx/teora/bento"
	"github.com/ongyx/teora/bento/anim"
)

const pixelScale = 4

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
)

type OpenWorld struct {
	bg, fg *ebiten.Image
	op     *ebiten.DrawImageOptions
}

func NewOpenWorld() bento.Scene {
	return &OpenWorld{}
}

func (w *OpenWorld) Update(stage *bento.Stage) error {
	return nil
}

func (w *OpenWorld) Draw(screen *ebiten.Image) {
	if w.op == nil {
		w.bg = assets.Demo.Render(bgmap)
		w.fg = assets.Demo.Render(fgmap)

		// scale to device dpi factor
		// TODO(ongyx): move to bento?
		sw, sh := screen.Size()
		ww, wh := w.bg.Size()
		scale := ebiten.DeviceScaleFactor() * pixelScale

		w.op = &ebiten.DrawImageOptions{}
		w.op.GeoM.Translate(-float64(ww)/2, -float64(wh)/2)
		w.op.GeoM.Scale(scale, scale)
		w.op.GeoM.Translate(float64(sw)/2, float64(sh)/2)
	}

	screen.DrawImage(w.bg, w.op)
	screen.DrawImage(w.fg, w.op)
}

func (w *OpenWorld) Enter() bento.Animation {
	return anim.NewFade(true, color.Black, 0.5)
}

func (w *OpenWorld) Exit() bento.Animation {
	return anim.NewFade(false, color.Black, 0.5)
}

func (w *OpenWorld) Entities() []*bento.Entity {
	return nil
}
