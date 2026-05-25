package truth

import (
	"github.com/fallofpheonix/phoenix-contracts"
	"sync"
)

// EvidenceStore is the central repository for validated forensic evidence.
type EvidenceStore struct {
	mu       sync.RWMutex
	Evidence map[int][]contracts.Evidence // Keyed by PID
}

func NewEvidenceStore() *EvidenceStore {
	return &EvidenceStore{
		Evidence: make(map[int][]contracts.Evidence),
	}
}

func (s *EvidenceStore) Add(e contracts.Evidence) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Evidence[e.GetPID()] = append(s.Evidence[e.GetPID()], e)
}

// GetByPID (B8) retrieves all evidence associated with a specific process.
func (s *EvidenceStore) GetByPID(pid int) []contracts.Evidence {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if evidence, ok := s.Evidence[pid]; ok {
		// Return a copy to prevent external modification
		return append([]contracts.Evidence(nil), evidence...)
	}
	return nil
}

// GetPIDs (B8) returns a list of all PIDs currently indexed.
func (s *EvidenceStore) GetPIDs() []int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	pids := make([]int, 0, len(s.Evidence))
	for pid := range s.Evidence {
		pids = append(pids, pid)
	}
	return pids
}

// TruthManager coordinates between Replay and the Evidence Store.
type TruthManager struct {
	Store  *EvidenceStore
	Ledger interface{} // Will point to unified ledger
}

func NewTruthManager(ledger interface{}) *TruthManager {
	return &TruthManager{
		Store:  NewEvidenceStore(),
		Ledger: ledger,
	}
}
