package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/ongyx/bento"
	"github.com/ongyx/bento/anim"
	"github.com/ongyx/teora/assets"
)

var (
	pixelScale float64

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

	arrow = make([]*ebiten.Image, 4)

	arrowKeys = []ebiten.Key{ebiten.KeyArrowUp, ebiten.KeyArrowDown, ebiten.KeyArrowLeft, ebiten.KeyArrowRight}
)

func init() {
	// scale to device dpi factor
	pixelScale = bento.DPIScale(4)

	for i := 0; i < 4; i++ {
		arrow[i] = assets.Arrow.Tile(i)
	}
}

type OpenWorld struct {
	bg, fg *ebiten.Image
	op     *ebiten.DrawImageOptions

	arrow *bento.Entity
}

func NewOpenWorld() bento.Scene {
	return &OpenWorld{
		bg:    assets.Demo.Render(bgmap),
		fg:    assets.Demo.Render(fgmap),
		arrow: bento.NewEntity(bento.NewCharacter(arrow[0:1])),
	}
}

func (w *OpenWorld) Update(stage *bento.Stage) error {
	a := w.arrow.Sprite.(*bento.Character)

	if err := w.arrow.Update(); err != nil {
		return err
	}

	for i, k := range arrowKeys {
		if inpututil.IsKeyJustPressed(k) {
			a.SetFrames(arrow[i : i+1])
			break
		}
	}

	return nil
}

func (w *OpenWorld) Draw(screen *ebiten.Image) {
	a := w.arrow.Sprite.(*bento.Character)

	if w.op == nil {
		c := bento.Center.Point(screen)

		g := bento.Geometry{}
		g.Align(bento.Center, w.bg.Bounds().Size())
		g.Scale(pixelScale)
		g.Translate(c)
		w.op = &ebiten.DrawImageOptions{GeoM: g.M}

		a.Align = bento.Center
		a.Scale = pixelScale
		a.Position = c
		w.arrow.Show(nil)
	}

	screen.DrawImage(w.bg, w.op)
	screen.DrawImage(w.fg, w.op)
	w.arrow.Draw(screen)
}

func (w *OpenWorld) Enter() bento.Animation {
	return anim.NewFade(true, color.Black, 0.5)
}

func (w *OpenWorld) Exit() bento.Animation {
	return anim.NewFade(false, color.Black, 0.5)
}
