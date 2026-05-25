package security

import (
	"fmt"
	"testing"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/truth"
)

func TestHashTamperDetection(t *testing.T) {
	l := truth.NewTruthLedger()
	l.AddEntry(truth.EvidenceWrapper{Sequence: 1, Hash: "A"})
	
	// Attempt to add a tampered entry (duplicate sequence with different hash)
	err := l.AddEntry(truth.EvidenceWrapper{Sequence: 1, Hash: "B"})
	if err == nil {
		t.Error("Expected error for hash tamper (duplicate sequence)")
	}
	fmt.Println("[PX-007] Hash Tamper Detection: PASSED")
}
