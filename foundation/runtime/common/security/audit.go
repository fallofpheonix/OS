/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
// =========================================================================
// WORKFLOW POSITION: UTILITY — CRYPTOGRAPHIC AUDIT TRAIL
//
// The CryptoAuditor logs cryptographic operations to the Ledger for
// forensic replay. Every hash computation and verification is recorded.
//
// WORKFLOW:
//   AuditHash(data, context) → SHA-256 hash → Ledger entry
//   AuditVerify(expected, actual, context) → compare → Ledger entry
//
// PURPOSE: Provides an audit trail for all cryptographic operations.
// If a hash collision or verification failure is detected, the Ledger
// entry provides the forensic evidence needed for investigation.
//
// SECURITY: The Ledger is append-only and tamper-evident.
// Audit entries cannot be modified or deleted without detection.
// =========================================================================
package security

import (
	"crypto/sha256"
	"fmt"

	ledger "github.com/fallofpheonix/phoenix/foundation/ledger/src"
)

// CryptoAuditor logs cryptographic operations to the ledger for forensic replay.
// This implements [SEC-010].
type CryptoAuditor struct {
	ledger *ledger.Ledger
}

func NewCryptoAuditor(l *ledger.Ledger) *CryptoAuditor {
	return &CryptoAuditor{ledger: l}
}

// AuditHash performs a SHA-256 hash and records the operation.
func (a *CryptoAuditor) AuditHash(data []byte, context string) string {
	hash := sha256.Sum256(data)
	hashStr := fmt.Sprintf("%x", hash)

	// Record to ledger
	_ = a.ledger.AddEntry("CRYPTO-OP", "SHA256", []byte(fmt.Sprintf("ctx:%s,hash:%s", context, hashStr)))

	return hashStr
}

// AuditVerify records a verification operation.
func (a *CryptoAuditor) AuditVerify(expected, actual, context string) bool {
	match := expected == actual
	_ = a.ledger.AddEntry("CRYPTO-VERIFY", context, []byte(fmt.Sprintf("match:%v", match)))
	return match
}
