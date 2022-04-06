package teora

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ongyx/teora/bento"
)

var (
	// TODO: add gradient?
	sbColor = color.NRGBA{250, 250, 250, 255}
	sbPad   = image.Pt(10, 10)

	confirmKeys = []ebiten.Key{ebiten.KeyEnter, ebiten.KeySpace}
)

// Scrollbox is a box with scrolling text inside.
type Scrollbox struct {
	scroll    *bento.Scroll
	scrollpos image.Point

	rect   image.Rectangle
	canvas *ebiten.Image

	stream bento.Stream
	done   bool
}

// NewScrollbox creates a new scrollbox from a stream.
func NewScrollbox(stream bento.Stream, font *bento.Font) *Scrollbox {
	s := bento.NewScroll(font, *stream.Read())

	return &Scrollbox{
		scroll: s,
		stream: stream,
	}
}

func (sb *Scrollbox) Init(entity *bento.Entity, size image.Point) {
	s := image.Pt(size.X/2, size.Y/16).Add(sbPad.Mul(2))

	// The total size of the scrollbox includes the semicircles on each side of the rectangle.
	b := bento.Bound(image.Point{}, s.Add(image.Pt(s.Y, 0)))

	sb.rect = bento.Bound(image.Pt(s.Y/2, 0), s)

	sb.scrollpos = bento.CenterLeft.Point(sb.rect)
	// offset scrollpos so it will be within the padding
	sb.scrollpos.X += sbPad.X
	sb.scrollpos.Y += sb.scroll.Size().Y / 2

	sb.canvas = bento.NewImageBound(b)

	p := image.Pt(size.X/2, int(float64(size.Y)*0.9))

	// TODO: enter transition
	entity.Position = bento.Center.Align(p, sb.canvas.Bounds().Size())
	entity.Show(nil)

	fmt.Println(p, entity.Position, sb.canvas.Bounds(), sb.rect)
}

func (sb *Scrollbox) Update(entity *bento.Entity) error {
	if bento.Keypress(confirmKeys) {
		// skip the text if it's still scrolling, otherwise go to the next text.
		if !sb.scroll.Done() {
			sb.scroll.Skip()
		} else {
			if t := sb.stream.Read(); t != nil {
				sb.scroll.SetText(*t)
			} else {
				sb.done = true
			}
		}
	}

	return nil
}

func (sb *Scrollbox) Render() *ebiten.Image {
	sb.canvas.Clear()

	r := sb.rect.Dy() / 2

	// draw the background of the scrollbox.
	var vec bento.Vec

	// rectangle
	vec.Rect(sb.rect)

	// semicircles
	vec.Arc(bento.CenterRight.Point(sb.rect), r, bento.Deg90, bento.Deg270)
	vec.Arc(bento.CenterLeft.Point(sb.rect), r, bento.Deg270, bento.Deg90)

	vec.Draw(sbColor, sb.canvas, nil)
	//vec.DrawShader(color.White, img, gradient, nil)

	// render scroll text
	sb.scroll.Draw(color.Black, sb.scrollpos, sb.canvas)

	return sb.canvas
}

// Done checks if the scrollbox has finished scrolling all text.
func (sb *Scrollbox) Done() bool {
	return sb.scroll.Done() && sb.done
}
