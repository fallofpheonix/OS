package security

import (
	"testing"
)

func TestSolver(t *testing.T) {
	m := NewDefaultMatrix()

	// High probability of attack -> Should Defend
	move1, _ := Solve(m, 0.9)
	if move1 != MoveDefend {
		t.Errorf("Expected DEFEND for high attack prob, got %s", move1)
	}

	// Low probability of attack -> Should Monitor
	move2, _ := Solve(m, 0.1)
	if move2 != MoveMonitor {
		t.Errorf("Expected MONITOR for low attack prob, got %s", move2)
	}
}
