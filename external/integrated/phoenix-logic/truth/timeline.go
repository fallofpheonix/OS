package truth

import (
	"sort"
	"time"
)

// Event represents a generic system event in a timeline.
type Event struct {
	Timestamp time.Time
	ID        string
	Type      string
	Payload   interface{}
}

// Timeline (B7) is a collection of events sorted by time.
type Timeline struct {
	Events []Event
}

// TimelineMerger (B7) merges multiple timelines into a single, unified view.
type TimelineMerger struct{}

func (m *TimelineMerger) Merge(t1, t2 *Timeline) *Timeline {
	merged := append(t1.Events, t2.Events...)

	// Sort by timestamp for consistency
	sort.Slice(merged, func(i, j int) bool {
		return merged[i].Timestamp.Before(merged[j].Timestamp)
	})

	return &Timeline{Events: merged}
}
