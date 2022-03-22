// assets handles loading of assets from the embedded filesystem.

package teora

import (
	"embed"
	"image/color"
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/ongyx/teora/assets"
)

const (
	dpi = 72
)

var (
	teoran = &Font{Color: color.White}
	hack   = &Font{Color: color.White}
)

func mustReadFile(fs embed.FS, name string) []byte {
	data, err := fs.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func mustLoadFont(font *Font, name string, o *opentype.FaceOptions) {
	data := mustReadFile(assets.Assets, name)
	if err := font.Load(data, o); err != nil {
		log.Fatal(err)
	}
}

func init() {
	mustLoadFont(
		teoran,
		"teoran.ttf",
		&opentype.FaceOptions{
			Size:    48,
			DPI:     dpi,
			Hinting: font.HintingFull,
		},
	)

	mustLoadFont(
		hack, "hack.ttf",
		&opentype.FaceOptions{
			Size:    24,
			DPI:     dpi,
			Hinting: font.HintingFull,
		},
	)
}
