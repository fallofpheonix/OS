package truth

import (
	"fmt"
	"github.com/fallofpheonix/phoenix-contracts"
)

// EvidenceGC (B3) is responsible for pruning old or low-importance evidence.
type EvidenceGC struct {
	Threshold float64
}

func NewEvidenceGC(threshold float64) *EvidenceGC {
	return &EvidenceGC{Threshold: threshold}
}

// CollectGarbage removes evidence that falls below the importance threshold.
func (gc *EvidenceGC) CollectGarbage(s *EvidenceStore) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	removedCount := 0
	for pid, evidenceList := range s.Evidence {
		var activeEvidence []contracts.Evidence
		for _, e := range evidenceList {
			if e.Score()*e.Confidence() >= gc.Threshold {
				activeEvidence = append(activeEvidence, e)
			} else {
				removedCount++
			}
		}
		s.Evidence[pid] = activeEvidence
	}
	return removedCount
}

// EvidenceCompressor (B4) handles the compression of evidence for long-term storage.
type EvidenceCompressor struct{}

func (c *EvidenceCompressor) Compress(e contracts.Evidence) ([]byte, error) {
	// Stub for evidence compression logic.
	// In a real implementation, this might use zlib or a custom binary format.
	return []byte(fmt.Sprintf("compressed:%d:%f", e.GetPID(), e.Score())), nil
}
