package ledger

import (
	"bytes"
	"fmt"
	"testing"

	"phoenix/common/resource"
)

func TestSnapshotRollback(t *testing.T) {
	l := NewLedger(resource.NewBoundedAllocator(1024*1024, 100))

	// 1. Add some entries
	l.AddEntry("E1", "START", []byte("data1"))
	l.AddEntry("E2", "E1", []byte("data2"))
	
	// 2. Take a checkpoint
	cp1, _ := l.Checkpoint()
	
	// 3. Add more entries
	l.AddEntry("E3", "E2", []byte("data3"))
	
	// 4. "Rollback" simulation: Verify that cp1 is a valid historical state
	// In our DAG model, we can verify if cp1 is a parent or a known hash.
	
	if _, ok := l.Entries[fmt.Sprintf("%x", cp1)]; !ok {
		t.Errorf("Checkpoint hash %x not found in ledger entries", cp1)
	}

	// 5. Hard Rollback: Reset heads to cp1
	l.Heads = [][]byte{cp1}
	l.Counter = 2 // Manual reset for test
	
	// 6. Add a divergent entry
	l.AddEntry("E3-divergent", "E2", []byte("divergent-data"))
	
	// 7. Verify Merkle DAG structure (E3 and E3-divergent both have E2 as parent)
	// Actually, our AddEntry automatically links to current heads.
	
	foundDivergence := false
	for _, entry := range l.Entries {
		if entry.EventID == "E3-divergent" {
			for _, p := range entry.ParentIDs {
				if bytes.Equal(p, cp1) {
					foundDivergence = true
				}
			}
		}
	}
	
	if !foundDivergence {
		t.Error("Failed to create divergent branch from checkpoint")
	}
}
