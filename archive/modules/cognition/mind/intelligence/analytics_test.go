package intelligence

import (
	"strings"
	"testing"
)

func TestAnalyticsHub_ProcessSimulatedActuation(t *testing.T) {
	ah := NewAnalyticsHub()

	evidence := &RichActuationEvidence{
		SignalTotal: 0.95,
		Confidence:  0.8,
	}

	ah.ProcessSimulatedActuation(evidence, "dec-1")

	if ah.TotalDecisions != 1 {
		t.Errorf("expected 1 total decision, got %d", ah.TotalDecisions)
	}

	outcome, ok := ah.Outcomes["dec-1"]
	if !ok || outcome.Label != LabelTruePositive {
		t.Errorf("expected TRUE_POSITIVE for high signal, got %s", outcome.Label)
	}

	// Test False Positive
	evidence2 := &RichActuationEvidence{
		SignalTotal: 0.5,
	}
	ah.ProcessSimulatedActuation(evidence2, "dec-2")

	outcome2 := ah.Outcomes["dec-2"]
	if outcome2.Label != LabelFalsePositive {
		t.Errorf("expected FALSE_POSITIVE for low signal, got %s", outcome2.Label)
	}

	// Verify Precision
	// 1 TP, 1 FP -> Precision = 0.5
	if ah.Precision != 0.5 {
		t.Errorf("expected precision 0.5, got %f", ah.Precision)
	}
}

func TestAnalyticsHub_ExportReport(t *testing.T) {
	ah := NewAnalyticsHub()
	ah.ProcessSimulatedActuation(&RichActuationEvidence{SignalTotal: 0.9}, "d1")

	report := ah.ExportReport()
	if !strings.Contains(report, "total_decisions") || !strings.Contains(report, "precision") {
		t.Errorf("report missing expected fields: %s", report)
	}
}
