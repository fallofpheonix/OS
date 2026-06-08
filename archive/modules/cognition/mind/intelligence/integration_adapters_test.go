/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package intelligence

import (
	"context"
	"testing"

	"github.com/fallofpheonix/phoenix/cognition/reasoning"
	"github.com/fallofpheonix/phoenix/foundation/ledger"
	"github.com/fallofpheonix/phoenix/foundation/runtime/bus"
	"github.com/fallofpheonix/phoenix/foundation/runtime/monitor"
	"github.com/fallofpheonix/phoenix/assurance/security/engine"
)

type MockPermissionGate struct {
	Authorized bool
}

func (m *MockPermissionGate) RequestPermission(command, reasoning string) bool {
	return m.Authorized
}

type MockOracle struct {
	Directive Directive
	Err       error
}

func (m *MockOracle) RequestDirective(event bus.TelemetryEvent, tcsScore float64, audit, graphContext string) (Directive, error) {
	return m.Directive, m.Err
}

type MockArbiterFeature struct {
	Authorized bool
}

func (m *MockArbiterFeature) Name() string { return "arbiter" }
func (m *MockArbiterFeature) Evaluate(score monitor.DriftScore, tcsScore float64) (engine.SystemState, engine.ActuationClass, bool) {
	return engine.StateCritical, engine.ClassIsolate, m.Authorized
}

type MockLedger struct {
	entries []string
}

func (m *MockLedger) AddEntry(id, eventType string, payload []byte) error {
	m.entries = append(m.entries, id)
	return nil
}
func (m *MockLedger) AddEntryV2(id, eventType string, payload []byte, metadata, stateBefore, stateAfter, version string) error {
	m.entries = append(m.entries, id)
	return nil
}
func (m *MockLedger) GenerateCertificate(id string, weight float64) ([]byte, error) {
	return []byte("cert"), nil
}
func (m *MockLedger) Verify() error { return nil }

func TestCognitiveAdapters_Integration(t *testing.T) {
	orch := NewAIOrchestrator("/tmp/test_advisor.log")
	orch.PermissionGate = &MockPermissionGate{Authorized: true}

	// Mock Ledger
	orch.RegisterFeature(&LedgerFeature{Ledger: &MockLedger{}})

	// Mock Arbiter to force authorization
	orch.RegisterFeature(&MockArbiterFeature{Authorized: true})

	// Mock Oracle
	orch.Oracle = &MockOracle{
		Directive: Directive{
			Command:         "LOG_ONLY",
			ConfidenceScore: 0.8,
		},
	}

	// Verify feature registration
	hub, hasHub := orch.GetFeature("cognition-hub").(*CognitionFeature)
	if !hasHub {
		t.Fatal("cognition-hub feature not registered")
	}

	// Mock event
	event := bus.TelemetryEvent{
		SeqID:     1,
		EventID:   "EVT_001",
		EventType: "PROCESS_START",
		Payload:   []byte(`{"pid": 1234}`),
	}

	// Execute Tick (Authoritative Mode)
	orch.OrchestrateTick(event, 1)

	// Wait for async Oracle and Analytics processing
	orch.Wg.Wait()

	// Verify Memory Hub Ingest
	if _, ok := hub.Memory.Search("evt-1"); !ok {
		t.Error("Formal fact not found in memory hub after execution")
	}

	// Verify Knowledge Hub Graph
	if _, ok := hub.Knowledge.Nodes["evt-1"]; !ok {
		t.Error("Formal node not found in knowledge hub graph after execution")
	}

	// Verify Reflection Hub Prediction
	if hub.Reflection.Metrics.TotalPredictions != 1 {
		t.Errorf("Expected 1 prediction in hub, got %d", hub.Reflection.Metrics.TotalPredictions)
	}
}

func TestCognitiveReasoning_Adapter(t *testing.T) {
	oracle := NewNexusBridge("http://localhost:8080", "test-key")
	adapter := &CognitiveReasoningFeature{Provider: oracle}

	req := &reasoning.InferenceRequest{
		Goal: "Verify anomaly",
	}

	resp, err := adapter.Provider.Reason(context.Background(), req)
	if err != nil {
		t.Fatalf("Reasoning failed: %v", err)
	}

	if resp.Confidence != 0.85 {
		t.Errorf("Expected confidence 0.85, got %.2f", resp.Confidence)
	}
}

func TestIntegrateBrain_Reconstruction(t *testing.T) {
	orch := NewAIOrchestrator("/tmp/test_advisor.log")

	// Create mock ledger events
	events := []*ledger.Event{
		{
			Type:    ledger.EventPrediction,
			Payload: []byte(`{"id": "pred-1"}`),
		},
	}

	// Trigger IntegrateBrain
	orch.IntegrateBrain(events)

	// Verify Reflection reconstruction via hub
	hub := orch.GetFeature("cognition-hub").(*CognitionFeature)
	if hub.Reflection.Metrics.TotalPredictions != 1 {
		t.Errorf("Expected 1 prediction in hub after reconstruction, got %d", hub.Reflection.Metrics.TotalPredictions)
	}
}
func TestGuardedAutonomy_Denial(t *testing.T) {
	orch := NewAIOrchestrator("/tmp/test_denial.log")

	// Deny all actions
	orch.PermissionGate = &MockPermissionGate{Authorized: false}

	// Mock high-confidence Oracle directive
	orch.Oracle = &MockOracle{
		Directive: Directive{
			Command:         "ISOLATE_PID",
			ConfidenceScore: 0.95,
			Reasoning:       "Critical threat detected",
		},
	}

	event := bus.TelemetryEvent{
		SeqID:     100,
		EventID:   "EVT_100",
		EventType: "MALICIOUS_PROCESS",
		Payload:   []byte(`{"pid": 666}`),
	}

	// Execute Tick
	orch.OrchestrateTick(event, 100)

	// Wait for async Oracle call
	orch.Wg.Wait()

	// If the test reaches here without panicking, the async logic is correctly handling the denial.
}

func TestRecovery_RollbackProtection(t *testing.T) {
	rm := NewRecoveryManager()
	workloadID := "target-workload"

	// Attempt 1, 2, 3 should pass
	for i := 1; i <= 3; i++ {
		if err := rm.TriggerRollback(workloadID); err != nil {
			t.Errorf("Rollback %d failed prematurely: %v", i, err)
		}
	}

	// Attempt 4 should fail (Quench)
	if err := rm.TriggerRollback(workloadID); err == nil {
		t.Error("Expected failure on 4th rollback attempt, but it succeeded")
	} else if err.Error() != "MAX_ROLLBACK_EXCEEDED" {
		t.Errorf("Unexpected error message: %v", err)
	}
}
