# PhoenixOS: Non-Goals

To maintain strict architectural boundaries, PhoenixOS explicitly avoids:

1. **General Purpose OS:** Security-control operating substrate, not a user OS.
2. **Non-Deterministic Actuation:** All actuation passes through Warden FSM.
3. **Lossy Telemetry:** Events are queued/throttled, never dropped.
4. **Mutable Evidence:** Truth Ledger is immutable once committed.
