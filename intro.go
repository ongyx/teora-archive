package teora

import (
	"fmt"
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

	IntroScene = &Intro{scroll: bento.NewEntity(sb)}
}

// Intro is the splash/startup screen.
type Intro struct {
	scroll *bento.Entity
}

func (i *Intro) Update(stage *bento.Stage) error {
	if err := i.scroll.Update(); err != nil {
		return err
	}

	if i.scroll.Sprite.(*Scrollbox).Done() && bento.Keypress(confirmKeys) {
		fmt.Println("intro: changing to start")
		stage.Change(StartScene)
	}

	return nil
}

func (i *Intro) Draw(screen *ebiten.Image) {
	i.scroll.Draw(screen)
}

func (i *Intro) Enter() bento.Animation {
	return nil
}

func (i *Intro) Exit() bento.Animation {
	// TODO: fix buggy scene changing when exit anim is nil
	//return nil

	return anim.NewFade(false, color.Black, 1)
}
