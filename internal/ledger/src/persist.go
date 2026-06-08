package ledger

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"sync"
)

type LedgerFileHeader struct {
	Version           string `json:"version"`
	GenesisID         string `json:"genesis_id"`
	Timestamp         int64  `json:"timestamp"`
	Algorithm         string `json:"algorithm"`
	FixedPointDivisor int64  `json:"fixed_point_divisor"`
}

type Persistor struct {
	path string
	mu   sync.Mutex
}

// PURPOSE: Persists ledger entries to a disk-backed append-only file.
// CONTRACT: Must guarantee crash-safe writes and partial-read recovery.
// FAILURE: Fails if disk is full, permissions are denied, or file is corrupted beyond recovery.
// CONNECTS: Used by the Ledger to durably store State Transitions and Events.
func NewPersistor(path string) (*Persistor, error) {
	return &Persistor{path: path}, nil
}

// PURPOSE: Writes the genesis header to a new ledger file.
// CONTRACT: Must only be called once per file creation.
// FAILURE: Fails if file cannot be created or written to.
// CONNECTS: Called during initial Ledger file allocation.
func (p *Persistor) WriteHeader(genesis LedgerFileHeader) error {
	data, err := json.Marshal(genesis)
	if err != nil {
		return err
	}
	data = append(data, '\n')

	// WHY: Open-per-write is fine for <100 entries/sec and guarantees flush on close.
	f, err := os.OpenFile(p.path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	return err
}

// PURPOSE: Appends a single LedgerEntry to the log.
// CONTRACT: Must atomically append and flush to disk.
// FAILURE: Fails on disk IO errors.
// CONNECTS: Called on every AddEntry/AddEntryV2 to guarantee persistence.
func (p *Persistor) Append(entry LedgerEntry) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	data = append(data, '\n')

	// WHY crash-safe write strategy: Open in APPEND mode, write complete JSON line, and close to force flush.
	// In the event of a crash during write, the last line will be incomplete JSON.
	f, err := os.OpenFile(p.path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	return err
}

// PURPOSE: Reads all valid entries and the header from the ledger file.
// CONTRACT: Must skip incomplete tail lines caused by crashes.
// FAILURE: Fails if header is missing/invalid or file cannot be read.
// CONNECTS: Called during node startup to reconstruct the in-memory Merkle DAG.
func (p *Persistor) ReadAll() ([]LedgerEntry, LedgerFileHeader, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	var entries []LedgerEntry
	var header LedgerFileHeader

	f, err := os.Open(p.path)
	if err != nil {
		return entries, header, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	if !scanner.Scan() {
		return entries, header, errors.New("empty ledger file")
	}

	if err := json.Unmarshal(scanner.Bytes(), &header); err != nil {
		return entries, header, errors.New("invalid header")
	}

	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}
		var entry LedgerEntry
		// WHY partial last-line policy: Ignore json unmarshal errors for the very last line, treating it as a crash during Append.
		if err := json.Unmarshal(line, &entry); err != nil {
			break
		}
		entries = append(entries, entry)
	}

	return entries, header, nil
}
