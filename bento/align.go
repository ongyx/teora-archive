package bento

import (
	"image"
)

const (
	// Right moves an image to the right of a point.
	Right Align = 1 << iota
	// HCenter moves an image to the horizontal center of a point.
	HCenter
	// Left moves an image to the left of a point.
	Left
	// Top moves an image above a point.
	Top
	// VCenter moves an image to the vertical center of a point.
	VCenter
	// Bottom moves an image below a point.
	Bottom

	// Default is the default alignment of an image (to the right and below).
	Default = BottomRight

	TopLeft   = Top | Left
	TopCenter = Top | HCenter
	TopRight  = Top | Right

	CenterLeft  = VCenter | Left
	Center      = VCenter | HCenter
	CenterRight = VCenter | Right

	BottomLeft   = Bottom | Left
	BottomCenter = Bottom | Center
	BottomRight  = Bottom | Right
)

// Align specifies the alignment to render an image at a point (x, y).
// Align must have at most one horizontal (AlignRight, AlignHCenter, AlignLeft) and vertical (AlignTop, AlignVCenter, AlignBottom) flag.
type Align int

// Has checks if the alignment flag is set.
func (a Align) Has(flag Align) bool {
	return (a & flag) != 0
}

// Align adjusts the point so that it will be the top-left point of an image given its size.
// The adjusted point can then be passed to ebiten.Image.DrawImage so the image will be in the correct position.
func (a Align) Align(point, size image.Point) image.Point {
	if a != Default {
		w := size.X
		h := size.Y

		// horizontal
		if a.Has(HCenter) {
			point.X -= w / 2
		} else if a.Has(Left) {
			point.X -= w
		}

		// vertical
		// NOTE: The top left of the screen is (0, 0)!
		if a.Has(VCenter) {
			point.Y -= h / 2
		} else if a.Has(Top) {
			point.Y -= h
		}
	}

	return point
}

// Point calculates a point in the bounds of an image.
func (a Align) Point(bounds image.Rectangle) image.Point {
	// top-left point
	p := bounds.Min

	w := bounds.Dx()
	h := bounds.Dy()

	if a.Has(Right) {
		p.X += w
	} else if a.Has(HCenter) {
		p.X += w / 2
	}

	if a.Has(VCenter) {
		p.Y += h / 2
	} else if a.Has(Bottom) {
		p.Y += h
	}

	return p
}
