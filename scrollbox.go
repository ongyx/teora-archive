package teora

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ongyx/teora/bento"
)

var (
	// TODO: add gradient?
	sbColor = color.NRGBA{250, 250, 250, 255}

	confirmKeys = []ebiten.Key{ebiten.KeyEnter, ebiten.KeySpace}
)

// Scrollbox is a box with scrolling text inside.
type Scrollbox struct {
	scroll    *bento.Scroll
	scrollpos image.Point

	vec    bento.Vec
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
	sb.canvas = ebiten.NewImage(int(float64(size.X)/1.5), size.Y/12)

	// Draw a transparent stadium on the canvas
	// https://en.wikipedia.org/wiki/Stadium_(geometry)
	b := sb.canvas.Bounds()
	d := b.Dy()
	r := d / 2
	a := b.Dx() - d

	rect := image.Rect(r, 0, r+a, d)
	rp := bento.CenterRight.Point(rect)
	lp := bento.CenterLeft.Point(rect)

	// rectangle
	sb.vec.Rect(rect)

	// semicircles
	sb.vec.Arc(rp, r, bento.Deg90, bento.Deg270)
	sb.vec.Arc(lp, r, bento.Deg270, bento.Deg90)

	sb.scrollpos = lp
	sb.scrollpos.Y += sb.scroll.Size().Y / 2

	p := image.Pt(size.X/2, int(float64(size.Y)*0.9))

	// TODO: enter transition
	entity.Position = bento.Center.Align(p, sb.canvas.Bounds().Size())
	entity.Show(nil)
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

	sb.vec.Draw(sbColor, sb.canvas, nil)

	// render scroll text
	sb.scroll.Draw(color.Black, sb.scrollpos, sb.canvas)

	return sb.canvas
}

// Done checks if the scrollbox has finished scrolling all text.
func (sb *Scrollbox) Done() bool {
	return sb.scroll.Done() && sb.done
}
