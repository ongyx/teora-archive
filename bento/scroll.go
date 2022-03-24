package bento

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

// Scroll allows several pieces of text to be scrolled across an image.
type Scroll struct {
	Font *Font
	Text []string

	index    int
	textpos  int
	textend  int
	textsize image.Point

	clock Clock
}

// NewScroll creates a new scroll at a point on an image.
func NewScroll(font *Font, tx []string) *Scroll {
	c := Clock{}
	c.Schedule(0.05)

	s := &Scroll{
		Font:  font,
		Text:  tx,
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
	s.textsize = size
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

// Size returns the total size of the scroll.
func (s *Scroll) Size() image.Point {
	return s.textsize
}

// Render renders the scroll on a new image.
func (s *Scroll) Render(clr color.Color) *ebiten.Image {
	p := s.Text[s.index]

	if s.textpos != s.textend {
		p = p[:s.textpos+1]

		if s.clock.Done() {
			s.textpos++
		}
	}

	s.clock.Tick()

	return s.Font.Draw(p, clr)
}
