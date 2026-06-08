package engine

import (
	"bytes"
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"

	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
	"github.com/fallofpheonix/phoenix/internal/consensus"
	"github.com/fallofpheonix/phoenix/internal/contracts"
	"github.com/fallofpheonix/phoenix/internal/protocol"
)

// OpCode defines the instructions for the pscript VM.
type OpCode uint8

const (
	OpPush OpCode = iota
	OpMove
	OpVerify
	OpHalt
)

// Instruction represents a single command for the VM.
type Instruction struct {
	Op   OpCode
	Args []interface{}
}

// Entity represents an entity in the game world.
type Entity struct {
	ID     string             `json:"id"`
	Pos    phxmath.FixedPoint `json:"pos"`
	Status string             `json:"status"`
}

// NewWorldState initializes a blank WorldState.
func NewWorldState(seed int64) *WorldState {
	return &WorldState{
		Seed:              seed,
		entities:          make(map[string]*Entity),
		LastSeenSequences: make(map[string]uint64),
		Validators:        [][]byte{},
		PendingValidators: [][]byte{},
	}
}

// Digest produces a canonical binary digest of the entity state.
func (e *Entity) Digest() []byte {
	h := sha256.New()
	b2 := make([]byte, 2)
	b8 := make([]byte, 8)

	binary.BigEndian.PutUint16(b2, uint16(len(e.ID)))
	h.Write(b2)
	h.Write([]byte(e.ID))

	binary.BigEndian.PutUint64(b8, uint64(e.Pos.V))
	h.Write(b8)

	binary.BigEndian.PutUint16(b2, uint16(len(e.Status)))
	h.Write(b2)
	h.Write([]byte(e.Status))

	return h.Sum(nil)
}

// WorldState represents the current state of the game world.
type WorldState struct {
	Tick              uint64             `json:"tick"`
	Epoch             uint64             `json:"epoch"`
	EventCount        uint64             `json:"event_count"`
	Seed              int64              `json:"seed"`
	entities          map[string]*Entity `json:"entities"`
	Validators        [][]byte           `json:"validators"`
	PendingValidators [][]byte           `json:"pending_validators"`
	LastSeenSequences map[string]uint64  `json:"last_seen_sequences"`
	LastEventHash     contracts.Hash     `json:"last_event_hash"`
	StateHash         contracts.Hash     `json:"state_hash"`
}

// MarshalJSON customizes the JSON representation of WorldState to include private fields.
func (ws WorldState) MarshalJSON() ([]byte, error) {
	type Alias WorldState
	return json.Marshal(&struct {
		Entities map[string]*Entity `json:"entities"`
		*Alias
	}{
		Entities: ws.entities,
		Alias:    (*Alias)(&ws),
	})
}

