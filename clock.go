package teora

import (
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

// Schedule sets the clock's timer to trigger every n seconds.
func (c *Clock) Schedule(n float64) {
	c.timer = int(n * tps)
	c.counter = 0
}

// ScheduleOnce sets the clock's timer to trigger once, after n seconds.
func (c *Clock) ScheduleOnce(n float64) {
	c.Schedule(n)
	c.once = true
}

// Done checks if the clock's timer has triggered on the current tick.
// If the timer was set with ScheduleOnce, this will only return true one time.
func (c *Clock) Done() bool {
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
