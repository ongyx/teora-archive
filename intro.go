package teora

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ongyx/bento"
	"github.com/ongyx/bento/anim"
	"github.com/ongyx/teora/assets"
)

// Intro is the splash/startup screen.
type Intro struct {
	scroll *Scrollbox
}

func NewIntro() bento.Scene {
	var s = bento.NewStream[string]()

	go func() {
		s.Write("This project is neither affiliated with nor endorsed by GeoEXE.")
		if Debug {
			s.Write("Also, this is a devbuild. There may be a lot of bugs.")
		}
		s.Close()
	}()

	return &Intro{scroll: NewScrollbox(s, assets.Hack)}
}

func (i *Intro) Update(stage *bento.Stage) error {
	if i.scroll.Done() && bento.Keypress(confirmKeys) {
		stage.Change(NewStart())
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
