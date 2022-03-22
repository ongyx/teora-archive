// data handles loading of assets from either embedded byte slices or the filesystem.

package teora

import (
	"image/color"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/ongyx/teora/data/fonts"
)

const (
	dpi = 72
)

var (
	teoran = &Font{color: color.White}
	hack   = &Font{color: color.White}
)

func init() {
	teoran.Load(
		fonts.TeoranStandard,
		&opentype.FaceOptions{
			Size:    48,
			DPI:     dpi,
			Hinting: font.HintingFull,
		},
	)

	hack.Load(
		fonts.Hack,
		&opentype.FaceOptions{
			Size:    24,
			DPI:     dpi,
			Hinting: font.HintingFull,
		},
	)
}
