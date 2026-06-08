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

	rootMemory "github.com/fallofpheonix/phoenix/cognition/memory"
	"github.com/fallofpheonix/phoenix/cognition/reasoning"
	"github.com/fallofpheonix/phoenix/foundation/runtime/bus"
	"github.com/fallofpheonix/phoenix/cognition/mind/advisory"
	"github.com/fallofpheonix/phoenix/cognition/mind/memory"
)

type MockInferenceProvider struct{}

func (m *MockInferenceProvider) Name() string { return "MockProvider" }

func (m *MockInferenceProvider) Reason(ctx context.Context, req *reasoning.InferenceRequest) (*reasoning.InferenceResponse, error) {
	return &reasoning.InferenceResponse{
		Logic:      "Mock reasoning: Anomaly detected.",
		Proposal:   "ISOLATE",
		Confidence: 0.85,
	}, nil
}

func TestEngine_ProcessTelemetry(t *testing.T) {
	b := bus.NewBus()
	pub := advisory.NewPublisher(b)
	mem := memory.NewCognitiveMemoryBridge(rootMemory.NewTieredMemory())
	prov := &MockInferenceProvider{}
	engine := NewEngine(mem, prov, pub)

	// Subscribe to advisory topic to verify containment
	ch := b.Subscribe("phoenix.advisories")

	event := bus.TelemetryEvent{
		SeqID:     123,
		PID:       456,
		EventType: "PROCESS_ANOMALY",
		Severity:  0.9,
	}

	ctx := context.Background()
	engine.ProcessTelemetry(ctx, event)

	select {
	case adv := <-ch:
		if adv.EventType != "ADVISORY_PROPOSAL" {
			t.Errorf("Expected ADVISORY_PROPOSAL, got %s", adv.EventType)
		}
		if adv.Severity != 0.85 {
			t.Errorf("Expected severity (confidence) 0.85, got %.2f", adv.Severity)
		}
	default:
		t.Error("No advisory proposal published to the bus")
	}
}
