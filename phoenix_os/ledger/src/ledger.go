package ledger

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

type Evidence struct {
	TraceHash    string  `json:"trace_hash"`
	SDI          float64 `json:"sdi"`
	PolicyID     string  `json:"policy_id"`
	Action       string  `json:"action"`
	Result       string  `json:"result"`
	Timestamp    int64   `json:"timestamp"`
	Confidence   float64 `json:"confidence"`
	ReplayID     string  `json:"replay_id"`
	ExperimentID string  `json:"experiment_id"`
	PrevHash     string  `json:"prev_hash"`
	EntryHash    string  `json:"entry_hash"`
}

type Ledger struct {
	mu       sync.RWMutex
	Entries  []Evidence
	LastHash string
}

func NewLedger() *Ledger {
	return &Ledger{
		Entries:  make([]Evidence, 0),
		LastHash: "GENESIS",
	}
}

func (l *Ledger) Commit(e Evidence) string {
	l.mu.Lock()
	defer l.mu.Unlock()

	e.Timestamp = time.Now().UnixNano()
	e.PrevHash = l.LastHash
	
	// Calculate Entry Hash for integrity
	data, _ := json.Marshal(e)
	h := sha256.New()
	h.Write(data)
	e.EntryHash = fmt.Sprintf("%x", h.Sum(nil))
	
	l.Entries = append(l.Entries, e)
	l.LastHash = e.EntryHash
	return e.EntryHash
}

func (l *Ledger) Verify() bool {
	l.mu.RLock()
	defer l.mu.RUnlock()

	expectedPrev := "GENESIS"
	for _, e := range l.Entries {
		if e.PrevHash != expectedPrev {
			return false
		}
		expectedPrev = e.EntryHash
	}
	return true
}
