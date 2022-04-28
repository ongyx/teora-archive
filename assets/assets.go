// assets handles loading of assets from the embedded filesystem.
package assets

import (
	"embed"
)

var (
	//go:embed fonts/*.ttf
	//go:embed shaders/*.go
	//go:embed sprites/*.png
	assets embed.FS
)
