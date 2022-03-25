// assets handles loading of assets from the embedded filesystem.

package teora

import (
	"embed"
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ongyx/teora/bento"
)

const (
	dpi = 72
)

var (
	//go:embed assets/fonts/teoran.ttf
	//go:embed assets/shaders/*.go
	assets embed.FS

	// hack font
	// NOTE: the hack font here is a special case because we need to load it before everything else.
	//go:embed assets/fonts/hack.ttf
	hackData []byte

	Hack   = &bento.Font{}
	teoran = &bento.Font{}

	gradient *ebiten.Shader
)

func init() {
	// hack
	try(Hack.Load(hackData, &opentype.FaceOptions{
		Size:    20,
		DPI:     dpi,
		Hinting: font.HintingFull,
	}))

	teoranData, err := assets.ReadFile("assets/fonts/teoran.ttf")
	try(err)

	try(teoran.Load(teoranData, &opentype.FaceOptions{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull,
	}))

	gradientCode, err := assets.ReadFile("assets/shaders/gradient.go")
	try(err)

	gradient, err = ebiten.NewShader(gradientCode)
	try(err)
}

func try(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
