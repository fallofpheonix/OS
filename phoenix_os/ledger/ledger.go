package ledger

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

type LedgerEntry struct {
	SeqID    int64
	ActionID string
	CauseID  string
	Payload  []byte
	PrevHash []byte
	Hash     []byte
}

type Ledger struct {
	Entries []LedgerEntry
	counter int64
}

func (l *Ledger) AddEntry(actionID, causeID string, payload []byte) error {
	prevHash := make([]byte, 32)
	if len(l.Entries) > 0 {
		prevHash = l.Entries[len(l.Entries)-1].Hash
	}

	entry := LedgerEntry{
		SeqID:    l.counter,
		ActionID: actionID,
		CauseID:  causeID,
		Payload:  payload,
		PrevHash: prevHash,
	}
	l.counter++

	h := sha256.New()
	binary.Write(h, binary.BigEndian, entry.SeqID)
	h.Write([]byte(entry.ActionID))
	h.Write([]byte(entry.CauseID))
	h.Write(entry.PrevHash)
	h.Write(entry.Payload)

	entry.Hash = h.Sum(nil)
	l.Entries = append(l.Entries, entry)
	return nil
}

func (l *Ledger) Verify() error {
	var prev []byte
	for i, entry := range l.Entries {
		if i > 0 {
			if string(entry.PrevHash) != string(prev) {
				return fmt.Errorf("hash chain broken at entry %d", i)
			}
		}

		h := sha256.New()
		binary.Write(h, binary.BigEndian, entry.SeqID)
		h.Write([]byte(entry.ActionID))
		h.Write([]byte(entry.CauseID))
		h.Write(entry.PrevHash)
		h.Write(entry.Payload)

		computed := h.Sum(nil)
		if string(computed) != string(entry.Hash) {
			return fmt.Errorf("invalid hash at entry %d", i)
		}
		prev = entry.Hash
	}
	return nil
}
