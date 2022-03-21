package data

import (
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/ongyx/teora/data/fonts"
)

const (
	fontSize = 24
	dpi      = 72
)

var (
	// TeoranStandard is a special fontface that is used occasionally in teora.
	TeoranStandard font.Face
)

func init() {
	tt, err := opentype.Parse(fonts.TeoranStandard)
	if err != nil {
		log.Fatal(err)
	}

	TeoranStandard, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	if err != nil {
		log.Fatal(err)
	}
}
