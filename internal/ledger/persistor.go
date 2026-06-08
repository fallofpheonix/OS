package ledger

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

// Persistor handles writing events to disk with durability guarantees.
type Persistor struct {
	mu   sync.Mutex
	file *os.File
	path string
}

// NewPersistor opens or creates the ledger file at the specified path.
func NewPersistor(path string) (*Persistor, error) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open persistor file: %w", err)
	}
	return &Persistor{
		file: f,
		path: path,
	}, nil
}

// Append persists a single event to disk and ensures it is synced.
func (p *Persistor) Append(e *Event) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	data, err := json.Marshal(e)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	if _, err := p.file.Write(append(data, '\n')); err != nil {
		return fmt.Errorf("failed to write event to disk: %w", err)
	}

	// Ensure durability
	if err := p.file.Sync(); err != nil {
		return fmt.Errorf("failed to sync file: %w", err)
	}

	return nil
}

// Load reads all events from disk into the provided Chain.
// It stops at the first corrupted line and returns an error.
func (p *Persistor) Load(c *Chain) (int, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	// Reset file pointer to beginning
	if _, err := p.file.Seek(0, 0); err != nil {
		return 0, fmt.Errorf("failed to seek to start of file: %w", err)
	}

	scanner := bufio.NewScanner(p.file)
	count := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}

		var e Event
		if err := json.Unmarshal(line, &e); err != nil {
			return count, fmt.Errorf("corruption detected at event %d: %w", count, err)
		}

		if err := c.Append(&e); err != nil {
			return count, fmt.Errorf("integrity check failed for event %d: %w", count, err)
		}
		count++
	}

	if err := scanner.Err(); err != nil {
		return count, fmt.Errorf("error reading persistor file: %w", err)
	}

	return count, nil
}

// Close closes the underlying file handle.
func (p *Persistor) Close() error {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.file.Close()
}
