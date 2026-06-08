package intelligence

import (
	"github.com/fallofpheonix/phoenix/foundation/runtime/bus"
	"github.com/fallofpheonix/phoenix/assurance/security/engine"
	"testing"
)

func TestOrchestrator_StagesIsolation(t *testing.T) {
	orch := NewAIOrchestrator("/tmp/test_orchestrator.log")

	// Disable all formal cognition flags to test basic substrate fallback
	orch.EnableFormalMemory = false
	orch.EnableFormalKnowledge = false
	orch.EnableFormalReasoning = false
	orch.EnableFormalReflection = false

	event := bus.TelemetryEvent{
		SeqID:     1,
		EventID:   "EVT_1",
		EventType: "TEST",
	}

	// 1. Ingress Stage without Hub
	// Should not panic even if hub is nil
	orch.ingressStage(event, nil)

	// 2. Sensory Stage without Monitor Feature
	// Replace monitor with nil or unregistered to test negative justification
	delete(orch.features, "monitor")
	score := orch.sensoryStage(event)
	if score.ZScore != 0.0 {
		t.Errorf("Expected zero drift score without monitor, got %v", score.ZScore)
	}

	// 3. Temporal Stage without TCS Feature
	delete(orch.features, "tcs")
	tcsScore := orch.temporalStage(event, false, nil)
	if tcsScore != 1.0 {
		t.Errorf("Expected default TCS score 1.0, got %v", tcsScore)
	}

	// 4. Strategic Stage without Arbiter Feature
	delete(orch.features, "arbiter")
	state, class, auth := orch.strategicStage(score, tcsScore)
	if state != engine.StateSafe || class != engine.ClassNone || auth != false {
		t.Errorf("Expected fallback strategic output (StateSafe, ClassNone, false), got %v, %v, %v", state, class, auth)
	}

	// 5. Cognitive Pipeline without Oracle or Hub
	// Should not panic, should just return early or handle gracefully
	orch.runCognitivePipeline(event, 1, tcsScore, score, state, class, nil, nil)
}

func TestOrchestrator_ReconcileGraph_Safe(t *testing.T) {
	orch := NewAIOrchestrator("/tmp/test_orchestrator.log")

	// Ensure calling ReconcileGraph without a ledger doesn't panic
	delete(orch.features, "ledger")
	orch.ReconcileGraph()
}
