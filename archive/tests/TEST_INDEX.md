# PhoenixOS Test Index

This index lists and maps the comprehensive validation suite for all nodes, servers, runtime components, and security bounds of PhoenixOS.

## Test Directory Structure

- **[unit/](file:///Users/fallofpheonix/os/tests/unit)** - Individual function-level logic checks.
- **[integration/](file:///Users/fallofpheonix/os/tests/integration)** - Subsystem boundary verification.
- **[nodes/](file:///Users/fallofpheonix/os/tests/nodes)** - Component verification (Telemetry, Replay, Truth, Arbiter, Warden, Containment, etc.).
- **[server/](file:///Users/fallofpheonix/os/tests/server)** - Lifecycle, crash resilience, and boundaries for server processes.
- **[soak/](file:///Users/fallofpheonix/os/tests/soak)** - Long-duration drift and soak stability tests.
- **[proofs/](file:///Users/fallofpheonix/os/tests/proofs)** - Formal invariant checks and repeatability proofs.
- **[regression/](file:///Users/fallofpheonix/os/tests/regression)** - Regression test suites.
- **[chaos/](file:///Users/fallofpheonix/os/tests/chaos)** - Fault injection and clock drift simulation.
- **[security/](file:///Users/fallofpheonix/os/tests/security)** - Attack vector resistance verification (Ledger tampering, exfiltration, beacon).
- **[observability/](file:///Users/fallofpheonix/os/tests/observability)** - Metrics emission and latency boundary validation.
- **[validation/](file:///Users/fallofpheonix/os/tests/validation)** - Topology rules and data flow verification.
- **[e2e/](file:///Users/fallofpheonix/os/tests/e2e)** - Complete pipeline integrations.

## Test Matrices

- [NODE_TEST_MATRIX.md](file:///Users/fallofpheonix/os/tests/NODE_TEST_MATRIX.md)
- [SERVER_TEST_MATRIX.md](file:///Users/fallofpheonix/os/tests/SERVER_TEST_MATRIX.md)
- [CHAOS_MATRIX.md](file:///Users/fallofpheonix/os/tests/CHAOS_MATRIX.md)
- [SECURITY_MATRIX.md](file:///Users/fallofpheonix/os/tests/SECURITY_MATRIX.md)
- [E2E_MATRIX.md](file:///Users/fallofpheonix/os/tests/E2E_MATRIX.md)
