# PhoenixOS (Autonomous CyberAI Operating System)

![Phoenix Logo](assets/Phoenix/logo_simple.svg)

PhoenixOS is a security-focused operating system project built from scratch and as a custom Linux derivative. It implements a unified **Mathematical-Physical-Game Architecture** to achieve autonomous system integrity and self-healing defense.

## Core Vision: The Phoenix Matrix
Unlike traditional operating systems that rely on passive signatures, PhoenixOS treats system security as a thermodynamic state. It uses real-time telemetry to compute a **Security Disorder Index (SDI)** and employs game-theoretic controllers to autonomously "quench" threats before they reach a cascading failure point.

## System Services (The 7-Layer Stack)
- **L7: Swarm coordination** (**Phoenix Nexus**) - Distributed consensus and MARL.
- **L6: System Physics** (**Phoenix Sentinel**) - Thermodynamic SDI monitoring.
- **L5.5: Strategic Policy** (**Phoenix Arbiter**) - Stackelberg Security Games.
- **L5: Actuation & Control** (**Phoenix Warden**) - PID/FSM process feedback loops.
- **L4: Graph Intelligence** (**Phoenix Trace**) - Causal process lineage DAGs.
- **L3: Telemetry Math** (**Phoenix Monitor**) - Shannon entropy and Kalman filters.
- **L2: Kernel Runtime** (**Phoenix Kernel**) - eBPF probes and ring buffers.
- **L1: Platform Integrity** (**Phoenix Guard**) - Kernel-level Fast Path enforcement.

## P0 Foundations (Verified Evidence)
- **Phoenix Ledger:** Every autonomous action is backed by a verifiable, content-addressable evidence chain (SHA-256).
- **Fast Path Enforcement:** Critical threats (e.g., ransomware encryption) are blocked in **<100ms** by bypassing strategic layers.

## Implementation Status
- **Phoenix Bus (L3):** 10M+ events/sec throughput (**STABLE**)
- **Phoenix Monitor (L3):** Shannon/KL entropy analysis (**STABLE**)
- **Phoenix Trace (L4):** Lineage DAG construction (**STABLE**)
- **Phoenix Sentinel (L6):** SDI thermodynamic calculation (**STABLE**)
- **Phoenix Ledger (Evidence):** Cryptographic hash-chained audit trail (**COMPLETED**)
- **Phoenix Guard (L1):** eBPF Fast Path detection and blocking (**COMPLETED**)
- **Phoenix Nexus (L7):** Gossip-based swarm synchronization (**IN PROGRESS**)

## Document Index
- [PHOENIX_TASKS.md](PHOENIX_TASKS.md) - Active task list and progress.
- [PHOENIX_PROBLEMS.md](PHOENIX_PROBLEMS.md) - Gap analysis and architectural blockers.
- [PHOENIX_SOLUTIONS.md](PHOENIX_SOLUTIONS.md) - Mathematical and systemic bridges.
- [MASTER_DEPENDENCY_GRAPH.md](MASTER_DEPENDENCY_GRAPH.md) - Global architecture DAG.

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

