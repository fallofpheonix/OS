# WorldState Specification (Authoritative)

**Status:** DRAFT (Vertical Slice Step 5)
**Confidence:** High
**Owner:** Phoenix.Nucleus / Game Layer

## 1. Structure
The `WorldState` is the authoritative snapshot of the game universe at a specific discrete time ($T$).

```go
type WorldState struct {
	Tick          uint64            `json:"tick"`           // Monotonic time step
	Seed          int64             `json:"seed"`           // Base entropy for the session
	Entities      map[string]*Entity `json:"entities"`       // Unordered storage (Map)
	LastEventHash string            `json:"last_event_hash"` // causal link to ledger
	StateHash     string            `json:"state_hash"`     // Cryptographic commitment to this state
}
```

## 2. Canonical Hashing (Determinism)
To ensure identical states produce identical hashes regardless of memory layout:

1. **Sort Entities:** Entities must be sorted by `ID` (string, ascending).
2. **Serialize:** Each entity is serialized to JSON with deterministic field ordering.
3. **Concatenate:** Tick + Seed + LastEventHash + SortedEntityHashes.
4. **Hash:** SHA-256 of the concatenated string.

## 3. State Transitions
A state transition $S_{T} \rightarrow S_{T+1}$ is ONLY valid if:
1. It is produced by a deterministic execution of the P-Script VM.
2. The `LastEventHash` at $T+1$ matches the hash of the event that triggered the transition.
3. The `StateHash` is re-calculated and verified.

## 4. Initialization
- **Genesis State ($T=0$):**
  - Tick: 0
  - Seed: Defined in Simulation Manifest
  - LastEventHash: Genesis Block Hash
  - Entities: Initial set defined in Manifest
