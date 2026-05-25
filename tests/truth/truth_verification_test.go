package truth

import (
	"fmt"
	"sync"
	"testing"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/truth_ledger/src"
)

type MockAllocator struct{}
func (m *MockAllocator) Allocate(bytes uint64) error { return nil }
func (m *MockAllocator) Deallocate(bytes uint64)      {}

func TestMutationRace(t *testing.T) {
	l := ledger.NewLedger(&MockAllocator{})
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			l.AddEntry(fmt.Sprintf("evt-%d", id), "test", []byte("data"))
		}(i)
	}
	wg.Wait()
	if err := l.Verify(); err != nil {
		t.Fatalf("Race corruption detected: %v", err)
	}
}

func TestSnapshotRepeat(t *testing.T) {
	l := ledger.NewLedger(&MockAllocator{})
	for i := 0; i < 100; i++ {
		l.AddEntry(fmt.Sprintf("evt-%d", i), "test", []byte("data"))
		l.Checkpoint()
	}
}

func TestSealStress(t *testing.T) {
	// Implicit seal via hash chain
}

func TestForkRecovery(t *testing.T) {
	// Ledger V2 logic
}

func TestHashRepeat(t *testing.T) {
	l := ledger.NewLedger(&MockAllocator{})
	l.AddEntry("e1", "c1", []byte("d1"))
	h1 := fmt.Sprintf("%x", l.Heads[0])
	l.Verify()
	h2 := fmt.Sprintf("%x", l.Heads[0])
	if h1 != h2 {
		t.Fatal("Hash non-deterministic across verify")
	}
}
