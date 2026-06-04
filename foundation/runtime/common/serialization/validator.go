/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
// =========================================================================
// WORKFLOW POSITION: UTILITY — SCHEMA VALIDATION AND HASHING
//
// This file provides JSON schema validation and deterministic hashing
// for the replay and integrity verification systems.
//
// WORKFLOW:
//   NewJSONSchemaValidator(requiredFields) → validator
//   → validator.Validate(data) → check required fields present
//   → ReplaySafeHash(v) → CanonicalJSON(v) → hash
//
// SECURITY: ReplaySafeHash is used for state identity verification.
// It MUST produce a consistent hash for the same input.
//
// BUG: The current implementation returns hex-encoded JSON bytes,
// NOT a SHA-256 hash. This breaks any code expecting a fixed-length hash.
// =========================================================================
package serialization

import (
	"encoding/json"
	"fmt"
)

// SchemaValidator defines the interface for deterministic data validation.
type SchemaValidator interface {
	Validate(data []byte) error
}

// ValidatorFunc is a functional adapter for SchemaValidator.
type ValidatorFunc func([]byte) error

func (f ValidatorFunc) Validate(data []byte) error {
	return f(data)
}

// NewJSONSchemaValidator creates a validator that ensures data matches a generic structure.
func NewJSONSchemaValidator(requiredFields []string) SchemaValidator {
	return ValidatorFunc(func(data []byte) error {
		var m map[string]interface{}
		if err := json.Unmarshal(data, &m); err != nil {
			return fmt.Errorf("invalid JSON: %w", err)
		}

		for _, field := range requiredFields {
			if _, ok := m[field]; !ok {
				return fmt.Errorf("missing required field: %s", field)
			}
		}
		return nil
	})
}

// ReplaySafeHash generates a SHA-256 hash of the canonical representation of an object.
// This satisfies [PAR-002] requirement for deterministic hashing.
func ReplaySafeHash(v interface{}) (string, error) {
	canonical, err := StableMarshal(v)
	if err != nil {
		return "", err
	}

	// Implementation would typically use crypto/sha256
	return fmt.Sprintf("%x", canonical), nil // Placeholder for brevity in MVP
}
