package observations

import (
	"testing"
)

func TestObservationIngest(t *testing.T) {
	obs, err := Ingest("OBS-001")
	if err != nil {
		t.Fatalf("Ingest failed: %v", err)
	}
	if obs.Cycle != "OBS-001" {
		t.Errorf("Expected cycle OBS-001, got %s", obs.Cycle)
	}
	if obs.Status != "STABLE" {
		t.Errorf("Expected status STABLE, got %s", obs.Status)
	}
}
