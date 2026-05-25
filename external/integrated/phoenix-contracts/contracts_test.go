package contracts

import (
	"testing"
)

// TestVersionFormatting validates semantic version string generation.
func TestVersionFormatting(t *testing.T) {
	v := Version{Major: 1, Minor: 2, Patch: 3}
	if v.String() != "1.2.3" {
		t.Errorf("Expected 1.2.3, got %s", v.String())
	}

	v2 := Version{Major: 0, Minor: 5, Patch: 0, PreRelease: "alpha"}
	if v2.String() != "0.5.0-alpha" {
		t.Errorf("Expected 0.5.0-alpha, got %s", v2.String())
	}
}

// TestCompatibilityCheck validates version compatibility rules.
func TestCompatibilityCheck(t *testing.T) {
	vMatch := Version{Major: 0, Minor: 2, Patch: 0}
	if !CompatibilityCheck(vMatch) {
		t.Errorf("Expected Major version 0 to be compatible")
	}

	vMismatch := Version{Major: 1, Minor: 0, Patch: 0}
	if CompatibilityCheck(vMismatch) {
		t.Errorf("Expected Major version 1 to be incompatible with 0")
	}
}
