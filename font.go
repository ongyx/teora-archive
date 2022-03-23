package teora

import (
	"image"
	"image/color"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const padding = 10

func center(img *ebiten.Image) image.Point {
	p := image.Pt(img.Size())
	p.X /= 2
	p.Y /= 2
	return p
}

// Font is a combination of a fontface and color, used to render text to images..
type Font struct {
	Face  font.Face
	Color color.Color
}

// Draw renders the text on an image at the point as-is (i.e with default alignment).
func (f *Font) Draw(
	str string,
	img *ebiten.Image,
	point image.Point,
) image.Rectangle {
	return f.Write(str, img, point, AlignDefault)
}

// Write renders the text on an image at the point with alignment and returns its bounds.
func (f *Font) Write(
	str string,
	img *ebiten.Image,
	point image.Point,
	align Align,
) image.Rectangle {
	s := text.BoundString(f.Face, str).Size()
	a := align.Adjust(point, s)

	text.Draw(img, str, f.Face, a.X, a.Y, f.Color)

	// ap is the bottom-left point of the bound.
	return image.Rect(
		a.X,
		a.Y-s.Y,
		a.X+s.X,
		a.Y,
	)
}

// WriteCenter renders the text in the center of an image.
func (f *Font) WriteCenter(str string, img *ebiten.Image) image.Rectangle {
	return f.Write(str, img, center(img), AlignCenter)
}

// Load loads an OpenType fontface from a source.
func (f *Font) Load(src []byte, opts *opentype.FaceOptions) error {
	tt, err := opentype.Parse(src)
	if err != nil {
		return err
	}

	face, err := opentype.NewFace(tt, opts)
	if err != nil {
		return err
	}

	f.Face = face
	return nil
}
