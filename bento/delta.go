package bento

import (
	"image"
	"math"

	"gonum.org/v1/gonum/floats"
)

const (
	// Linear specifies a delta is constant.
	Linear DeltaAlgorithm = iota
	// Exponential specifies a delta in exponential (e^x) space.
	Exponential
	// Log specifies a delta in log (ln x) space.
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
	case Log:
		algo = expSpan
	}

	algo(dx, 1, x)
	algo(dy, 1, y)

	return &Delta{
		delta: delta,
		clock: c,
		limit: l,
		dx:    dx,
		dy:    dy,
	}
}

// Update updates the delta.
// NOTE: If you're using this in a sprite/transition, only call this _after_ any call of d.Delta()!
func (d *Delta) Update() {
	if d.clock.Done() {
		d.idx = d.limit
	}

	if d.idx < (d.limit - 1) {
		d.idx++
	}

	d.clock.Tick()
}

// Delta returns the current delta.
func (d *Delta) Delta() image.Point {
	var x, y float64

	// special case: if x/y delta is 0, return 0 here too
	// otherwise it will return NaN
	if d.delta.X != 0 {
		x = d.dx[d.idx] - 1
	}

	if d.delta.Y != 0 {
		y = d.dy[d.idx] - 1
	}

	return image.Pt(int(x), int(y))
}

// Done checks if the current delta is equal to the total delta.
func (d *Delta) Done() bool {
	return d.idx == d.limit
}

func expSpan(dst []float64, l, u float64) []float64 {
	floats.Span(dst, l, u)

	for i := range dst {
		dst[i] = math.Log(dst[i])
	}

	// NOTE: assuming u > l > 0!
	floats.Scale(u/floats.Max(dst), dst)

	return dst
}
