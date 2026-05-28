package truth_test

import (
	"testing"
	"github.com/fallofpheonix/phoenixmind-validator/truth/evidence"
)

func TestEvidenceRegistry(t *testing.T) {
	registry := evidence.NewEvidenceRegistry()
	
	e := &evidence.Evidence{
		EntityID: "arbiter",
		State:    evidence.VALIDATED,
	}

	registry.Add(e)

	got, ok := registry.Get("arbiter")
	if !ok {
		t.Fatal("expected evidence to be found")
	}

	if got.State != evidence.VALIDATED {
		t.Errorf("expected state %s, got %s", evidence.VALIDATED, got.State)
	}
}
