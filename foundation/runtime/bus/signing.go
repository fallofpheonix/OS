/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
// =========================================================================
// WORKFLOW POSITION: STUB — EVENT SIGNING (NOT IMPLEMENTED)
//
// This file is a STUB for event signing functionality.
// The intended purpose is to provide HMAC-SHA256 signing for all events
// published to the Bus, ensuring event integrity and non-repudiation.
//
// PLANNED WORKFLOW:
//   Bus.Publish() → SignEvent(payload, key) → HMAC-SHA256 signature
//     → TelemetryEvent.Hash = signature
//     → Consumer verifies: SignEvent(payload, key) == event.Hash
//
// CURRENT STATE: Returns payload unchanged (NO SIGNING).
// This means ALL events on the Bus have NO integrity protection.
// An attacker who compromises any subscriber can modify events freely.
//
// SECURITY IMPACT: CRITICAL — the entire audit trail can be tampered with.
// The Ledger, Truth, and Replay systems all depend on event integrity.
// Without signing, these systems provide NO security guarantees.
//
// FIX REQUIRED: Implement HMAC-SHA256 with a key loaded from a vault.
// Add signature verification at the consumer side.
// =========================================================================
package bus

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

var (
	// defaultKey is used if no other key is provided.
	// In production, this MUST be replaced by a vault-managed key.
	defaultKey = []byte("PHOENIX_MATRIX_INSECURE_DEFAULT_KEY")
)

// SetSigningKey sets the global key used for event signing.
// This should be called during system initialization with a key from a secure vault.
func SetSigningKey(key []byte) {
	if len(key) > 0 {
		defaultKey = key
	}
}

// SignEvent computes the HMAC-SHA256 signature for the given payload.
func SignEvent(payload, key []byte) []byte {
	if len(key) == 0 {
		key = defaultKey
	}
	h := hmac.New(sha256.New, key)
	h.Write(payload)
	return h.Sum(nil)
}

// VerifyEvent checks if the Hash in the TelemetryEvent matches the signature of its content.
func VerifyEvent(event TelemetryEvent, key []byte) bool {
	if event.Hash == "" {
		return false
	}

	sig, err := hex.DecodeString(event.Hash)
	if err != nil {
		return false
	}

	expectedSig := ComputeEventSignature(event, key)
	return hmac.Equal(sig, expectedSig)
}

// ComputeEventSignature serializes the event (without the Hash) and signs it.
func ComputeEventSignature(event TelemetryEvent, key []byte) []byte {
	// Clear hash to ensure we sign the content only
	event.Hash = ""
	data, _ := json.Marshal(event)
	return SignEvent(data, key)
}
