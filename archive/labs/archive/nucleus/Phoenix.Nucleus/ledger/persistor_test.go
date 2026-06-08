package ledger

import (
	"os"
	"testing"
)

func TestPersistor(t *testing.T) {
	path := "test_ledger.log"
	defer os.Remove(path)

	p, err := NewPersistor(path)
	if err != nil {
		t.Fatalf("Failed to create persistor: %v", err)
	}

	e1 := NewEvent(0, EventGenesis, []byte("{}"), "", "SYSTEM")
	if err := p.Append(e1); err != nil {
		t.Errorf("Failed to append event: %v", err)
	}

	p.Close()

	// Verify persistence
	c := NewChain()
	p2, _ := NewPersistor(path)
	if err := p2.Load(c); err != nil {
		t.Errorf("Failed to load ledger: %v", err)
	}

	if c.GetHead().Sequence != 0 {
		t.Errorf("Expected sequence 0, got %d", c.GetHead().Sequence)
	}
}
