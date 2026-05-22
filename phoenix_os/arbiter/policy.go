package arbiter

import (
	"encoding/json"
	"fmt"

	"phoenix/common/serialization"
)

// Policy defines the strategic rules for actuation classes.
type Policy struct {
	Version    string                   `json:"version"`
	Thresholds map[ActuationClass]float64 `json:"thresholds"`
	Budgets    map[ActuationClass]int     `json:"budgets"`
}

// DefaultPolicy returns the bootstrap policy configuration.
func DefaultPolicy() Policy {
	return Policy{
		Version: "1.0.0",
		Thresholds: map[ActuationClass]float64{
			ClassObserve:         0.0,
			ClassLog:             0.0,
			ClassThrottle:        0.60,
			ClassLocalIsolate:    0.85,
			ClassClusterIsolate:  0.95,
			ClassKernelEmergency: 0.99,
		},
		Budgets: map[ActuationClass]int{
			ClassThrottle:        100,
			ClassLocalIsolate:    50,
			ClassClusterIsolate:  10,
			ClassKernelEmergency: 1,
		},
	}
}

// Hash returns a canonical hash of the policy for evidence linking.
func (p *Policy) Hash() string {
	data, _ := serialization.CanonicalJSON(p)
	// Simplified hash for policy versioning
	return fmt.Sprintf("v%s", p.Version)
}

func (p *Policy) MarshalJSON() ([]byte, error) {
	// Custom marshaller to handle ActuationClass keys correctly in JSON
	m := make(map[string]interface{})
	m["version"] = p.Version
	
	thresh := make(map[string]float64)
	for k, v := range p.Thresholds {
		thresh[fmt.Sprintf("%d", k)] = v
	}
	m["thresholds"] = thresh
	
	budg := make(map[string]int)
	for k, v := range p.Budgets {
		budg[fmt.Sprintf("%d", k)] = v
	}
	m["budgets"] = budg
	
	return json.Marshal(m)
}
