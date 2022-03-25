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

	text  string
	tpos  int
	tend  int
	tsize image.Point

	clock Clock
}

// NewScroll creates a new scroll.
func NewScroll(font *Font, tx string) *Scroll {
	c := Clock{}
	c.Schedule(0.05)

	s := &Scroll{
		Font:  font,
		clock: c,
	}
	s.SetText(tx)

	return s
}

// Speed changes the speed of scrolling text,
// where secs is the number of seconds to wait between scrolling each character and must be positive.
func (s *Scroll) Speed(secs float64) {
	s.clock.Schedule(secs)
}

// Text returns the current text in the scroll.
func (s *Scroll) Text() string {
	return s.text
}

// SetText changes the text currently scrolling.
func (s *Scroll) SetText(tx string) {
	size := text.BoundString(s.Font.Face, tx).Size()

	s.text = tx
	s.tpos = 0
	s.tend = len(tx)
	s.tsize = size
}

// Skip causes the next render to render the whole text instead of waiting for scrolling.
func (s *Scroll) Skip() {
	s.tpos = s.tend
}

// Done checks if the scrolling has finished.
func (s *Scroll) Done() bool {
	return s.tpos == s.tend
}

// Size returns the total size of the scroll.
func (s *Scroll) Size() image.Point {
	return s.tsize
}

// Render renders the scroll on a new image.
func (s *Scroll) Render(
	clr color.Color,
	point image.Point,
	img *ebiten.Image,
) {
	t := s.text

	if s.tpos != s.tend {
		t = t[:s.tpos+1]

		if s.clock.Done() {
			s.tpos++
		}
	}

	s.clock.Tick()

	s.Font.Write(t, clr, img, point, Default)
}
