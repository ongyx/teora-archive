package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ongyx/teora"
)

func main() {
	game := teora.NewGame()

	ebiten.SetWindowSize(256, 256)
	ebiten.SetWindowTitle("Teora")

	if err := ebiten.RunGame(game); err != nil {
		fmt.Errorf(err.Error())
	}
}
