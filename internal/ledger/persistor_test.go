package ledger

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPersist_AppendAndLoad(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "ledger.jsonl")

	p, err := NewPersistor(path)
	if err != nil {
		t.Fatalf("Failed to create persistor: %v", err)
	}
	defer p.Close()

	// 1. Create Genesis
	e0 := NewEvent(0, EventGenesis, []byte("{}"), "", "SYSTEM")
	if err := p.Append(e0); err != nil {
		t.Fatalf("Failed to append genesis: %v", err)
	}

	// 2. Add some events
	e1 := NewEvent(1, EventEnforce, []byte(`{"state":"WATCH"}`), e0.Hash, "WARDEN")
	if err := p.Append(e1); err != nil {
		t.Fatalf("Failed to append e1: %v", err)
	}

	// 3. Load into new chain
	chain := NewChain()
	count, err := p.Load(chain)
	if err != nil {
		t.Fatalf("Failed to load events: %v", err)
	}

	if count != 2 {
		t.Errorf("Expected 2 events loaded, got %d", count)
	}

	if head := chain.GetHead(); head == nil || head.Hash != e1.Hash {
		t.Errorf("Head hash mismatch. Expected %s, got %v", e1.Hash, head)
	}
}

func TestPersist_SyncAndLoad(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "ledger_sync.jsonl")

	// Session 1: Write and Close
	p1, _ := NewPersistor(path)
	e0 := NewEvent(0, EventGenesis, []byte("{}"), "", "SYSTEM")
	p1.Append(e0)
	p1.Close()

	// Session 2: Open, Write, and Close
	p2, _ := NewPersistor(path)
	e1 := NewEvent(1, EventEnforce, []byte(`{"state":"CRITICAL"}`), e0.Hash, "WARDEN")
	p2.Append(e1)
	p2.Close()

	// Session 3: Replay
	p3, _ := NewPersistor(path)
	defer p3.Close()
	chain := NewChain()
	count, err := p3.Load(chain)
	if err != nil {
		t.Fatalf("Failed to load: %v", err)
	}

	if count != 2 {
		t.Errorf("Expected 2 events, got %d", count)
	}
	if chain.GetHead().Type != EventEnforce {
		t.Errorf("Wrong head type: %s", chain.GetHead().Type)
	}
}

func TestPersist_Corruption(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "ledger_corrupt.jsonl")

	p, _ := NewPersistor(path)
	e0 := NewEvent(0, EventGenesis, []byte("{}"), "", "SYSTEM")
	p.Append(e0)
	p.Close()

	// Manually corrupt the file by appending garbage
	f, _ := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	f.Write([]byte("THIS IS NOT JSON\n"))
	f.Close()

	p2, _ := NewPersistor(path)
	defer p2.Close()
	chain := NewChain()
	_, err := p2.Load(chain)
	if err == nil {
		t.Error("Expected error loading corrupted ledger, got nil")
	}
}
