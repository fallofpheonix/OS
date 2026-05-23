package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"phoenix/bus"
	"phoenix/monitor"
	"phoenix/warden"
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
}

// BuildPrompt formats the telemetry signals into a natural language prompt for PhoenixMind.
func (pm *PhoenixMind) BuildPrompt(score monitor.DriftScore, tcsScore float64) string {
	return fmt.Sprintf(`You are PhoenixMind, the AI advisor for PhoenixOS.
Analyze the following telemetry signals and output a JSON response.

Signals:
- Event: %s (PID: %d, UID: %d)
- Entropy/Severity: %.2f
- Z-Score: %.2f (Threshold: 3.0)
- TCS Confidence: %.2f
- Frequency ($F_A$): %.4f

Output Schema:
{
  "command": "ONE_OF: [ISOLATE_PID, THROTTLE_NETWORK, REVOKE_UID, LOG_ONLY]",
  "confidence_score": 0.0-1.0,
  "reasoning": "Brief technical justification"
}

JSON Response:`, score.EventType, score.PID, score.UID, score.Severity, score.ZScore, tcsScore, score.Frequency)
}

// GenerateAdvice calls the local Ollama LLM and parses the advisory output.
func (pm *PhoenixMind) GenerateAdvice(score monitor.DriftScore, tcsScore float64) (MindAdvice, error) {
	prompt := pm.BuildPrompt(score, tcsScore)

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
func NewPhoenixMind(url, model string) *PhoenixMind {
	pm := &PhoenixMind{
		OllamaURL: url,
		Model:     model,
		Client:    &http.Client{Timeout: 10 * time.Second},
		requests:  make(chan adviceRequest, 100),
	}
	go pm.runAdvisor()
	return pm
}

// runAdvisor is a dedicated go-routine that aggregates signals and calls the LLM.
// It implements 5-second batching to respect resource limits (Step 7 & 8).
func (pm *PhoenixMind) runAdvisor() {
	ticker := time.NewTicker(5 * time.Second)
	var latest *adviceRequest

	for {
		select {
		case req := <-pm.requests:
			// Aggregation strategy: in the 5s window, prioritize the highest Z-Score (most chaotic).
			if latest == nil || req.score.ZScore > latest.score.ZScore {
				latest = &req
			}
		case <-ticker.C:
			if latest != nil {
				advice, err := pm.GenerateAdvice(latest.score, latest.tcsScore)
				if err != nil {
					log.Printf("[AI Orchestrator][PhoenixMind Error] %v", err)
				} else {
					log.Printf("[AI Orchestrator][PhoenixMind Advice] Seq: %d, Command: %s, Confidence: %.2f, Reasoning: %s",
						latest.score.EventID, advice.Command, advice.ConfidenceScore, advice.Reasoning)
				}
				latest = nil
			}
		}
	}
}

// RequestAdvice queues a new signal for AI evaluation.
func (pm *PhoenixMind) RequestAdvice(score monitor.DriftScore, tcsScore float64) {
	select {
	case pm.requests <- adviceRequest{score, tcsScore}:
	default:
		// Drop request if channel is full to prevent backpressure on telemetry fast-path.
	}
}

// AIOrchestrator is the central coordinator directing all system subsystems as features.
type AIOrchestrator struct {
	features map[string]Feature
	Mind     *PhoenixMind
}

// NewAIOrchestrator creates a new instance of the orchestrator.
func NewAIOrchestrator() *AIOrchestrator {
	return &AIOrchestrator{
		features: make(map[string]Feature),
		Mind:     NewPhoenixMind("http://localhost:11434/api/generate", "llama3.2:3b"),
	}
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
	// In an AI-First architecture, the AI Orchestrator controls the lifecycle and sequence of all features.

	// 1. Resolve Ledger Feature first (essential for recording system state changes)
	ledgerFeature, ok := o.features["ledger"].(*LedgerFeature)
	if !ok || ledgerFeature.Ledger == nil {
		log.Fatal("[AI Orchestrator] Ledger feature not registered")
	}

	// 2. Trace Feature: Ingest event into the process lineage DAG
	if traceFeature, ok := o.features["trace"].(*TraceFeature); ok && traceFeature.Store != nil {
		if err := traceFeature.Write(event); err != nil {
			log.Printf("[AI Orchestrator][Trace Feature Error] %v", err)
		}
	}

	// 3. Monitor Feature: Perform mathematical telemetry/signal analysis (L3)
	var score monitor.DriftScore
	if monitorFeature, ok := o.features["monitor"].(*MonitorFeature); ok && monitorFeature.Service != nil {
		score = monitorFeature.Process(event)
	}

	// 4. TCS Feature: Add telemetry and evaluate system confidence/degradation boundaries
	tcsScore := 1.0
	if tcsFeature, ok := o.features["tcs"].(*TCSFeature); ok && tcsFeature.Window != nil {
		tcsScore = tcsFeature.AddAndEvaluate(event)
		tcsFeature.EvaluateDegradation(tcsScore, ledgerFeature.Ledger)
	}

	// 5. Arbiter Feature: Evaluate game-theoretic security policies
	authorized := false
	var targetState warden.SystemState
	var class warden.ActuationClass
	if arbiterFeature, ok := o.features["arbiter"].(*ArbiterFeature); ok && arbiterFeature.Arb != nil {
		targetState, class, authorized = arbiterFeature.Evaluate(score, tcsScore)
	}

	// 5.5 PhoenixMind Advisory: Trigger AI evaluation on chaotic signals (Z > 3.0)
	if score.ZScore > 3.0 && o.Mind != nil {
		o.Mind.RequestAdvice(score, tcsScore)
	}

	// 6. Warden Feature: Perform tactical state machine actuation
	isDegraded := false
	if tcsFeature, ok := o.features["tcs"].(*TCSFeature); ok && tcsFeature.DegMon != nil {
		isDegraded = tcsFeature.DegMon.IsDegraded()
	}

	if authorized && !isDegraded {
		if wardenFeature, ok := o.features["warden"].(*WardenFeature); ok && wardenFeature.Warden != nil {
			wardenFeature.Actuate(targetState, class, tcsScore, event.SeqID, event.WallTimeUnix, logicalTick, ledgerFeature.Ledger)
		}
	}
}
