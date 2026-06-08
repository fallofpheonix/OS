package consensus

import (
	"testing"

	"github.com/fallofpheonix/phoenix/internal/contracts"
	"github.com/fallofpheonix/phoenix/internal/protocol"
)

func TestEvent_CanonicalDigest(t *testing.T) {
	e := contracts.Event{
		Version: 1,
		Type:    EventMove,
		Payload: []byte{0xAA, 0xBB},
	}

	digest, err := protocol.DigestEvent(e)
	if err != nil {
		t.Fatalf("DigestEvent failed: %v", err)
	}
	got := digest.String()

	// Golden Hash (PROTOCOL-015: 01000100010002AABB)
	want := "6f413efab66c6735513bb8e089ce2cdd2b87839f1d64de50bf3910fe645fe278"

	if got != want {
		t.Errorf("Canonical Digest Mismatch!\nGot:  %s\nWant: %s", got, want)
	}
}
