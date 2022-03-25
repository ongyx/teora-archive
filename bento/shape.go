package bento

import (
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// Bound calculates the bounds of an image, given its top-left point and its size.
func Bound(point, size image.Point) image.Rectangle {
	return image.Rectangle{Min: point, Max: point.Add(size)}
}

// Pad pads the bounds of an image by a fixed amount.
func Pad(bounds image.Rectangle, pad image.Point) image.Rectangle {
	return image.Rectangle{
		Min: bounds.Min.Sub(pad),
		Max: bounds.Max.Add(pad),
	}
}

// Rect draws a rectangle at point with size and color.
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

// Radian converts an angle in degrees to radians.
func Radian(degree float64) float64 {
	return degree * (math.Pi / 180)
}
