# PhoenixOS: System Architecture

PhoenixOS is a multi-layered, deterministic security substrate designed for high-assurance environments. Rather than acting as a general-purpose operating system, it functions as an autonomous, closed-loop security controller that monitors, verifies, and quenches system disorder.

---

## 1. The 4-Layer Substrate (Phoenix Matrix)

The system operates as a layered hierarchy where lower layers enforce invariants and feed high-integrity telemetry to higher layers.

```mermaid
graph TD
  subgraph L4: Cognitive [Layer 4: Cognitive & Swarm]
    direction TB
    AI["PhoenixMind (Local Advisor LLM)"]
    Swarm["Nexus Consensus (PoA + Reputation)"]
  end

  subgraph L3: Control [Layer 3: Control & Actuation]
    direction TB
    Arbiter["Arbiter (Cost-Benefit Engine)"]
    Warden["Warden FSM (SAFE -> CONTAIN)"]
  end

  subgraph L2: Truth [Layer 2: State & Time Integrity]
    direction TB
    Ledger["TruthLedger (Merkle DAG Hash Chain)"]
    Registry["State Registry (Transition Audit)"]
  end

  subgraph L1: Substrate [Layer 1: Platform & Kernel]
    direction TB
    eBPF["eBPF Telemetry Probes"]
    XDP["XDP Packet Ingress/Egress Filters"]
    Guard["Guard Python Ingestion Daemon"]
  end

  L1 -->|HMAC-signed telemetry stream| L2
  L2 -->|Verified historical replay| L3
  L3 -->|Action commands / cgroup bounds| L1
  L4 -.->|Read-only advisory insight (LOCKED)| L3

  style L4 fill:#1a1010,stroke:#5c2222,stroke-width:2px,stroke-dasharray: 5 5
  style L3 fill:#111a1a,stroke:#225c5c,stroke-width:2px
  style L2 fill:#11111a,stroke:#22225c,stroke-width:2px
  style L1 fill:#161616,stroke:#444444,stroke-width:2px
```

### Layer 1: Substrate (Platform Integrity)
- **Responsibility:** Raw telemetry collection and fast-path filter enforcement.
- **Components:**
  - `xdp_ingress` & `egress_policy`: eBPF/XDP hooks executing under 100ns to block/throttle network packets.
  - `guard_runtime_py`: Fast-path daemon running inside an isolated namespace, handling initial socket collection and HMAC evidence signing.

### Layer 2: Truth (State & Time Integrity)
- **Responsibility:** Cryptographic non-repudiation and state history tracking.
- **Components:**
  - `TruthLedger`: Merkle-style DAG hash chain where each block links to parent hashes.
  - `State Registry`: Central registry validating state transitions, handling snapshots, and performing deterministic rollbacks.

### Layer 3: Control (Autonomous Actuation)
- **Responsibility:** Scoring system health and executing containment actions.
- **Components:**
  - `Arbiter`: A cost-benefit engine that compares Attack Cost (AC) vs. Containment Cost (CC).
  - `Warden FSM`: Finite-State Machine with 5 states (`SAFE`, `WATCH`, `ALERT`, `CONTAIN`, `RECOVERY`) enforcing hysteresis (30 ticks) and recovery budgets.

### Layer 4: Cognitive (Strategic AI & Consensus)
- **Responsibility:** Distributed agreement and explainable forensics.
- **Status:** **LOCKED** under Stage F0 constraints to ensure single-node determinism is fully validated first.

---

## 2. Key System Invariants

> [!IMPORTANT]
> The substrate guarantees **100% deterministic replayability** across all active nodes:
> $$\text{Replay}(S_0, E_{1..N}) \implies S_N$$

1. **Zero Wall-Clock Dependency:** No calls to `time.Now()` are permitted in the active path. Time progresses strictly via monotonic logical ticks carried in telemetry payloads.
2. **Deterministic Serialization:** Field ordering is strictly canonicalized (Canonical JSON/Bencode) before hashing or persistence to prevent map-iteration drift.
3. **Evidence-First Actuation:** No Warden state change can execute without a matching cryptographic entry appended to the ledger.

---

## 3. Validation & Testing Architecture

PhoenixOS integrates a dedicated validation test layer (`tests/validation/`) designed to verify the core invariants under stress, chaos, and adversarial mutation before code moves to staging or runtime phases:
1. **Replay Identity & Determinism:** Evaluates cross-run hash equivalence (`Run 1 == Run 100 == Run 1000`) and verifies that replaying identical trace inputs yields identical state outcomes.
2. **Mutation and Tampering Resistance:** Asserts ledger integrity by checking that modifying payloads, tampering with validation hashes, splitting timelines, or attempting hash forks is immediately detected and rejected.
3. **FSM State Invariants:** Validates transition rules, hysteresis dwell limits (30 ticks), cooldowns (10 ticks), and recovery budgets to prevent containment oscillation.

