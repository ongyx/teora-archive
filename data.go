// data handles loading of assets from either embedded byte slices or the filesystem.

package teora

import (
	"image/color"
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/ongyx/teora/data/fonts"
)

const (
	dpi = 72
)

var (
	teoran = &Font{Color: color.White}
	hack   = &Font{Color: color.White}
)

func init() {
	if err := teoran.Load(
		fonts.TeoranStandard,
		&opentype.FaceOptions{
			Size:    48,
			DPI:     dpi,
			Hinting: font.HintingFull,
		},
	); err != nil {
		log.Fatal(err)
	}

	if err := hack.Load(
		fonts.Hack,
		&opentype.FaceOptions{
			Size:    24,
			DPI:     dpi,
			Hinting: font.HintingFull,
		},
	); err != nil {
		log.Fatal(err)
	}
}
