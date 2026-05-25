package replay

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
)

// CalculateHash computes the SHA-256 hash of a TelemetryEvent.
// It includes the prev_hash to form a cryptographic chain.
func CalculateHash(ev bus.TelemetryEvent) string {
	h := sha256.New()

	// Monotonic fields first
	binary.Write(h, binary.BigEndian, ev.SeqID)
	binary.Write(h, binary.BigEndian, ev.MonotonicNs)

	// Identity fields
	h.Write([]byte(ev.Source))
	h.Write([]byte(ev.HostID))
	binary.Write(h, binary.BigEndian, int64(ev.PID))
	binary.Write(h, binary.BigEndian, int64(ev.TID))
	binary.Write(h, binary.BigEndian, int64(ev.UID))
	binary.Write(h, binary.BigEndian, int64(ev.GID))

	// Event specific
	h.Write([]byte(ev.EventType))
	h.Write(ev.Payload)

	// The Chain Link
	h.Write([]byte(ev.PrevHash))

	return hex.EncodeToString(h.Sum(nil))
}

// VerifyChain checks if a sequence of events forms a valid hash chain.
func VerifyChain(events []bus.TelemetryEvent) error {
	if len(events) == 0 {
		return nil
	}

	for i, ev := range events {
		// 1. Verify self-hash
		computedHash := CalculateHash(ev)
		if computedHash != ev.Hash {
			return fmt.Errorf("invalid hash at sequence %d: expected %s, got %s", ev.SeqID, ev.Hash, computedHash)
		}

		// 2. Verify link to previous
		if i > 0 {
			if ev.PrevHash != events[i-1].Hash {
				return fmt.Errorf("hash chain broken at sequence %d: prev_hash %s != previous event hash %s",
					ev.SeqID, ev.PrevHash, events[i-1].Hash)
			}
		}
	}

	return nil
}

// SignEvent computes and sets the Hash for an event, given the previous event's hash.
func SignEvent(ev *bus.TelemetryEvent, prevHash string) {
	ev.PrevHash = prevHash
	ev.Hash = CalculateHash(*ev)
}
