//go:generate file2byteslice -package=fonts -input=./fonts/teoran.ttf -output=./fonts/teoran.go -var=TeoranStandard
//go:generate file2byteslice -package=fonts -input=./fonts/hack.ttf -output=./fonts/hack.go -var=Hack

package data

import (
	// NOTE: currently all data is embedded into the final executable. Maybe store it somewhere else in the filesystem?
	_ "github.com/hajimehoshi/file2byteslice"
)
