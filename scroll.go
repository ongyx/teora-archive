// Scroll provides text scrolling across a screen.

package teora

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

// Scroll allows text to be scrolled across an image, starting from a point.
type Scroll struct {
	Font *Font
	Text string

	index, end int
	point      image.Point
}

// NewScroll creates a new scroll at a point on an image.
func NewScroll(font *Font, txt string, x, y int, al Align) *Scroll {
	size := text.BoundString(font.Face, txt).Size()
	al.Adjust(&x, &y, size)

	return &Scroll{
		Font:  font,
		Text:  txt,
		end:   len(txt),
		point: image.Pt(x, y),
	}
}

// Render renders the scroll onto an image.
// If advance is true, the next character will be drawn.
func (s *Scroll) Render(img *ebiten.Image, advance bool) {
	if s.index == s.end {
		s.Font.Draw(s.Text, img, s.point)
	} else {
		part := s.Text[:s.index+1]

		s.Font.Draw(part, img, s.point)

		if advance {
			s.index++
		}
	}
}
