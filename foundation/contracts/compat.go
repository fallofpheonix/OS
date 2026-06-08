/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package contracts

const (
	StateSafe        = "SAFE"
	StateWatch       = "WATCH"
	StateSuspicious  = "SUSPICIOUS"
	StateCritical    = "CRITICAL"
	StateCompromised = "COMPROMISED"
)

// ILedger defines the canonical interface for the Phoenix Ledger.
type ILedger interface {
	AddEntry(eventID, causeID string, tick uint64, payload []byte) error
	AddEntryV2(eventID, causeID string, tick uint64, payload []byte, traceHash string, stateBefore, stateAfter []byte, policyVersion string) error
	GenerateCertificate(eventID string, weight float64) ([]byte, error)
	Verify() error
}
