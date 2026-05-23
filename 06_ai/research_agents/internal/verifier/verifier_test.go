package verifier

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBoundaryRules(t *testing.T) {
	// Root is /Users/fallofpheonix/os, which is 3 directories up from agents/internal/verifier
	rootPath, err := filepath.Abs("../../../")
	if err != nil {
		t.Fatalf("Failed to get absolute path of workspace root: %v", err)
	}

	// Verify that we are indeed in the workspace root by checking for go.work or similar
	if _, err := os.Stat(filepath.Join(rootPath, "go.work")); err != nil {
		t.Logf("Warning: go.work not found at %s, fallback to current dir as scan root", rootPath)
		rootPath, err = filepath.Abs(".")
		if err != nil {
			t.Fatalf("Failed to get current dir path: %v", err)
		}
	}

	t.Logf("Scanning workspace path for architectural boundary violations: %s", rootPath)

	violations, err := VerifyImports(rootPath)
	if err != nil {
		t.Fatalf("Error scanning imports: %v", err)
	}

	if len(violations) > 0 {
		t.Errorf("Found %d architectural boundary violations:", len(violations))
		for i, v := range violations {
			t.Errorf("Violation %d:\n  File: %s\n  Import: %s\n  Rule: %s\n", i+1, v.FilePath, v.Import, v.Rule)
		}
	} else {
		t.Log("All architectural boundary checks passed! No import violations found.")
	}
}
