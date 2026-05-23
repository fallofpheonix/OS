package security

import (
	"crypto/sha256"
	"fmt"
	"phoenix/ledger/src"
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
	a.ledger.AddEntry("CRYPTO-OP", "SHA256", []byte(fmt.Sprintf("ctx:%s,hash:%s", context, hashStr)))
	
	return hashStr
}

// AuditVerify records a verification operation.
func (a *CryptoAuditor) AuditVerify(expected, actual string, context string) bool {
	match := expected == actual
	a.ledger.AddEntry("CRYPTO-VERIFY", context, []byte(fmt.Sprintf("match:%v", match)))
	return match
}
