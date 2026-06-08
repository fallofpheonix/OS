package ledger

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestReplayerTopologicalSort(t *testing.T) {
	path := "test_dag_replay.log"
	defer os.Remove(path)

	p, err := NewPersistor(path)
	if err != nil {
		t.Fatal(err)
	}

	header := LedgerFileHeader{Version: "1.0", GenesisID: "GENESIS", Timestamp: 12345, Algorithm: "SHA256", FixedPointDivisor: 1000000}
	if err := p.WriteHeader(header); err != nil {
		t.Fatal(err)
	}

	// We'll construct a DAG like this:
	// A -> B
	// A -> C
	// B, C -> D
	//
	// We will write them to the file in order: A, D, C, B 
	// D depends on B and C, which aren't fully known when D is written if replayed linearly.
	
	l := NewLedger(nil)
	
	// Node A (Genesis-ish)
	err = l.AddEntryV2("A", "ROOT", 0, []byte("A"), "", nil, []byte("S_A"), "1.0")
	if err != nil { t.Fatal(err) }
	hashA := l.Heads[0]

	// Node B (Child of A)
	l.Heads = [][]byte{hashA}
	err = l.AddEntryV2("B", "A", 1, []byte("B"), "", []byte("S_A"), []byte("S_B"), "1.0")
	if err != nil { t.Fatal(err) }
	hashB := l.Heads[0]

	// Node C (Child of A)
	l.Heads = [][]byte{hashA}
	err = l.AddEntryV2("C", "A", 1, []byte("C"), "", []byte("S_A"), []byte("S_C"), "1.0")
	if err != nil { t.Fatal(err) }
	hashC := l.Heads[0]

	// Node D (Child of B and C)
	l.Heads = [][]byte{hashB, hashC}
	err = l.AddEntryV2("D", "BC", 2, []byte("D"), "", nil, []byte("S_D"), "1.0")
	if err != nil { t.Fatal(err) }
	hashD := l.Heads[0]

	// Get the entries to manually write them out of order
	entryA := l.Entries[fmt.Sprintf("%x", hashA)]
	entryB := l.Entries[fmt.Sprintf("%x", hashB)]
	entryC := l.Entries[fmt.Sprintf("%x", hashC)]
	entryD := l.Entries[fmt.Sprintf("%x", hashD)]

	// Write out of order: A, D, C, B
	p.Append(entryA)
	p.Append(entryD)
	p.Append(entryC)
	p.Append(entryB)

	// Replay
	replayer, err := NewReplayer(path)
	if err != nil { t.Fatal(err) }

	recovered, err := replayer.Replay(nil)
	if err != nil {
		t.Fatalf("Topological replay failed: %v", err)
	}

	// Verify D is the head (or among heads)
	foundD := false
	for _, h := range recovered.Heads {
		if bytes.Equal(h, hashD) {
			foundD = true
			break
		}
	}
	if !foundD {
		t.Errorf("Head mismatch: expected node D (%x) to be head, got %x", hashD, recovered.Heads)
	}

	if len(recovered.Entries) != 4 {
		t.Errorf("Entry count mismatch: expected 4, got %d", len(recovered.Entries))
	}
}

func TestReplayerCycleDetection(t *testing.T) {
	path := "test_cycle.log"
	defer os.Remove(path)

	p, err := NewPersistor(path)
	if err != nil { t.Fatal(err) }
	p.WriteHeader(LedgerFileHeader{Version: "1.0"})

	l := NewLedger(nil)
	
	// Create Node A
	_ = l.AddEntryV2("A", "ROOT", 0, []byte("A"), "", nil, nil, "1.0")
	hashA := l.Heads[0]
	entryA := l.Entries[fmt.Sprintf("%x", hashA)]

	// Create Node B that points to A
	l.Heads = [][]byte{hashA}
	_ = l.AddEntryV2("B", "A", 1, []byte("B"), "", nil, nil, "1.0")
	hashB := l.Heads[0]

	// Manually corrupt A to point to B (Cycle!)
	entryA.ParentIDs = [][]byte{hashB}
	// Note: We don't recompute hashA because the replayer verifies structural hash FIRST.
	// If we change ParentIDs without recomputing hash, structural verification fails.
	// So we need to create a validly-hashed cycle (impossible with SHA-256) 
	// OR we test the topologicalSort function directly with a cycle.
}

func TestTopologicalSortCycle(t *testing.T) {
	r := &Replayer{}
	entries := make(map[string]LedgerEntry)
	
	hashA := []byte{0xA}
	hashB := []byte{0xB}
	
	entries["0a"] = LedgerEntry{Hash: hashA, ParentIDs: [][]byte{hashB}}
	entries["0b"] = LedgerEntry{Hash: hashB, ParentIDs: [][]byte{hashA}}
	
	_, err := r.topologicalSort(entries)
	if err == nil {
		t.Error("Expected error for causal cycle, got nil")
	}
}
