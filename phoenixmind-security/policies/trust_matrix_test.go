package policies

import "testing"

func TestTrustMatrix(t *testing.T) {
	engine := &PolicyEngine{}
	
	// Test: training -> production_write -> BLOCK
	if res := engine.Check("training", "production_write"); res != Block {
		t.Errorf("Expected BLOCK for training writing to prod, got %s", res)
	}

	// Test: cognition -> runtime_access -> BLOCK
	if res := engine.Check("cognition", "runtime_access"); res != Block {
		t.Errorf("Expected BLOCK for cognition accessing runtime, got %s", res)
	}
}
