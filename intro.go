package teora

import (
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
	var s bento.Stream
	s.Source(func(c chan<- string) {
		c <- "This project is neither affiliated with nor endorsed by GeoEXE."
		if Debug {
			c <- "Also, this is a devbuild. There may be a lot of bugs."
		}
	})

	IntroScene = &Intro{
		scroll: NewScrollbox(s, Hack),
	}
}

type msg struct {
	text string
}

func (m msg) String() string {
	return m.text
}

// Intro is the splash/startup screen.
type Intro struct {
	scroll *Scrollbox
}

func (i *Intro) Init(screen *ebiten.Image) {
	i.scroll.Init(screen)
}

func (i *Intro) Update(stage *bento.Stage) error {
	i.scroll.Update()

	if i.scroll.Done() && bento.Keypress(confirmKeys) {
		stage.Change(StartScene)
	}

	return nil
}

func (i *Intro) Draw(screen *ebiten.Image) {
	i.scroll.Draw(screen)
}

func (i *Intro) Enter() bento.Transition {
	return anim.NewFade(true, color.Black, 1)
}

func (i *Intro) Exit() bento.Transition {
	return anim.NewFade(false, color.Black, 1)
}
