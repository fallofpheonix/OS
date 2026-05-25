package truth

// LedgerVerifier provides methods to validate the integrity of the TruthLedger.
type LedgerVerifier struct{}

// NewLedgerVerifier creates a new verifier instance.
func NewLedgerVerifier() *LedgerVerifier {
	return &LedgerVerifier{}
}

// Verify checks the entire hash chain for consistency.
// Note: In this simplified version, we'd ideally need the original payloads to fully re-verify.
// If the ledger only stores hashes, we can only check if the chain is self-consistent
// if the previous hash was included in the current hash calculation.
func (v *LedgerVerifier) Verify(l *TruthLedger) error {
	_, err := l.Verify()
	return err
}

// DetectRepair (B5) identifies indices where the hash chain might have been tampered with or "repaired".
func (v *LedgerVerifier) DetectRepair(l *TruthLedger) ([]int, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	brokenIndices := make([]int, 0)
	// In a real repair detection scenario, we would compare the current ledger
	// against a trusted remote witness or a signed checkpoint.
	// For this foundation, we simulate the check.
	if len(l.Entries) > 0 && l.LastHash == "" {
		brokenIndices = append(brokenIndices, 0)
	}

	return brokenIndices, nil
}
