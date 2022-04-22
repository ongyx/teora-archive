package bento

// Component is a state that updates every tick.
// Update must be called before using the component's state, i.e
//
// type myComponent struct {
// 	sub Component
// }
//
// func (mc *myComponent) Update() error {
// 	if err := mc.sub.Update(); err != nil {
// 		return err
// 	}
//
//	// use mc.sub here
// }
//
// or else accessing the component's state may result in undefined behaviour/panics.
type Component interface {
	Update() error
}
