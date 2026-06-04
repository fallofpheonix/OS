package ledger

import (
	"bufio"
	"encoding/json"
	"os"
)

// Persistor handles writing events to disk.
type Persistor struct {
	file *os.File
}

// NewPersistor opens or creates the ledger file.
func NewPersistor(path string) (*Persistor, error) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &Persistor{file: f}, nil
}

// Append persists a single event to disk.
func (p *Persistor) Append(e *Event) error {
	data, err := json.Marshal(e)
	if err != nil {
		return err
	}
	_, err = p.file.Write(append(data, '\n'))
	return err
}

// Load reads all events from disk into a Chain.
func (p *Persistor) Load(c *Chain) error {
	// Re-open in read mode
	f, err := os.Open(p.file.Name())
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var e Event
		if err := json.Unmarshal(scanner.Bytes(), &e); err != nil {
			return err
		}
		if err := c.Append(&e); err != nil {
			return err
		}
	}
	return scanner.Err()
}

func (p *Persistor) Close() error {
	return p.file.Close()
}
