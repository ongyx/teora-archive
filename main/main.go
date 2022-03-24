package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/silbinarywolf/preferdiscretegpu"

	"github.com/ongyx/teora"
)

func main() {
	// ebiten.SetWindowSize(256, 256)
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Teora")

	if err := ebiten.RunGame(teora.Stage); err != nil {
		log.Fatal(err)
	}
}
