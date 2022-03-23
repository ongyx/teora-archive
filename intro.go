package teora

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/ongyx/teora/bento"
)

var msgs = []string{
	"This is a demo build.",
	"project teora is neither affiliated with nor endorsed by GeoEXE.",
}

func init() {
	Game.Scenes["intro"] = &Intro{}
	Game.SetScene("intro")
}

// Intro is the splash/startup screen.
type Intro struct {
	scroll *bento.Scroll
}

// Update updates the scroll.
func (i *Intro) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		// skip the text if it's still scrolling, otherwise go to the next text.
		if !i.scroll.Done() {
			i.scroll.Skip()
		} else {
			i.scroll.Next()
		}
	}

	return nil
}

// Render renders the intro sequence to the screen.
func (i *Intro) Render(screen *ebiten.Image) {
	// draw tps/fps at the top left of the screen
	hack.Write(
		fmt.Sprintf("tps: %0.2f", ebiten.CurrentTPS()),
		screen,
		image.Pt(0, 0),
		bento.AlignRight|bento.AlignBottom,
	)

	bounds := teoran.WriteCenter("Hello World!", screen)

	// we can only init scroll here because we need to know the screen size
	if i.scroll == nil {
		i.scroll = bento.NewScroll(
			hack,
			msgs,
			image.Pt(bento.Center(screen).X, bounds.Max.Y+20),
			bento.AlignCenter,
		)
	}

	i.scroll.Render(screen)
}
