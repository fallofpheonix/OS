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

	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
)

// Directive represents the actionable command from the G0DM0D3 Oracle.
type Directive struct {
	Command         string   `json:"command"`
	ConfidenceScore float64  `json:"confidence_score"`
	Reasoning       string   `json:"reasoning"`
	GraphProof      []string `json:"graph_proof"`
}

// NexusBridge handles communication between the PhoenixOS Master and the G0DM0D3 Oracle.
type NexusBridge struct {
	OracleURL string
	APIKey    string
	Client    *http.Client
}

// NewNexusBridge creates a new instance of the bridge.
func NewNexusBridge(url string, apiKey string) *NexusBridge {
	return &NexusBridge{
		OracleURL: url,
		APIKey:    apiKey,
		Client:    &http.Client{Timeout: 60 * time.Second},
	}
}

// RequestDirective queries the G0DM0D3 ULTRAPLINIAN API for a strategic decision.
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
