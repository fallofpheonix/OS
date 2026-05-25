# PhoenixOS: State Model

FSM managed by `state.Registry`.

## States
- SAFE
- WATCH
- ALERT
- CONTAIN
- RECOVERY

## Transitions
All transitions via `Registry.Transition()`, generating `StateAuditEntry`.
Rollbacks via `Registry.Rollback()`, appending new audit entries.
