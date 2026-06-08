/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: INTEGRATED - Formal Cognition Authoritative
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 *
 * [PRAS HEADER]
 * Purpose: Defines the modular cognitive features used by the AI Orchestrator. Includes Graph, Trace, Monitor, TCS, Arbiter, Warden, Reality, and Ledger integrations.
 * Subsystem: Cognition Intelligence
 * Dependencies: arbiter, guardEngine, monitor, tcs, traceEngine, bus, serialization, ledger, contracts, process_graphs
 * Dependents: AIOrchestrator
 * Security Considerations: High. WardenFeature handles system actuation; ArbiterFeature makes strategic decisions; LedgerFeature maintains the source of truth.
 * Performance Considerations: Graph mutations use RWMutex for concurrency. TCS uses sliding windows for efficient evaluation.
 * Labels: feature-modular, causal-graph, system-actuation, truth-ledger
 */
package intelligence

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fallofpheonix/phoenix/cognition/knowledge"
	rootMemory "github.com/fallofpheonix/phoenix/cognition/memory"
	"github.com/fallofpheonix/phoenix/cognition/reasoning"
	"github.com/fallofpheonix/phoenix/cognition/reflection"
	"github.com/fallofpheonix/phoenix/foundation/runtime/arbiter"
	"github.com/fallofpheonix/phoenix/foundation/runtime/bus"
	"github.com/fallofpheonix/phoenix/foundation/runtime/common/serialization"
	"github.com/fallofpheonix/phoenix/foundation/contracts"
	"github.com/fallofpheonix/phoenix/foundation/runtime/monitor"
	"github.com/fallofpheonix/phoenix/foundation/runtime/tcs"
	guardEngine "github.com/fallofpheonix/phoenix/assurance/security/engine"
)

// Feature is the interface that all cognitive capabilities must implement.
type Feature interface {
	Name() string
}

// MonitorFeatureInterface allows for multiple monitor implementations.
type MonitorFeatureInterface interface {
	Feature
	Process(event bus.TelemetryEvent) monitor.DriftScore
}

// ArbiterFeatureInterface allows for multiple policy engine implementations.
type ArbiterFeatureInterface interface {
	Feature
	Evaluate(score monitor.DriftScore, tcsScore float64) (guardEngine.SystemState, guardEngine.ActuationClass, bool)
}

// WardenFeatureInterface allows for multiple actuation implementations.
type WardenFeatureInterface interface {
	Feature
	Actuate(targetState guardEngine.SystemState, class guardEngine.ActuationClass, tcsScore float64, seqID int64, targetPID int, targetTgid uint32, lamportClock uint64, evLedger contracts.ILedger, proof *guardEngine.GraphProof, evidence *RichActuationEvidence)
	IsSimulation() bool
}

type MonitorFeature struct {
	Service *monitor.MonitorService
}

func (mf *MonitorFeature) Name() string { return "monitor" }

/*
 * [FUNCTION HEADER]
 * Purpose: Processes a telemetry event through the monitor service.
 * Responsibilities: Compute a drift score based on incoming telemetry.
 * Inputs: event (bus.TelemetryEvent)
 * Outputs: monitor.DriftScore
 * Complexity: O(Monitor_Logic)
 */
func (mf *MonitorFeature) Process(event bus.TelemetryEvent) monitor.DriftScore {
	return mf.Service.Process(event)
}

type TCSFeature struct {
	Window *tcs.SlidingWindow
	DegMon *tcs.DegradationMonitor
}

func (tf *TCSFeature) Name() string { return "tcs" }

/*
 * [FUNCTION HEADER]
 * Purpose: Adds an event to the TCS sliding window and evaluates confidence.
 * Responsibilities: Convert bus event to TCS event and compute new TCS score.
 * Inputs: event (bus.TelemetryEvent)
 * Outputs: float64 (TCS Score)
 * Complexity: O(Sliding_Window_Size)
 */
