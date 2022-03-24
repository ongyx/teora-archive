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

// Center returns the approximate center coordinates of an image.
func Center(img *ebiten.Image) image.Point {
	p := image.Pt(img.Size())
	p.X /= 2
	p.Y /= 2
	return p
}

// Font is a wrapper around a fontface, used to render text to images..
type Font struct {
	Face font.Face
}

// Draw renders the text on an image at the point as-is (i.e with default alignment) on a new image.
func (f *Font) Draw(
	str string,
	clr color.Color,
) *ebiten.Image {
	s := text.BoundString(f.Face, str).Size()
	i := ebiten.NewImage(s.X, s.Y)

	// NOTE: ebiten's text module treats the point as the bottom-left bound of the text drawn,
	// so we have to shift the y point here.
	text.Draw(i, str, f.Face, 0, s.Y, clr)

	return i
}

// Write renders the text on an image at the point with alignment and returns its bounds.
func (f *Font) Write(
	str string,
	clr color.Color,
	img *ebiten.Image,
	point image.Point,
	align Align,
) image.Rectangle {
	i := f.Draw(str, clr)
	s := image.Pt(i.Size())

	// a is now the top-left bound of the text.
	a := align.Adjust(point, s)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(a.X), float64(a.Y))

	img.DrawImage(i, op)

	return image.Rectangle{
		Min: a,
		Max: a.Add(s),
	}
}

// WriteCenter renders the text in the center of an image.
func (f *Font) WriteCenter(str string, clr color.Color, img *ebiten.Image) image.Rectangle {
	return f.Write(str, clr, img, Center(img), AlignCenter)
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
