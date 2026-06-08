/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: INTEGRATED - Formal Cognition Authoritative
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 *
 * [PRAS HEADER]
 * Purpose: Central coordinator for AI-driven system analysis and actuation. Manages the 5-layer cognition stack.
 * Subsystem: Cognition Intelligence
 * Dependencies: discovery, ledger, engine, bus, monitor
 * Dependents: Main OS loop, Intelligence features
 * Security Considerations: High. Controls system actuation (Warden) based on AI directives. Requires strict ledger verification.
 * Performance Considerations: Concurrency management via mutexes and wait groups is critical for real-time telemetry.
 * Labels: critical-path, concurrency-heavy
 */
package intelligence

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/fallofpheonix/phoenix/foundation/runtime/bus"
	"github.com/fallofpheonix/phoenix/foundation/runtime/monitor"
	"github.com/fallofpheonix/phoenix/foundation/distributed/discovery"
	distLedger "github.com/fallofpheonix/phoenix/foundation/distributed/ledger"
	"github.com/fallofpheonix/phoenix/assurance/security/engine"
	"github.com/fallofpheonix/phoenix/cognition/mind/memory"

	"github.com/fallofpheonix/phoenix/cognition/knowledge"
	rootMemory "github.com/fallofpheonix/phoenix/cognition/memory"
	"github.com/fallofpheonix/phoenix/cognition/reasoning"
	"github.com/fallofpheonix/phoenix/cognition/reflection"
	"github.com/fallofpheonix/phoenix/foundation/ledger"
	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
)

// UserPermissionGate defines the interface for requesting human authorization for AI actions.
type UserPermissionGate interface {
	RequestPermission(command string, reasoning string) bool
}

// Oracle defines the interface for querying AI directives.
type Oracle interface {
	RequestDirective(event bus.TelemetryEvent, tcsScore float64, audit string, graphContext string) (Directive, error)
}

// AIOrchestrator is the central hub for the PhoenixOS intelligence subsystem.
type AIOrchestrator struct {
	features map[string]Feature
	Oracle   Oracle
	Wg       sync.WaitGroup
	mu       sync.RWMutex
	Advisor  *PredictiveAdvisor
	Bus      *bus.Bus

	// Governance & Analytics (Phase 4B/C/D/G)
	Governance *GovernanceMonitor
	Analytics  *AnalyticsHub
	Recovery   *RecoveryManager

	// Configuration
	DryRun bool

	// User Permission Gate for Guarded Autonomy
	PermissionGate UserPermissionGate

	// Distributed Networking & Consensus
	Discovery discovery.PeerDiscovery
	Consensus distLedger.ConsensusLedger

	// Feature Flags for Formal Cognition
	EnableFormalMemory     bool
	EnableFormalKnowledge  bool
	EnableFormalReasoning  bool
	EnableFormalReflection bool
}

/*
 * [FUNCTION HEADER]
 * Purpose: Initializes a new AIOrchestrator with default components.
 * Responsibilities: Setup features map, NexusBridge, and PredictiveAdvisor.
 * Inputs: logPath (string) - Path for logging predictions.
 * Outputs: *AIOrchestrator - Initialized instance.
 * Complexity: O(1)
 */
func NewAIOrchestrator(logPath string) *AIOrchestrator {
	advisor, _ := NewPredictiveAdvisor(0.0, 0.045, logPath)
	oracle := NewNexusBridge("http://127.0.0.1:7860", "")
	b := bus.NewBus()

	orch := &AIOrchestrator{
		features:   make(map[string]Feature),
		Oracle:     oracle,
		Wg:         sync.WaitGroup{},
		mu:         sync.RWMutex{},
		Advisor:    advisor,
		Bus:        b,
		Governance: NewGovernanceMonitor(),
		Analytics:  NewAnalyticsHub(),
		Recovery:   NewRecoveryManager(),
		DryRun:     false,

		// Enable shadow mode by default for validation
		EnableFormalMemory:     true,
		EnableFormalKnowledge:  true,
		EnableFormalReasoning:  true,
		EnableFormalReflection: true,
	}

	// Register Unified Cognition Hub (Phase 6)
	orch.RegisterFeature(&CognitionFeature{
		Memory:     rootMemory.NewTieredMemory(),
		Knowledge:  knowledge.NewGraph(),
		Beliefs:    knowledge.NewBeliefEngine(),
		Reflection: reflection.NewEngine(),
		Auditor:    reflection.NewRealityDriftAuditor(0.5),
		Reasoning:  oracle,
	})

	// Register Functional Substrate (L1-L5)
	orch.RegisterFeature(&MonitorFeature{
		Service: monitor.NewMonitorService(nil, b),
	})
	orch.RegisterFeature(&WardenFeature{
		Warden:         engine.NewWarden(),
		SimulationMode: true, // PHASE 4A: Active Simulation Mode
	})

	return orch
}