func (tf *TCSFeature) AddAndEvaluate(event bus.TelemetryEvent) float64 {
	tf.Window.AddEvent(tcs.TelemetryEvent{
		Timestamp:  time.Unix(int64(event.LamportClock), 0),
		SequenceID: uint64(event.SeqID),
		Payload:    event.Payload,
	})
	return tf.Window.Evaluate()
}

/*
 * [FUNCTION HEADER]
 * Purpose: Monitors for TCS confidence degradation.
 * Responsibilities: Trigger a ledger entry if a confidence violation is detected.
 * Inputs: tcsScore (float64), seqID (int64), evLedger (contracts.ILedger)
 * Outputs: None
 * Complexity: O(1)
 */
func (tf *TCSFeature) EvaluateDegradation(tcsScore float64, seqID int64, evLedger contracts.ILedger) {
	oldDegraded := tf.DegMon.IsDegraded()
	tf.DegMon.Evaluate(tcsScore)
	if tf.DegMon.IsDegraded() != oldDegraded {
		action := "NORMAL"
		if tf.DegMon.IsDegraded() {
			action = "DEGRADED"
		}
		payload, _ := serialization.CanonicalJSON(map[string]interface{}{"action": action, "score": tcsScore})
		if err := evLedger.AddEntry(fmt.Sprintf("TCS-ANOMALY-%d", seqID), "TCS-VIOLATION", uint64(seqID), payload); err != nil {
			fmt.Fprintf(os.Stderr, "Arbiter: Failed to add TCS anomaly to ledger: %v\n", err)
		}
	}
}

type ArbiterFeature struct {
	Arb *arbiter.Arbiter
}

func (af *ArbiterFeature) Name() string { return "arbiter" }

/*
 * [FUNCTION HEADER]
 * Purpose: Evaluates drift and TCS scores to make a strategic decision.
 * Responsibilities: Delegate to the underlying Arbiter service for decision logic.
 * Inputs: score (monitor.DriftScore), tcsScore (float64)
 * Outputs: targetState (guardEngine.SystemState), class (guardEngine.ActuationClass), authorized (bool)
 * Complexity: O(Arbiter_Logic)
 */
func (af *ArbiterFeature) Evaluate(score monitor.DriftScore, tcsScore float64) (guardEngine.SystemState, guardEngine.ActuationClass, bool) {
	state, class, authorized := af.Arb.Evaluate(score, tcsScore)
	return guardEngine.SystemState(state), guardEngine.ActuationClass(class), authorized
}

// RichActuationEvidence captures the full context of a (simulated) actuation.
type RichActuationEvidence struct {
	TargetWorkload string               `json:"target_workload"`
	SignalTotal    float64              `json:"signal_total"`
	Components     map[string]float64   `json:"signal_components"`
	ReasonPath     reasoning.ReasonPath `json:"reason_path"`
	Counterfactual string               `json:"counterfactual"`
	Confidence     float64              `json:"confidence"`
}

type WardenFeature struct {
	Warden         *guardEngine.Warden
	SimulationMode bool
}

func (wf *WardenFeature) Name() string { return "warden" }

func (wf *WardenFeature) IsSimulation() bool { return wf.SimulationMode }

/*
 * [FUNCTION HEADER]
 * Purpose: Executes or simulates a system actuation based on authorized directives.
 * Responsibilities: Generate certificates, log rich evidence, and optionally trigger state transitions.
 * Inputs: targetState, class, tcsScore, seqID, targetPID, targetTgid, lamportClock, evLedger, proof, evidence
 * Outputs: None
 * Complexity: O(Warden_Actuate_Time)
 */
