/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 *
 * [PRAS HEADER]
 * Purpose: Communication bridge between PhoenixOS and the G0DM0D3 Oracle (L7 AI model) for high-level directives.
 * Subsystem: Cognition Intelligence
 * Dependencies: net/http, encoding/json, regexp, bus
 * Dependents: AIOrchestrator
 * Security Considerations: High. Handles sensitive directives that control system state. Protect API keys and validate Oracle responses.
 * Performance Considerations: Subject to network and LLM latency. Uses a 60-second timeout.
 * Labels: api-bridge, external-dependency, security-critical
 */
package intelligence

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/fallofpheonix/phoenix/cognition/reasoning"
	"github.com/fallofpheonix/phoenix/foundation/runtime/bus"
)

// Directive represents the actionable command from the G0DM0D3 Oracle.
type Directive struct {
	Command         string   `json:"command"`
	ConfidenceScore float64  `json:"confidence_score"`
	Reasoning       string   `json:"reasoning"`
	GraphProof      []string `json:"graph_proof"`
}

// jsonExtractRE extracts JSON blocks from LLM responses that may be wrapped in markdown.
var jsonExtractRE = regexp.MustCompile(`(?s)\{.*\}`)

// NexusBridge handles communication between the PhoenixOS Master and the G0DM0D3 Oracle.
type NexusBridge struct {
	OracleURL string
	APIKey    string
	Client    *http.Client
}

/*
 * [FUNCTION HEADER]
 * Purpose: Initializes a new NexusBridge with the specified configuration.
 * Responsibilities: Setup HTTP client with a 60-second timeout.
 * Inputs: url (string), apiKey (string)
 * Outputs: *NexusBridge
 * Complexity: O(1)
 */
// NewNexusBridge creates a new instance of the bridge.
func NewNexusBridge(url, apiKey string) *NexusBridge {
	return &NexusBridge{
		OracleURL: url,
		APIKey:    apiKey,
		Client:    &http.Client{Timeout: 60 * time.Second},
	}
}

// Name returns the provider name.
func (n *NexusBridge) Name() string {
	return "G0DM0D3-Nexus-Bridge"
}

// Reason implements the reasoning.Provider interface.
func (n *NexusBridge) Reason(ctx context.Context, req *reasoning.InferenceRequest) (*reasoning.InferenceResponse, error) {
	log.Printf("[Nexus Provider] Reasoning for goal: %s", req.Goal)

	// Bridging InferenceRequest to the G0DM0D3 Oracle format
	// In a real implementation, we'd marshal Beliefs and Knowledge into the prompt.
	prompt := fmt.Sprintf("GOAL: %s\nBELIEFS: %d\nKNOWLEDGE_EDGES: %d", req.Goal, len(req.Beliefs), len(req.Knowledge))
	log.Printf("[DEBUG] Shadow Reason prompt construction: %s", prompt)

	// For shadow mode validation, we'll return a simulated response grounded in the request.
	return &reasoning.InferenceResponse{
		Logic:      fmt.Sprintf("Analyzed %d beliefs and %d knowledge edges for goal: %s", len(req.Beliefs), len(req.Knowledge), req.Goal),
		Proposal:   "PROCEED_WITH_CAUTION",
		Confidence: 0.85,
	}, nil
}

/*
 * [FUNCTION HEADER]
 * Purpose: Queries the G0DM0D3 Oracle for a strategic directive based on system state.
 * Responsibilities: Construct prompt, perform HTTP POST to Oracle API, and parse JSON response into a Directive.
 * Inputs: event (bus.TelemetryEvent), tcsScore (float64), audit (string), graphContext (string)
 * Outputs: Directive, error
 * Complexity: O(Network_Latency + LLM_Inference_Time)
 */
// RequestDirective queries the G0DM0D3 ULTRAPLINIAN API for a strategic decision.
func (n *NexusBridge) RequestDirective(event bus.TelemetryEvent, tcsScore float64, audit, graphContext string) (Directive, error) {
	log.Printf("[Nexus Bridge] Requesting Directive from G0DM0D3 Oracle...")

	// Build the prompt for ULTRAPLINIAN racing
	prompt := fmt.Sprintf(`Analyze this PhoenixOS anomaly:
Event: %s (Severity: %.2f)
TCS Confidence: %.2f
System Audit: %s
Causal Context: %s

You must justify your action using a causal path (GraphProof) derived from the Causal Context.
Your directive will be verified against formal invariants.

Output JSON Directive:
{
  "command": "[ISOLATE_PID | THROTTLE_NETWORK | REVOKE_UID | LOG_ONLY]",
  "confidence_score": 0.0-1.0,
  "reasoning": "Technical justification.",
  "graph_proof": ["node1", "node2", "nodeN"]
}`, event.EventType, event.Severity, tcsScore, audit, graphContext)

	// G0DM0D3 API follows OpenAI format but optimized for ULTRAPLINIAN
	requestBody, _ := json.Marshal(map[string]interface{}{
		"model": "ultraplinian/standard",
		"messages": []map[string]string{
			{"role": "system", "content": "You are the Phoenix Nexus L7 Oracle. Provide deterministic security directives anchored to causal lineage."},
			{"role": "user", "content": prompt},
		},
		"stream": false,
	})

	req, err := http.NewRequest("POST", n.OracleURL+"/v1/chat/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		return Directive{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	if n.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+n.APIKey)
	}

	resp, err := n.Client.Do(req)
	if err != nil {
		return Directive{}, fmt.Errorf("Oracle connection failed: %w", err)
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			log.Printf("[Nexus Bridge] Error closing response body: %v", closeErr)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Directive{}, err
	}

	// Parse G0DM0D3 / OpenAI compatible response
	var g0dm0d3Resp struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.Unmarshal(body, &g0dm0d3Resp); err != nil {
		return Directive{}, fmt.Errorf("failed to parse Oracle response: %w", err)
	}

	if len(g0dm0d3Resp.Choices) == 0 {
		return Directive{}, fmt.Errorf("Oracle returned no choices")
	}

	// The Oracle might return JSON inside its content string, sometimes wrapped in markdown
	content := g0dm0d3Resp.Choices[0].Message.Content

	// Try to find JSON block if it exists
	match := jsonExtractRE.FindString(content)
	if match != "" {
		content = match
	}

	var directive Directive
	if err := json.Unmarshal([]byte(content), &directive); err != nil {
		return Directive{}, fmt.Errorf("failed to parse Directive JSON from Oracle (Content: %s): %w", content, err)
	}

	return directive, nil
}
