package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fallofpheonix/phoenix-control/warden"
	"github.com/fallofpheonix/phoenix-logic/monitor"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
)

// MindAdvice represents the structured output from PhoenixMind.
type MindAdvice struct {
	Command         string  `json:"command"`
	ConfidenceScore float64 `json:"confidence_score"`
	Reasoning       string  `json:"reasoning"`
}

// PhoenixMind is the advisory LLM loop for PhoenixOS.
type PhoenixMind struct {
	OllamaURL string
	Model     string
	Client    *http.Client
	requests  chan adviceRequest
}

// adviceRequest holds signals for the advisory loop.
type adviceRequest struct {
	score    monitor.DriftScore
	tcsScore float64
	audit    string
}

// BuildPrompt formats the telemetry signals and self-audit into a natural language prompt.
func (pm *PhoenixMind) BuildPrompt(score monitor.DriftScore, tcsScore float64, audit string) string {
	return fmt.Sprintf(`You are PhoenixMind, the AI advisor for PhoenixOS.
Analyze the following telemetry signals AND the Current System Reality Audit, then output a JSON response.

Current System Reality:
%s

Signals:
- Event: %s (PID: %d, UID: %d)
- Entropy/Severity: %.2f
- Z-Score: %.2f (Threshold: 3.0)
- TCS Confidence: %.2f

Output Schema:
{
  "command": "ONE_OF: [ISOLATE_PID, THROTTLE_NETWORK, REVOKE_UID, LOG_ONLY]",
  "confidence_score": 0.0-1.0,
  "reasoning": "Brief technical justification. IDENTIFY if any 'EMPTY' or 'PARTIAL' modules in the Audit are contributing to risk."
}

JSON Response:`, audit, score.EventType, score.PID, score.UID, score.Severity, score.ZScore, tcsScore)
}

