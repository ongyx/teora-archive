package assets

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	Gradient *ebiten.Shader
)

func init() {
	if g, err := loadShader("shaders/gradient.go"); err != nil {
		panic(err)
	} else {
		Gradient = g
	}
}

func loadShader(path string) (*ebiten.Shader, error) {
	if c, err := assets.ReadFile(path); err != nil {
		return nil, err
	} else {
		if s, err := ebiten.NewShader(c); err != nil {
			return nil, err
		} else {
			return s, nil
		}
	}
}
