package marl

import (
	"sync"
	"time"
)

type StabilityController struct {
	mu           sync.Mutex
	actionDebt   map[string]int
	cooldowns    map[string]time.Time
	cooldownDur  time.Duration
}

func NewStabilityController(cooldown time.Duration) *StabilityController {
	return &StabilityController{
		actionDebt:  make(map[string]int),
		cooldowns:   make(map[string]time.Time),
		cooldownDur: cooldown,
	}
}

// CanAct returns true if the node is not in cooldown and has no excessive debt
func (c *StabilityController) CanAct(nodeID string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if lastAction, ok := c.cooldowns[nodeID]; ok {
		if time.Since(lastAction) < c.cooldownDur {
			return false
		}
	}

	if debt, ok := c.actionDebt[nodeID]; ok && debt > 5 {
		return false
	}

	return true
}

func (c *StabilityController) RegisterAction(nodeID string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cooldowns[nodeID] = time.Now()
	c.actionDebt[nodeID]++
}

func (c *StabilityController) ReduceDebt(nodeID string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.actionDebt[nodeID] > 0 {
		c.actionDebt[nodeID]--
	}
}
