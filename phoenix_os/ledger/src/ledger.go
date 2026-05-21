package main

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
	ModelVersion string  `json:"model_version"`
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

func (l *Ledger) calculateHash(e Evidence) string {
	// We must clear hashes before re-calculating to ensure consistency
	temp := e
	temp.EntryHash = ""
	data, _ := json.Marshal(temp)
	h := sha256.New()
	h.Write(data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (l *Ledger) Commit(e Evidence) string {
	l.mu.Lock()
	defer l.mu.Unlock()

	e.Timestamp = time.Now().UnixNano()
	e.PrevHash = l.LastHash
	e.EntryHash = l.calculateHash(e)
	
	l.Entries = append(l.Entries, e)
	l.LastHash = e.EntryHash
	return e.EntryHash
}

func (l *Ledger) Verify() bool {
	l.mu.RLock()
	defer l.mu.RUnlock()

	expectedPrev := "GENESIS"
	for _, e := range l.Entries {
		// 1. Check Linkage
		if e.PrevHash != expectedPrev {
			return false
		}
		// 2. Check Data Integrity
		if e.EntryHash != l.calculateHash(e) {
			return false
		}
		expectedPrev = e.EntryHash
	}
	return true
}

func (l *Ledger) Print() {
	l.mu.RLock()
	defer l.mu.RUnlock()
	
	fmt.Println("--- PHOENIX LEDGER ---")
	for _, e := range l.Entries {
		fmt.Printf("[%d] Trace: %s | SDI: %.2f | Action: %s | Hash: %s...\n", 
			e.Timestamp, e.TraceHash, e.SDI, e.Action, e.EntryHash[:8])
	}
	fmt.Println("----------------------")
}

func main() {
	fmt.Println("Phoenix Ledger starting...")
	l := NewLedger()

	// Commit evidence
	l.Commit(Evidence{
		TraceHash:    "0xABC123",
		SDI:          0.85,
		PolicyID:     "POL-001",
		Action:       "THROTTLE",
		Result:       "SUCCESS",
		ModelVersion: "v1.2.3",
		Confidence:   0.98,
		ReplayID:     "RUN-42",
		ExperimentID: "EXP-ALPHA",
	})

	l.Commit(Evidence{
		TraceHash:    "0xDEF456",
		SDI:          0.92,
		PolicyID:     "POL-001",
		Action:       "ISOLATE",
		Result:       "SUCCESS",
		ModelVersion: "v1.2.3",
		Confidence:   0.99,
		ReplayID:     "RUN-42",
		ExperimentID: "EXP-ALPHA",
	})

	l.Print()
	fmt.Printf("Ledger Verification: %v\n", l.Verify())
}
