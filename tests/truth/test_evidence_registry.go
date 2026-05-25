package truth_test

import (
	"testing"
	"github.com/fallofpheonix/os/phoenix_os/truth/evidence"
)

func TestEvidenceRegistry(t *testing.T) {
	registry := evidence.NewEvidenceRegistry()
	
	e := &evidence.Evidence{
		Entity:     "arbiter",
		Status:     evidence.VALIDATED,
		Confidence: 0.91,
	}

	registry.Add(e)

	got, ok := registry.Get("arbiter")
	if !ok {
		t.Fatal("expected evidence to be found")
	}

	if got.Status != evidence.VALIDATED {
		t.Errorf("expected status %s, got %s", evidence.VALIDATED, got.Status)
	}
}
