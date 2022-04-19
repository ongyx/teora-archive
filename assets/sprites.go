package assets

import (
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ongyx/teora/bento"
)

var (
	Demo *bento.Tileset
)

func init() {
	if d, err := loadTileset("sprites/demo.png", 16); err != nil {
		panic(err)
	} else {
		Demo = d
	}
}

func loadTileset(path string, tilesize int) (*bento.Tileset, error) {
	if img, err := loadSprite(path); err != nil {
		return nil, err
	} else {
		return bento.NewTileset(img, tilesize), nil
	}
}

func loadSprite(path string) (*ebiten.Image, error) {
	f, err := assets.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if img, _, err := image.Decode(f); err != nil {
		return nil, err
	} else {
		return ebiten.NewImageFromImage(img), nil
	}
}
