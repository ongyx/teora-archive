package teora

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/ongyx/teora/bento"
)

var msgs = []string{
	"This is a demo build.",
	"project teora is neither affiliated with nor endorsed by GeoEXE.",
}

var introScene = &Intro{}

// Intro is the splash/startup screen.
type Intro struct {
	scroll *bento.Scroll // Scrollbox
}

// Update updates the scroll.
func (i *Intro) Update() (bento.Scene, error) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		// skip the text if it's still scrolling, otherwise go to the next text.
		if !i.scroll.Done() {
			i.scroll.Skip()
		} else {
			i.scroll.Next()
		}
	}

	return nil, nil
}

// Render renders the intro sequence to the screen.
func (i *Intro) Render(screen *ebiten.Image) {
	// draw tps/fps at the top left of the screen
	hack.Write(
		fmt.Sprintf("tps: %0.2f", ebiten.CurrentTPS()),
		color.White,
		screen,
		image.Pt(0, 0),
		bento.AlignRight|bento.AlignBottom,
	)

	bounds := teoran.WriteCenter("Hello World!", color.White, screen)

	// we can only init scroll here because we need to know the screen size
	if i.scroll == nil {
		/*
			i.scroll = &Scrollbox{
				Scroll: bento.NewScroll(hack, msgs),
			}
		*/
		i.scroll = bento.NewScroll(hack, msgs)
	}

	render := i.scroll.Render(color.White)

	point := bento.AlignCenter.Adjust(
		image.Pt(bento.Center(screen).X, bounds.Max.Y+40),
		image.Pt(render.Size()),
	)

	// TODO: find a nicer way to set draw options?
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(point.X), float64(point.Y))

	screen.DrawImage(render, op)
}
