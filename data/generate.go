//go:generate file2byteslice -package=fonts -input=./fonts/teoran.ttf -output=./fonts/teoran.go -var=TeoranStandard

package data

import (
	_ "github.com/hajimehoshi/file2byteslice"
)
