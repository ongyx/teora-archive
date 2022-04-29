package main

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/silbinarywolf/preferdiscretegpu"

	"github.com/ongyx/bento"
	"github.com/ongyx/teora/assets"
)

const (
	logfile = "log.txt"
)

func main() {
	// start pprof
	go pprof()

	// setup logs
	if f, err := os.Create(logfile); err != nil {
		log.Fatal("can't create log: ", err)
	} else {
		defer f.Close()

		log.SetFlags(log.Ltime | log.Lshortfile)
		log.SetOutput(f)
	}

	// setup stage
	var op bento.StageOptions

	op.HiDPI = true
	if Debug {
		op.Font = assets.Hack
	}

	stage := bento.NewStage(NewIntro())
	stage.Op = op

	// run stage
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Teora")

	if err := ebiten.RunGame(stage); err != nil {
		log.Fatal(err)
	}
}
