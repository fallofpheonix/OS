package reasoning

import (
	"strings"
	"testing"
)

func TestExplainer_Basic(t *testing.T) {
	e := NewExplainer("DEC_001")

	e.AddLink("Sensory", "SDI_ZScore", 8.5, 0.0, 0.4)
	e.AddLink("Temporal", "TCS_Score", 0.95, 1.0, 0.3)

	pathStr := e.Path.String()
	if !strings.Contains(pathStr, "Sensory") || !strings.Contains(pathStr, "Temporal") {
		t.Errorf("path string missing expected nodes: %s", pathStr)
	}

	if !strings.Contains(pathStr, "Dev: +10000.0%") {
		t.Errorf("deviation calculation for 0.0 expected expected +10000.0%%, got: %s", pathStr)
	}
}

func TestExplainer_Counterfactual(t *testing.T) {
	e := NewExplainer("DEC_002")
	e.AddLink("S1", "Signal1", 1.0, 0.0, 0.5)
	e.AddLink("S2", "Signal2", 1.0, 0.0, 0.5)

	// Total signal = (1.0 * 0.5) + (1.0 * 0.5) = 1.0
	// Counterfactual: S1 = 0.0 -> Total = (0.0 * 0.5) + (1.0 * 0.5) = 0.5

	cf := e.GenerateCounterfactual("S1", 0.0)
	if !strings.Contains(cf, "would be 0.50") {
		t.Errorf("unexpected counterfactual result: %s", cf)
	}

	if !strings.Contains(cf, "Decision change: true") {
		t.Errorf("expected decision change for 0.50 (< 0.8), got: %s", cf)
	}
}

func TestExplainer_ExplainIgnored(t *testing.T) {
	e := NewExplainer("DEC_003")
	e.AddLink("S1", "Signal1", 0.1, 0.0, 0.5)

	ignored := e.ExplainIgnored("Low confidence")
	if !strings.Contains(ignored, "IGNORE") || !strings.Contains(ignored, "Low confidence") {
		t.Errorf("unexpected ignore explanation: %s", ignored)
	}
}
