package arbiter

import (
	"encoding/json"
	"fmt"
	"phoenix/warden"
)

type FrequencyMode int

const (
	FrequentThreat FrequencyMode = iota
	RareThreat
	AdaptiveThreat
)

// Policy defines the strategic rules for actuation classes.
type Policy struct {
	Version         string                            `json:"version"`
	Thresholds      map[warden.ActuationClass]float64 `json:"thresholds"`
	Budgets         map[warden.ActuationClass]int     `json:"budgets"`
	NodeCriticality map[int]float64                  `json:"node_criticality"` // CN
	FrequencyK      float64                           `json:"frequency_k"`      // k tuning constant
	FrequencyMode   FrequencyMode                     `json:"frequency_mode"`
}

// DefaultPolicy returns the bootstrap policy configuration.
func DefaultPolicy() Policy {
	return Policy{
		Version: "1.1.0",
		Thresholds: map[warden.ActuationClass]float64{
			warden.ClassObserve:         0.0,
			warden.ClassLog:             0.0,
			warden.ClassThrottle:        0.60,
			warden.ClassLocalIsolate:    0.85,
			warden.ClassClusterIsolate:  0.95,
			warden.ClassKernelEmergency: 0.99,
		},
		Budgets: map[warden.ActuationClass]int{
			warden.ClassThrottle:        100,
			warden.ClassLocalIsolate:    50,
			warden.ClassClusterIsolate:  10,
			warden.ClassKernelEmergency: 1,
		},
		NodeCriticality: map[int]float64{
			0:   1.0, // root
			1:   0.9, // bin
			100: 0.8, // system users
		},
		FrequencyK:    2.0,
		FrequencyMode: AdaptiveThreat,
	}
}

// Hash returns a canonical hash of the policy for evidence linking.
func (p *Policy) Hash() string {
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
