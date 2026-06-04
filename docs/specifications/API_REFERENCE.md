---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# PhoenixOS API Reference

## Central Event Bus (`runtime/bus`)

### `TelemetryEvent` Struct
The canonical telemetry payload schema:
```go
type TelemetryEvent struct {
	SeqID        int64           `json:"seq_id"`
	LogicalTick  uint64          `json:"logical_tick"`
	MonotonicNs  int64           `json:"monotonic_ns"`
	WallTimeUnix int64           `json:"wall_time_unix"`
	Source       string          `json:"source"`
	HostID       string          `json:"host_id"`
	PID          int             `json:"pid"`
	TID          int             `json:"tid"`
	UID          int             `json:"uid"`
	GID          int             `json:"gid"`
	EventType    string          `json:"event_type"`
	Severity     float64         `json:"severity"`
	Payload      json.RawMessage `json:"payload"`
	PrevHash     string          `json:"prev_hash"`
	Hash         string          `json:"hash"`

	// Legacy & Compatibility Fields
	LamportClock uint64 `json:"lamport_clock"`
	Nsproxy      uint32 `json:"nsproxy"`
	Tgid         uint32 `json:"tgid"`

	// Hardening Fields
	EventID     string `json:"event_id"`
	CausalID    string `json:"causal_id"`
	SequenceNo  int64  `json:"sequence_no"`
	SourceEpoch int64  `json:"source_epoch"`
}
```

### `Bus` Struct & Methods
- `NewBus() *Bus`: Instantiates the central pub/sub bus.
- `Subscribe(topic string) chan TelemetryEvent`: Subscribes to a topic. Channels are buffered.
- `Publish(topic string, event TelemetryEvent)`: Publishes event to topic. Enforces queue pressure rules.
- `QueuePressure(topic string) float64`: Returns current queue pressure (fill ratio, 0.0 to 1.0).

---

## Evidence Ledger (`ledger/src`)

### `Ledger` Struct & Methods
- `NewLedger(alloc ResourceAllocator) *Ledger`: Creates an append-only evidence Merkle DAG.
- `AddEntry(eventID, causeID string, payload []byte) error`: Appends a new V1 entry.
- `AddEntryV2(eventID, causeID string, payload []byte, traceHash, stateBefore, stateAfter, policyVersion string) error`: Appends a new V2 entry with state transition metadata.
- `Verify() error`: Cryptographically traverses the Merkle DAG to verify all hashes and state transitions.
- `GenerateCertificate(eventID string, weight float64) ([]byte, error)`: Generates a signed certificate proof.
- `VerifyCertificate(eventID string, weight float64, cert []byte) bool`: Verifies a certificate matches the ledger record.

---

## Warden FSM Engine (`security/warden`)

### `Warden` Struct & Methods
- `NewWarden(b *bus.Bus) *Warden`: Instantiates the FSM engine in SAFE state.
- `RegisterInvariant(inv Invariant)`: Registers a formal proof-gate.
- `RegisterActuator(act Actuator)`: Registers an enforcement mechanism (Process, eBPF).
- `Actuate(req AuthorityEscalationRequest, seq int, lamportClock uint64) bool`: Entry point for executing FSM state changes and physical actuations.

### `AuthorityEscalationRequest` Struct
```go
type AuthorityEscalationRequest struct {
	EventID        string
	TargetPID      int
	TargetTgid     int
	TargetNsproxy  uint32
	TargetState    SystemState
	ActuationClass ActuationClass
	EvidenceWeight float64
	Certificate    []byte
	GraphProof     *GraphProof
}
```

## Contract Note
The API surface described here must be implemented through versioned contracts and adapters. Public interfaces belong to the contract packages, not to implementation packages.
