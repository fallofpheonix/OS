# PhoenixOS: Non-Goals

To maintain strict architectural boundaries, PhoenixOS explicitly avoids the following:

1. **General Purpose OS:** We are a security-control operating substrate running on Linux, not a general-purpose user OS.
2. **Non-Deterministic Actuation:** AI and heuristics cannot directly actuate without passing through the Warden FSM.
3. **Lossy Telemetry:** We do not drop critical events for performance; we queue, throttle, or isolate.
4. **Mutable Evidence:** Logs and events can never be altered once written to the Truth Ledger.
