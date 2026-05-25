package unit

import (
	"testing"
)

func TestGovernanceValidation(t *testing.T) {
	root := "/Users/fallofpheonix/os"

	// Test Roadmap Integrity
	missing, err := ValidateRoadmapIntegrity(root)
	if err != nil {
		t.Fatalf("Validation failed: %v", err)
	}
	if len(missing) > 0 {
		t.Errorf("Missing subsystems: %v", missing)
	}

	// Test Axioms
	if err := CheckMandatoryAxioms(root); err != nil {
		t.Errorf("Axiom check failed: %v", err)
	}
}
