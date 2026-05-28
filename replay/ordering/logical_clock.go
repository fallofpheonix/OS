package ordering

type LogicalClock interface {
	Tick() uint64
	Now() uint64
	AdvanceTo(v uint64)
}

type MonotonicClock struct {
	counter uint64
}

func (c *MonotonicClock) Tick() uint64 {
	c.counter++
	return c.counter
}

func (c *MonotonicClock) Now() uint64 {
	return c.counter
}

func (c *MonotonicClock) AdvanceTo(v uint64) {
	if v > c.counter {
		c.counter = v
	}
}
