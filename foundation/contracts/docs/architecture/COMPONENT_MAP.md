---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Contracts — Component Map

> Last verified: 2026-06-04

The `foundation/contracts` subproject defines all canonical interfaces and constant values for the PhoenixOS system. It serves as the single source of truth for communication and state-transition contracts across all layers.

## Component Breakdown

```
foundation/contracts/
├── compat.go                  # Global interfaces (ILedger) and states
├── go.mod                     # Module configuration (no external dependencies)
├── events/
│   └── v1/
│       ├── doc.go             # Package documentation
│       ├── event.go           # Event and Serializer interfaces
│       └── envelope.go        # EventEnvelope interface
├── security/
│   └── v1/
│       ├── doc.go             # Package documentation
│       └── actuator.go        # Actuator, Containment, Escalation interfaces
├── replay/
│   └── v1/
│       ├── doc.go             # Package documentation
│       └── replay.go          # Snapshot, Reconstructor, ReplayEngine interfaces
└── fsm/
    ├── warden.yaml            # Containment FSM schema
    ├── replay-engine.yaml     # Replay state machine schema
    ├── arbiter.yaml           # Arbiter consensus FSM schema
    └── distributed.yaml       # Distributed sync FSM schema
```

### Component Details

1. **Global Contracts (`contracts/`)**
   - **`ILedger`**: Defines the interface for ledger appends, verification, and proof generation. Used by `foundation/ledger` for implementation and other modules for ingestion.
   - **States**: Constants defining the operational health state of the system (`StateSafe`, `StateWatch`, `StateSuspicious`, `StateCritical`, `StateCompromised`).

2. **Event Schema Contracts (`contracts/events/v1`)**
   - **`Event`**: Interface representing raw event payloads, logical timestamps, signatures, and cryptographic evidence.
   - **`EventEnvelope`**: Wraps raw events with metadata (source component, correlation ID, trust scores, causal chain) required for replay and security routing.
   - **`Serializer`**: Declares format-agnostic marshaling/unmarshaling contracts.

3. **Security Contracts (`contracts/security/v1`)**
   - **`Actuator`**: The interface implemented by security wardens to execute containment policies and process killing.
   - **`Containment` & `Escalation`**: Interfaces specifying containment level states and trigger reasons.

4. **Replay Contracts (`contracts/replay/v1`)**
   - **`ReplayEngine`**: Core replay engine contract ensuring execution determinism.
   - **`Reconstructor`**: Builds memory snapshots from an envelope sequence.
   - **`Snapshot`**: Holds serializable, hashed states for checkpoints.
