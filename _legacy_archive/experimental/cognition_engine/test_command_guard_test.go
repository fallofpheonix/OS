package main

import (
	"phoenixmind/cognition_engine/sandbox"
	"testing"
)

func TestCommandGuard(t *testing.T) {
	guard := &sandbox.CommandGuard{}
	decision := guard.Allow("kernel_access")
	if decision != "REJECT" {
		t.Errorf("Expected REJECT, got %s", decision)
	}
}
