package reflection

import (
	"testing"
	"github.com/fallofpheonix/phoenix/foundation/ledger"
)

func TestEngine_PredictAndVerify(t *testing.T) {
	e := NewEngine()

	pID := "P1"
	fID := "F1"
	expected := []byte("sunny")
	actual := []byte("sunny")

	e.Predict(pID, "ACTION_001", expected, 1000)

	err := e.Verify(pID, fID, actual)
	if err == nil {
		t.Fatal("expected reflection error, got nil")
	}

	if err.Divergence != 0.0 {
		t.Errorf("expected 0.0 divergence for identical payloads, got %f", err.Divergence)
	}

	// Test mismatch
	pID2 := "P2"
	actual2 := []byte("cloudy")
	e.Predict(pID2, "ACTION_002", expected, 1001)

	err2 := e.Verify(pID2, fID, actual2)
	if err2.Divergence == 0.0 {
		t.Error("expected non-zero divergence for differing payloads")
	}
}

func TestRealityDriftAuditor_Quarantine(t *testing.T) {
	auditor := NewRealityDriftAuditor(0.5)

	// Case 1: Low drift
	auditor.RecordError(&ReflectionError{Divergence: 0.1})
	if auditor.IsQuarantined {
		t.Error("system quarantined incorrectly at low drift")
	}

	// Case 2: High drift triggers quarantine
	auditor.RecordError(&ReflectionError{Divergence: 0.9})
	// Average = (0.1 + 0.9) / 2 = 0.5 (Not > 0.5)
	if auditor.IsQuarantined {
		t.Error("system quarantined incorrectly at boundary")
	}

	auditor.RecordError(&ReflectionError{Divergence: 1.0})
	// Average = (0.1 + 0.9 + 1.0) / 3 = 0.66 (> 0.5)
	if !auditor.IsQuarantined {
		t.Error("system failed to quarantine at high drift")
	}

	// Case 3: Reset
	auditor.Reset()
	if auditor.IsQuarantined || auditor.Cumulative != 0 {
		t.Error("auditor failed to reset")
	}
}

func TestCalculateDivergence(t *testing.T) {
	tests := []struct {
		name     string
		expected []byte
		actual   []byte
		want     float64
	}{
		{"identical", []byte("abc"), []byte("abc"), 0.0},
		{"totally different", []byte("abc"), []byte("xyz"), 1.0},
		{"partial match", []byte("abc"), []byte("abz"), 1.0 / 3.0},
		{"empty both", []byte(""), []byte(""), 0.0},
		{"one empty", []byte("abc"), []byte(""), 1.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateDivergence(tt.expected, tt.actual)
			if got != tt.want {
				t.Errorf("calculateDivergence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccuracyMetric_Operations(t *testing.T) {
	am := &AccuracyMetric{}

	if am.CalculateAccuracy() != 1.0 {
		t.Errorf("expected 1.0 accuracy for 0 predictions, got %f", am.CalculateAccuracy())
	}

	am.TotalPredictions = 10
	am.CorrectCount = 5
	if am.CalculateAccuracy() != 0.5 {
		t.Errorf("expected 0.5 accuracy, got %f", am.CalculateAccuracy())
	}

	// Test Reconstruction
	events := []*ledger.Event{
		{Type: ledger.EventPrediction},
		{Type: ledger.EventPrediction},
		{Type: ledger.EventFact}, // Ignored
	}
	
	am2 := &AccuracyMetric{}
	am2.ReconstructFromLedger(events)
	if am2.TotalPredictions != 2 {
		t.Errorf("expected 2 total predictions after reconstruction, got %d", am2.TotalPredictions)
	}
}

