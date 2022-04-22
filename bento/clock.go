package bento

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

const tps = ebiten.DefaultTPS

// Clock acts as a basic scheduler for operations.
// One tick is equivalent to a single frame, where 1 second is 60 frames.
type Clock struct {
	tick    int
	timer   int
	counter int

	once bool
}

// NewClock creates a new clock that triggers every n seconds.
func NewClock(n float64) *Clock {
	c := newClock()
	c.Schedule(n)

	return c
}

// NewClockOnce creates a new clock that triggers once after n seconds.
func NewClockOnce(n float64) *Clock {
	c := newClock()
	c.ScheduleOnce(n)

	return c
}

func newClock() *Clock {
	return &Clock{tick: -1}
}

// Ticks returns the current ticks of the clock.
func (c *Clock) Ticks() int {
	return c.tick
}

// Limit returns the clock's timer in ticks.
func (c *Clock) Limit() int {
	return c.timer
}

// Schedule sets the clock's timer to trigger every n seconds.
// This will panic if n is too small to calculate ticks for.
func (c *Clock) Schedule(n float64) {
	c.timer = SecondToTick(n)
	c.counter = 0

	if c.timer == 0 {
		panic(fmt.Sprintf("clock: duration of n too small (%f)", n))
	}
}

// ScheduleOnce sets the clock's timer to trigger once, after n seconds.
// This will panic if n is too small to calculate ticks for.
func (c *Clock) ScheduleOnce(n float64) {
	c.Schedule(n)
	c.once = true
}

// Done checks if the clock's timer has triggered on the current tick.
// If the timer was set with ScheduleOnce, this will only return true one time.
func (c *Clock) Done() bool {
	switch c.tick {
	case -1:
		panic(&InitError{"clock", "invalid tick"})
	case 0:
		// NOTE: clock should not trigger when tick == 0!
		return false
	}

	trigger := (c.tick % c.timer) == 0

	if trigger {
		if c.once {
			// timer has triggered, so make sure this is only the first trigger.
			return c.counter == 0
		}

		c.counter++
	}

	return trigger
}

// Tick increments the clock's tick count.
// This must be called once every tick in ebiten's game loop.
func (c *Clock) Tick() {
	c.tick++
}

// SecondToTick converts seconds to ticks.
func SecondToTick(seconds float64) int {
	return int(seconds * tps)
}

// TickToSecond converts ticks to seconds.
func TickToSecond(ticks int) float64 {
	return float64(ticks) / tps
}
