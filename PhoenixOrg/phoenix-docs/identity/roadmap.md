# MASTER ROADMAP

## Status: UNIFIED

### Source: 00_program_management/roadmap/main_roadmap.md

# PhoenixOS: Master Execution Roadmap

## Project Definition
**PhoenixOS** is a **Deterministic Cybernetic Security Runtime** running on Linux. It focuses on mathematically reproducible replay, bounded control, and telemetry-first validation.

## Core Axioms
1. **Determinism is sacred.**
2. **Replay is authoritative.**
3. **AI is advisory.**
4. **Control must remain bounded.**
5. **Telemetry correctness > AI sophistication.**
6. **Never scale instability.**

---

## Phase A: Core Product (Stages 0–4)
*Focus: Deterministic single-node cybernetic security runtime.*

### Stage 0: Foundation Stabilization
* **Goal:** Architectural coherence.
* **Deliverables:** RFC system, repository structure, FSM model, `DETERMINISM.md`.
* **Exit Criteria:** No major architectural contradictions remain.

### Stage 1: Deterministic Replay Runtime (COMPLETED)
* **Goal:** Mathematically reproducible single-node replay.
* **Systems:** Replay engine, Event bus, Ledger, FSM, TCS, Drift engine.
* **Features:** Logical time, Canonical serialization, Stable hashing, Bounded queues.
* **Exit Criteria:** Replay hashes are byte-for-byte identical across runs.

### Stage 2: Real Telemetry Runtime (ACTIVE)
* **Goal:** Replace simulated data with real Linux eBPF/XDP telemetry.
* **Build:** Phoenix Guard, Kernel adapters, Telemetry collectors.
* **Exit Criteria:** Stable single-node telemetry replay.

### Stage 3: Immutable Runtime Image
* **Goal:** Boot Phoenix as a minimal immutable environment.
* **Build:** LinuxKit appliance, Minimal kernel, Immutable initrd, PID1 Warden.
* **Exit Criteria:** Phoenix runtime boots as a standalone immutable appliance.

### Stage 4: Constrained Actuation (COMPLETED / HARDENED)
* **Goal:** Safe, bounded autonomous response with full concurrency safety and rate limits.
* **Systems:** Warden FSM, Actuation budgets, Cooldowns, Hysteresis, Mutex synchronizers.
* **Allowed Actions:** Log, Throttle, Socket block, Process isolate, Operator reset.
* **Exit Criteria:** No state oscillation under adversarial stress, and thread-safe SOC API interaction.

---

## Phase B: Advanced Platform (Stages 5–8)
*Focus: Distributed, forensic-grade enterprise observability.*

### Stage 5: Evidence & Replay Infrastructure
* **Goal:** Forensic-grade replay DAGs and cross-session provenance.
* **Features:** Causal linking, Event provenance, Hash chains.
* **Exit Criteria:** Replay becomes the authoritative system truth.

### Stage 6: Graph Runtime
* **Goal:** Transform telemetry into temporal causality structures.
* **Systems:** Process DAG, Node graph, Attack chains.
* **Exit Criteria:** Stable attack graph reconstruction.

### Stage 7: Distributed Telemetry
* **Goal:** Multi-node observability using vector clocks.
* **Features:** Event replication, Global replay, Distributed evidence.
* **Exit Criteria:** Cross-node deterministic replay succeeds.

### Stage 8: Cloud Runtime
* **Goal:** Cluster orchestration and immutable deployment fleet.
* **Build:** Scheduler, Policy distribution, Image rollout.
* **Exit Criteria:** Phoenix nodes are manageable as a fleet.

---

## Phase C: Cybernetic Research (Stages 9–12)
*Focus: Advanced reasoning, state estimation, and automated strategic defense.*

### Stage 9: Advisory AI Layer
* **Goal:** Explainable AI correlation and ranking (Advisory only).
* **Allowed Tasks:** Summaries, Policy suggestions, Replay explanation.
* **Exit Criteria:** AI outputs are strictly bounded and explainable.
* **Detailed Execution:** See [COGNITION_ENGINE_ROADMAP.md](00_program_management/roadmap/COGNITION_ENGINE_ROADMAP.md).

### Stage 10: Physics + Control Theory
* **Goal:** Formal system-state estimation.
* **Concepts:** Entropy, Variance (Welford's), Kalman filters, Signal processing.
* **Exit Criteria:** Physics layer measurably improves detection.

### Stage 11: Game Theory Runtime
* **Goal:** Strategic bounded defense planning.
* **Build:** Resource allocation games, Stackelberg defense, Decoy placement.
* **Exit Criteria:** Algorithmic defense decisions outperform heuristics.

### Stage 12: Cybernetic Runtime
* **Goal:** Full closed-loop bounded defense system.
* **Properties:** Observable, Bounded, Reversible, Deterministic.
* **Exit Criteria:** Stable and autonomous under adversarial stress.

---

## Phase D: Optional Evolution (Stage 13)
### Stage 13: Research OS
* **Focus:** Custom schedulers, hardware isolation (SGX/SEV), microkernel experiments.
* **Status:** Pure OS research, not required for core project success.

---

## Very Long-Term Evolution

For the multi-year master roadmap detailing Stage 14 through Stage 20 (including Proof-of-Anomaly swarm consensus, vector clocks, and CFS scheduler patches), see the detailed [very_long_term_plan.md](file:///Users/fallofpheonix/os/00_program_management/roadmap/very_long_term_plan.md).


### Source: 00_program_management/roadmap/TASK_BREAKDOWN.md

# Task Breakdown

## Learning Tasks

- Stage_00_Foundations
- Stage_01_System_Internals
- Stage_03_Networking foundations
- Stage_04_Security_Fundamentals foundations
- Stage_11_AI_ML_Core

## Research Tasks

- Stage_20_Research_and_Future
- Research-only files under `docs/research/` and `research/00_foundations/`
- Papers/datasets/repos roots when populated

## Coding Tasks

- Stage_01 labs: boot, allocator, logging, syscall/userland primitives
- Stage_02/15: build images, packages, release artifacts
- Stage_03/05/06/09/10/12/14/16/18: collectors, detectors, parsers, models, kernel/user modules

## Contribution Tasks

- Repository-aligned work in `REPOSITORY_ALIGNMENT.md`
- Candidate upstreams: Linux, Zeek, Suricata, Falco, Tracee, Wazuh, TheHive, OpenCTI, MISP, QEMU, LLVM/GCC, OSDev-adjacent docs

## Integration Tasks

- Stage_08 observability schema into Stage_10 SOC
- Stage_09 telemetry into Stage_10 SOC and Stage_12 Security AI
- Stage_12 model outputs into Stage_14 SOAR
- Stage_15 distribution packaging into Stage_17 Hybrid OS
- Stage_16 kernel extensions into Stage_18 Custom OS


### Source: UPDATED_IMPLEMENTATION_PLAN.md

# Updated Implementation Plan

## Roadmap
- **F0 Foundation**: Core substrate stabilization
- **F1 Runtime**: Deterministic userspace execution
- **F2 Formal Verification**: TLA+ models and proofs
- **F3 OS Primitives**: VFS and Process Management
- **F4 Memory**: Safe memory substrate
- **F5 Advisor**: Strategic policy engine
- **F6 Distributed**: Multi-node consensus
- **F7 Cognitive**: Advanced AI integration

## Parallel Research Tracks
- Replay Math
- Control Theory
- Optimization
- Simulation
- Formal Systems


