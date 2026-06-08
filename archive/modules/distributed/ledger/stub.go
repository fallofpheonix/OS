/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package ledger

import (
	"context"
	"errors"
	"sync"
	"time"
)

// StubConsensusLedger is a simplified implementation of ConsensusLedger for development.
type StubConsensusLedger struct {
	mu       sync.RWMutex
	entries  []Entry
	watchers []chan Entry
}

func NewStubConsensusLedger() *StubConsensusLedger {
	return &StubConsensusLedger{
		entries:  make([]Entry, 0),
		watchers: make([]chan Entry, 0),
	}
}

func (s *StubConsensusLedger) Propose(ctx context.Context, eventID string, payload []byte) (uint64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	entry := Entry{
		Index:     uint64(len(s.entries)),
		Term:      1,
		Timestamp: time.Now(),
		Payload:   payload,
		EventID:   eventID,
	}

	s.entries = append(s.entries, entry)

	for _, ch := range s.watchers {
		select {
		case ch <- entry:
		default:
		}
	}

	return entry.Index, nil
}

func (s *StubConsensusLedger) Commit(ctx context.Context, index uint64) error {
	// Stubs are auto-committed for now
	return nil
}

func (s *StubConsensusLedger) Get(ctx context.Context, index uint64) (*Entry, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if index >= uint64(len(s.entries)) {
		return nil, errors.New("index out of bounds")
	}
	return &s.entries[index], nil
}

func (s *StubConsensusLedger) Subscribe(ctx context.Context) (<-chan Entry, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	ch := make(chan Entry, 10)
	s.watchers = append(s.watchers, ch)
	return ch, nil
}

func (s *StubConsensusLedger) Verify(ctx context.Context) error {
	return nil
}
