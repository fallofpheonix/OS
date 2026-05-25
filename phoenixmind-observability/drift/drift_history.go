package drift

import "sync"

type DriftRecord struct {
	Module     string
	Drift      float64
	Timestamp  uint64
	Variance   float64
	Confidence float64
}

type DriftHistory struct {
	mu      sync.RWMutex
	Records []DriftRecord
}

func (h *DriftHistory) Add(rec DriftRecord) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.Records = append(h.Records, rec)
}
