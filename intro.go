package teora

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ongyx/teora/bento"
)

var (
	IntroScene *Intro
	IntroMsg   []string
)

func init() {
	IntroMsg = []string{"This project is neither affiliated with nor endorsed by GeoEXE."}

	if Debug {
		IntroMsg = append(IntroMsg, "Also, this is a devbuild. There may be a lot of bugs.")
	}

	IntroScene = &Intro{
		scroll: NewScrollbox(
			IntroMsg,
			hack,
		),
	}
}

// Intro is the splash/startup screen.
type Intro struct {
	scroll *Scrollbox
}

// Update updates the scroll.
func (i *Intro) Update(stage *bento.Stage) error {
	i.scroll.Update()

	return nil
}

// Render renders the intro sequence to the screen.
func (i *Intro) Render(screen *ebiten.Image) {
	// draw tps/fps at the top left of the screen
	hack.Write(
		fmt.Sprintf("tps: %0.2f", ebiten.CurrentTPS()),
		color.White,
		screen,
		image.Pt(0, 0),
		bento.Default,
	)

	teoran.WriteCenter("Hello World!", color.White, screen)

	b := screen.Bounds()
	p := image.Point{
		X: bento.Center.Point(b).X,
		Y: int(float64(b.Dy()) * 0.9),
	}

	i.scroll.Render(p, screen)
}
