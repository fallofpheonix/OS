package ledger

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"testing"
)

func TestReplayerHeadHashContinuity(t *testing.T) {
	path := "test_replay.log"
	defer os.Remove(path)

	// 1. Create Persistor, write header, append 10 entries via Ledger+Persistor
	p, err := NewPersistor(path)
	if err != nil {
		t.Fatal(err)
	}

	header := LedgerFileHeader{Version: "1.0", GenesisID: "GENESIS", Timestamp: 12345, Algorithm: "SHA256", FixedPointDivisor: 1000000}
	if err := p.WriteHeader(header); err != nil {
		t.Fatal(err)
	}

	liveLedger := NewLedger(nil).WithPersistor(p)
	for i := 0; i < 10; i++ {
		err := liveLedger.AddEntryV2("EVT", "CAUSE", uint64(i), []byte("payload"), "", []byte("PRE"), []byte("POST"), "1.0")
		if err != nil {
			t.Fatal(err)
		}
	}

	// 2. Record the live ledger's Head Hash
	if len(liveLedger.Heads) != 1 {
		t.Fatalf("Expected exactly 1 head, got %d", len(liveLedger.Heads))
	}
	recordedHeadHash := liveLedger.Heads[0]

	// 3. Create Replayer with same file path
	replayer, err := NewReplayer(path)
	if err != nil {
		t.Fatal(err)
	}

	// 4. Call Replay(), get fresh ledger
	freshLedger, err := replayer.Replay()
	if err != nil {
		t.Fatal(err)
	}

	// 5. Assert fresh ledger Head Hash == recorded Head Hash
	if len(freshLedger.Heads) != 1 {
		t.Fatalf("Expected exactly 1 head on fresh ledger, got %d", len(freshLedger.Heads))
	}

	if !bytes.Equal(freshLedger.Heads[0], recordedHeadHash) {
		t.Fatalf("Head Hash mismatch: expected %x, got %x", recordedHeadHash, freshLedger.Heads[0])
	}

	// 6. Assert len(entries) == 10
	if len(freshLedger.Entries) != 10 {
		t.Fatalf("Expected 10 entries in fresh ledger, got %d", len(freshLedger.Entries))
	}
}

func TestReplayAfterConcurrentWrites(t *testing.T) {
	path := "test_concurrent_writes.log"
	defer os.Remove(path)

	p, err := NewPersistor(path)
	if err != nil {
		t.Fatal(err)
	}

	header := LedgerFileHeader{Version: "1.0", GenesisID: "GENESIS", Timestamp: 12345, Algorithm: "SHA256", FixedPointDivisor: 1000000}
	if err := p.WriteHeader(header); err != nil {
		t.Fatal(err)
	}

	liveLedger := NewLedger(nil).WithPersistor(p)

	// Write 50 entries concurrently (10 goroutines * 5 entries)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(routineID int) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				// Event ID includes routine to be unique, tick is somewhat arbitrary for this test
				eventID := fmt.Sprintf("EVT-R%d-J%d", routineID, j)
				tick := uint64(routineID*5 + j)
				_ = liveLedger.AddEntryV2(eventID, "CAUSE", tick, []byte("payload"), "", []byte("PRE"), []byte("POST"), "1.0")
			}
		}(i)
	}
	wg.Wait()

	// Ensure there is only one head (since AddEntryV2 uses a simple head replacement mechanism right now,
	// it will be the last one written, but because of concurrency, it's non-deterministic which one is last.
	// However, the replayer must produce the exact same single head hash).
	// NOTE: In a true Merkle DAG, concurrent writes would produce multiple heads until merged.
	// Our current AddEntryV2 overwrites l.Heads = [][]byte{entry.Hash}, so there is always exactly 1 head.

	if len(liveLedger.Heads) != 1 {
		t.Fatalf("Expected exactly 1 head, got %d", len(liveLedger.Heads))
	}

	// [RECTIFIED]: Step A - Canonical Ordering.
	// In a concurrent write scenario, the live ledger's "Heads" is non-deterministic
	// (it's whoever got the lock last). However, the Replayer produces a deterministic
	// result by sorting by (Tick, Sequence). We assert that the Replayer matches
	// the canonical head (lexicographical max of the sorted entries).
	entries := liveLedger.SortedEntries()
	canonicalHeadHash := entries[len(entries)-1].Hash

	// Replay
	replayer, err := NewReplayer(path)
	if err != nil {
		t.Fatal(err)
	}

	freshLedger, err := replayer.Replay()
	if err != nil {
		t.Fatal(err)
	}

	// Assert Head Hash matches canonical head
	if len(freshLedger.Heads) != 1 {
		t.Fatalf("Expected exactly 1 head on fresh ledger, got %d", len(freshLedger.Heads))
	}

	if !bytes.Equal(freshLedger.Heads[0], canonicalHeadHash) {
		t.Fatalf("Head Hash mismatch: expected canonical %x, got %x", canonicalHeadHash, freshLedger.Heads[0])
	}

	// Assert entry count
	if len(freshLedger.Entries) != 50 {
		t.Fatalf("Expected 50 entries in fresh ledger, got %d", len(freshLedger.Entries))
	}
}
