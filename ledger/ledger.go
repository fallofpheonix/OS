package ledger

type LedgerEntry struct {
	LogicalTick   uint64
	EventID       string
	CauseID       string
	Source        string
	PolicyVersion string
	StateBefore   []byte
	StateAfter    []byte
	ParentIDs     [][]byte
	Payload       []byte
	PrevHash      []byte
	ReplayHash    []byte
	Hash          []byte
}

type Ledger struct {
	Entries map[string]LedgerEntry
	Heads   [][]byte
	Counter uint64
}

func NewLedger() *Ledger {
	return &Ledger{Entries: make(map[string]LedgerEntry)}
}

func (l *Ledger) AddEntry(eventID, causeID string, payload []byte) error {
	return nil
}
