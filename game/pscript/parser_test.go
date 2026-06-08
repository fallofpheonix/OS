package pscript

import (
	"testing"
)

func TestParser_ProximityProbe(t *testing.T) {
	input := `
		let target = 100
		move("agent_01", target)
		verify("agent_01")
	`

	l := NewLexer(input)
	p := NewParser(l)

	_, err := p.Parse()
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	// This test is currently verifying parser structure.
	// VM integration tests should be in engine or replay package.
}
