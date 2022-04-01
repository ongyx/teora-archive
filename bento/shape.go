package bento

import (
	"image"
	"math"
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

// Radian converts an angle in degrees to radians.
func Radian(degree float64) float64 {
	return degree * (math.Pi / 180)
}
