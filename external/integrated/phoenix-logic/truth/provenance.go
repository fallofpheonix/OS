package truth

import (
	"fmt"
	"time"
)

// EventProvenance (B9) tracks the origin and causal history of a system event.
type EventProvenance struct {
	EventID     string
	SourceLayer string
	SourceNode  string
	CausalID    string // Reference to the parent/triggering event
	Timestamp   time.Time
}

// ProvenanceTracker (B9) maintains the lineage of events.
type ProvenanceTracker struct {
	Lineage map[string]*EventProvenance
}

func NewProvenanceTracker() *ProvenanceTracker {
	return &ProvenanceTracker{
		Lineage: make(map[string]*EventProvenance),
	}
}

func (t *ProvenanceTracker) Register(p *EventProvenance) {
	t.Lineage[p.EventID] = p
}

func (t *ProvenanceTracker) GetLineage(eventID string) ([]*EventProvenance, error) {
	var path []*EventProvenance
	currentID := eventID

	for {
		p, ok := t.Lineage[currentID]
		if !ok {
			break
		}
		path = append(path, p)
		if p.CausalID == "" {
			break
		}
		currentID = p.CausalID

		// Prevent infinite loops in cyclic dependencies (though they shouldn't happen)
		if len(path) > 1000 {
			return path, fmt.Errorf("lineage too deep or cyclic")
		}
	}
	return path, nil
}
