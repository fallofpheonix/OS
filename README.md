# PhoenixOS v1.0: Sovereign Operating Substrate

[![Maturity: v1.0-Verifiable](https://img.shields.io/badge/Maturity-v1.0--Verifiable-blue.svg)](./docs/roadmap/MASTER_PHOENIX_ROADMAP.md)
[![Verification: PASSED](https://img.shields.io/badge/Verification-PASSED-green.svg)](./core/Phoenix.Nucleus/PhoenixValidation/proofs/)

PhoenixOS is a deterministic, constitutional operating substrate designed for autonomous systems where security and state integrity are absolute.

## 1. Project Mission
To provide a mathematically verified and forensic-quality runtime for secure execution and intelligent understanding. PhoenixOS guarantees deterministic state reconstruction, absolute authority constraint, and containment of untrusted code.

## 2. Core Substrate Structure
The repository is strictly organized into the following verified domains:

- **[`/core`](./core):** The Sovereign Runtime
  - **`Phoenix.Nucleus`:** The Foundation (Execution, Authority, Ledger, Event Bus).
  - **`Phoenix.Cognition`:** The Intelligence (Memory, Knowledge, Reasoning, Reflection).
  - **`Phoenix.Arbiter`:** The Governance (Policy Mapping, Scanning, Authority Oversight).
  - **`Phoenix.Guard`:** Tactical Enforcement (Warden FSM, Actuators, Invariants).
  - **`Phoenix.Terminus`:** The Interface (Operator CLI, Dashboard, LLM Oracle).
  - **`Phoenix.Crucible`:** The Evolution (Simulation, Adversarial Testing).
  - **`Phoenix.UI`:** The Aesthetics (UX Research, GPU-Accelerated Interfaces).

- **[`/docs`](./docs):** The Knowledge Base
  - **[`MASTER_SPECIFICATION.md`](./docs/architecture/MASTER_SPECIFICATION.md)**: Single source of truth for design and flow.
  - **[`MASTER_PHOENIX_ROADMAP.md`](./docs/roadmap/MASTER_PHOENIX_ROADMAP.md)**: Execution phases and future goals.
  - **[`REPOSITORY_CONSTITUTION.md`](./docs/governance/REPOSITORY_CONSTITUTION.md)**: Mandatory engineering standards.
  - **[`GLOSSARY.md`](./docs/governance/GLOSSARY.md)**: Standardized terminology.

- **[`/tools`](./tools):** Engineering Utilities
  - Repository maintenance, deployment, and auditing scripts.

- **[`/archive`](./archive):** Historical Substrate
  - Research prototypes, dead simulations, and classification reports.

## 3. The Constitutional Engine
PhoenixOS v1.0 is governed by a **Machine-Executable Constitution**.
1. **Boot Validation:** Node refuses boot unless Constitution and Ledger integrity are verified.
2. **Deterministic Replay:** State is a bit-perfect projection of history (Ledger + Artifacts).
3. **Perfect Recovery:** Node resurrection produces identical authoritative state hashes.
4. **Containment Ladder:** Real-time isolation (Observe -> Warn -> Throttle -> Freeze -> Isolate -> Kill).

## 4. Verified Proofs
Every claim in PhoenixOS is backed by an automated verification suite:
- **Proof 1: Replay** — 10 runs, 3 verifiers, 0 divergence.
- **Proof 2: Recovery** — Destroy, Recover, Verify (State Hash Match).
- **Proof 3: Containment** — Attack, Detect, Contain, Verify.

## 5. Development
- **Mandatory Standards:** See [REPOSITORY_CONSTITUTION.md](./docs/REPOSITORY_CONSTITUTION.md).
- **Glossary:** See [GLOSSARY.md](./docs/GLOSSARY.md) for terminology.
- **Roadmap:** See [MASTER_PHOENIX_ROADMAP.md](./docs/MASTER_PHOENIX_ROADMAP.md) for execution phases.

## 6. Launching the Node
Initialize the sovereign node using the master Makefile:

```bash
make ignite  # Starts the Docker-based Phoenix mesh
make test    # Executes the verification proofs
```

The system defaults to **Shadow Mode** for safety, evaluating containment decisions without physical enforcement until explicitly authorized.
