/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package contracts

// ILedger defines the interface for an evidence ledger (local or distributed).
type ILedger interface {
	AddEntry(eventID, causeID string, payload []byte) error
	AddEntryV2(eventID, causeID string, payload []byte, traceHash, stateBefore, stateAfter, policyVersion string) error
	Verify() error
	GenerateCertificate(eventID string, weight float64) ([]byte, error)
}
