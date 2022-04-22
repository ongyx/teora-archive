package bento

import (
	"image"

	"gonum.org/v1/gonum/floats"
)

const (
	// Linear specifies a delta is constant.
	Linear DeltaAlgorithm = iota
	// Exponential specifies a delta in exponential (e^x) space.
	Exponential
)

// DeltaAlgorithm specifies the algorithm to use when generating deltas.
type DeltaAlgorithm int

// Delta is a delta that changes over time.
type Delta struct {
	delta image.Point
	clock *Clock

	index, limit int

	dx, dy []float64
}

// NewDelta creates an delta with the total delta, and the period over which to increase the current delta.
func NewDelta(
	dalgo DeltaAlgorithm,
	delta image.Point,
	period float64,
) *Delta {
	c := NewClockOnce(period)

	l := c.Limit()
	// TODO(ongyx): a more efficient way to store the delta per tick?
	dx := make([]float64, l)
	dy := make([]float64, l)

	// some algorithms require the start to be at least 1, so add 1 to the delta here.
	delta.X += 1
	delta.Y += 1

	x := float64(delta.X)
	y := float64(delta.Y)

	var algo func([]float64, float64, float64) []float64

	switch dalgo {
	case Linear:
		algo = floats.Span
	case Exponential:
		algo = floats.LogSpan
	}

	algo(dx, 1, x)
	algo(dy, 1, y)

	return &Delta{
		delta: delta,
		clock: c,
		index: -1,
		limit: l - 1,
		dx:    dx,
		dy:    dy,
	}
}

// Update updates the delta.
func (d *Delta) Update() {
	d.clock.Tick()

	if d.clock.Done() {
		d.index = d.limit
	}

	if d.index < d.limit {
		d.index++
	}
}

// Delta returns the current delta.
// This will panic if delta has not been updated yet.
func (d *Delta) Delta() (x, y float64) {
	if d.index == -1 {
		panic(&InitError{"delta", "invalid index"})
	}

	// special case: if x/y delta is 0, return 0 here too
	// otherwise it will return NaN
	if d.delta.X != 0 {
		x = d.dx[d.index] - 1
	}

	if d.delta.Y != 0 {
		y = d.dy[d.index] - 1
	}

	return x, y
}

func (d *Delta) DeltaPt() image.Point {
	x, y := d.Delta()
	return image.Pt(int(x), int(y))
}

// Done checks if the current delta is equal to the total delta.
func (d *Delta) Done() bool {
	return d.index == d.limit
}
