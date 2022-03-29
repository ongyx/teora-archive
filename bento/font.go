package bento

import (
	"image"
	"image/color"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const padding = 10

// Font is a wrapper around a fontface, used to render text to images.
type Font struct {
	Face font.Face
}

// Draw renders the text on an image at the point as-is (i.e without any alignment).
// NOTE: point is the bottom-left point of the text.
func (f *Font) Draw(
	str string,
	clr color.Color,
	img *ebiten.Image,
	point image.Point,
) {
	text.Draw(img, str, f.Face, point.X, point.Y, clr)
}

// Write renders the text on an image at the point with alignment and returns its bounds.
// point is the top-left point of the text.
func (f *Font) Write(
	str string,
	clr color.Color,
	img *ebiten.Image,
	point image.Point,
	align Align,
) image.Rectangle {
	s := text.BoundString(f.Face, str).Size()
	a := align.Align(point, s)

	// adjust Y axis because we need the bottom-left point, not the top-left point.
	f.Draw(str, clr, img, a.Add(image.Pt(0, s.Y)))

	return image.Rectangle{
		Min: a,
		Max: a.Add(s),
	}
}

// WriteCenter renders the text in the center of an image.
func (f *Font) WriteCenter(str string, clr color.Color, img *ebiten.Image) image.Rectangle {
	return f.Write(str, clr, img, Center.Point(img.Bounds()), Center)
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
