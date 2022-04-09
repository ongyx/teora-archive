package bento

import (
	"image"

	"gonum.org/v1/gonum/floats"
)

const (
	// Linear specifies a delta is constant (a linear graph).
	Linear DeltaAlgorithm = iota
	// Log specifies a delta over log space (a log graph).
	Log
)

// DeltaAlgorithm specifies the algorithm to use when generating deltas.
type DeltaAlgorithm int

// Delta is a delta that changes over time.
type Delta struct {
	delta image.Point
	clock *Clock

	idx, limit int

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
	// TODO: is there a more efficient way to store the delta per tick?
	dx := make([]float64, l)
	dy := make([]float64, l)

	x := float64(delta.X)
	y := float64(delta.Y)

	switch dalgo {
	case Linear:
		floats.Span(dx, 1, x)
		floats.Span(dy, 1, y)
	case Log:
		floats.LogSpan(dx, 1, x)
		floats.LogSpan(dy, 1, y)
	}

	return &Delta{
		delta: delta,
		clock: c,
		limit: l - 1,
		dx:    dx,
		dy:    dy,
	}
}

// Update updates the delta.
func (d *Delta) Update() {
	if d.clock.Done() {
		d.idx = d.limit
	} else if d.idx < d.limit {
		d.idx++
	}

	d.clock.Tick()
}

// Delta returns the current delta.
func (d *Delta) Delta() image.Point {
	x := d.dx[d.idx]
	y := d.dy[d.idx]
	return image.Pt(int(x), int(y))
}

// Done checks if the current delta is equal to the total delta.
func (d *Delta) Done() bool {
	return d.idx == d.limit
}
