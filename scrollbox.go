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
	size, pos image.Point
	bound     image.Rectangle

	stream bento.Stream
	done   bool
}

// NewScrollbox creates a new scrollbox from a stream.
func NewScrollbox(stream bento.Stream, font *bento.Font) *Scrollbox {
	s := bento.NewScroll(font, *stream.Read())

	return &Scrollbox{
		scroll: s,
		// the initial size of the scroll
		size:   s.Size(),
		stream: stream,
	}
}

func (sb *Scrollbox) Init(img *ebiten.Image) {
	b := img.Bounds()
	p := image.Point{
		X: bento.Center.Point(b).X,
		Y: int(float64(b.Dy()) * 0.9),
	}

	// initalise size relative to the screen's size.
	s := image.Pt(b.Dx()/2, b.Dy()/16)
	ap := bento.Center.Align(p, s)

	sb.bound = bento.Bound(ap, s)

	// align scroll to the left edge and vertically center of the scrollbox.
	sb.pos = bento.CenterLeft.Point(sb.bound)
	// adjust y manually so the text doesn't jitter becuase of changing height.
	sb.pos.Y += sb.size.Y / 2
}

func (sb *Scrollbox) Update() {
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
}

func (sb *Scrollbox) Render(img *ebiten.Image) {
	// draw the background of the scrollbox.
	pb := bento.Pad(sb.bound, image.Pt(10, 10))

	r := pb.Dy() / 2

	var vec bento.Vec

	// rectangle
	vec.Rect(pb)

	// semicircles
	vec.Arc(bento.CenterRight.Point(pb), r, bento.Deg90, bento.Deg270)
	vec.Arc(bento.CenterLeft.Point(pb), r, bento.Deg270, bento.Deg90)

	vec.Draw(sbColor, img, nil)
	//vec.DrawShader(color.White, img, gradient, nil)

	// render scroll text
	sb.scroll.Render(color.Black, sb.pos, img)
}

// Done checks if the scrollbox has finished scrolling all text.
func (sb *Scrollbox) Done() bool {
	return sb.scroll.Done() && sb.done
}
