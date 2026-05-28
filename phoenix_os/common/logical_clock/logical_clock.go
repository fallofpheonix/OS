package logical_clock

import "sync/atomic"

type Clock struct {
	tick uint64
}

func NewClock() *Clock {
	return &Clock{tick: 0}
}

func (c *Clock) Tick() uint64 {
	return atomic.AddUint64(&c.tick, 1)
}

func (c *Clock) Current() uint64 {
	return atomic.LoadUint64(&c.tick)
}
