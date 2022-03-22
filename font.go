package teora

import (
	"image/color"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

// Font is a wrapper around a fontface for easier rendering to an image.
type Font struct {
	face  font.Face
	color color.Color
}

// Draw renders the text on an image at the x and y coordinates.
func (f *Font) Draw(txt string, img *ebiten.Image, x, y int, al Align) {
	size := text.BoundString(f.face, txt).Size()
	width := size.X
	height := size.Y

	// horizontal alignment
	if al.Has(AlignHCenter) {
		x -= width / 2
	} else if al.Has(AlignLeft) {
		x -= width
	}

	// vertical alignment
	// NOTE: The top left of the screen is (0, 0)!
	if al.Has(AlignVCenter) {
		y += height / 2
	} else if al.Has(AlignBottom) {
		y += height
	}

	text.Draw(img, txt, f.face, x, y, f.color)
}

// DrawCenter renders the text in the center of an image.
func (f *Font) DrawCenter(txt string, img *ebiten.Image) {
	x, y := img.Size()
	f.Draw(txt, img, x/2, y/2, AlignCenter)
}

// Load loads an OpenType fontface from a source.
func (f *Font) Load(src []byte, o *opentype.FaceOptions) error {
	tt, err := opentype.Parse(src)
	if err != nil {
		return err
	}

	face, err := opentype.NewFace(tt, o)
	if err != nil {
		return err
	}

	f.face = face
	return nil
}
