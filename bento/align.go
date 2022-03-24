package bento

import (
	"image"
)

const (
	// AlignRight moves text to the right of a point.
	AlignRight Align = 1 << iota
	// AlignHCenter moves text to the horizontal center of a point.
	AlignHCenter
	// AlignLeft moves text to the left of a point.
	AlignLeft
	// AlignTop moves text above a point.
	AlignTop
	// AlignVCenter moves text to the vertical center of a point.
	AlignVCenter
	// AlignBottom moves text below a point.
	AlignBottom
	// AlignDefault is the default alignment of text (to the right and below).
	AlignDefault = AlignRight | AlignBottom
	// AlignCenter horizontally and vertically centers text on a point.
	AlignCenter = AlignHCenter | AlignVCenter
)

// Align specifies the alignment to render text at a point (x, y).
// Align must have at most one horizontal (AlignRight, AlignHCenter, AlignLeft) and vertical (AlignTop, AlignVCenter, AlignBottom) flag.
type Align int

// Has checks if the alignment flag is set.
func (a Align) Has(flag Align) bool {
	return (a & flag) != 0
}

// Adjust changes the point according to the text size,
// depending on the alignment flags set.
func (a Align) Adjust(point, size image.Point) image.Point {

	if a != AlignDefault {

		w := size.X
		h := size.Y

		// horizontal alignment
		if a.Has(AlignHCenter) {
			point.X -= w / 2
		} else if a.Has(AlignLeft) {
			point.X -= w
		}

		// vertical alignment
		// NOTE: The top left of the screen is (0, 0)!
		if a.Has(AlignVCenter) {
			point.Y -= h / 2
		} else if a.Has(AlignTop) {
			point.Y -= h
		}

	}

	return point
}
