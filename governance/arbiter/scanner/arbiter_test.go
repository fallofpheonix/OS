/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestEngine_Verify(t *testing.T) {
	engine := NewEngine()

	// 1. Test with a fully compliant mock file
	compliantContent := `
/**
 * FILE: mock.go
 * PURPOSE: Test
 * SUBSYSTEM: Tests
 * [STATUS: VERIFIED]
 */
package main
// BEGINNER EXPLANATION:
// EXPERT EXPLANATION:
`
	tmpCompliant := filepath.Join(t.TempDir(), "mock.go")
	os.WriteFile(tmpCompliant, []byte(compliantContent), 0o644)

	report, err := engine.Verify(tmpCompliant)
	if err != nil {
		t.Fatalf("verification failed: %v", err)
	}
	if report.CompletenessScore != 100.0 {
		t.Errorf("expected 100%% completeness, got %v (missing: %v)", report.CompletenessScore, report.MissingFields)
	}

	// 2. Test with a non-compliant file
	badContent := `package main`
	tmpBad := filepath.Join(t.TempDir(), "bad.go")
	os.WriteFile(tmpBad, []byte(badContent), 0o644)

	badReport, _ := engine.Verify(tmpBad)
	if badReport.CompletenessScore == 100.0 {
		t.Errorf("expected failure for non-compliant file")
	}
	if len(badReport.MissingFields) != 6 {
		t.Errorf("expected 6 missing fields, got %d", len(badReport.MissingFields))
	}
}
