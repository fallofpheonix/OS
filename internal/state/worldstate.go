package state

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"sort"

	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
	"github.com/fallofpheonix/phoenix/internal/contracts"
)

// Entity represents an entity in the game world.
type Entity struct {
	ID     string             `json:"id"`
	Pos    phxmath.FixedPoint `json:"pos"`
	Status string             `json:"status"`
}

// Digest produces a canonical binary digest of the entity state.
func (e *Entity) Digest() []byte {
	h := sha256.New()
	b2 := make([]byte, 2)
	b8 := make([]byte, 8)

	// ID (Length-prefixed)
	binary.BigEndian.PutUint16(b2, uint16(len(e.ID)))
	h.Write(b2)
	h.Write([]byte(e.ID))

	// Position (FixedPoint raw value)
	binary.BigEndian.PutUint64(b8, uint64(e.Pos.V))
	h.Write(b8)

	// Status (Length-prefixed)
	binary.BigEndian.PutUint16(b2, uint16(len(e.Status)))
	h.Write(b2)
	h.Write([]byte(e.Status))

	return h.Sum(nil)
}

// worldState is the unexported internal state of the OS physics.
// INV-004: All mutations must pass through the StateGuard.
type worldState struct {
	Tick              uint64             `json:"tick"`
	Epoch             uint64             `json:"epoch"`
	EventCount        uint64             `json:"event_count"`
	Seed              int64              `json:"seed"`
	Entities          map[string]*Entity `json:"entities"`
	Validators        [][]byte           `json:"validators"`
	PendingValidators [][]byte           `json:"pending_validators"`
	LastSeenSequences map[string]uint64  `json:"last_seen_sequences"`
	LastEventHash     contracts.Hash     `json:"last_event_hash"`
	StateHash         contracts.Hash     `json:"state_hash"`
}

func (ws *worldState) MarshalJSON() ([]byte, error) {
	type Alias worldState
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(ws),
	})
}

func (ws *worldState) UnmarshalJSON(data []byte) error {
	type Alias worldState
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(ws),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}

func newWorldState(seed int64) *worldState {
	return &worldState{
		Seed:              seed,
		Entities:          make(map[string]*Entity),
		LastSeenSequences: make(map[string]uint64),
		Validators:        [][]byte{},
		PendingValidators: [][]byte{},
	}
}

// calculateHash produces the deterministic cryptographic hash of the state.
// PROTOCOL-020: Deterministic Execution Law - Map iteration is sorted.
func (ws *worldState) calculateHash() contracts.Hash {
	h := sha256.New()
	b2 := make([]byte, 2)
	b8 := make([]byte, 8)

	// 1. Tick & Epoch
	binary.BigEndian.PutUint64(b8, ws.Tick)
	h.Write(b8)
	binary.BigEndian.PutUint64(b8, ws.Epoch)
	h.Write(b8)

	// 2. Seed
	binary.BigEndian.PutUint64(b8, uint64(ws.Seed))
	h.Write(b8)

	// 3. LastEventHash
	h.Write(ws.LastEventHash[:])

	// 4. Validators (Sorted)
	binary.BigEndian.PutUint16(b2, uint16(len(ws.Validators)))
	h.Write(b2)
	vKeys := make([]string, len(ws.Validators))
	for i, v := range ws.Validators {
		vKeys[i] = string(v)
	}
	sort.Strings(vKeys)
	for _, k := range vKeys {
		binary.BigEndian.PutUint16(b2, uint16(len(k)))
		h.Write(b2)
		h.Write([]byte(k))
	}

	// 5. PendingValidators (Sorted)
	binary.BigEndian.PutUint16(b2, uint16(len(ws.PendingValidators)))
	h.Write(b2)
	pKeys := make([]string, len(ws.PendingValidators))
	for i, v := range ws.PendingValidators {
		pKeys[i] = string(v)
	}
	sort.Strings(pKeys)
	for _, k := range pKeys {
		binary.BigEndian.PutUint16(b2, uint16(len(k)))
		h.Write(b2)
		h.Write([]byte(k))
	}

	// 6. Sequences (Sorted)
	binary.BigEndian.PutUint16(b2, uint16(len(ws.LastSeenSequences)))
	h.Write(b2)
	sKeys := make([]string, 0, len(ws.LastSeenSequences))
	for k := range ws.LastSeenSequences {
		sKeys = append(sKeys, k)
	}
	sort.Strings(sKeys)
	for _, k := range sKeys {
		binary.BigEndian.PutUint16(b2, uint16(len(k)))
		h.Write(b2)
		h.Write([]byte(k))
		binary.BigEndian.PutUint64(b8, ws.LastSeenSequences[k])
		h.Write(b8)
	}

	// 7. Entities (Sorted by ID)
	binary.BigEndian.PutUint16(b2, uint16(len(ws.Entities)))
	h.Write(b2)
	eIDs := make([]string, 0, len(ws.Entities))
	for id := range ws.Entities {
		eIDs = append(eIDs, id)
	}
	sort.Strings(eIDs)
	for _, id := range eIDs {
		h.Write(ws.Entities[id].Digest())
	}

	var res contracts.Hash
	copy(res[:], h.Sum(nil))
	return res
}
