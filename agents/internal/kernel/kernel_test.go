package kernel

import (
	"testing"
)

func TestKernelAgent(t *testing.T) {
	agent := NewKernelAgent()
	if agent == nil {
		t.Fatal("failed to create agent")
	}

	if !agent.IsLocked() {
		t.Error("expected kernel agent to be locked by default")
	}

	err := agent.RegisterLSMHook("test_hook")
	if err == nil {
		t.Error("expected error when registering hook while locked")
	}

	agent.Unlock()
	if agent.IsLocked() {
		t.Error("expected kernel agent to be unlocked")
	}

	err = agent.RegisterLSMHook("test_hook")
	if err != nil {
		t.Errorf("expected no error when registering hook while unlocked: %v", err)
	}
}
