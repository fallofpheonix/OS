package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestPhoenixMindReasoning(t *testing.T) {
	// 1. Prepare telemetry data
	telemetry := `{"Source": "kernel.ebpf", "Entropy": 8.5, "Action": "unexpected_syscall", "Process": "unauthorized_exec"}`
	prompt := fmt.Sprintf("Analyze this telemetry and provide a security verdict (Normal/Suspicious/Critical): %s", telemetry)

	// 2. Prepare payload for Ollama
	payload := map[string]interface{}{
		"model":  "phoenix-mind",
		"prompt": prompt,
		"stream": false,
	}
	data, _ := json.Marshal(payload)

	// 3. Call local LLM
	resp, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewBuffer(data))
	if err != nil {
		t.Fatalf("Failed to connect to local LLM: %v", err)
	}
	defer resp.Body.Close()

	// 4. Parse response
	var result struct {
		Response string `json:"response"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("Failed to decode LLM response: %v", err)
	}

	// 5. Assert reasoning capability
	fmt.Printf("[AI TEST] LLM Verdict: %s\n", result.Response)
	if !strings.Contains(strings.ToUpper(result.Response), "SUSPICIOUS") && !strings.Contains(strings.ToUpper(result.Response), "CRITICAL") {
		t.Errorf("LLM failed to identify the threat in the telemetry. Response: %s", result.Response)
	}
}
