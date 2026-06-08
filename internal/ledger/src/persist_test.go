package ledger

import (
	"os"
	"testing"
)

func TestPersistorWriteAndRead(t *testing.T) {
	path := "test_persist.log"
	defer os.Remove(path)

	p, err := NewPersistor(path)
	if err != nil {
		t.Fatal(err)
	}

	header := LedgerFileHeader{Version: "1.0", GenesisID: "GENESIS", Timestamp: 12345, Algorithm: "SHA256", FixedPointDivisor: 1000000}
	if err := p.WriteHeader(header); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 5; i++ {
		entry := LedgerEntry{EventID: "EVT", LogicalTick: uint64(i)}
		if err := p.Append(entry); err != nil {
			t.Fatal(err)
		}
	}

	entries, readHeader, err := p.ReadAll()
	if err != nil {
		t.Fatal(err)
	}

	if readHeader.GenesisID != "GENESIS" {
		t.Fatal("Header mismatch")
	}

	if len(entries) != 5 {
		t.Fatalf("Expected 5 entries, got %d", len(entries))
	}
}

func TestPersistorCrashRecovery(t *testing.T) {
	path := "test_crash.log"
	defer os.Remove(path)

	p, _ := NewPersistor(path)
	p.WriteHeader(LedgerFileHeader{Version: "1.0", FixedPointDivisor: 1000000})

	for i := 0; i < 5; i++ {
		p.Append(LedgerEntry{EventID: "EVT", LogicalTick: uint64(i)})
	}

	// Truncate last line to simulate crash
	f, _ := os.OpenFile(path, os.O_RDWR, 0600)
	stat, _ := f.Stat()
	f.Truncate(stat.Size() - 5) // truncate some bytes to break JSON
	f.Close()

	entries, _, err := p.ReadAll()
	if err != nil {
		t.Fatal(err)
	}

	if len(entries) != 4 {
		t.Fatalf("Expected 4 entries after crash recovery, got %d", len(entries))
	}
}

func TestPersistorHeaderValidation(t *testing.T) {
	path := "test_invalid_header.log"
	defer os.Remove(path)

	os.WriteFile(path, []byte("{invalid json}\n"), 0600)

	p, _ := NewPersistor(path)
	_, _, err := p.ReadAll()
	if err == nil {
		t.Fatal("Expected error on invalid header")
	}
}
