# PhoenixOS

![Phoenix Logo](assets/Phoenix/logo_simple.svg)

**PhoenixOS** is a **Deterministic Cybernetic Security Runtime** running on Linux. It implements a unified **Mathematical-Physical-Game Architecture** to achieve autonomous system integrity and self-healing defense through mathematically reproducible replay and bounded control.

## Core Axioms
1. **Determinism is sacred.**
2. **Replay is authoritative.**
3. **AI is advisory.**
4. **Control must remain bounded.**
5. **Telemetry correctness > AI sophistication.**
6. **Never scale instability.**

## The Phoenix Matrix (7-Layer Stack)
- **L7: Swarm Coordination (Phoenix Nexus):** Distributed consensus and replication.
- **L6: System Physics (Phoenix Sentinel):** Thermodynamic SDI monitoring.
- **L5.5: Strategic Policy (Phoenix Arbiter):** Game-theoretic planning (Stackelberg).
- **L5: Actuation & Control (Phoenix Warden):** Bounded FSM/PID controllers.
- **L4: Graph Intelligence (Phoenix Trace):** Causal process lineage DAGs.
- **L3: Telemetry Math (Phoenix Monitor):** Signal processing and drift detection.
- **L2: Kernel Runtime (Phoenix Kernel):** eBPF/XDP telemetry ingestion.
- **L1: Platform Integrity (Phoenix Guard):** <100ms Fast-Path enforcement.

## P0 Foundations (Verified Evidence)
- **Phoenix Ledger:** Every action is backed by a verifiable Merkle-DAG evidence chain.
- **Deterministic Replay:** Mathematically identical execution traces across boot sessions.

## Project Resources
- [**Master Execution Roadmap**](00_program_management/roadmap/main_roadmap.md) - The 12-stage path to a mature runtime.
- [**Algorithm & Model Map**](ALGORITHM_MAP.md) - Mathematical primitives by stage.
- [**GEMINI.md**](GEMINI.md) - Core instructions and engineering axioms.
- [PHOENIX_TASKS.md](PHOENIX_TASKS.md) - Active task list and progress.

## Quick Start (Simulation)
```sh
make boot
# Optional, if qemu-system-i386 is installed:
make run
```
# Development

Testing and CI
- Root build validation requires `nasm` and `gcc` on PATH.
- In this workspace, `make all` now completes after installing `nasm` and using the portable kernel prototype.
- Run unit tests locally in a virtualenv:

```bash
python3 -m venv .venv
. .venv/bin/activate
python -m pip install -r requirements.txt
PYTHONPATH=$(pwd) pytest -q
```

- The repository includes a GitHub Actions workflow at `.github/workflows/ci.yml` which runs the same tests on push/PR to `main`.

API
- The orchestrator FastAPI app is `agents/surface/orchestrator/api.py`. Run locally with:

```bash
PYTHONPATH=$(pwd) uvicorn agents.surface.orchestrator.api:app --reload --port 8000
```

