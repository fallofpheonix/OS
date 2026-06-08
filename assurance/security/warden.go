// Package security implements the tactical enforcement FSM for PhoenixOS.
// Domain Logic: Orchestrates state transition validation against formal invariants and triggers physical actuators.
// Responsibility: Acts as the central enforcement engine to maintain system integrity and containment.
package security

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	securityv1 "github.com/fallofpheonix/phoenix/foundation/contracts/security/v1"
	ledger "github.com/fallofpheonix/phoenix/foundation/ledger/src"
	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
	"github.com/fallofpheonix/phoenix/foundation/runtime/bus"
	"github.com/fallofpheonix/phoenix/foundation/runtime/common/config"
)

const AuditBufferSize = 1024 // Power of two for bitwise & 1023

// SystemState represents the FSM level.
type SystemState uint8

const (
	StateInvalid     SystemState = 0x00
	StateSafe        SystemState = 0x01
	StateWatch       SystemState = 0x02
	StateSuspicious  SystemState = 0x03
	StateCritical    SystemState = 0x04
	StateCompromised SystemState = 0x05
)

// ActuationClass represents the physical action to be taken.
type ActuationClass uint8

const (
	ClassInvalid  ActuationClass = 0x00
	ClassNone     ActuationClass = 0x01
	ClassLog      ActuationClass = 0x02
	ClassThrottle ActuationClass = 0x03
	ClassIsolate  ActuationClass = 0x04
	ClassKill     ActuationClass = 0x05
)

// AuditEntry represents a forensically-anchored record in the STATE-002 ring buffer.
type AuditEntry struct {
	Index     uint64
	Timestamp int64
	Content   string
	PrevHash  [32]byte
	Hash      [32]byte
}

// PostureSnapshot ensures decision consistency within a single execution round.
type PostureSnapshot struct {
	State       SystemState
	ShadowState SystemState
	ShadowMode  bool
	Policy      *config.RedLines
	PolicyHash  [32]byte
}

type GraphProof struct {
	Path            []string
	ExpectedNsproxy uint32
}

type AuthorityEscalationRequest struct {
	EventID        string
	TargetPID      uint32
	TargetTgid     uint32
	TargetNsproxy  uint32
	TargetState    SystemState
	ActuationClass ActuationClass
	EvidenceWeight phxmath.FixedPoint
	Certificate    []byte
	PolicyHash     [32]byte
	GraphProof     *GraphProof
}

// ActuationTask wraps a request for the async drainer.
type ActuationTask struct {
	req     AuthorityEscalationRequest
	tick    uint64
	context context.Context
	cancel  context.CancelFunc
}

type pendingEntry struct {
	startTick     uint64
	lastHeartbeat uint64
}

// Invariant defines the interface for formal proof-gates.
type Invariant interface {
	Verify(req AuthorityEscalationRequest, snap PostureSnapshot) error
}

// GraphProvider verifies causal paths extracted from GraphProof.
type GraphProvider interface {
	VerifyPath(path []string) (bool, error)
}

// Warden is the hardened tactical enforcement engine.
type Warden struct {
	// 1. Fast-Path Posture (Lock-Free)
	state       atomic.Uint32 // Under-the-hood SystemState
	shadowState atomic.Uint32 // Ghost state for divergence detection
	shadowMode  atomic.Bool
	locked      atomic.Bool
	policies    atomic.Pointer[config.RedLines]
	stateMu     sync.Mutex // Protects transition logic serialization

	// 2. Audit Substrate (STATE-002 Hash-Chain)
	auditHead   atomic.Uint64
	auditMu     sync.Mutex // Protects hash-chain RMW
	auditBuffer [AuditBufferSize]AuditEntry
	lastHash    atomic.Value // Stores [32]byte

	// 3. Async Actuator Queues (NP + HP)
	hpQueue  chan ActuationTask
	npQueue  chan ActuationTask
	stopChan chan struct{}

	// 4. Dependencies (Unexported for Nil-Safety)
	bus    *bus.Bus
	ledger *ledger.Ledger

	// 5. Configuration (Mutex Protected - Setup Only)
	mu         sync.RWMutex
	invariants []Invariant
	actuators  []securityv1.Actuator

	// 6. Causal Integrity (RECTIFIED: Debt 1)
	pendingMu  sync.Mutex
	pendingMap map[string]pendingEntry // CauseID -> Metadata
	epoch      atomic.Uint64           // Current Consensus Epoch

	// 7. I/O Isolation (RECTIFIED: Debt 2)
	ledgerMu sync.Mutex // Dedicated lock for ledger writes
}

