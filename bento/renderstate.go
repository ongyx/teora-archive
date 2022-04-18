package bento

// RenderState represents the rendering state of an entity on the stage.
//
// The render lifecycle has several states:
// Hidden (default): The entity is not rendering to the screen.
// Entering: An enter transition is being rendered on the entity.
// Visible: The entity is rendering normally.
// Exiting: An exit transition is being rendered on the entity.
//
// After an entity has exited the stage, the render state will change back to Hidden.
//go:generate stringer -type=RenderState
type RenderState int

const (
	Hidden RenderState = iota
	Entering
	Visible
	Exiting
)