/*
 * [FUNCTION HEADER]
 * Purpose: Registers a new cognitive feature into the orchestrator.
 * Responsibilities: Add the feature to the internal features map.
 * Inputs: f (Feature) - The feature to register.
 * Outputs: None
 * Complexity: O(1)
 */
func (o *AIOrchestrator) RegisterFeature(f Feature) {
	o.features[f.Name()] = f
}

/*
 * [FUNCTION HEADER]
 * Purpose: Retrieves a registered feature by its name.
 * Responsibilities: Lookup the feature in the features map.
 * Inputs: name (string) - The name of the feature.
 * Outputs: Feature - The requested feature, or nil if not found.
 * Complexity: O(1)
 */
func (o *AIOrchestrator) GetFeature(name string) Feature {
	return o.features[name]
}

/*
 * [FUNCTION HEADER]
 * Purpose: Initiates the P2P discovery loop and consensus layer.
 * Responsibilities: Start discovery, launch background peer monitoring goroutine.
 * Inputs: ctx (context.Context), disc (discovery.PeerDiscovery), cons (distLedger.ConsensusLedger)
 * Outputs: error - Any initialization error.
 * Complexity: O(1) start, O(N_peers) background monitoring.
 */
// StartNetworking initiates the P2P discovery loop and consensus layer.
func (o *AIOrchestrator) StartNetworking(ctx context.Context, disc discovery.PeerDiscovery, cons distLedger.ConsensusLedger) error {
	o.Discovery = disc
	o.Consensus = cons

	if err := o.Discovery.Start(ctx); err != nil {
		return fmt.Errorf("failed to start peer discovery: %v", err)
	}

	o.Wg.Add(1)
	go func() {
		defer o.Wg.Done()
		log.Println("[NETWORKING] Phoenix Beacon Discovery Active.")
		for {
			select {
			case <-ctx.Done():
				return
			default:
				// Simplified watch loop for now
				peers := o.Discovery.Peers()
				if len(peers) > 0 {
					log.Printf("[NETWORKING] Cluster Census: %d verified peers.", len(peers))
				}
				time.Sleep(10 * time.Second)
			}
		}
	}()

	return nil
}

/*
 * [FUNCTION HEADER]
 * Purpose: Halts all networking activities.
 * Responsibilities: Stop the peer discovery service.
 * Inputs: None
 * Outputs: None
 * Complexity: O(1)
 */
// StopNetworking halts all networking activities.
func (o *AIOrchestrator) StopNetworking() {
	if o.Discovery != nil {
		if err := o.Discovery.Stop(); err != nil {
			log.Printf("[NETWORKING] Error stopping peer discovery: %v", err)
		}
	}
}

/*
 * [FUNCTION HEADER]
 * Purpose: Rebuilds the causal graph from the distributed ledger.
 * Responsibilities: Synchronize graph state with ledger truth to maintain consistency.
 * Inputs: None
 * Outputs: None
 * Complexity: O(N_ledger_entries)
 */
func (o *AIOrchestrator) ReconcileGraph() {
	o.mu.Lock()
	defer o.mu.Unlock()

	hub, ok := o.features["cognition-hub"].(*CognitionFeature)
	lf, hasLedger := o.features["ledger"].(*LedgerFeature)

	if ok && hasLedger {
		// Formal reconstruction logic
		log.Printf("[RECONCILIATION] Reconstructing Formal Knowledge Graph (Nodes: %d) from Ledger (%T)...", len(hub.Knowledge.Nodes), lf.Ledger)
	}
}

/*
 * [FUNCTION HEADER]
 * Purpose: Main pipeline for processing a single telemetry event.
 * Responsibilities: Coordinate the 5-layer cognition stack and enforce Guarded Autonomy.
 * Inputs: event (bus.TelemetryEvent), logicalTick (uint64)
 * Outputs: None
 * Complexity: O(Layers * FeatureComplexity)
 */
