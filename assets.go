// assets handles loading of assets from the embedded filesystem.

package teora

import (
	"embed"
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/ongyx/teora/bento"
)

const (
	dpi = 72
)

var (
	//go:embed assets/fonts/teoran.ttf
	assets embed.FS

	// hack font
	// NOTE: the hack font here is a special case because we need to load it before everything else.
	//go:embed assets/fonts/hack.ttf
	hackData []byte

	hack   = &bento.Font{}
	teoran = &bento.Font{}
)

func init() {
	// hack
	try(hack.Load(hackData, &opentype.FaceOptions{
		Size:    24,
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
}

func try(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