// GenerateAdvice calls the local Ollama LLM and parses the advisory output.
func (pm *PhoenixMind) GenerateAdvice(score monitor.DriftScore, tcsScore float64, audit string) (MindAdvice, error) {
	log.Printf("[AI Orchestrator] Sending request to Ollama (Model: %s)...", pm.Model)
	prompt := pm.BuildPrompt(score, tcsScore, audit)

	requestBody, _ := json.Marshal(map[string]interface{}{
		"model":  pm.Model,
		"prompt": prompt,
		"stream": false,
		"format": "json",
	})

	resp, err := pm.Client.Post(pm.OllamaURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return MindAdvice{}, fmt.Errorf("Ollama request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return MindAdvice{}, fmt.Errorf("failed to read Ollama response: %w", err)
	}
	log.Printf("[AI Orchestrator] Received raw response from Ollama (%d bytes)", len(body))

	var ollamaResp struct {
		Response string `json:"response"`
	}
	if err := json.Unmarshal(body, &ollamaResp); err != nil {
		return MindAdvice{}, fmt.Errorf("failed to unmarshal Ollama response: %w", err)
	}

	var advice MindAdvice
	if err := json.Unmarshal([]byte(ollamaResp.Response), &advice); err != nil {
		return MindAdvice{}, fmt.Errorf("failed to parse AI advice JSON: %w", err)
	}

	return advice, nil
}

// NewPhoenixMind creates a new PhoenixMind instance and starts the background advisor.
func NewPhoenixMind(url, model string, o *AIOrchestrator) *PhoenixMind {
	pm := &PhoenixMind{
		OllamaURL: url,
		Model:     model,
		Client:    &http.Client{Timeout: 120 * time.Second},
		requests:  make(chan adviceRequest, 100),
	}
	go pm.runAdvisor(o)
	return pm
}

// runAdvisor is a dedicated go-routine that aggregates signals and calls the LLM.
// It implements 5-second batching to respect resource limits (Step 7 & 8).
func (pm *PhoenixMind) runAdvisor(o *AIOrchestrator) {
	log.Printf("[AI Orchestrator] PhoenixMind Advisor loop started")
	ticker := time.NewTicker(5 * time.Second)
	var latest *adviceRequest

	for {
		select {
		case req := <-pm.requests:
			log.Printf("[AI Orchestrator] Received evaluation request for event %s", req.score.EventType)
			// Aggregation strategy: in the 5s window, prioritize the highest Z-Score (most chaotic).
			if latest == nil || req.score.ZScore > latest.score.ZScore {
				latest = &req
			}
		case <-ticker.C:
			if latest != nil {
				advice, err := pm.GenerateAdvice(latest.score, latest.tcsScore, latest.audit)
				if err != nil {
					log.Printf("[AI Orchestrator][PhoenixMind Error] %v", err)
				} else {
					log.Printf("[AI Orchestrator][PhoenixMind Advice] Seq: %d, Command: %s, Confidence: %.2f, Reasoning: %s",
						latest.score.EventID, advice.Command, advice.ConfidenceScore, advice.Reasoning)
				}
				if o != nil {
					o.Wg.Done()
				}
				latest = nil
			}
		}
	}
}

// RequestAdvice queues a new signal for AI evaluation.
func (pm *PhoenixMind) RequestAdvice(score monitor.DriftScore, tcsScore float64, audit string, wg *sync.WaitGroup) {
	select {
	case pm.requests <- adviceRequest{score, tcsScore, audit}:
		if wg != nil {
			wg.Add(1)
		}
	default:
		// Drop request if channel is full to prevent backpressure on telemetry fast-path.
	}
}

// MetaLearningMonitor tracks the system's awareness of its own limitations.
type MetaLearningMonitor struct {
	AccuracyHistory []float64
	LimitationFlags map[string]bool
	mu              sync.Mutex
}

func NewMetaLearningMonitor() *MetaLearningMonitor {
	return &MetaLearningMonitor{
		LimitationFlags: make(map[string]bool),
	}
}

func (m *MetaLearningMonitor) RecordPerformance(confidence, accuracy float64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.AccuracyHistory = append(m.AccuracyHistory, accuracy)

	// Self-identify limitations: if average accuracy drops, flag the module
	if len(m.AccuracyHistory) > 10 {
		avg := 0.0
		for _, v := range m.AccuracyHistory[len(m.AccuracyHistory)-10:] {
			avg += v
		}
		avg /= 10.0
		if avg < 0.6 {
			m.LimitationFlags["HIGH_UNCERTAINTY_DETECTED"] = true
		}
	}
}

// AIOrchestrator is the central coordinator directing all system subsystems as features.
type AIOrchestrator struct {
	features map[string]Feature
	Mind     *PhoenixMind
	Meta     *MetaLearningMonitor
	Wg       sync.WaitGroup
}

// NewAIOrchestrator creates a new instance of the orchestrator.
func NewAIOrchestrator() *AIOrchestrator {
	o := &AIOrchestrator{
		features: make(map[string]Feature),
		Meta:     NewMetaLearningMonitor(),
		Wg:       sync.WaitGroup{},
	}
	o.Mind = NewPhoenixMind("http://localhost:11434/api/generate", "llama3.2:3b", o)
	return o
}
// RegisterFeature registers a modular security/telemetry feature.
func (o *AIOrchestrator) RegisterFeature(f Feature) {
	o.features[f.Name()] = f
}

// GetFeature retrieves a registered feature by name.
func (o *AIOrchestrator) GetFeature(name string) Feature {
	return o.features[name]
}

// OrchestrateTick runs a logical tick on the event, coordinating the registered features.
func (o *AIOrchestrator) OrchestrateTick(event bus.TelemetryEvent, logicalTick uint64) {
	// PHOENIX BRAIN: The 5-Layer Thinking Path

	// 0. Ensure Ledger (Memory Foundation)
	ledgerFeature, ok := o.features["ledger"].(*LedgerFeature)
	if !ok || ledgerFeature.Ledger == nil {
		log.Fatal("[AI Orchestrator] Ledger feature not registered")
	}

	// 1. The Sensory Brain (Phoenix Sentinel - L6) -> Raw Signal
	var sdiScore float64
	var score monitor.DriftScore
	if monitorFeature, ok := o.features["monitor"].(*MonitorFeature); ok && monitorFeature.Service != nil {
		score = monitorFeature.Process(event)
		sdiScore = score.Severity // Entropy mapped to SDI
		if sdiScore > 5.0 {
			log.Printf("[Sentinel Brain] Sensed raw chaotic signal. SDI: %.2f", sdiScore)
		}
	}

	// 2. The Causal Brain (Phoenix Trace - L4) -> Causal Context
	var lineage string
	if graphFeature, ok := o.features["graph"].(*GraphFeature); ok && graphFeature.Graph != nil {
		graphFeature.AddEvent(event)
	}

	if traceFeature, ok := o.features["trace"].(*TraceFeature); ok && traceFeature.Store != nil {
		if err := traceFeature.Write(event); err != nil {
			log.Printf("[Trace Brain] Causal Context Error: %v", err)
		} else {
			lineage = fmt.Sprintf("EventID: %d", event.SeqID)
		}
	}

	// Calculate TCS (Telemetry Confidence Score)
	tcsScore := 1.0
	if tcsFeature, ok := o.features["tcs"].(*TCSFeature); ok && tcsFeature.Window != nil {
		tcsScore = tcsFeature.AddAndEvaluate(event)
		tcsFeature.EvaluateDegradation(tcsScore, ledgerFeature.Ledger)
	}

	// 3. The Strategic Brain (Phoenix Arbiter - L5.5) -> Strategic Decision
	authorized := false
	var targetState warden.SystemState
	var class warden.ActuationClass
	if arbiterFeature, ok := o.features["arbiter"].(*ArbiterFeature); ok && arbiterFeature.Arb != nil {
		targetState, class, authorized = arbiterFeature.Evaluate(score, tcsScore)
		if authorized {
			log.Printf("[Arbiter Brain] Strategic Decision: Target=%s, Cost-Aware Authorization=TRUE", targetState)
		}
	}

	// 4. The Swarm Brain (Phoenix Nexus - L7) -> Consensus
	// Mock Nexus consensus for now, assuming local node is authoritative.
	consensus := true

	// 5. The Advisory Brain (Cognition Engine - F7 / PhoenixMind) -> AI Insight
	if o.Mind != nil {
		audit := ""
		if realityFeature, ok := o.features["reality"].(*RealityFeature); ok {
			audit = realityFeature.ReadAudit()
		}

		// REFLECTIVE FEEDBACK: Include meta-learning flags in the audit
		o.Meta.mu.Lock()
		if o.Meta.LimitationFlags["HIGH_UNCERTAINTY_DETECTED"] {
			audit += "\n[META-LEARNING ALERT] High uncertainty detected in recent decisions. Advisory confidence downgraded."
		}
		o.Meta.mu.Unlock()

		// Send context (Signal, Causal Context, Strategic Decision) to Ollama for Insight
		o.Mind.RequestAdvice(score, tcsScore, audit, &o.Wg)
	}

	// 6. Action (Warden)
	isDegraded := false
	if tcsFeature, ok := o.features["tcs"].(*TCSFeature); ok && tcsFeature.DegMon != nil {
		isDegraded = tcsFeature.DegMon.IsDegraded()
	}

	if authorized && !isDegraded && consensus {
		if wardenFeature, ok := o.features["warden"].(*WardenFeature); ok && wardenFeature.Warden != nil {
			wardenFeature.Actuate(targetState, class, tcsScore, event.SeqID, event.WallTimeUnix, logicalTick, ledgerFeature.Ledger)
			_ = lineage // lineage is captured in the trace DB

			// FEEDBACK LOOP: Record outcome back to meta-learning
			// In a real system, this would be updated after the next tick's SDI change
			o.Meta.RecordPerformance(tcsScore, 1.0) // Mocking success for now
		}
	}

}
