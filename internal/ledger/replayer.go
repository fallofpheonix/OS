package ledger

import (
	"fmt"
)

// Replayer orchestrates the recovery of system state from persisted logs.
type Replayer struct {
	persistor *Persistor
}

// NewReplayer creates a new Replayer instance.
func NewReplayer(p *Persistor) *Replayer {
	return &Replayer{
		persistor: p,
	}
}

// Reconstruct rebuilding a Chain from disk and verifies its integrity.
// It returns the head event's hash, representing the authoritative system state.
func (r *Replayer) Reconstruct() (string, error) {
	chain := NewChain()

	count, err := r.persistor.Load(chain)
	if err != nil {
		return "", fmt.Errorf("failed to load chain: %w", err)
	}

	if count == 0 {
		return "", fmt.Errorf("ledger is empty, cannot reconstruct state")
	}

	// Full cryptographic audit
	if err := chain.VerifyChain(); err != nil {
		return "", fmt.Errorf("cryptographic verification failed: %w", err)
	}

	head := chain.GetHead()
	if head == nil {
		return "", fmt.Errorf("no head event found after load")
	}

	return head.Hash, nil
}
