// assets handles loading of assets from the embedded filesystem.

package teora

import (
	"embed"
	"io/fs"
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/ongyx/teora/bento"
)

const (
	dpi = 72
)

var (
	//go:embed assets/fonts/*.ttf
	assets embed.FS
	fonts  fs.FS

	teoran = &bento.Font{}
	hack   = &bento.Font{}
)

func init() {
	var err error

	fonts, err = fs.Sub(assets, "assets/fonts")
	if err != nil {
		log.Fatal(err)
	}
}

func mustReadFile(fsys fs.FS, name string) []byte {
	data, err := fs.ReadFile(fsys, name)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func mustLoadFont(font *bento.Font, name string, o *opentype.FaceOptions) {
	data := mustReadFile(fonts, name)
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
