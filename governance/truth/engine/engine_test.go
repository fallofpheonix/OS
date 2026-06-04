/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package engine_test

import (
	"testing"

	ledgerpb "github.com/fallofpheonix/phoenix/foundation/runtime/proto/v1/ledger"
	"github.com/fallofpheonix/phoenix/governance/truth/engine"
)

func TestTruthAssessment(t *testing.T) {
	eval := engine.NewEvaluator(0.85)

	// High confidence evidence
	e1 := &ledgerpb.EvidenceRecord{
		EvidenceId: "e1",
		TrustScore: 0.9,
		SourceRef:  "sensor_a",
	}

	// Low confidence evidence
	e2 := &ledgerpb.EvidenceRecord{
		EvidenceId: "e2",
		TrustScore: 0.4,
		SourceRef:  "sensor_b",
	}

	eval.IngestEvidence(e1)
	eval.IngestEvidence(e2)

	// Test 1: Verified state
	assessment, err := eval.Assess("a1", []string{"e1"})
	if err != nil {
		t.Fatalf("Assessment failed: %v", err)
	}
	if assessment.Conclusion != "VERIFIED" {
		t.Errorf("Expected VERIFIED, got %s", assessment.Conclusion)
	}

	// Test 2: Insufficient evidence (average < 0.85)
	assessment2, err := eval.Assess("a2", []string{"e1", "e2"})
	if err != nil {
		t.Fatalf("Assessment failed: %v", err)
	}
	if assessment2.Conclusion != "INSUFFICIENT_EVIDENCE" {
		t.Errorf("Expected INSUFFICIENT_EVIDENCE, got %s", assessment2.Conclusion)
	}
}

func TestContradictionDetection(t *testing.T) {
	e1 := &ledgerpb.EvidenceRecord{
		EvidenceId: "e1",
		TrustScore: 0.9,
		SourceRef:  "identity_check",
	}

	e2 := &ledgerpb.EvidenceRecord{
		EvidenceId: "e2",
		TrustScore: 0.2, // HIGH DRIFT
		SourceRef:  "identity_check",
	}

	contradiction := engine.DetectContradiction(e1, e2)
	if contradiction == nil {
		t.Fatal("Expected contradiction, got nil")
	}
	if contradiction.Severity != "HIGH" {
		t.Errorf("Expected HIGH severity, got %s", contradiction.Severity)
	}
}
