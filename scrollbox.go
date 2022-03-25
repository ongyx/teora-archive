package teora

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/ongyx/teora/bento"
)

// Scrollbox is a box with scrolling text inside.
// Point is the center of the scrollbox.
type Scrollbox struct {
	*bento.Scroll

	msg   []string
	index int
	size  *image.Point
}

// NewScrollbox creates a new scrollbox, where msg are the texts to scroll.
func NewScrollbox(msg []string, font *bento.Font) *Scrollbox {
	return &Scrollbox{
		Scroll: bento.NewScroll(font, msg[0]),
		msg:    msg,
	}
}

// Update updates the scrollbox's state.
func (sb *Scrollbox) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		// skip the text if it's still scrolling, otherwise go to the next text.
		if !sb.Done() {
			sb.Skip()
		} else {
			sb.index++
			if sb.index < len(sb.msg) {
				sb.SetText(sb.msg[sb.index])
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

	p := bento.Center.Adjust(point, *sb.size)

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

	vec.Draw(color.White, img)

	// align scroll to the left edge and vertically center of the scrollbox.
	sp := bento.CenterLeft.Point(b)
	asp := bento.CenterRight.Adjust(sp, sb.Scroll.Size())

	sb.Scroll.Render(color.Black, asp, img)
}
