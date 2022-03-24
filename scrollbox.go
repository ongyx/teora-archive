package teora

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ongyx/teora/bento"
)

// Scrollbox is a box with scrolling text inside.
type Scrollbox struct {
	*bento.Scroll

	Pad image.Point
}

// Render renders the scrollbox into a new image.
func (sb *Scrollbox) Render() *ebiten.Image {
	// the size of this scrollbox is the scroll's size + padding for the background
	s := sb.Scroll.Size()
	s.Add(sb.Pad)

	// render the scrollbox background, and then draw the scroll on top.
	c := ebiten.NewImage(s.X, s.Y)
	c.Fill(color.White)

	sr := sb.Scroll.Render(color.Black)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(sb.Pad.X), float64(sb.Pad.Y))

	c.DrawImage(sr, op)

	return c
}
