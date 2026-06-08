/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
/*
 * @file nexus_bridge.go
 * @package ai
 * @subsystem Terminus-AI
 *
 * @description Implements the NexusBridge, which interfaces PhoenixOS with the 
 * G0DM0D3 ULTRAPLINIAN Oracle for strategic AI-driven security directives.
 *
 * @status 18-Repository Substrate Consolidated
 * @future Needs HDF5 vector optimizations for causal graph representation.
 *
 * @dependencies
 * - github.com/fallofpheonix/phoenix/foundation/runtime/bus
 * - net/http
 * - encoding/json
 *
 * @dependents
 * - phoenix_os/ai/orchestrator.go
 *
 * @security
 * - Critical: Oracle directives can trigger system-wide isolations and revocations.
 * - API Key management (Authorization header).
 * - [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal.
 *
 * @performance
 * - Subject to network latency and Oracle processing time (60s timeout).
 * - [ERROR PRONE AREA]: Potential concurrency bottlenecks if many requests are queued.
 *
 * @labels ai-integration, networking, phase-2-complete
 */
package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/fallofpheonix/phoenix/foundation/runtime/bus"
)

/*
 * @struct Directive
 * @description Represents an actionable security command issued by the AI Oracle.
 * @fields Command, ConfidenceScore, Reasoning, GraphProof.
 */
type Directive struct {
	Command         string   `json:"command"`
	ConfidenceScore float64  `json:"confidence_score"`
	Reasoning       string   `json:"reasoning"`
	GraphProof      []string `json:"graph_proof"`
}

/*
 * @class NexusBridge
 * @description Facilitates the secure communication channel between PhoenixOS and the G0DM0D3 Oracle.
 * @responsibilities HTTP request management, prompt construction, and response parsing.
 */
type NexusBridge struct {
	OracleURL string
	APIKey    string
	Client    *http.Client
}

/*
 * @function NewNexusBridge
 * @description Constructor for NexusBridge.
 * @params {string} url - The URL of the Oracle API.
 * @params {string} apiKey - Optional API key for authentication.
 * @returns {*NexusBridge}
 * @complexity O(1)
 */
func NewNexusBridge(url string, apiKey string) *NexusBridge {
	return &NexusBridge{
		OracleURL: url,
		APIKey:    apiKey,
		Client:    &http.Client{Timeout: 60 * time.Second},
	}
}

/*
 * @function RequestDirective
 * @memberof NexusBridge
 * @description Queries the Oracle for a strategic security decision based on system telemetry and causal context.
 * @params {bus.TelemetryEvent} event - The anomaly event.
 * @params {float64} tcsScore - Telemetry Confidence Score.
 * @params {string} audit - Runtime reality audit data.
 * @params {string} graphContext - Serialized causal graph context.
 * @returns {Directive} The parsed directive from the Oracle.
 * @returns {error}
 * @complexity O(N) where N is the size of the prompt and response.
 */
func (n *NexusBridge) RequestDirective(event bus.TelemetryEvent, tcsScore float64, audit string, graphContext string) (Directive, error) {
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
	defer resp.Body.Close()

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
	re := regexp.MustCompile("(?s)\\{.*\\}")
	match := re.FindString(content)
	if match != "" {
		content = match
	}

	var directive Directive
	if err := json.Unmarshal([]byte(content), &directive); err != nil {
		return Directive{}, fmt.Errorf("failed to parse Directive JSON from Oracle (Content: %s): %w", content, err)
	}

	return directive, nil
}
