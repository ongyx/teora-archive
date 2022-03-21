//go:generate file2byteslice -package=fonts -input=./fonts/teoran.ttf -output=./fonts/teoran.go -var=TeoranStandard

package data

import (
	// NOTE: currently all data is embedded into the final executable. Maybe store it somewhere else in the filesystem?
	_ "github.com/hajimehoshi/file2byteslice"
)
