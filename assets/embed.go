package assets

import (
	"embed"
)

// Assets contains all game assets (fonts, sprites, etc.) needed by teora.
//go:embed fonts/*.ttf
var Assets embed.FS
