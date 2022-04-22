package bento

import (
	"fmt"
)

// InitError indicates that the component has not been updated yet with .Update.
type InitError struct {
	name, msg string
}

func (ie *InitError) Error() string {
	return fmt.Sprintf("%s: %s (maybe you forgot to call .Update beforehand?)", ie.name, ie.msg)
}
