package teora

import (
	"image/color"
	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

// Align specifies the alignment to render text at a point (x, y).
type Align int

const (
	// AlignRight aligns text to the right of a point.
	AlignRight Align = iota
	// AlignCenter aligns text in the middle of a point.
	AlignCenter
	// AlignLeft aligns text to the left of a point.
	AlignLeft
)

// Center returns the xy coordinates of the center in an image.
func Center(img *ebiten.Image) (cx, cy int) {
	x, y := img.Size()
	return x / 2, y / 2
}

// Font is a combination of a fontface and a color.
type Font struct {
	face font.Face
	color color.Color
}

// NewFont creates a new font.
func NewFont(face font.Face, color color.Color) *Font {
	return &Font{
		face: face,
		color: color,
	}
}

// Draw renders the text on an image at the x and y coordinates.
func (f *Font) Draw(txt string, img *ebiten.Image, x, y int, al Align) {
	width := text.BoundString(f.face, txt).Size().X

	// adjust the point xy to align the text.		
	switch al {
		case AlignCenter:
			x -= width / 2
		case AlignLeft:
			x -= width
		default:
			// AlignRight, so do nothing
	}

	text.Draw(img, txt, f.face, x, y, f.color)
}

// DrawCenter renders the text in the center of an image.
func (f *Font) DrawCenter(txt string, img *ebiten.Image, al Align) {
	x, y := Center(img)
	f.Draw(txt, img, x, y, al)
}
