# Phoenix Validator System

The Phoenix Validator system provides the rules by which the Replay Engine judges execution traces. It enforces dynamic invariants and security policies against the evidence recorded in the `TruthLedger`.

## Core Components

### Validator Interface
All security rule-checkers implement the `Validator` interface:
- `Validate(entry *ledger.LedgerEntry) ValidationResult`: Checks an entry against invariants.
- `Name() string`: Returns the validator name.
- `Reset()`: Resets any internal state.

### Core Validators
- **SequenceValidator**: Ensures that `LogicalTick` values are strictly monotonic, detecting sequence regressions or gaps.
- **TransitionValidator**: Enforces the Warden's Finite State Machine (FSM) rules (e.g., `SAFE -> WATCH -> SUSPICIOUS`).
- **EntropyValidator**: Implements Dynamic Information Flow Control (DIFC) by calculating Shannon entropy on payloads to detect potential encrypted exfiltration or ransomware activity.

### ValidatorRegistry
A central hub where validators are registered and executed sequentially against traces.

## Usage

```go
registry := validators.NewValidatorRegistry()
registry.Register(validators.NewSequenceValidator())
registry.Register(validators.NewTransitionValidator())
registry.Register(validators.NewEntropyValidator(7.5))

for _, entry := range trace {
    failures := registry.ValidateAll(&entry)
    if len(failures) > 0 {
        // Handle validation failures
    }
}
```