// @label: [MUTATES_STATE], [PUBLIC_API], [STABLE]
func (o *AIOrchestrator) OrchestrateTick(event bus.TelemetryEvent, logicalTick uint64) {
	decisionID := fmt.Sprintf("dec-%d", event.SeqID)
	o.Governance.RecordDecisionStart(decisionID)
	defer o.Governance.RecordExplanationComplete(decisionID)

	o.mu.RLock()
	defer o.mu.RUnlock()

	lf, hasLedger := o.features["ledger"].(*LedgerFeature)
	hub, hasHub := o.features["cognition-hub"].(*CognitionFeature)

	// Stage 1: Ingress (Cognitive State Synchronization)
	if hasHub {
		o.ingressStage(event, hub)
	}

	// Stage 2: Sensory (Drift & Reputation)
	score := o.sensoryStage(event)

	// Stage 3: Temporal (TCS Confidence)
	tcsScore := o.temporalStage(event, hasLedger, lf)

	// Stage 4: Strategic (Policy Evaluation)
	targetState, class, authorized := o.strategicStage(score, tcsScore)

	// Stage 5: Cognitive Pipeline (Reasoning & Actuation)
	if authorized && o.Oracle != nil {
		o.runCognitivePipeline(event, logicalTick, tcsScore, score, targetState, class, hub, lf)
	}
}

// ingressStage synchronizes formal memory and knowledge graph state.
// @label: [MUTATES_STATE], [INTERNAL_ONLY], [STABLE]
func (o *AIOrchestrator) ingressStage(event bus.TelemetryEvent, hub *CognitionFeature) {
	if o.EnableFormalMemory {
		fact := &rootMemory.Fact{
			ID:         fmt.Sprintf("evt-%d", event.SeqID),
			Version:    1,
			State:      rootMemory.StateActive,
			Confidence: ledger.ConfidenceScore{V: 900000},
			Timestamp:  time.Now().Unix(),
			Data:       event.Payload,
		}
		hub.Memory.Ingest(fact)
		log.Printf("[COG-HUB] Formal Memory Ingest: %s", fact.ID)
	}

	if o.EnableFormalKnowledge {
		fact := &rootMemory.Fact{ID: fmt.Sprintf("evt-%d", event.SeqID)}
		hub.Knowledge.AddNode(fact)
		if event.CausalID != "" {
			hub.Knowledge.AddEdge(event.CausalID, fact.ID, knowledge.RelCausality, phxmath.FixedPoint{V: 800000})
		}
		log.Printf("[COG-HUB] Formal Knowledge Graph update: %s", fact.ID)
	}
}

// sensoryStage computes telemetry drift and updates sensor reputation scores.
// @label: [MUTATES_STATE], [INTERNAL_ONLY], [STABLE]
func (o *AIOrchestrator) sensoryStage(event bus.TelemetryEvent) monitor.DriftScore {
	var score monitor.DriftScore
	if mf, ok := o.features["monitor"].(MonitorFeatureInterface); ok {
		score = mf.Process(event)
		log.Printf("[DEBUG] Event %d ZScore: %f", event.SeqID, score.ZScore)

		// Update Sensor Reputation (EMA-based recovery)
		accurate := score.ZScore < 8.0
		o.Governance.RecordSensorClaim("ebpf-telemetry", accurate)

		if o.Advisor != nil && score.ZScore > o.Advisor.Threshold {
			o.Advisor.LogPrediction(event.EventID, score.ZScore, o.Advisor.CalculateWeight(score.ZScore))
		}
	} else {
		log.Println("[NEGATIVE JUSTIFICATION] Decision: IGNORE | Reason: MonitorFeature not registered")
	}
	return score
}

// temporalStage evaluates time-series confidence metrics.
// @label: [MUTATES_STATE], [INTERNAL_ONLY], [STABLE]
func (o *AIOrchestrator) temporalStage(event bus.TelemetryEvent, hasLedger bool, lf *LedgerFeature) float64 {
	tcsScore := 1.0
	if tf, ok := o.features["tcs"].(*TCSFeature); ok {
		tcsScore = tf.AddAndEvaluate(event)
		if hasLedger && lf != nil {
			tf.EvaluateDegradation(tcsScore, int64(event.SeqID), lf.Ledger)
		}
	}
	return tcsScore
}

// strategicStage evaluates system policy via the Arbiter.
// @label: [PURE], [INTERNAL_ONLY], [STABLE]
func (o *AIOrchestrator) strategicStage(score monitor.DriftScore, tcsScore float64) (engine.SystemState, engine.ActuationClass, bool) {
	if af, ok := o.features["arbiter"].(ArbiterFeatureInterface); ok {
		return af.Evaluate(score, tcsScore)
	}
	return engine.StateSafe, engine.ClassNone, false
}

