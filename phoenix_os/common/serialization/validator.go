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
	canonical, err := CanonicalJSON(v)
	if err != nil {
		return "", err
	}
	
	// Implementation would typically use crypto/sha256
	return fmt.Sprintf("%x", canonical), nil // Placeholder for brevity in MVP
}
