---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# PhoenixOS Testing Guide

## Test Architecture
Tests in PhoenixOS verify both positive and negative execution paths, edge cases, and deterministic invariants. We utilize standard Go unit tests, integration tests, and simulator-driven stress testing.

Contract compatibility tests are part of the core test surface and must gate boundary changes.

---

## Running the Test Suite

### Running All Tests (Multi-Module Workspace)
Since the repository is organized into a Go workspace with a `go.work` file, you can run all tests sequentially across all modules using:

```bash
for d in ./Phoenix.Nucleus/PhoenixCore \
         ./Phoenix.Nucleus/PhoenixDistributed \
         ./Phoenix.Nucleus/PhoenixGuard \
         ./Phoenix.Nucleus/PhoenixKernel \
         ./Phoenix.Nucleus/PhoenixTrace \
         ./Phoenix.Nucleus/PhoenixTruth \
         ./Phoenix.Nucleus/PhoenixValidation \
         ./Phoenix.Cognition/PhoenixMind \
         ./Phoenix.Crucible/PhoenixSimulation \
         ./Phoenix.Terminus/PhoenixOS; do
    go test -C "$d" ./... || exit 1
done
```

### Running with the Race Detector
To check for concurrency bottlenecks and data races (as required by the system rules), run the sequential test loop with the `-race` flag:

```bash
for d in ./Phoenix.Nucleus/PhoenixCore \
         ./Phoenix.Nucleus/PhoenixDistributed \
         ./Phoenix.Nucleus/PhoenixGuard \
         ./Phoenix.Nucleus/PhoenixKernel \
         ./Phoenix.Nucleus/PhoenixTrace \
         ./Phoenix.Nucleus/PhoenixTruth \
         ./Phoenix.Nucleus/PhoenixValidation \
         ./Phoenix.Cognition/PhoenixMind \
         ./Phoenix.Crucible/PhoenixSimulation \
         ./Phoenix.Terminus/PhoenixOS; do
    go test -C "$d" -race ./... || exit 1
done
```

---

## Red Team Attack Simulators (`Phoenix.Crucible/PhoenixRedteam`)
We have developed four attack simulators to stress-test the FSM and event bus:

1. **byzantine_subversion**: Simulates network/consensus attacks.
2. **drowning_man**: Simulates memory pressure and resource depletion.
3. **replay_stress**: Simulates high-frequency event replays to verify duplicate detection.
4. **system_pulse**: Simulates high-velocity event bursts to test the bus queue watermarks.

To build the simulators:
```bash
go build -C ./Phoenix.Crucible/PhoenixRedteam byzantine_subversion.go
go build -C ./Phoenix.Crucible/PhoenixRedteam drowning_man.go
go build -C ./Phoenix.Crucible/PhoenixRedteam replay_stress.go
go build -C ./Phoenix.Crucible/PhoenixRedteam system_pulse.go
```
