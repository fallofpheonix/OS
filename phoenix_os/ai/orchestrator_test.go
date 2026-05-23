package ai

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"phoenix/monitor"
)

func TestPromptBuilder(t *testing.T) {
	pm := NewPhoenixMind("http://localhost:11434/api/generate", "llama3.2:3b")
	score := monitor.DriftScore{
		EventType: "process.exec",
		PID:       1234,
		UID:       0,
		Severity:  0.9,
		ZScore:    4.5,
		Frequency: 0.001,
	}
	tcsScore := 0.95

	prompt := pm.BuildPrompt(score, tcsScore)
	if prompt == "" {
		t.Fatal("Prompt should not be empty")
	}
	
	expected := []string{"PhoenixMind", "process.exec", "1234", "0.90", "4.50", "0.95"}
	for _, s := range expected {
		if !strings.Contains(prompt, s) {
			t.Errorf("Prompt missing expected string: %s", s)
		}
	}
}

func TestGenerateAdvice_Integration(t *testing.T) {
	// Mock Ollama server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		advice := MindAdvice{
			Command:         "ISOLATE_PID",
			ConfidenceScore: 0.9,
			Reasoning:       "Highly anomalous root process execution.",
		}
		adviceJSON, _ := json.Marshal(advice)
		
		resp := map[string]interface{}{
			"response": string(adviceJSON),
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	pm := NewPhoenixMind(server.URL, "llama3.2:3b")
	pm.Client = server.Client() // Use the test server's client

	score := monitor.DriftScore{
		EventType: "process.exec",
		PID:       1234,
		UID:       0,
		Severity:  0.9,
		ZScore:    5.0,
		Frequency: 0.001,
	}
	
	advice, err := pm.GenerateAdvice(score, 1.0)
	if err != nil {
		t.Fatalf("GenerateAdvice failed: %v", err)
	}

	if advice.Command != "ISOLATE_PID" {
		t.Errorf("Expected command ISOLATE_PID, got %s", advice.Command)
	}
	if advice.ConfidenceScore != 0.9 {
		t.Errorf("Expected confidence 0.9, got %.2f", advice.ConfidenceScore)
	}
	if advice.Reasoning == "" {
		t.Error("Expected reasoning to be non-empty")
	}
}
