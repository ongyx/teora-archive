package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/silbinarywolf/preferdiscretegpu"

	"github.com/ongyx/teora"
	"github.com/ongyx/teora/bento"
)

func main() {
	stage := bento.NewStage(teora.IntroScene)

	if teora.Debug {
		stage.Debug = &bento.DebugOptions{Font: teora.Hack}
	}

	// ebiten.SetWindowSize(256, 256)
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Teora")

	if err := ebiten.RunGame(stage); err != nil {
		log.Fatal(err)
	}
}
