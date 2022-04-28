package assets

import (
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/ongyx/bento"
)

var (
	dpi = bento.DPIScale(96)

	Hack   = &bento.Font{}
	Teoran = &bento.Font{}
)

func init() {
	if err := loadFont(
		"fonts/hack.ttf",
		Hack,
		&opentype.FaceOptions{
			Size:    12,
			DPI:     dpi,
			Hinting: font.HintingFull,
		},
	); err != nil {
		panic(err)
	}

	if err := loadFont(
		"fonts/teoran.ttf",
		Teoran,
		&opentype.FaceOptions{
			Size:    24,
			DPI:     dpi,
			Hinting: font.HintingFull,
		},
	); err != nil {
		panic(err)
	}
}

func loadFont(path string, font *bento.Font, o *opentype.FaceOptions) error {
	if d, err := assets.ReadFile(path); err != nil {
		return err
	} else {
		if err := font.Load(d, o); err != nil {
			return err
		}
	}

	return nil
}
