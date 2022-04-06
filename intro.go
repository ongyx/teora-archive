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
	var s bento.Stream
	s.Source(func(c chan<- string) {
		c <- "This project is neither affiliated with nor endorsed by GeoEXE."
		if Debug {
			c <- "Also, this is a devbuild. There may be a lot of bugs."
		}
	})

	sb := NewScrollbox(s, Hack)

	IntroScene = &Intro{
		scroll:  sb,
		entites: bento.NewEntitySlice(sb),
	}
}

// Intro is the splash/startup screen.
type Intro struct {
	scroll  *Scrollbox
	entites []*bento.Entity
}

func (i *Intro) Init(size image.Point) {}

func (i *Intro) Update(stage *bento.Stage) error {
	if i.scroll.Done() && bento.Keypress(confirmKeys) {
		stage.Change(StartScene)
	}

	return nil
}

func (i *Intro) Draw(screen *ebiten.Image) {}

func (i *Intro) Entities() []*bento.Entity {
	return i.entites
}

func (i *Intro) Enter() bento.Transition {
	return anim.NewFade(true, color.Black, 0.5)
}

func (i *Intro) Exit() bento.Transition {
	return anim.NewFade(false, color.Black, 0.5)
}