// NewWarden initializes the enforcement substrate with fail-closed defaults.
func NewWarden(b *bus.Bus, l *ledger.Ledger) *Warden {
	// [RECTIFIED]: Constraint 1 - Nil-safe approach.
	// Ledger and Bus are optional. Tests and staging environments may
	// operate without durable storage. This is not an error.

	w := &Warden{
		bus:        b,
		ledger:     l,
		hpQueue:    make(chan ActuationTask, 128),
		npQueue:    make(chan ActuationTask, 1024),
		stopChan:   make(chan struct{}),
		pendingMap: make(map[string]pendingEntry),
	}

	w.state.Store(uint32(StateSafe))
	w.shadowState.Store(uint32(StateSafe))
	w.shadowMode.Store(true) // AXIOM 3: Guarded Autonomy
	w.lastHash.Store([32]byte{})
	w.policies.Store(&config.RedLines{})

	return w
}

func (w *Warden) Start() {
	go w.supervisor()
}

func (w *Warden) Stop() {
	close(w.stopChan)
}

func (w *Warden) supervisor() {
	var restarts int
	for restarts < 3 {
		func() {
			defer func() {
				if r := recover(); r != nil {
					restarts++
					if restarts >= 3 {
						w.Lock()
					}
				}
			}()
			w.drainLoop()
		}()

		select {
		case <-w.stopChan:
			return
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (w *Warden) drainLoop() {
	for {
		var task ActuationTask
		// Priority Select Pattern
		select {
		case task = <-w.hpQueue:
		case <-w.stopChan:
			w.flushAndExit()
			return
		default:
			select {
			case task = <-w.hpQueue:
			case task = <-w.npQueue:
			case <-w.stopChan:
				w.flushAndExit()
				return
			}
		}
		w.executeActuation(task)
	}
}

func (w *Warden) flushAndExit() {
	for {
		select {
		case t := <-w.hpQueue:
			w.executeActuation(t)
		case t := <-w.npQueue:
			w.executeActuation(t)
		default:
			return
		}
	}
}

func (w *Warden) executeActuation(t ActuationTask) {
	defer t.cancel()

	var err error
	actuators := w.getActuators()
	for _, act := range actuators {
		// [CYCLE 9] Physical Actuation
		switch t.req.ActuationClass {
		case ClassKill:
			err = act.Kill(t.context, int(t.req.TargetPID))
		case ClassIsolate:
			err = act.Actuate(t.context, localContainment{level: securityv1.LevelIsolate})
		}
	}

	// [CYCLE 10] Record COMPLETION
	status := "SUCCESS"
	if err != nil {
		status = fmt.Sprintf("FAILED:%v", err)
		if errors.Is(err, syscall.ESRCH) {
			status = "SUCCESS_GONE"
		}
	}

	compID := fmt.Sprintf("COMP-%s", t.req.EventID)
	stateBytes := []byte{uint8(t.req.TargetState)}

	// Ledger Write Serialization
	w.ledgerMu.Lock()
	if w.ledger != nil {
		_ = w.ledger.AddEntryV2(
			compID,
			t.req.EventID,
			t.tick,
			[]byte(status),
			"",
			stateBytes,
			stateBytes,
			fmt.Sprintf("%x", t.req.PolicyHash),
		)
	}
	w.ledgerMu.Unlock()

	// 2. Clear Pending Map
	w.pendingMu.Lock()
	delete(w.pendingMap, t.req.EventID)
	w.pendingMu.Unlock()

	w.appendAudit(fmt.Sprintf("COMPLETION|ID:%s|STATUS:%s", t.req.EventID, status), t.tick)
}

func (w *Warden) getActuators() []securityv1.Actuator {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.actuators
}

type localContainment struct {
	level securityv1.ContainmentLevel
}

func (c localContainment) Level() securityv1.ContainmentLevel { return c.level }
func (c localContainment) Target() string                     { return "SYSTEM" }
func (c localContainment) Reason() string                     { return "WARDEN_FSM_ESCALATION" }

func (s SystemState) String() string {
	switch s {
	case StateInvalid:
		return "INVALID"
	case StateSafe:
		return "SAFE"
	case StateWatch:
		return "WATCH"
	case StateSuspicious:
		return "SUSPICIOUS"
	case StateCritical:
		return "CRITICAL"
	case StateCompromised:
		return "COMPROMISED"
	default:
		return fmt.Sprintf("UNKNOWN(0x%02x)", uint8(s))
	}
}

var stateThresholds = map[SystemState]phxmath.FixedPoint{
	StateWatch:       phxmath.NewFixedPointRaw(200000), // 0.20
	StateSuspicious:  phxmath.NewFixedPointRaw(500000), // 0.50
	StateCritical:    phxmath.NewFixedPointRaw(800000), // 0.80
	StateCompromised: phxmath.NewFixedPointRaw(950000), // 0.95
}

func GetThreshold(s SystemState) phxmath.FixedPoint {
	if t, ok := stateThresholds[s]; ok {
		return t
	}
	return phxmath.NewFixedPointRaw(1000000) // Default 1.0
}

// allowedTransitions defines the static topology (Fast-Up, Slow-Down).
var allowedTransitions = [6][6]bool{
	StateInvalid: {},
	StateSafe: {
		StateSafe: true, StateWatch: true, StateSuspicious: true, StateCritical: true, StateCompromised: true,
	},
	StateWatch: {
		StateSafe: true, StateWatch: true, StateSuspicious: true, StateCritical: true, StateCompromised: true,
	},
	StateSuspicious: {
		StateWatch: true, StateSuspicious: true, StateCritical: true, StateCompromised: true,
	},
	StateCritical: {
		StateSuspicious: true, StateCritical: true, StateCompromised: true,
	},
	StateCompromised: {}, // Handled by validateTransition override
}

// serializePayload produces the authoritative 43-byte binary representation.
func serializePayload(req AuthorityEscalationRequest, shadow bool, shadowCurr SystemState) ([]byte, error) {
	/*
	   LAYOUT SPEC (BFT-SPEC-004):
	   [0]    : Version (0x01)
	   [1]    : ActuationClass (uint8)
	   [2-5]  : TargetPID (uint32, BigEndian, capped at 0x7FFFFFFF)
	   [6-37] : PolicyHash (32 bytes)
	   [38]   : TargetState (uint8)
	   [39]   : IsShadow (0x01=true, 0x00=false)
	   [40]   : ShadowCurr (uint8)
	   [41-42]: Reserved (Zeroed)
	*/
	buf := make([]byte, 43)
	buf[0] = 0x01
	buf[1] = uint8(req.ActuationClass)

	pid := req.TargetPID
	if pid > 0x7FFFFFFF {
		pid = 0x7FFFFFFF
	}
	binary.BigEndian.PutUint32(buf[2:6], pid)

	copy(buf[6:38], req.PolicyHash[:])

	buf[38] = uint8(req.TargetState)
	if shadow {
		buf[39] = 0x01
	}
	buf[40] = uint8(shadowCurr)

	return buf, nil
}

// deserializePayload reconstructs the request from the ledger.
func deserializePayload(data []byte) (req AuthorityEscalationRequest, shadow bool, shadowCurr SystemState, err error) {
	if len(data) != 43 {
		return req, false, 0, fmt.Errorf("invalid payload length: expected 43, got %d", len(data))
	}
	if data[0] != 0x01 {
		return req, false, 0, fmt.Errorf("unsupported payload version: 0x%02x", data[0])
	}

	// 1. Range Validation
	class := ActuationClass(data[1])
	targetState := SystemState(data[38])
	isShadowByte := data[39]
	currState := SystemState(data[40])

	if class < ClassNone || class > ClassKill {
		return req, false, 0, fmt.Errorf("semantic violation: invalid actuation class 0x%02x", data[1])
	}
	if targetState < StateSafe || targetState > StateCompromised {
		return req, false, 0, fmt.Errorf("semantic violation: invalid target state 0x%02x", data[38])
	}
	if isShadowByte > 0x01 {
		return req, false, 0, fmt.Errorf("semantic violation: invalid isShadow byte 0x%02x", data[39])
	}
	if currState < StateSafe || currState > StateCompromised {
		return req, false, 0, fmt.Errorf("semantic violation: invalid current state 0x%02x", data[40])
	}

	req.ActuationClass = class
	req.TargetPID = binary.BigEndian.Uint32(data[2:6])
	copy(req.PolicyHash[:], data[6:38])
	req.TargetState = targetState
	shadow = isShadowByte == 0x01
	shadowCurr = currState

	return req, shadow, shadowCurr, nil
}

// ActuateRequest is the core decision gate for system escalation.
func (w *Warden) ActuateRequest(req AuthorityEscalationRequest, seq int, lamport uint64) (bool, error) {
	/*
	   CANONICAL ENFORCEMENT FLOW (BFT-SPEC-005):
	   1. [SNAPSHOT]: Freeze (State, ShadowState, ShadowMode, Policy)
	   2. [LOCKOUT]: Check atomic kill-switch
	   3. [WEIGHT]: Verify EvidenceWeight vs GetThreshold(req.TargetState)
	   4. [MATRIX]: Validate LADDER_VIOLATION (Fast-Up, Slow-Down)
	   5. [INVARIANTS]: Execute Fail-Fast invariant loop
	   6. [LEDGER]: Record 43-byte payload (WRITE-AHEAD INTENT)
	   7. [APPLY]: Atomic state transition + Logic Serialization
	   8. [SHADOW]: Bypass physical actuation if snap.ShadowMode == true
	*/

	// 1. Snapshot Posture
	snap := PostureSnapshot{
		State:       SystemState(w.state.Load()),
		ShadowState: SystemState(w.shadowState.Load()),
		ShadowMode:  w.shadowMode.Load(),
		Policy:      w.policies.Load(),
	}
	snap.PolicyHash = computePolicyHash(snap.Policy)
	if snap.PolicyHash != req.PolicyHash {
		return false, fmt.Errorf("ERR_POLICY_MISMATCH")
	}

	// 2. Lockout
	if w.locked.Load() {
		return false, fmt.Errorf("ERR_FSM_LOCKED")
	}

	// 3. Weight Check
	if req.EvidenceWeight.V < GetThreshold(req.TargetState).V {
		return false, fmt.Errorf("ERR_INSUFFICIENT_WEIGHT: %v < %v",
			req.EvidenceWeight, GetThreshold(req.TargetState))
	}

	// 4. Matrix Validation
	hasProof := len(req.Certificate) > 0
	if err := w.validateTransition(snap, req.TargetState, hasProof); err != nil {
		return false, err
	}

	// 5. Invariants
	invs := w.getInvariants()
	for _, inv := range invs {
		if err := inv.Verify(req, snap); err != nil {
			w.emergencyHaltLocked(err, lamport)
			return false, fmt.Errorf("ERR_INVARIANT_BREACH: %v", err)
		}
	}

	// 6. Ledger Write-Ahead (PENDING)
	payload, _ := serializePayload(req, snap.ShadowMode, snap.ShadowState)
	if err := w.recordLedgerEvent(req.EventID, "ROOT", lamport, payload, snap, req.TargetState); err != nil {
		return false, fmt.Errorf("ERR_LEDGER_WRITE_FAILURE: %v", err)
	}

	// 7. Apply Transition
	if err := w.applyTransition(req.TargetState, hasProof); err != nil {
		return false, err
	}

	// 8. Queue Actuation
	if snap.ShadowMode {
		w.appendAudit(fmt.Sprintf("SHADOW_ONLY: Transition to %v", req.TargetState), lamport)
		return true, nil
	}

	// Causal tracking
	w.pendingMu.Lock()
	w.pendingMap[req.EventID] = pendingEntry{startTick: lamport, lastHeartbeat: lamport}
	w.pendingMu.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	task := ActuationTask{req: req, tick: lamport, context: ctx, cancel: cancel}

	if req.ActuationClass == ClassKill || req.ActuationClass == ClassIsolate {
		select {
		case w.hpQueue <- task:
		case <-time.After(100 * time.Millisecond):
			w.executeActuation(task)
		}
	} else {
		// Non-blocking NP with Shedding per Q23
		select {
		case w.npQueue <- task:
		default:
			w.appendAudit(fmt.Sprintf("NP_QUEUE_FULL: Shedding task %s", req.EventID), lamport)
			cancel()
		}
	}

	return true, nil
}

func (w *Warden) validateTransition(snap PostureSnapshot, target SystemState, hasProof bool) error {
	if w.locked.Load() {
		return fmt.Errorf("FSM_LOCKED")
	}
	if snap.State == StateInvalid || !isStateValid(target) {
		return fmt.Errorf("INVALID_CONSTANTS")
	}

	// COMPROMISED Exit Override (ROOT-005)
	if snap.State == StateCompromised {
		if (target == StateCritical || target == StateCompromised) && hasProof {
			return nil // Authorized Recovery or Lease Renewal
		}
		return fmt.Errorf("TERMINAL_STATE_VIOLATION: COMPROMISED requires proof to exit or renew")
	}

	if !allowedTransitions[snap.State][target] {
		return fmt.Errorf("LADDER_VIOLATION: %v -> %v", snap.State, target)
	}

	return nil
}

func (w *Warden) applyTransition(target SystemState, hasProof bool) error {
	w.stateMu.Lock()
	defer w.stateMu.Unlock()

	if w.locked.Load() {
		return fmt.Errorf("FSM_LOCKED")
	}

	w.state.Store(uint32(target))
	return nil
}

func (w *Warden) getInvariants() []Invariant {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.invariants
}

func (w *Warden) appendAudit(content string, tick uint64) [32]byte {
	w.auditMu.Lock()
	defer w.auditMu.Unlock()

	prevHash := w.lastHash.Load().([32]byte)
	index := w.auditHead.Add(1)

	h := sha256.New()
	h.Write(prevHash[:])
	h.Write([]byte(content))
	binary.Write(h, binary.BigEndian, index)
	binary.Write(h, binary.BigEndian, tick)

	var currentHash [32]byte
	copy(currentHash[:], h.Sum(nil))

	w.auditBuffer[index&(AuditBufferSize-1)] = AuditEntry{
		Index:     index,
		Timestamp: int64(tick),
		Content:   content,
		PrevHash:  prevHash,
		Hash:      currentHash,
	}
	w.lastHash.Store(currentHash)
	return currentHash
}

func (w *Warden) emergencyHaltLocked(err error, tick uint64) {
	w.locked.Store(true)
	w.state.Store(uint32(StateCompromised))
	w.appendAudit(fmt.Sprintf("FATAL:INVARIANT_BREACH:%v", err), tick)
}

func (w *Warden) recordLedgerEvent(id, cause string, tick uint64, payload []byte, snap PostureSnapshot, target SystemState) error {
	w.ledgerMu.Lock() // Fixed: Isolated I/O lock
	defer w.ledgerMu.Unlock()

	if w.ledger == nil {
		// [RECTIFIED]: Constraint 1 - Silent no-op.
		// WHY: Ledger is optional. Tests and staging environments may
		// operate without durable storage. This is not an error.
		return nil
	}

	// RECTIFIED: Debt 3 - Raw byte IDs for bit-perfect matching
	stateBefore := []byte{uint8(snap.State)}
	stateAfter := []byte{uint8(target)}

	return w.ledger.AddEntryV2(
		id,
		cause,
		tick,
		payload,
		"",
		stateBefore,
		stateAfter,
		fmt.Sprintf("%x", snap.PolicyHash),
	)
}

// TODO(BFT-SPEC-004): Replace json.Marshal with DeterministicHash before Stage B
func computePolicyHash(p *config.RedLines) [32]byte {
	if p == nil {
		return [32]byte{}
	}
	b, err := json.Marshal(p)
	if err != nil {
		return [32]byte{0xFF}
	}
	return sha256.Sum256(b)
}

func isStateValid(s SystemState) bool    { return s >= StateSafe && s <= StateCompromised }
func isClassValid(c ActuationClass) bool { return c >= ClassNone && c <= ClassKill }

// Lock freezes the FSM in place.
func (w *Warden) Lock() {
	w.locked.Store(true)
}

// GetState returns the current FSM state.
func (w *Warden) GetState() SystemState {
	return SystemState(w.state.Load())
}
