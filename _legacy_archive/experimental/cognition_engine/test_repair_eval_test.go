package main

import (
	"phoenixmind/cognition_engine/repair"
	"testing"
)

func TestRepairEval(t *testing.T) {
	eval := &repair.Evaluator{}
	decision := eval.Evaluate(repair.Dangerous)
	if decision != "HUMAN_ONLY" {
		t.Errorf("Expected HUMAN_ONLY, got %s", decision)
	}
}
