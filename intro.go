package teora

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ongyx/teora/bento"
	"github.com/ongyx/teora/bento/anim"
)

var (
	IntroScene bento.Scene
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
			Hack,
		),
	}
}

// Intro is the splash/startup screen.
type Intro struct {
	scroll *Scrollbox
}

func (i *Intro) Update(stage *bento.Stage) error {
	i.scroll.Update()

	if i.scroll.Done() && bento.Keypress(confirmKeys) {
		stage.Change(StartScene)
	}

	return nil
}

func (i *Intro) Render(screen *ebiten.Image) {
	b := screen.Bounds()
	p := image.Point{
		X: bento.Center.Point(b).X,
		Y: int(float64(b.Dy()) * 0.9),
	}

	i.scroll.Render(p, screen)
}

func (i *Intro) Enter() bento.Transition {
	return anim.NewFade(true, color.Black, 1)
}

func (i *Intro) Exit() bento.Transition {
	return anim.NewFade(false, color.Black, 1)
}
