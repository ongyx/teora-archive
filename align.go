package teora

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
	// AlignDefault is the default alignment of text (to the right and above).
	AlignDefault = AlignRight | AlignTop
	// AlignCenter horizontally and vertically centers text on a point.
	AlignCenter = AlignHCenter | AlignVCenter
)

// Align specifies the alignment to render text at a point (x, y).
// Align must have at most one horizontal (AlignRight, AlignHCenter, AlignLeft) and vertical (AlignTop, AlignVCenter, AlignBottom) flag.
type Align int

// Has checks if the alignment flag is set.
func (al Align) Has(flag Align) bool {
	return (al & flag) != 0
}

// Adjust changes the values of x and y according to the text size,
// depending on the alignment flags set.
func (al Align) Adjust(x, y *int, size image.Point) {
	width := size.X
	height := size.Y

	// horizontal alignment
	if al.Has(AlignHCenter) {
		*x -= width / 2
	} else if al.Has(AlignLeft) {
		*x -= width
	}

	// vertical alignment
	// NOTE: The top left of the screen is (0, 0)!
	if al.Has(AlignVCenter) {
		*y += height / 2
	} else if al.Has(AlignBottom) {
		*y += height
	}
}