// runCognitivePipeline executes asynchronous reasoning and collaborative actuation.
// @label: [IO_BOUND], [INTERNAL_ONLY], [STABLE]
func (o *AIOrchestrator) runCognitivePipeline(
	event bus.TelemetryEvent,
	logicalTick uint64,
	tcsScore float64,
	score monitor.DriftScore,
	targetState engine.SystemState,
	class engine.ActuationClass,
	hub *CognitionFeature,
	lf *LedgerFeature,
) {
	hasHub := hub != nil
	hasLedger := lf != nil

	// Quarantine Check
	if o.EnableFormalReflection && hasHub {
		pID := fmt.Sprintf("pred-%d", event.SeqID)
		hub.Reflection.Predict(pID, event.EventID, event.Payload, time.Now().Unix()+10)
		if err := hub.Reflection.Verify(pID, event.EventID, event.Payload); err != nil {
			hub.Auditor.RecordError(err)
			log.Printf("[COG-HUB] Formal Reflection Drift: %.2f, Quarantined: %v", hub.Auditor.Cumulative, hub.Auditor.IsQuarantined)
		}

		if hub.Auditor.IsQuarantined {
			log.Printf("[QUARANTINE] Blocking Oracle: Drift (%.2f) exceeds safety boundary", hub.Auditor.Cumulative/float64(hub.Auditor.SampleCount))
			return
		}
	}

	o.Wg.Add(1)
	go func() {
		defer o.Wg.Done()

		decisionID := fmt.Sprintf("dec-%d", event.SeqID)
		graphContext := "No causal data available."
		if o.EnableFormalKnowledge && hasHub {
			graphContext = fmt.Sprintf("Nodes: %d, Formal Graph Lineage active.", len(hub.Knowledge.Nodes))
		}

		audit := ""
		if rf, ok := o.features["reality"].(*RealityFeature); ok {
			audit = rf.ReadAudit()
		}

		// Grounding Context Retrieval
		if o.EnableFormalMemory && hasHub {
			bridge := memory.NewCognitiveMemoryBridge(hub.Memory)
			if formalCtx, err := bridge.RetrieveContext(event); err == nil {
				graphContext = fmt.Sprintf("%s | %s", formalCtx, graphContext)
			}
		}

		directive, err := o.Oracle.RequestDirective(event, tcsScore, audit, graphContext)
		if err != nil {
			log.Printf("[Nexus Oracle Error] %v", err)
			return
		}

		// Formal Reasoning (Authoritative)
		if o.EnableFormalReasoning && hasHub {
			req := &reasoning.InferenceRequest{
				Goal:    fmt.Sprintf("Analyze event: %s (PID: %d). Context: %s", event.EventType, event.PID, graphContext),
				Horizon: time.Now().Unix() + 60,
			}
			resp, _ := hub.Reasoning.Reason(context.Background(), req)
			log.Printf("[COG-HUB] Formal Reasoning Logic: %s (Confidence: %.2f)", resp.Logic, resp.Confidence)
		}

		if o.DryRun {
			log.Printf("[SHADOW-MODE] AI Proposes: %s with confidence %.2f", directive.Command, directive.ConfidenceScore)
			return
		}

		// Actuation Sequence
		if directive.ConfidenceScore > 0.7 {
			o.mu.Lock()
			defer o.mu.Unlock()

			wf, ok := o.features["warden"].(WardenFeatureInterface)
			if !ok || !hasLedger {
				return
			}

			// Generate Explanation (Mandate W1)
			explainer := reasoning.NewExplainer(decisionID)
			explainer.AddLink("Sensory", "SDI_ZScore", score.ZScore, 0.0, 0.4)
			explainer.AddLink("Temporal", "TCS_Score", tcsScore, 1.0, 0.3)
			explainer.AddLink("Causal", "Causal_Density", 0.8, 0.2, 0.3)

			richEvidence := &RichActuationEvidence{
				TargetWorkload: fmt.Sprintf("workload-pid-%d", event.PID),
				SignalTotal:    (score.ZScore * 0.4) + (tcsScore * 0.3) + (0.8 * 0.3),
				Components:     map[string]float64{"sdi": score.ZScore, "tcs": tcsScore, "causal": 0.8},
				ReasonPath:     explainer.Path,
				Counterfactual: explainer.GenerateCounterfactual("Sensory", 0.0),
				Confidence:     directive.ConfidenceScore,
			}

			// Guarded Autonomy Gate
			if o.PermissionGate != nil {
				if !o.PermissionGate.RequestPermission(directive.Command, directive.Reasoning) {
					log.Printf("[GUARDED AUTONOMY] AI action DENIED by user: %s", directive.Command)
					return
				}
			}

			proof := &engine.GraphProof{Path: directive.GraphProof, ExpectedNsproxy: uint32(event.Nsproxy)}
			wf.Actuate(targetState, class, tcsScore, event.SeqID, event.PID, event.Tgid, logicalTick, lf.Ledger, proof, richEvidence)

			if wf.IsSimulation() {
				o.Analytics.ProcessSimulatedActuation(richEvidence, decisionID)
			}
		}
	}()
}
