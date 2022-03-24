package bento

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var empty *ebiten.Image

func init() {
	empty = ebiten.NewImage(1, 1)
	empty.Fill(color.White)
}

// Rect draws a rectangle at point with size and color in an image.
func Rect(
	rect image.Rectangle,
	clr color.Color,
	img *ebiten.Image,
) {
	rs := rect.Size()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(float64(rs.X), float64(rs.Y))
	op.GeoM.Translate(float64(rect.Min.X), float64(rect.Min.Y))
	op.ColorM.Apply(clr)

	img.DrawImage(empty, op)
}
