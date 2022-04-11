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

type sbAnim struct {
	rect image.Rectangle

	delta *bento.Delta

	vec  *bento.Vec
	mask *ebiten.Image
}

func (sba *sbAnim) Update() error {
	return nil
}

func (sba *sbAnim) Draw(img *ebiten.Image) {
	if sba.delta == nil {
		b := img.Bounds()
		d := b.Dy()
		r := d / 2
		a := b.Dx() - d

		sba.rect = image.Rect(r, 0, r, d)

		sba.delta = bento.NewDelta(bento.Exponential, image.Pt(a, 0), 0.5)

		sba.mask = ebiten.NewImage(img.Size())
	}

	rect := image.Rectangle{
		Min: sba.rect.Min,
		Max: sba.rect.Max.Add(sba.delta.Delta()),
	}

	// Draw a stadium on the canvas
	// https://en.wikipedia.org/wiki/Stadium_(geometry)
	sba.vec = stadium(rect)

	sba.mask.Fill(color.Alpha{0})

	sba.vec.Draw(color.Alpha{255}, sba.mask, nil)

	mask(img, sba.mask, nil)

	sba.delta.Update()
}

func (sba *sbAnim) Done() bool {
	return sba.delta.Done()
}

// Scrollbox is a box with scrolling text inside.
type Scrollbox struct {
	scroll *bento.Scroll
	stream bento.Stream

	anim      *sbAnim
	canvas    *ebiten.Image
	scrollpos image.Point

	done bool
}

// NewScrollbox creates a new scrollbox from a stream.
func NewScrollbox(stream bento.Stream, font *bento.Font) *Scrollbox {
	s := bento.NewScroll(font, *stream.Read())

	return &Scrollbox{
		scroll: s,
		stream: stream,
		anim:   &sbAnim{},
	}
}

func (sb *Scrollbox) Update() error {
	sb.scroll.Update()

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

func (sb *Scrollbox) Render(entity *bento.Entity, size image.Point) *ebiten.Image {
	if sb.canvas == nil {
		sb.canvas = ebiten.NewImage(int(float64(size.X)/1.5), size.Y/12)

		r := sb.canvas.Bounds().Dy() / 2

		sb.scrollpos = image.Pt(r, r)
		sb.scrollpos.Y += sb.scroll.Size().Y / 2

		p := image.Pt(size.X/2, int(float64(size.Y)*0.9))

		entity.Position = bento.Center.Align(p, sb.canvas.Bounds().Size())
		entity.Show(sb.anim)
	}

	sb.canvas.Fill(sbColor)

	// render scroll text
	sb.scroll.Draw(color.Black, sb.scrollpos, sb.canvas)

	if entity.RenderState() == bento.Normal {
		mask(sb.canvas, sb.anim.mask, nil)
	}

	return sb.canvas
}

// Done checks if the scrollbox has finished scrolling all text.
func (sb *Scrollbox) Done() bool {
	return sb.scroll.Done() && sb.done
}

func stadium(rect image.Rectangle) *bento.Vec {
	var vec bento.Vec

	r := rect.Dy() / 2
	rp := bento.CenterRight.Point(rect)
	lp := bento.CenterLeft.Point(rect)

	// rectangle
	vec.Rect(rect)

	// semicircles
	vec.Arc(rp, r, bento.Deg90, bento.Deg270)
	vec.Arc(lp, r, bento.Deg270, bento.Deg90)

	return &vec
}

func mask(dest, source *ebiten.Image, op *ebiten.DrawImageOptions) {
	if op == nil {
		op = &ebiten.DrawImageOptions{}
	}

	op.CompositeMode = ebiten.CompositeModeDestinationIn
	dest.DrawImage(source, op)
}
