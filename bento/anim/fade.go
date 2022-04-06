package anim

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ongyx/teora/bento"
)

const (
	alphaMin = 0
	alphaMax = 0xffff
)

// Fade is a transition that fades from/into a color.
type Fade struct {
	in    bool
	color color.NRGBA64
	clock *bento.Clock

	alpha, apt float64
	overlay    *ebiten.Image
}

// NewFade creates a new fade transition with a duration.
// If in is true, the transition fades from solid to transparent,
// otherwise the transition fades from transparent to solid.
func NewFade(in bool, clr color.Color, duration float64) *Fade {
	c := bento.NewClockOnce(duration)

	var a float64
	if in {
		a = alphaMax
	} else {
		a = alphaMin
	}

	return &Fade{
		in:    in,
		color: color.NRGBA64Model.Convert(clr).(color.NRGBA64),
		clock: c,
		alpha: a,
		apt:   alphaMax / float64(c.Limit()),
	}
}

func (f *Fade) Update() error {
	var d, m float64

	if f.in {
		d = -f.apt
		m = alphaMin
	} else {
		d = f.apt
		m = alphaMax
	}

	if !f.clock.Done() {
		f.alpha += d
	} else {
		f.alpha = m
	}

	f.clock.Tick()

	return nil
}

func (f *Fade) Draw(screen *ebiten.Image) {
	if f.overlay == nil {
		f.overlay = ebiten.NewImage(screen.Size())
	}

	f.color.A = uint16(f.alpha)
	f.overlay.Fill(f.color)

	screen.DrawImage(f.overlay, nil)
}

func (f *Fade) Done() bool {
	return f.clock.Done()
}
