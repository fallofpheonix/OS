/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: INTEGRATED - Formal Cognition Authoritative
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 *
 * [PRAS HEADER]
 * Purpose: Core reasoning engine that orchestrates inference using telemetry data and episodic memory context.
 * Subsystem: Cognition Intelligence
 * Dependencies: bus, advisory, memory
 * Dependents: AIOrchestrator, system-wide reasoning consumers
 * Security Considerations: Medium. Influences system-wide advisories and potential containment actions.
 * Performance Considerations: LLM/Inference latency is the primary bottleneck.
 * Labels: reasoning-engine, inference-bridge
 */
/*
 * REPOSITORY: PhoenixMind
 * ARCHITECTURAL JUSTIFICATION: Core reasoning engine. Oracle orchestration.
 * DEPENDENCY BOUNDARY: Bounded by memory retrieval and advisory publishing.
 * DETERMINISTIC CONSIDERATIONS: Probabilistic scoring mapped to deterministic contracts.
 */

package intelligence

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fallofpheonix/phoenix/cognition/reasoning"
	"github.com/fallofpheonix/phoenix/foundation/runtime/bus"
	"github.com/fallofpheonix/phoenix/cognition/mind/advisory"
	"github.com/fallofpheonix/phoenix/cognition/mind/memory"
)

// Engine handles the high-level reasoning orchestration.
type Engine struct {
	Memory    memory.MemoryBridge
	Provider  reasoning.Provider
	Publisher *advisory.Publisher
}

/*
 * [FUNCTION HEADER]
 * Purpose: Initializes a new Engine for reasoning orchestration.
 * Responsibilities: Bind memory bridge, reasoning provider, and advisory publisher.
 * Inputs: m (memory.MemoryBridge), p (reasoning.Provider), pub (*advisory.Publisher)
 * Outputs: *Engine - Initialized instance.
 * Complexity: O(1)
 */
func NewEngine(m memory.MemoryBridge, p reasoning.Provider, pub *advisory.Publisher) *Engine {
	return &Engine{
		Memory:    m,
		Provider:  p,
		Publisher: pub,
	}
}

/*
 * [FUNCTION HEADER]
 * Purpose: Processes a telemetry event to generate a reasoning-based advisory.
 * Responsibilities: Retrieve memory context, construct prompt, perform formal reasoning, and publish advisory.
 * Inputs: ctx (context.Context), event (bus.TelemetryEvent)
 * Outputs: None
 * Complexity: O(Inference_Time + Context_Retrieval_Time)
 */
// ProcessTelemetry ingests events, retrieves context, and dispatches an advisory via formal reasoning.
func (e *Engine) ProcessTelemetry(ctx context.Context, event bus.TelemetryEvent) {
	// 1. Contextual Grounding (Episodic Retrieval)
	memContext, err := e.Memory.RetrieveContext(event)
	if err != nil {
		log.Printf("[PhoenixMind] Context retrieval failed: %v", err)
		memContext = "Context retrieval unavailable."
	}

	// 2. Formal Reasoning Request
	req := &reasoning.InferenceRequest{
		Goal:    fmt.Sprintf("Analyze event: %s (PID: %d, Severity: %.2f). Context: %s", event.EventType, event.PID, event.Severity, memContext),
		Horizon: time.Now().Unix() + 30,
	}

	// 3. Oracle Reasoning
	resp, err := e.Provider.Reason(ctx, req)
	if err != nil {
		log.Printf("[PhoenixMind] Oracle reasoning failure: %v", err)
		return
	}

	// 4. Advisory Publishing (The Containment Valve)
	evidence := []string{fmt.Sprintf("event:%d", event.SeqID)}
	actionScope := "ISOLATE_PROCESS_BOUNDARY"

	if err := e.Publisher.Publish(resp.Confidence, resp.Logic, evidence, actionScope); err != nil {
		log.Printf("[PhoenixMind] Advisory publication failure: %v", err)
	} else {
		log.Printf("[PhoenixMind] Oracle advisory issued. ID: %d, Confidence: %.2f", event.SeqID, resp.Confidence)
	}
}
