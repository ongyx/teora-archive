package main

import (
	"log"
	"os"
	//"runtime"
	//"runtime/pprof"

	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/silbinarywolf/preferdiscretegpu"

	"github.com/ongyx/bento"
	"github.com/ongyx/teora"
	"github.com/ongyx/teora/assets"
)

const (
	logfile = "log.txt"
	//cpuprof = "cpu.prof"
	//memprof = "mem.prof"
)

func main() {
	// setup logs
	if f, err := os.Create(logfile); err != nil {
		log.Fatal("can't create log: ", err)
	} else {
		defer f.Close()

		log.SetFlags(log.Ltime | log.Lshortfile)
		log.SetOutput(f)
	}

	/*
		// setup pprof
		if f, err := os.Create(cpuprof); err != nil {
			log.Fatal("can't create cpu profile: ", err)
		} else {
			defer f.Close()

			if err := pprof.StartCPUProfile(f); err != nil {
				log.Fatal("can't start cpu profile: ", err)
			} else {
				defer pprof.StopCPUProfile()
			}
		}
	*/

	// setup stage
	var op bento.StageOptions

	op.HiDPI = true
	if teora.Debug {
		op.Font = assets.Hack
	}

	stage := bento.NewStage(teora.NewIntro())
	stage.Op = op

	// run stage
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Teora")

	if err := ebiten.RunGame(stage); err != nil {
		log.Fatal(err)
	}

	/*
		if f, err := os.Create(memprof); err != nil {
			log.Fatal("could not create memory profile: ", err)
		} else {
			defer f.Close()
			runtime.GC()
			if err := pprof.WriteHeapProfile(f); err != nil {
				log.Fatal("could not write memory profile: ", err)
			}
		}
	*/
}
