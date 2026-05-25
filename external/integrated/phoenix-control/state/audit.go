package state

import (
	"sync"
	"time"
)

// AuditLog records a single state transition event.
type AuditEntry struct {
	Timestamp time.Time
	From      RuntimeState
	To        RuntimeState
	Reason    string
	Tick      int64
}

// AuditLogger manages state transition logs.
type AuditLogger struct {
	mu      sync.Mutex
	entries []AuditEntry
}

func NewAuditLogger() *AuditLogger {
	return &AuditLogger{
		entries: make([]AuditEntry, 0),
	}
}

func (l *AuditLogger) Log(from, to RuntimeState, reason string, tick int64) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.entries = append(l.entries, AuditEntry{
		Timestamp: time.Now(),
		From:      from,
		To:        to,
		Reason:    reason,
		Tick:      tick,
	})
}

func (l *AuditLogger) GetEntries() []AuditEntry {
	l.mu.Lock()
	defer l.mu.Unlock()
	// Return a copy to avoid race conditions
	entriesCopy := make([]AuditEntry, len(l.entries))
	copy(entriesCopy, l.entries)
	return entriesCopy
}
