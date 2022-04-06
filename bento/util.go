package bento

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// Bound calculates a bound, given its top-left point and its size.
func Bound(point, size image.Point) image.Rectangle {
	return image.Rectangle{Min: point, Max: point.Add(size)}
}

// Pad adds padding to the bound by a fixed amount.
func Pad(bound image.Rectangle, pad image.Point) image.Rectangle {
	return image.Rectangle{
		Min: bound.Min.Sub(pad),
		Max: bound.Max.Add(pad),
	}
}

// Unpad removes padding from a bound by a fixed amount.
func Unpad(bound image.Rectangle, pad image.Point) image.Rectangle {
	return image.Rectangle{
		Min: bound.Min.Add(pad),
		Max: bound.Max.Sub(pad),
	}
}

// NewImageBound creates an image from a bound.
func NewImageBound(bound image.Rectangle) *ebiten.Image {
	size := bound.Size()
	return ebiten.NewImage(size.X, size.Y)
}

// Radian converts an angle in degrees to radians.
func Radian(degree float64) float64 {
	return degree * (math.Pi / 180)
}

// DPIScale scales the given resolution by the device's scale factor.
// This allows high-DPI rendering.
func DPIScale(res int) int {
	return int(float64(res) * ebiten.DeviceScaleFactor())
}
