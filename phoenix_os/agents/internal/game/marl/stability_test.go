package marl

import (
	"testing"
	"time"
)

func TestStabilityController(t *testing.T) {
	cooldown := 100 * time.Millisecond
	c := NewStabilityController(cooldown)
	nodeID := "node1"

	// Should be able to act initially
	if !c.CanAct(nodeID) {
		t.Error("Expected CanAct to be true initially")
	}

	// Register action, should now be in cooldown
	c.RegisterAction(nodeID)
	if c.CanAct(nodeID) {
		t.Error("Expected CanAct to be false during cooldown")
	}

	// Wait for cooldown
	time.Sleep(cooldown + 10*time.Millisecond)
	if !c.CanAct(nodeID) {
		t.Error("Expected CanAct to be true after cooldown")
	}

	// Test Debt
	for i := 0; i < 6; i++ {
		c.RegisterAction(nodeID)
		time.Sleep(cooldown + 10*time.Millisecond) // Bypass cooldown to test debt
	}
	
	if c.CanAct(nodeID) {
		t.Error("Expected CanAct to be false due to high debt")
	}
}