func (wf *WardenFeature) Actuate(
	targetState guardEngine.SystemState,
	class guardEngine.ActuationClass,
	tcsScore float64,
	seqID int64,
	targetPID int,
	targetTgid uint32,
	lamportClock uint64,
	evLedger contracts.ILedger,
	proof *guardEngine.GraphProof,
	evidence *RichActuationEvidence,
) {
	stateBefore := string(wf.Warden.GetState())
	eventID := fmt.Sprintf("TCS-ANOMALY-%d", seqID)
	cert, _ := evLedger.GenerateCertificate(eventID, tcsScore)

	req := guardEngine.AuthorityEscalationRequest{
		EventID:        eventID,
		TargetPID:      targetPID,
		TargetTgid:     int(targetTgid),
		TargetNsproxy:  proof.ExpectedNsproxy,
		TargetState:    targetState,
		ActuationClass: class,
		EvidenceWeight: tcsScore,
		Certificate:    cert,
		GraphProof:     proof,
	}

	if wf.SimulationMode {
		// Log rich SIMULATED_ACTUATION (Phase 4A)
		payload, _ := serialization.CanonicalJSON(evidence)
		_ = evLedger.AddEntryV2(
			fmt.Sprintf("SIM-ACTUATION-%d", seqID),
			"SIMULATED-ACTUATION",
			uint64(seqID),
			payload,
			"",
			[]byte(stateBefore),
			[]byte(targetState),
			"1.0.0",
		)
		log.Printf("[SIMULATION] Recorded rich evidence for simulated actuation on workload: %s", evidence.TargetWorkload)
		return
	}

	if transitioned := wf.Warden.ActuateRequest(req, int(seqID), lamportClock); transitioned {
		payload, _ := serialization.CanonicalJSON(map[string]interface{}{"state": string(wf.Warden.GetState())})
		if err := evLedger.AddEntryV2(fmt.Sprintf("WARDEN-ACTION-%d", seqID), "POLICY-ACTUATION", uint64(seqID), payload, "", []byte(stateBefore), []byte(wf.Warden.GetState()), "1.0.0"); err != nil {
			fmt.Fprintf(os.Stderr, "Warden: Failed to add actuation to ledger: %v\n", err)
		}
	}
}

type RealityFeature struct {
	AuditPath string
}

func (rf *RealityFeature) Name() string { return "reality" }

/*
 * [FUNCTION HEADER]
 * Purpose: Reads the system audit report from disk.
 * Responsibilities: Provide ground truth context for AI directives.
 * Inputs: None
 * Outputs: string (Audit Content)
 * Complexity: O(File_Read)
 */
func (rf *RealityFeature) ReadAudit() string {
	content, err := os.ReadFile(rf.AuditPath)
	if err != nil {
		return "Audit report missing"
	}
	return string(content)
}

type LedgerFeature struct {
	Ledger contracts.ILedger
}

func (lf *LedgerFeature) Name() string { return "ledger" }

// CognitiveKnowledgeFeature wraps formal knowledge packages.
type CognitiveKnowledgeFeature struct {
	Graph  *knowledge.Graph
	Belief *knowledge.BeliefEngine
}

func (cf *CognitiveKnowledgeFeature) Name() string { return "cognitive-knowledge" }

// CognitiveReflectionFeature wraps formal reflection packages.
type CognitiveReflectionFeature struct {
	Engine  *reflection.Engine
	Auditor *reflection.RealityDriftAuditor
}

func (cf *CognitiveReflectionFeature) Name() string { return "cognitive-reflection" }

// CognitiveReasoningFeature wraps formal reasoning packages.
type CognitiveReasoningFeature struct {
	Provider reasoning.Provider
}

func (cf *CognitiveReasoningFeature) Name() string { return "cognitive-reasoning" }

// CognitiveMemoryFeature wraps formal tiered memory package.
type CognitiveMemoryFeature struct {
	Store *rootMemory.TieredMemory
}

func (cf *CognitiveMemoryFeature) Name() string { return "cognitive-memory" }

// CognitionFeature is the unified hub for all formal cognitive layers.
// It implements the Feature interface and provides direct access to
// Memory, Knowledge, Beliefs, Reflection, and Reasoning subsystems.
type CognitionFeature struct {
	Memory     *rootMemory.TieredMemory
	Knowledge  *knowledge.Graph
	Beliefs    *knowledge.BeliefEngine
	Reflection *reflection.Engine
	Auditor    *reflection.RealityDriftAuditor
	Reasoning  reasoning.Provider
}

func (cf *CognitionFeature) Name() string { return "cognition-hub" }
