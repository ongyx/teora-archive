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

// Font is a wrapper around a fontface for easier rendering to an image.
type Font struct {
	Face  font.Face
	Color color.Color
}

// Draw renders the text on an image at the point as is.
func (f *Font) Draw(txt string, img *ebiten.Image, p image.Point) {
	text.Draw(img, txt, f.Face, p.X, p.Y, f.Color)
}

// Write renders the text on an image at the x and y coordinates with alignment.
func (f *Font) Write(txt string, img *ebiten.Image, p image.Point, al Align) {
	size := text.BoundString(f.Face, txt).Size()

	al.Adjust(&p.X, &p.Y, size)

	f.Draw(txt, img, p)
}

// WriteCenter renders the text in the center of an image.
func (f *Font) WriteCenter(txt string, img *ebiten.Image) {
	p := image.Pt(img.Size())
	p.X /= 2
	p.Y /= 2
	f.Write(txt, img, p, AlignCenter)
}

// Log renders the text at the bottom right of an image.
func (f *Font) Log(txt string, img *ebiten.Image) {
	width, height := img.Size()
	f.Write(
		txt,
		img,
		image.Pt(width-padding, height-padding),
		AlignLeft|AlignTop,
	)
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

	f.Face = face
	return nil
}
