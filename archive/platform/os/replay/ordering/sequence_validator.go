/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package ordering

import "fmt"

// SequenceValidator defines the interface for validating sequences of events.
type SequenceValidator interface {
	// Validate checks if a slice of events is strictly ordered by their SequenceID.
	// It returns an error if any event is out of sequence or if there are duplicate IDs.
	Validate(events []*Event) error
}

// NewSequenceValidator creates a new instance of SequenceValidator.
func NewSequenceValidator() SequenceValidator {
	return &simpleSequenceValidator{}
}

type simpleSequenceValidator struct{}

// Validate checks if a slice of events is strictly ordered by their SequenceID.
func (ssv *simpleSequenceValidator) Validate(events []*Event) error {
	if len(events) == 0 {
		return nil // An empty sequence is valid
	}

	seenIDs := make(map[uint64]bool)
	for i, event := range events {
		if event == nil {
			return fmt.Errorf("event at index %d is nil", i)
		}
		if event.SequenceID == 0 {
			return fmt.Errorf("event at index %d has a zero sequence ID", i)
		}

		if seenIDs[event.SequenceID] {
			return fmt.Errorf("duplicate sequence ID %d found at index %d", event.SequenceID, i)
		}
		seenIDs[event.SequenceID] = true

		if i > 0 {
			if event.SequenceID <= events[i-1].SequenceID {
				return fmt.Errorf("sequence regression or duplication detected: event at index %d (ID %d) is not strictly greater than event at index %d (ID %d)",
					i, event.SequenceID, i-1, events[i-1].SequenceID)
			}
		}
	}
	return nil
}
