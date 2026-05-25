package control

import (
	"github.com/fallofpheonix/phoenix-os/phoenix_os/agents/internal/types"
	"testing"
	"time"
)

func TestControlAgent(t *testing.T) {
	agent := NewControlAgent(1.0, 0.1, 0.05, 2.0)
	if agent == nil {
		t.Fatal("failed to create agent")
	}

	strategy := types.Strategy{
		ContainmentLevel: 1,
		TargetPIDs:       []uint32{123},
	}

	now := time.Now()
	// First update
	err := agent.EnforceStrategy(strategy, 5.0, now) // 5.0 is above 2.0 setpoint
	if err != nil {
		t.Fatalf("failed to enforce strategy: %v", err)
	}

	metrics := agent.GetPIDMetrics()
	if metrics.Output <= 0 {
		t.Errorf("expected positive PID output for error, got %f", metrics.Output)
	}

	history := agent.GetActionHistory()
	if len(history) == 0 {
		t.Error("action history should not be empty")
	}
}
