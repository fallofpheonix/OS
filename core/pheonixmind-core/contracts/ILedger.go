package contracts

// ILedger defines the interface for an evidence ledger (local or distributed).
type ILedger interface {
	AddEntry(eventID, causeID string, payload []byte) error
	AddEntryV2(eventID, causeID string, payload []byte, traceHash, stateBefore, stateAfter, policyVersion string) error
	Verify() error
	GenerateCertificate(eventID string, weight float64) ([]byte, error)
}
