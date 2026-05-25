# PhoenixOS

**PhoenixOS** is a **Deterministic Cybernetic Security Runtime**. It implements a unified **Mathematical-Physical-Game Architecture** to achieve autonomous system integrity through mathematically reproducible replay and bounded control.

## 1. Project Dimensions

### Structural Roots (The 8 Directories)
*   `phoenix_os/`: Active runtime core (Telemetry, Kernel, Security).
*   `tests/`: Centralized verification labs.
*   `tools/`: Build systems and CLI interface.
*   `02_docs/`: Canonical governance and architecture.
*   `external/`: Categorized third-party repositories.
*   `research/`: Verified studies and theoretical foundations.
*   `experimental/`: Isolated sandboxes for future-phase systems.
*   `archive/`: Historical artifacts and deprecated code.

### The Phoenix Matrix (Functional Stack)
*   **L7: Phoenix Nexus**: Swarm Coordination.
*   **L6: Phoenix Sentinel**: System Physics (Entropy/SDI).
*   **L5.5: Phoenix Arbiter**: Strategic Policy (Game Theory).
*   **L5: Phoenix Warden**: Actuation & Control (FSM).
*   **L4: Phoenix Trace**: Graph Intelligence (Causal DAG).
*   **L3: Phoenix Monitor**: Telemetry Math (Kalman/Signal).
*   **L2: Phoenix Kernel**: Runtime Probes (eBPF).
*   **L1: Phoenix Guard**: Platform Integrity (Fast-Path).

## 2. Quick Start (Workable System)

### Prerequisites
- Go 1.25+
- Linux with eBPF support (for kernel probes)

### Build the Runtime
```sh
go build ./phoenix_os/...
```

### Run the Replay Identity Lab
```sh
go test ./tests/replay/...
```

### Documentation
See the [Documentation Index](02_docs/README.md) for detailed architecture and models.

## 3. Core Axioms
1. **Determinism is sacred.**
2. **Replay is authoritative.**
3. **AI is advisory.**
4. **Control must remain bounded.**
5. **Telemetry correctness > AI sophistication.**
6. **Never scale instability.**

---
*For development instructions and contribution guidelines, see [CLAUDE.md](CLAUDE.md).*
