package guards

import (
	"testing"
	"github.com/fallofpheonix/phoenix-os/phoenixmind-security/policies"
)

func TestAllGuards(t *testing.T) {
	// Runtime Guard
	runtimeGuard := &RuntimeGuard{Engine: &policies.PolicyEngine{}}
	if res := runtimeGuard.CheckAction("cognition", "runtime_write"); res != policies.Block {
		t.Errorf("RuntimeGuard: Expected BLOCK for cognition, got %s", res)
	}

	// Merge Guard
	mergeGuard := &MergeGuard{}
	if res := mergeGuard.CanMerge("sandbox"); res != policies.Block {
		t.Errorf("MergeGuard: Expected BLOCK for sandbox, got %s", res)
	}

	// Phase Guard
	phaseGuard := &PhaseGuard{}
	if res := phaseGuard.CanTransition("F1", "proposal"); res != policies.Block {
		t.Errorf("PhaseGuard: Expected BLOCK for F1->proposal, got %s", res)
	}
}
