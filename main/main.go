package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/silbinarywolf/preferdiscretegpu"

	"github.com/ongyx/bento"
	"github.com/ongyx/teora"
	"github.com/ongyx/teora/assets"
)

const (
	logfile = "log.txt"
)

func main() {
	f, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(fmt.Sprintf("main: can't open logfile: %v", err))
	}
	defer f.Close()

	log.SetFlags(log.Ltime | log.Lshortfile)
	log.SetOutput(f)

	var dbg *bento.Debug
	if teora.Debug {
		dbg = &bento.Debug{Font: assets.Hack}
	}

	stage := bento.NewStage(teora.NewIntro(), dbg)

	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Teora")

	if err := ebiten.RunGame(stage); err != nil {
		log.Fatal(err)
	}
}
