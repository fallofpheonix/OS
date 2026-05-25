package truth

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sync"
	"time"
)

// EvidenceWrapper encapsulates payload with hash chain metadata.
type EvidenceWrapper struct {
	ID        string
	Sequence  uint64
	Timestamp time.Time
	Payload   interface{}
	PrevHash  string
	Hash      string
}

// TruthLedger implements the immutable hash chain evidence store.
type TruthLedger struct {
	mu       sync.RWMutex
	Entries  []EvidenceWrapper
	LastHash string
	seq      uint64
}

// NewTruthLedger creates a fresh evidence chain.
func NewTruthLedger() *TruthLedger {
	return &TruthLedger{
		Entries:  make([]EvidenceWrapper, 0),
		LastHash: "genesis",
		seq:      0,
	}
}

// Append adds a new evidence entry and calculates the hash chain.
func (l *TruthLedger) Append(payload interface{}) (string, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.seq++
	ts := time.Now()

	// Create hash
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%d-%d-%s-%v", l.seq, ts.UnixNano(), l.LastHash, payload)))
	hashStr := hex.EncodeToString(h.Sum(nil))

	entry := EvidenceWrapper{
		ID:        fmt.Sprintf("ev_%d", l.seq),
		Sequence:  l.seq,
		Timestamp: ts,
		Payload:   payload,
		PrevHash:  l.LastHash,
		Hash:      hashStr,
	}

	l.Entries = append(l.Entries, entry)
	l.LastHash = hashStr
	return hashStr, nil
}

// Verify performs a full integrity pass on the hash chain.
func (l *TruthLedger) Verify() (bool, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	expectedPrevHash := "genesis"
	for i, entry := range l.Entries {
		if entry.PrevHash != expectedPrevHash {
			return false, fmt.Errorf("integrity break at sequence %d: hash mismatch", entry.Sequence)
		}

		// Re-calculate hash
		h := sha256.New()
		h.Write([]byte(fmt.Sprintf("%d-%d-%s-%v", entry.Sequence, entry.Timestamp.UnixNano(), entry.PrevHash, entry.Payload)))
		actualHash := hex.EncodeToString(h.Sum(nil))

		if entry.Hash != actualHash {
			return false, fmt.Errorf("corruption at sequence %d: hash mutation detected", entry.Sequence)
		}

		expectedPrevHash = entry.Hash
		_ = i // placeholder
	}
	return true, nil
}

// Export returns the full ledger for replay or audit.
func (l *TruthLedger) Export() ([]EvidenceWrapper, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.Entries, nil
}
