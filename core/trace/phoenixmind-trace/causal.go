package trace

import (
	"time"
)

// CausalEvent represents an event that can cause or be caused by other events.
type CausalEvent struct {
	ID        string
	Timestamp time.Time
	Type      string // e.g., "EVIDENCE_INGESTED", "STATE_CHANGE", "DECISION_MADE"
	Payload   map[string]interface{}
}

// CausalLink represents a directional causal relationship between two events.
type CausalLink struct {
	CauseID  string
	EffectID string
	Type     string // e.g., "PRECEDES", "CAUSED_BY"
	Strength float64 // How strong is the causal link, 0.0 to 1.0
}

// CausalGraph tracks the cause-and-effect relationships within the system.
type CausalGraph struct {
	Links []CausalLink
}

// AddLink records a new causal relationship.
func (cg *CausalGraph) AddLink(causeID, effectID, linkType string, strength float64) {
	cg.Links = append(cg.Links, CausalLink{
		CauseID:  causeID,
		EffectID: effectID,
		Type:     linkType,
		Strength: strength,
	})
}

// GetEffectsOf retrieves all direct effects caused by a given event.
func (cg *CausalGraph) GetEffectsOf(causeID string) []CausalLink {
	var effects []CausalLink
	for _, link := range cg.Links {
		if link.CauseID == causeID {
			effects = append(effects, link)
		}
	}
	return effects
}
