package trace

import (
	"time"
	"github.com/fallofpheonix/phoenixmind-validator/truth/evidence"
)

// TimelineEvent represents a significant event in the system's history.
type TimelineEvent struct {
	ID        string
	Timestamp time.Time
	EventType string // e.g., "EVIDENCE_INGESTED", "STATE_UPDATED", "DECISION_TAKEN"
	EntityID  string
	State     evidence.TruthState
	Metadata  map[string]interface{}
}

// Timeline maintains an ordered sequence of all significant events.
type Timeline struct {
	Events []TimelineEvent
}

// AddEvent adds a new event to the timeline, maintaining chronological order.
func (tl *Timeline) AddEvent(event TimelineEvent) {
	// For simplicity, we just append. In a real system, this would be a sorted insert
	// or events would be guaranteed to arrive in order.
	tl.Events = append(tl.Events, event)
}

// GetEventsForEntity retrieves all events related to a specific entity.
func (tl *Timeline) GetEventsForEntity(entityID string) []TimelineEvent {
	var entityEvents []TimelineEvent
	for _, event := range tl.Events {
		if event.EntityID == entityID {
			entityEvents = append(entityEvents, event)
		}
	}
	return entityEvents
}
