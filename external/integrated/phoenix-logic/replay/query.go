package replay

import (
	"fmt"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
)

// ReplayIndex holds an indexed representation of a replay run.
type ReplayIndex struct {
	Events []bus.TelemetryEvent
	ByUID  map[int][]int    // Map UID to indices in Events
	ByType map[string][]int // Map EventType to indices in Events
	ByPID  map[int][]int    // Map PID to indices in Events
}

// NewReplayIndex creates a new index from a slice of events.
func NewReplayIndex(events []bus.TelemetryEvent) *ReplayIndex {
	idx := &ReplayIndex{
		Events: events,
		ByUID:  make(map[int][]int),
		ByType: make(map[string][]int),
		ByPID:  make(map[int][]int),
	}

	for i, ev := range events {
		idx.ByUID[ev.UID] = append(idx.ByUID[ev.UID], i)
		idx.ByType[ev.EventType] = append(idx.ByType[ev.EventType], i)
		idx.ByPID[ev.PID] = append(idx.ByPID[ev.PID], i)
	}

	return idx
}

// SearchUID returns all events matching the given UID.
func (idx *ReplayIndex) SearchUID(uid int) []bus.TelemetryEvent {
	indices := idx.ByUID[uid]
	result := make([]bus.TelemetryEvent, len(indices))
	for i, pos := range indices {
		result[i] = idx.Events[pos]
	}
	return result
}

// Diff identifies the first point of divergence between two replay indices.
func Diff(idxA, idxB *ReplayIndex) (int, error) {
	lenA := len(idxA.Events)
	lenB := len(idxB.Events)

	minLen := lenA
	if lenB < minLen {
		minLen = lenB
	}

	for i := 0; i < minLen; i++ {
		evA := idxA.Events[i]
		evB := idxB.Events[i]

		if evA.SeqID != evB.SeqID || evA.EventType != evB.EventType || evA.UID != evB.UID {
			return i, fmt.Errorf("divergence at index %d: EventA(Seq:%d, Type:%s) != EventB(Seq:%d, Type:%s)",
				i, evA.SeqID, evA.EventType, evB.SeqID, evB.EventType)
		}
	}

	if lenA != lenB {
		return minLen, fmt.Errorf("replay length mismatch: RunA has %d events, RunB has %d events", lenA, lenB)
	}

	return -1, nil
}
