//go:build tools
// +build tools

package bento

// This file contains dependencies on tools used by go:generate.

import (
	_ "golang.org/x/tools/cmd/stringer"
)
