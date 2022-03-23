// Scroll provides text scrolling across a screen.

package teora

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

// Scroll allows several pieces of text to be scrolled across an image, starting from a point.
type Scroll struct {
	Font  *Font
	Text  []string
	Point image.Point
	Align Align

	index   int
	textpos int
	textend int
	// the adjusted point
	point image.Point

	clock Clock
}

// NewScroll creates a new scroll at a point on an image.
func NewScroll(font *Font, tx []string, point image.Point, align Align) *Scroll {
	c := Clock{}
	c.Schedule(0.05)

	s := &Scroll{
		Font:  font,
		Text:  tx,
		Point: point,
		Align: align,
		clock: c,
	}
	s.SetIndex(0)

	return s
}

// Speed changes the speed of scrolling text,
// where secs is the number of seconds to wait between scrolling each character and must be positive.
func (s *Scroll) Speed(secs float64) {
	s.clock.Schedule(secs)
}

// Index gets the current index of the scroll.
func (s *Scroll) Index() int {
	return s.index
}

// SetIndex changes the piece of text currently scrolling.
// index must be a valid index in the scroll's list of text.
func (s *Scroll) SetIndex(index int) {
	t := s.Text[index]
	size := text.BoundString(s.Font.Face, t).Size()

	s.index = index
	s.textpos = 0
	s.textend = len(t)
	s.point = s.Align.Adjust(s.Point, size)
}

// Next changes to the next piece of text to scroll.
func (s *Scroll) Next() {
	if s.index != (len(s.Text) - 1) {
		s.SetIndex(s.index + 1)
	}
}

// Prev changes to the previous piece of text to scroll.
func (s *Scroll) Prev() {
	if s.index != 0 {
		s.SetIndex(s.index - 1)
	}
}

// Skip causes the next render to render the whole text instead of waiting for scrolling.
func (s *Scroll) Skip() {
	s.textpos = s.textend
}

// Done checks if the scrolling has finished.
func (s *Scroll) Done() bool {
	return s.textpos == s.textend
}

// Render renders the scroll onto an image and returns its current bounds.
func (s *Scroll) Render(img *ebiten.Image) image.Rectangle {
	p := s.Text[s.index]

	if s.textpos != s.textend {
		p = p[:s.textpos+1]

		if s.clock.Done() {
			s.textpos++
		}
	}

	s.clock.Tick()

	return s.Font.Draw(p, img, s.point)
}
