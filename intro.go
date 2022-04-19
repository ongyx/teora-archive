package teora

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ongyx/teora/assets"
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

	IntroScene = &Intro{scroll: NewScrollbox(s, assets.Hack)}
}

// Intro is the splash/startup screen.
type Intro struct {
	scroll *Scrollbox
}

func (i *Intro) Update(stage *bento.Stage) error {
	if i.scroll.Done() && bento.Keypress(confirmKeys) {
		stage.Change(StartScene)
	}

	return nil
}

func (i *Intro) Draw(screen *ebiten.Image) {}

func (i *Intro) Enter() bento.Animation {
	return anim.NewFade(true, color.Black, 0.5)
}

func (i *Intro) Exit() bento.Animation {
	return anim.NewFade(false, color.Black, 0.5)
}

func (i *Intro) Entities() []*bento.Entity {
	return bento.NewEntities(i.scroll)
}
