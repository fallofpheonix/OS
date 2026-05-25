package state

import "fmt"

// MigrationMap maps legacy state identifiers to current RuntimeState values.
var MigrationMap = map[string]RuntimeState{
	"NORMAL":      Safe,
	"SUSPICIOUS":  Watch,
	"CONTAINED":   Contain,
	"CRITICAL":    Alert,   // Mapping legacy CRITICAL to ALERT
	"COMPROMISED": Contain, // Mapping legacy COMPROMISED to CONTAIN
}

// CompatRegistry provides lookup capabilities for legacy state names.
type CompatRegistry struct {
	migrations map[string]RuntimeState
}

func NewCompatRegistry() *CompatRegistry {
	return &CompatRegistry{
		migrations: MigrationMap,
	}
}

// Lookup resolves a state name (legacy or current) to a RuntimeState.
func (c *CompatRegistry) Lookup(name string) (RuntimeState, error) {
	// Check if it's already a valid current state
	s := RuntimeState(name)
	switch s {
	case Safe, Watch, Alert, Contain, Recovery:
		return s, nil
	}

	// Check migrations
	if target, ok := c.migrations[name]; ok {
		return target, nil
	}

	return "", fmt.Errorf("unknown state name: %s", name)
}
