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
// Point is the center of the scrollbox.
type Scrollbox struct {
	scroll *bento.Scroll
	// The inital height of the scroll.
	scrollh int

	msg    []string
	msglen int
	index  int
	size   *image.Point
}

// NewScrollbox creates a new scrollbox, where msg are the texts to scroll.
func NewScrollbox(msg []string, font *bento.Font) *Scrollbox {
	s := bento.NewScroll(font, msg[0])

	return &Scrollbox{
		scroll:  s,
		scrollh: s.Size().Y,
		msg:     msg,
		msglen:  len(msg),
	}
}

// Update updates the scrollbox's state.
func (sb *Scrollbox) Update() {
	if bento.Keypress(confirmKeys) {
		// skip the text if it's still scrolling, otherwise go to the next text.
		if !sb.scroll.Done() {
			sb.scroll.Skip()
		} else {
			sb.index++
			if sb.index < len(sb.msg) {
				sb.scroll.SetText(sb.msg[sb.index])
			}
		}
	}
}

// Render renders the scrollbox onto a image.
// point is the center point of the scrollbox.
func (sb *Scrollbox) Render(point image.Point, img *ebiten.Image) {
	if sb.size == nil {
		// initalise size relative to the screen's size.
		w, h := img.Size()
		sb.size = &image.Point{X: w / 2, Y: h / 16}
	}

	p := bento.Center.Align(point, *sb.size)

	// draw the background of the scrollbox.
	b := bento.Bound(p, *sb.size)
	pb := bento.Pad(b, image.Pt(10, 10))

	r := pb.Dy() / 2

	var vec bento.Vec

	// rectangle
	vec.Rect(pb)

	// semicircles
	vec.Arc(bento.CenterRight.Point(pb), r, bento.Deg90, bento.Deg270)
	vec.Arc(bento.CenterLeft.Point(pb), r, bento.Deg270, bento.Deg90)

	vec.Draw(sbColor, img, nil)
	//vec.DrawShader(color.White, img, gradient, nil)

	// align scroll to the left edge and vertically center of the scrollbox.
	sp := bento.CenterLeft.Point(b)
	// adjust y manually so the text doesn't jitter becuase of changing height.
	sp.Y += sb.scrollh / 2

	// render scroll text
	sb.scroll.Render(color.Black, sp, img)
}

// Done checks if the scrollbox has finished scrolling all text.
func (sb *Scrollbox) Done() bool {
	return sb.scroll.Done() && sb.index == sb.msglen
}