// UnmarshalJSON customizes the JSON unmarshaling of WorldState to populate private fields.
func (ws *WorldState) UnmarshalJSON(data []byte) error {
	type Alias WorldState
	aux := &struct {
		Entities map[string]*Entity `json:"entities"`
		*Alias
	}{
		Alias: (*Alias)(ws),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	ws.entities = aux.Entities
	return nil
}

// GetEntities returns a copy of the entities map for read-only access.
func (ws *WorldState) GetEntities() map[string]*Entity {
	copy := make(map[string]*Entity, len(ws.entities))
	for k, v := range ws.entities {
		copy[k] = v
	}
	return copy
}

// GetEntity returns a specific entity by ID.
func (ws *WorldState) GetEntity(id string) (*Entity, bool) {
	e, ok := ws.entities[id]
	return e, ok
}

// InjectEntity allows tests to inject an entity directly into the state.
func InjectEntity(ws *WorldState, e *Entity) {
	if ws.entities == nil {
		ws.entities = make(map[string]*Entity)
	}
	ws.entities[e.ID] = e
}

// CalculateHash produces a deterministic cryptographic hash of the entire world state.
func (ws *WorldState) CalculateHash() contracts.Hash {
	h := sha256.New()
	b2 := make([]byte, 2)
	b8 := make([]byte, 8)

	binary.BigEndian.PutUint64(b8, ws.Tick)
	h.Write(b8)
	binary.BigEndian.PutUint64(b8, ws.Epoch)
	h.Write(b8)
	binary.BigEndian.PutUint64(b8, uint64(ws.Seed))
	h.Write(b8)
	h.Write(ws.LastEventHash[:])

	binary.BigEndian.PutUint16(b2, uint16(len(ws.Validators)))
	h.Write(b2)
	vKeys := make([]string, len(ws.Validators))
	for i, k := range ws.Validators {
		vKeys[i] = string(k)
	}
	sort.Strings(vKeys)
	for _, k := range vKeys {
		binary.BigEndian.PutUint16(b2, uint16(len(k)))
		h.Write(b2)
		h.Write([]byte(k))
	}

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

	binary.BigEndian.PutUint16(b2, uint16(len(ws.entities)))
	h.Write(b2)
	ids := make([]string, 0, len(ws.entities))
	for id := range ws.entities {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	for _, id := range ids {
		h.Write(ws.entities[id].Digest())
	}

	var res contracts.Hash
	copy(res[:], h.Sum(nil))
	return res
}

const (
	EpochInterval = 100
)

// ApplyEvent applies a single event to the world state.
func (ws *WorldState) ApplyEvent(applied contracts.AppliedEvent, rules map[string]VerificationRule) error {
	event := applied.Event

	if applied.Epoch > ws.Epoch {
		ws.processEpochTransition(applied.Epoch)
	}

	// Physics / Semantic Logic
	switch event.Type {
	case contracts.EventSpawn:
		var payload struct {
			ID  string             `json:"id"`
			Pos phxmath.FixedPoint `json:"pos"`
		}
		json.Unmarshal(event.Payload, &payload)
		ws.entities[payload.ID] = &Entity{
			ID:     payload.ID,
			Pos:    payload.Pos,
			Status: "SPAWNED",
		}
	case contracts.EventMove:
		var payload struct {
			ID  string             `json:"id"`
			Pos phxmath.FixedPoint `json:"pos"`
		}
		json.Unmarshal(event.Payload, &payload)
		entity, ok := ws.entities[payload.ID]
		if !ok {
			return fmt.Errorf("entity %s not found", payload.ID)
		}
		entity.Pos = payload.Pos
	case contracts.EventVerify:
		var payload struct {
			ID string `json:"id"`
		}
		json.Unmarshal(event.Payload, &payload)
		entity, ok := ws.entities[payload.ID]
		if !ok {
			return fmt.Errorf("entity %s not found", payload.ID)
		}
		if rule, ok := rules[payload.ID]; ok {
			success, reason := rule.Verify(ws, payload.ID)
			if success {
				entity.Status = "VERIFIED"
			} else {
				entity.Status = "FAILED:" + reason
			}
		}
	case contracts.EventUpdateValidator:
		var payload struct {
			ID     string `json:"id"`
			Status string `json:"status"`
		}
		json.Unmarshal(event.Payload, &payload)
		vKey, _ := hex.DecodeString(payload.ID)
		if len(ws.PendingValidators) == 0 && len(ws.Validators) > 0 {
			ws.PendingValidators = make([][]byte, len(ws.Validators))
			copy(ws.PendingValidators, ws.Validators)
		}
		if payload.Status == "ADD" {
			ws.PendingValidators = append(ws.PendingValidators, vKey)
		} else if payload.Status == "REMOVE" {
			for i, v := range ws.PendingValidators {
				if bytes.Equal(v, vKey) {
					ws.PendingValidators = append(ws.PendingValidators[:i], ws.PendingValidators[i+1:]...)
					break
				}
			}
		}
	}

	ws.Tick = applied.Height
	ws.EventCount++

	d, _ := protocol.DigestEvent(event)
	ws.LastEventHash = d
	ws.StateHash = ws.CalculateHash()
	return nil
}

func (ws *WorldState) ApplyEnvelope(env *contracts.SignedEnvelope) {
	if ws.LastSeenSequences == nil {
		ws.LastSeenSequences = make(map[string]uint64)
	}
	pubKeyHex := hex.EncodeToString(env.Validator[:])
	ws.LastSeenSequences[pubKeyHex] = env.Sequence
	ws.StateHash = ws.CalculateHash()
}

func (ws *WorldState) VerifyEnvelope(env *contracts.SignedEnvelope) error {
	pubKeyHex := hex.EncodeToString(env.Validator[:])
	lastSeq, ok := ws.LastSeenSequences[pubKeyHex]
	if ok && env.Sequence <= lastSeq {
		return fmt.Errorf("replay detected: sequence %d <= %d", env.Sequence, lastSeq)
	}
	return nil
}

func (ws *WorldState) processEpochTransition(newEpoch uint64) {
	if len(ws.PendingValidators) > 0 {
		ws.Validators = make([][]byte, len(ws.PendingValidators))
		for i, v := range ws.PendingValidators {
			ws.Validators[i] = make([]byte, len(v))
			copy(ws.Validators[i], v)
		}
	}
	ws.Epoch = newEpoch
}

func (ws *WorldState) Snapshot() ([]byte, error) {
	return json.Marshal(ws)
}

func (ws *WorldState) Restore(data []byte) error {
	return json.Unmarshal(data, ws)
}

const (
	MinValidatorsForBFT = 4
)

func (ws *WorldState) CheckQuorum(signatures []contracts.SignatureEntry, digest contracts.Hash) (bool, error) {
	if len(ws.Validators) < MinValidatorsForBFT && len(ws.Validators) > 0 {
		return false, fmt.Errorf("insufficient validators for BFT: %d (min %d)", len(ws.Validators), MinValidatorsForBFT)
	}
	return consensus.CheckQuorum(ws.Validators, signatures, digest)
}

func Replay(initial *WorldState, applied []contracts.AppliedEvent, rules map[string]VerificationRule) (*WorldState, error) {
	ws := initial
	if ws == nil {
		ws = NewWorldState(0)
	}
	for _, a := range applied {
		if err := ws.ApplyEvent(a, rules); err != nil {
			return nil, err
		}
	}
	return ws, nil
}

type VM struct {
	State      WorldState
	Events     []contracts.Event
	Stack      []interface{}
	PC         int
	Code       []Instruction
	Rules      map[string]VerificationRule
	PrivateKey ed25519.PrivateKey
}

func NewVM(code []Instruction) *VM {
	_, priv, _ := ed25519.GenerateKey(nil)
	return &VM{
		State: WorldState{
			Tick:              0,
			entities:          make(map[string]*Entity),
			LastSeenSequences: make(map[string]uint64),
		},
		Events:     []contracts.Event{},
		Stack:      []interface{}{},
		PC:         0,
		Code:       code,
		Rules:      make(map[string]VerificationRule),
		PrivateKey: priv,
	}
}

func (vm *VM) Step() error {
	if vm.PC >= len(vm.Code) {
		return fmt.Errorf("PC out of bounds")
	}
	instr := vm.Code[vm.PC]
	vm.PC++

	var event contracts.Event
	event.Version = 1

	switch instr.Op {
	case OpPush:
		vm.Stack = append(vm.Stack, instr.Args[0])
		return nil
	case OpMove:
		pos := vm.Stack[len(vm.Stack)-1].(phxmath.FixedPoint)
		id := vm.Stack[len(vm.Stack)-2].(string)
		vm.Stack = vm.Stack[:len(vm.Stack)-2]
		event.Type = contracts.EventMove
		event.Payload, _ = json.Marshal(map[string]interface{}{"id": id, "pos": pos})
	case OpVerify:
		id := vm.Stack[len(vm.Stack)-1].(string)
		vm.Stack = vm.Stack[:len(vm.Stack)-1]
		event.Type = contracts.EventVerify
		event.Payload, _ = json.Marshal(map[string]interface{}{"id": id})
	default:
		return fmt.Errorf("unknown opcode: %v", instr.Op)
	}

	applied := contracts.AppliedEvent{
		Height: vm.State.Tick + 1,
		Epoch:  vm.State.Epoch,
		Event:  event,
	}

	if err := vm.State.ApplyEvent(applied, vm.Rules); err != nil {
		return err
	}
	vm.Events = append(vm.Events, event)
	return nil
}

func (vm *VM) Run() error {
	for vm.PC < len(vm.Code) {
		if err := vm.Step(); err != nil {
			return err
		}
	}
	return nil
}

func (vm *VM) ToJSON() (string, error) {
	b, err := json.MarshalIndent(vm.State, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
