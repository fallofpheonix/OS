/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 *
 * [PRAS FILE HEADER]
 * PURPOSE: Publishes AI-generated advisories to the system bus in a structured envelope.
 * SUBSYSTEM: Phoenix.Cognition/PhoenixMind/Advisory
 * DEPENDENCIES: encoding/json, fmt, time, github.com/fallofpheonix/phoenix/foundation/runtime/bus
 * DEPENDENTS: Phoenix.Terminus/PhoenixOS (Warden/Orchestrator), Monitoring Agents
 * SECURITY CONSIDERATIONS: Enforces AdvisoryEnvelope constraints; requires warden consent.
 * PERFORMANCE CONSIDERATIONS: Low latency; performance bound by bus throughput.
 *
 * [LABELS]: stable, core-integration
 */
/*
 * REPOSITORY: PhoenixMind
 * ARCHITECTURAL JUSTIFICATION: Containment valve ensuring AI remains an advisory oracle.
 * DEPENDENCY BOUNDARY: Strictly formats output to canonical AdvisoryEnvelope.
 * DETERMINISTIC CONSIDERATIONS: Enforces bounded action scope and expiration.
 */

package advisory

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/fallofpheonix/phoenix/foundation/runtime/bus"
	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
)

// AdvisoryEnvelope mirrors the canonical contract defined in PhoenixCore.
// [PRAS CLASS HEADER]
// PURPOSE: Defines the canonical structure for AI advisories.
// RESPONSIBILITIES: Data transfer object for structured advisory metadata.
// INPUTS: N/A (Struct)
// OUTPUTS: N/A (Struct)
// COMPLEXITY: O(1)
type AdvisoryEnvelope struct {
	AdvisoryID          string             `json:"advisory_id"`
	Confidence          phxmath.FixedPoint `json:"confidence"`
	EvidenceReferences  []string           `json:"evidence_references"`
	RecommendationType  string             `json:"recommendation_type"`
	BoundedActionScope  string             `json:"bounded_action_scope"`
	Expiration          int64              `json:"expiration"`
	ApprovalRequirement string             `json:"approval_requirement"`
	ForbiddenActions    []string           `json:"forbidden_actions"`
	Reasoning           string             `json:"reasoning"`
}

// Publisher handles the distribution of advisories.
// [PRAS CLASS HEADER]
// PURPOSE: Handles the distribution of advisories to the PhoenixOS bus.
// RESPONSIBILITIES: Formatting reasoning output into the AdvisoryEnvelope and publishing.
// INPUTS: bus.Bus (Dependency)
// OUTPUTS: *Publisher
// COMPLEXITY: O(1)
type Publisher struct {
	Bus *bus.Bus
}

// [PRAS FUNCTION HEADER]
// PURPOSE: Constructor for Publisher.
// RESPONSIBILITIES: Initializes Publisher with a shared bus instance.
// INPUTS: b (*bus.Bus)
// OUTPUTS: *Publisher
// COMPLEXITY: O(1)
func NewPublisher(b *bus.Bus) *Publisher {
	return &Publisher{Bus: b}
}

// Publish takes reasoning output and formats it into the strict containment envelope.
// [PRAS FUNCTION HEADER]
// PURPOSE: Formats and publishes an advisory event.
// RESPONSIBILITIES: Wrapping reasoning in an envelope, marshaling to JSON, and sending to bus.
// INPUTS: confidence (float64), reasoning (string), evidence ([]string), actionScope (string)
// OUTPUTS: error
// COMPLEXITY: O(1) space, O(N) serialization where N is payload size.
func (p *Publisher) Publish(confidence float64, reasoning string, evidence []string, actionScope string) error {
	now := time.Now()
	fpConfidence := phxmath.FixedPoint{V: int64(confidence * 1000000)}

	env := AdvisoryEnvelope{
		AdvisoryID:          fmt.Sprintf("adv-%d", now.UnixNano()),
		Confidence:          fpConfidence,
		EvidenceReferences:  evidence,
		RecommendationType:  "CONTAINMENT_PROPOSAL",
		BoundedActionScope:  actionScope,
		Expiration:          now.Add(5 * time.Minute).Unix(),
		ApprovalRequirement: "WARDEN_DETERMINISTIC_CONSENT",
		ForbiddenActions:    []string{"UNBOUNDED_KERNEL_ACCESS", "DIRECT_MEMORY_WRITES"},
		Reasoning:           reasoning,
	}

	payload, err := json.Marshal(env)
	if err != nil {
		return err
	}

	// Wrap in canonical TelemetryEvent for the bus distribution.
	telem := bus.TelemetryEvent{
		SeqID:        now.UnixNano(),
		MonotonicNs:  now.UnixNano(),
		WallTimeUnix: now.Unix(),
		Source:       "PhoenixMind:Oracle",
		EventType:    "ADVISORY_PROPOSAL",
		Severity:     fpConfidence,
		Payload:      payload,
	}

	p.Bus.Publish("phoenix.advisories", telem)
	return nil
}
