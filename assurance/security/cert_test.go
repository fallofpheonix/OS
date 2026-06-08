/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 */
package security

import (
	"testing"

	ledger "github.com/fallofpheonix/phoenix/foundation/ledger/src"
	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
	"github.com/fallofpheonix/phoenix/foundation/runtime/common/resource"
)

func TestCertificateInvariant(t *testing.T) {
	alloc := resource.NewBoundedAllocator(1024, 1000)
	l := ledger.NewLedger(alloc)

	eventID := "TEST-EVENT"
	l.AddEntry(eventID, "CAUSE", 1, []byte("payload"))

	weight := phxmath.NewFixedPointRaw(900000)
	cert, err := l.GenerateCertificate(eventID, weight)
	if err != nil {
		t.Fatalf("GenerateCertificate failed: %v", err)
	}

	inv := &CertificateInvariant{Validator: l}
	req := AuthorityEscalationRequest{
		EventID:        eventID,
		EvidenceWeight: weight,
		Certificate:    cert,
	}

	snap := PostureSnapshot{State: StateSafe}

	if err := inv.Verify(req, snap); err != nil {
		t.Errorf("Certificate verification failed: %v", err)
	}
}
