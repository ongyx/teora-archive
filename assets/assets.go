// assets handles loading of assets from the embedded filesystem.
package assets

import (
	"embed"
)

const (
	dpi = 72
)

var (
	//go:embed fonts/*.ttf
	//go:embed shaders/*.go
	//go:embed sprites/*.png
	assets embed.FS
)
