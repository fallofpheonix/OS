import os
import sys

docs_root = "/Users/fallofpheonix/os/docs"

# Define the folder layers and files
layers = {
    "governance": [
        "PROJECT_VISION.md",
        "ROADMAP.md",
        "MASTER_STATUS.md",
        "CHANGELOG.md",
        "DECISION_LOG.md",
        "PROJECT_HEALTH.md"
    ],
    "architecture": [
        "SYSTEM_ARCHITECTURE.md",
        "COMPONENT_MAP.md",
        "DATAFLOW_MAP.md",
        "CONTROL_FLOW.md",
        "MODULE_BOUNDARIES.md",
        "DEPENDENCY_GRAPH.md",
        "CAPABILITY_MATRIX.md",
        "SERVICE_REGISTRY.md"
    ],
    "integration": [
        "EXTERNAL_REPOS.md",
        "MERGE_POLICY.md",
        "FORK_STRATEGY.md",
        "PATCH_HISTORY.md",
        "UPSTREAM_SYNC.md",
        "INTEGRATION_STATUS.md",
        "API_COMPATIBILITY.md",
        "CONFLICT_LOG.md",
        "DEPRECATION_TRACKER.md"
    ],
    "agents": [
        "AGENT_REGISTRY.md",
        "AGENT_INTERACTION_MAP.md",
        "AGENT_CHAINING.md",
        "AGENT_MEMORY_POLICY.md",
        "AGENT_TRAINING_PIPELINE.md",
        "AGENT_FAILURE_CASES.md",
        "AGENT_REWARD_SYSTEM.md",
        "AGENT_VALIDATION.md",
        "AGENT_ALIGNMENT.md",
        "AGENT_EVOLUTION_LOG.md",
        "MULTI_AGENT_PROTOCOL.md",
        "SELF_MODIFICATION_RULES.md"
    ],
    "security": [
        "THREAT_MODEL.md",
        "ATTACK_SURFACE.md",
        "SECURITY_BOUNDARIES.md",
        "ZERO_TRUST_MODEL.md",
        "SANDBOX_POLICY.md",
        "PERMISSION_MODEL.md",
        "PRIVILEGE_ESCALATION.md",
        "CRYPTO_DESIGN.md",
        "INCIDENT_RESPONSE.md",
        "RED_TEAM_CASES.md",
        "DEFENSE_PLAYBOOK.md",
        "FORENSICS_GUIDE.md",
        "KILL_SWITCH.md",
        "FAIL_SAFE_MODES.md",
        "RECOVERY_PROTOCOL.md"
    ],
    "kernel": [
        "BOOT_SEQUENCE.md",
        "PROCESS_MODEL.md",
        "MEMORY_MODEL.md",
        "SCHEDULER_DESIGN.md",
        "FILESYSTEM_DESIGN.md",
        "DEVICE_MODEL.md",
        "DRIVER_REGISTRY.md",
        "IPC_MODEL.md",
        "RESOURCE_MANAGER.md",
        "POWER_MANAGEMENT.md",
        "REALTIME_SUPPORT.md",
        "KERNEL_LIMITS.md",
        "PERFORMANCE_TARGETS.md"
    ],
    "research": [
        "RESEARCH_INDEX.md",
        "PAPERS_TRACKER.md",
        "EXPERIMENT_LOG.md",
        "FAILED_IDEAS.md",
        "THEORY_NOTES.md",
        "MATH_MODELS.md",
        "GAME_THEORY.md",
        "CONTROL_SYSTEMS.md",
        "PHYSICS_REFERENCES.md",
        "ML_NOTES.md",
        "OPTIMIZATION_LOG.md"
    ],
    "validation": [
        "TEST_PLAN.md",
        "UNIT_TEST_MATRIX.md",
        "INTEGRATION_TESTS.md",
        "SYSTEM_TESTS.md",
        "CHAOS_TESTING.md",
        "STRESS_TESTS.md",
        "FAULT_INJECTION.md",
        "LATENCY_REPORT.md",
        "BENCHMARKS.md",
        "BUG_DATABASE.md",
        "KNOWN_LIMITATIONS.md",
        "VERIFICATION_STATUS.md",
        "MODEL_ACCURACY.md"
    ],
    "deployment": [
        "BUILD_PIPELINE.md",
        "CI_CD.md",
        "RELEASE_PROCESS.md",
        "VERSION_POLICY.md",
        "ROLLBACK.md",
        "INSTALLATION.md",
        "UPDATE_SYSTEM.md",
        "MIGRATION_GUIDE.md",
        "BACKUP_POLICY.md",
        "RESTORE_GUIDE.md"
    ],
    "operations": [
        "TELEMETRY.md",
        "LOGGING_SCHEMA.md",
        "METRICS.md",
        "EVENT_STREAMS.md",
        "OBSERVABILITY.md",
        "ALERTING_RULES.md",
        "HEALTH_MONITORING.md",
        "ANOMALY_DETECTION.md"
    ],
    "github": [
        "ISSUE_TEMPLATES.md",
        "PR_RULES.md",
        "CODE_REVIEW.md",
        "MERGE_CHECKLIST.md",
        "BRANCH_POLICY.md",
        "LABELS.md",
        "AUTOMATION_RULES.md",
        "BOT_BEHAVIOR.md",
        "REPO_HYGIENE.md",
        "DOCUMENTATION_POLICY.md"
    ],
    "runtime_ai": [
        "MEMORY_SYSTEM.md",
        "KNOWLEDGE_GRAPH.md",
        "PERSONALIZATION.md",
        "LEARNING_POLICY.md",
        "MODEL_SELECTION.md",
        "LOCAL_MODEL_REGISTRY.md",
        "MODEL_LIFECYCLE.md",
        "MODEL_PRUNING.md",
        "SELF_TRAINING.md",
        "ADAPTATION_RULES.md",
        "DRIFT_DETECTION.md",
        "MODEL_ROLLBACK.md"
    ],
    "emergency": [
        "DISASTER_RECOVERY.md",
        "CORRUPTION_RECOVERY.md",
        "BOOT_FAILURES.md",
        "AI_FAILURE_PROTOCOL.md",
        "SAFE_MODE.md",
        "MANUAL_OVERRIDE.md",
        "EMERGENCY_OPERATIONS.md"
    ],
    "diagrams": [],
    "experiments": [],
    "decisions": [],
    "archives": []
}

# Create base directories
for layer in layers:
    path = os.path.join(docs_root, layer)
    os.makedirs(path, exist_ok=True)
    print(f"Created directory: {path}")

# Initialize placeholder content for all 134 files
for layer, files in layers.items():
    for f in files:
        file_path = os.path.join(docs_root, layer, f)
        title = f.replace(".md", "").replace("_", " ").title()
        
        content = f"""# {title}

This document is part of the **{layer.title()}** layer of the PhoenixOS Structured Documentation Architecture.

## 1. Overview
Placeholder overview for {title}.

## 2. Technical Specification
Detailed specifications and boundaries for this module.

## 3. Operations & Auditing
Traceability, logs, and alignment verification requirements.

---
*Created: 2026-05-23 | Status: DRAFT / INCREMENTAL GROWTH*
"""
        with open(file_path, "w") as fh:
            fh.write(content)

print("Generated 134 placeholder document stubs.")

# ──────────────────────────────────────────────────────────────────────────
# Define rich, specific content for the 12 P0 Core Documents
# ──────────────────────────────────────────────────────────────────────────

core_docs = {}

# 1. PROJECT_VISION.md
core_docs["governance/PROJECT_VISION.md"] = """# Project Vision: PhoenixOS

PhoenixOS is a **Deterministic Cybernetic Security Runtime** designed to operate on Linux. The core objective of PhoenixOS is to achieve system security as a thermodynamic state of low entropy, autonomously "quenching" disorder via the Phoenix Matrix.

## Core Philosophy
In traditional operating systems, security is reactive and fragmented. PhoenixOS models the entire system state space as a thermodynamic physics system where anomalies represent high entropy (disorder). Through deterministic telemetry and closed-loop control, PhoenixOS continuously pushes the system back into a low-entropy state of safety.

## Six Immutable Axioms
1. **Determinism is sacred:** No reliance on non-deterministic primitives (unordered maps, race-dependent scheduling, or timestamp-only ordering).
2. **Replay is authoritative:** Causal replay governs security semantics. System logs, metrics, and AI recommendations are secondary to reproducible execution.
3. **AI is advisory:** AI modules inform and assist but never directly control or bypass the kernel or actuation finite state machine.
4. **Control must remain bounded:** Actuation is rate-limited, isolated, and strictly reversible to avoid denial-of-service and state oscillation.
5. **Telemetry correctness > AI sophistication:** Precise, replayable telemetry is the foundation of cybernetic control.
6. **Never scale instability:** Single-node stability and determinism must be mathematically validated before distributing execution.

## Target Users
- High-assurance cloud operations (Zero-Trust nodes).
- Mission-critical edge infrastructure.
- Automated security operations centers (SOC) needing deterministic threat hunting.

## Non-Goals
- Replacing general-purpose Linux distributions (PhoenixOS boots on top of Linux as a security runtime).
- Bypassing human override (operator manual control budget resets are always prioritized).

## Constraints
- **Resource Constraints:** Under 1GB memory allocator limits, strict eBPF telemetry processing latency (<1ms).
- **Latency Constraints:** Fast-path containment execution in under 100ms.
"""

# 2. ROADMAP.md
core_docs["governance/ROADMAP.md"] = """# Master Execution Roadmap: PhoenixOS

This roadmap charts the progress of PhoenixOS from single-node userspace validation to distributed swarm-level autonomous containment.

```
[Phase I: Hardening] -> [Phase II: eBPF & App] -> [Phase III: Graph & Swarm] -> [Phase IV: Strategic Policy] -> [Phase V: Custom Kernel]
      (Completed)            (Months 6-12)             (Months 12-24)               (Months 24-36)              (Months 36+)
```

## Phase I: Single-Node Stabilization & Hardening (COMPLETED)
- **Goal:** Architectural coherence and deterministic single-node replay validation.
- **Key Milestones:**
  - Standardized deterministic sequence reordering window in [guard.go](file:///Users/fallofpheonix/os/phoenix_os/guard/guard.go).
  - Implemented Evidence Reserve Lane (85%) and Overflow Snapshots (95%) in the event [bus.go](file:///Users/fallofpheonix/os/phoenix_os/bus/bus.go).
  - Hardened [warden.go](file:///Users/fallofpheonix/os/phoenix_os/warden/warden.go) state controller with dwell limits (30 ticks), cooldowns (10 ticks), and recovery budgets (3 de-escalations).
  - Created cryptographic Ledger V2 in [ledger.go](file:///Users/fallofpheonix/os/phoenix_os/ledger/src/ledger.go) supporting parent-hash chain verification.
  - Resolved Warden/Ledger concurrency race conditions and TCS sliding window sequence math.

## Phase II: Real Telemetry, Kernel space & Immutable Appliance (Months 6–12) (ACTIVE)
- **Goal:** Replace mock logs with real Linux eBPF/XDP telemetry probes.
- **Deliverables:**
  - eBPF probe collection (process exec, socket bind, file open).
  - Standalone bootable BusyBox/Initrd appliance image.
  - cgroups v2 containment integration.

## Phase III: Graph Intelligence & Swarm Observability (Months 12–24)
- **Goal:** Vector clock multi-node syncing and causal process lineage graph engine.
- **Deliverables:**
  - 3-tier DAG database traversals (<7ms for 100k nodes).
  - Consensus engine (Proof-of-Anomaly + Reputation).

## Phase IV: Strategic Policy & Advisor AI (Months 24–36)
- **Goal:** Game-theoretic defense solver and PhoenixMind LLM advisory loop.
- **Deliverables:**
  - Stackelberg policy solver (<1ms execution).
  - Kalman filtering on telemetry drift.

## Phase V: Autonomous Swarm OS (Months 36+)
- **Goal:** Bare-metal kernel scheduler patches and closed-loop self-repair.
- **Deliverables:**
  - CFS scheduler patches mapping CPU time slices to game-theoretic payoffs.
  - Swarm-level autonomous peer isolation.
"""

# 3. SYSTEM_ARCHITECTURE.md
core_docs["architecture/SYSTEM_ARCHITECTURE.md"] = """# System Architecture: PhoenixOS

PhoenixOS is structured around a 7-layer cybernetic stack known as the **Phoenix Matrix**.

```
                   ┌──────────────────────────────────────┐
                   │ L7: Swarm Coordination (Nexus)       │
                   └──────────────────┬───────────────────┘
                                      ▼
                   ┌──────────────────────────────────────┐
                   │ L6: System Physics (Sentinel)        │
                   └──────────────────┬───────────────────┘
                                      ▼
                   ┌──────────────────────────────────────┐
                   │ L5.5: Strategic Policy (Arbiter)     │
                   └──────────────────┬───────────────────┘
                                      ▼
                   ┌──────────────────────────────────────┐
                   │ L5: Actuation & Control (Warden)     │
                   └──────────────────┬───────────────────┘
                                      ▼
                   ┌──────────────────────────────────────┐
                   │ L4: Graph Intelligence (Trace)       │
                   └──────────────────┬───────────────────┘
                                      ▼
                   ┌──────────────────────────────────────┐
                   │ L3: Telemetry Math (Monitor/TCS)     │
                   └──────────────────┬───────────────────┘
                                      ▼
                   ┌──────────────────────────────────────┐
                   │ L2: Kernel Runtime (Probes)          │
                   └──────────────────┬───────────────────┘
                                      ▼
                   ┌──────────────────────────────────────┐
                   │ L1: Platform Integrity (Guard)       │
                   └──────────────────────────────────────┘
```

## Layer Breakdown

1. **L7: Swarm Coordination (Phoenix Nexus):** Proof-of-Anomaly consensus and multi-node replication.
2. **L6: System Physics (Phoenix Sentinel):** Continuous SDI signal physics modeling.
3. **L5.5: Strategic Policy (Phoenix Arbiter):** Game-theoretic decision engine solving Stackelberg policy bounds.
4. **L5: Actuation & Control (Phoenix Warden):** 5-State Discrete FSM (SAFE -> WATCH -> SUSPICIOUS -> CRITICAL -> COMPROMISED) enforcing hysteresis constraints.
5. **L4: Graph Intelligence (Phoenix Trace):** Causal process lineage DAGs stored across HOT/WARM/COLD tiers.
6. **L3: Telemetry Math (Phoenix Monitor):** Signal processing engine utilizing Kalman filters and sliding window sequence verification.
7. **L2: Kernel Runtime (Phoenix Kernel):** eBPF telemetry collectors hook process, socket, and file boundaries.
8. **L1: Platform Integrity (Phoenix Guard):** Microsecond-latency fast-path syscall enforcement maps.

## Inter-Process Communication (IPC) & Flow Paths
- **Raw Telemetry Ingestion:** Syscalls -> L1 (Guard) -> L2 (eBPF Probes) -> L3 (Event Bus).
- **Processing Loop:** L3 Event Bus -> L4 (Trace Lineage Graph) -> L3 Monitor (Entropy & Kalman calculations) -> L5.5 Arbiter (Payoff & Threshold gating) -> L5 Warden (FSM State changes).
- **Advisory AI Loop:** Chaotic signals (Z > 3.0) queue requests to PhoenixMind LLM (running as an asynchronous advisory thread) which provides offline audit suggestions.
"""

# 4. COMPONENT_MAP.md
core_docs["architecture/COMPONENT_MAP.md"] = """# Subsystem Component Map

Every module within PhoenixOS is assigned ownership, directory structure, and dependencies.

| Subsystem | Folder | Owner | Core Dependencies | Purpose |
| :--- | :--- | :--- | :--- | :--- |
| **Bus** | [bus/](file:///Users/fallofpheonix/os/phoenix_os/bus) | Platform Team | Go channels | Ingestion and queue-pressure scheduling |
| **Guard** | [guard/](file:///Users/fallofpheonix/os/phoenix_os/guard) | Kernel Team | eBPF (Stage 2) | Telemetry sorting and sequence proof hashing |
| **Ledger** | [ledger/](file:///Users/fallofpheonix/os/phoenix_os/ledger) | Security Team | SHA-256 / Allocator | Verifiable state evidence Merkle-DAG logging |
| **Monitor** | [monitor/](file:///Users/fallofpheonix/os/phoenix_os/monitor) | Research Team | Kalman / EWMA | Entropy drift modeling and frequency metrics |
| **TCS** | [tcs/](file:///Users/fallofpheonix/os/phoenix_os/tcs) | Research Team | Event Bus | Jitter and sequence-loss confidence scores |
| **Warden** | [warden/](file:///Users/fallofpheonix/os/phoenix_os/warden) | Platform Team | Event Bus | Finite state containment actuation |
| **Arbiter** | [arbiter/](file:///Users/fallofpheonix/os/phoenix_os/arbiter) | Research Team | Policy maps | Strategic game-theoretic decision solver |
| **AI Orchestrator** | [ai/](file:///Users/fallofpheonix/os/phoenix_os/ai) | AI Team | Subsystem Features | Asynchronous advisory loop and LLM driver |

## Subsystem Boundaries
Modules are strictly decoupled. No strategic layer (e.g. Arbiter) should directly manipulate Kernel states; all strategic modifications must route through the Warden FSM.
"""

# 5. THREAT_MODEL.md
core_docs["security/THREAT_MODEL.md"] = """# Threat Model: PhoenixOS

This threat model defines our assets, attacker vectors, entry points, impacts, and defenses under a Zero-Trust architecture.

## Asset Catalog
1. **Evidence Ledger:** Cryptographic proof chain documenting process actions and FSM state transitions.
2. **Telemetry Stream:** Live syscall records driving entropy monitoring.
3. **Warden FSM State:** Current system containment level (NORMAL -> CONTAINED).
4. **Policy Configuration:** Gating thresholds and de-escalation budgets.

## Attacker Profile & Scenarios

### Threat 1: Denial-of-Service via Telemetry Noise (Mimicry Attack)
- **Attacker:** Compromised low-privilege userspace process generating massive normal-severity traffic.
- **Entry Point:** Syscall interface (L1/L2).
- **Impact:** Ring-buffer saturation leading to dropping of actual critical threat alerts.
- **Defense:** *Evidence Reserve Lane* blocks low-severity events at 85% queue pressure. *Priority Pre-emption Shield* drops oldest low-severity events to accommodate critical events at 100% capacity.

### Threat 2: Time Warp State Poisoning
- **Attacker:** Root process attempting to inject fake historical events to trigger rollback or reset de-escalation history.
- **Entry Point:** Replay injection endpoint.
- **Impact:** FSM escapes containment or resets de-escalation budget without approval.
- **Defense:** Bounded sequence window verification in [tcs.go](file:///Users/fallofpheonix/os/phoenix_os/tcs/tcs.go) rejects future/past anomalies, and Ledger validation hashes enforce state progression integrity sequentially.

### Threat 3: Concurrent State Manipulation (SOC Operator Bypass)
- **Attacker:** Malicious script invoking SOC `/game/action` API concurrently during high-load telemetry runs.
- **Entry Point:** Game Server HTTP API.
- **Impact:** Warden state corruption, double-decrementing budgets.
- **Defense:** Enforced Mutex locks on `Warden` and RWMutex on `Ledger` ensure atomic, serialized state transitions.
"""

# 6. SECURITY_BOUNDARIES.md
core_docs["security/SECURITY_BOUNDARIES.md"] = """# Security Boundaries & Isolation Zones

PhoenixOS operates strict isolation zones to contain anomalies and protect platform integrity.

## 1. Fast-Path Isolation (L1 Guard)
- Syscalls that trigger high-confidence alerts (entropy > 7.9) immediately trigger containment in under 100ms via L1 Guard kernel hooks, bypassing Userspace, Strategic Policy, and AI layers.

## 2. Warden State Containment (L5)
- **CONTAINED State:** Suspicious PIDs are frozen using cgroups v2 freezer controller. Sockets are isolated via XDP eBPF filters.
- **Privilege Limits:** Warden runs with minimal caps (CAP_SYS_ADMIN for cgroups control and eBPF maps) and is locked to PID 1 in the final OS appliance.

## 3. Cryptographic Tamper-Proofing (Ledger)
- The evidence log is append-only.
- Every entry hash incorporates the previous block's hash.
- Gaps or clock tampering immediately cause `Ledger.Verify()` to fail, preventing node boot or peer gossip verification.
"""

# 7. EXTERNAL_REPOS.md
core_docs["integration/EXTERNAL_REPOS.md"] = """# External Repository Registry & Merge Policies

This document registry tracks all external code imported and modified under the PhoenixOS project.

## Registry

1. **go-sqlite3**
   - **Purpose:** SQL database storage for L4 Trace process graphs.
   - **Version:** v1.14.22 (pinned).
   - **License:** MIT.
   - **Fork Status:** Unmodified upstream dependency.
   - **Risk Level:** Low.
   - **Replacement Plan:** Fallback to raw flat files or badger key-value store if memory overhead becomes restrictive.

## Merge Policy
- Upstream packages must be vendored or pinned to exact commit SHA-256 hashes.
- Direct code changes to external libraries are strictly prohibited without an RFC documenting safety impact and fork strategy.
"""

# 8. AGENT_REGISTRY.md
core_docs["agents/AGENT_REGISTRY.md"] = """# AI Agent Registry

This document records the identity, capabilities, and boundaries of all AI agents operating within PhoenixOS.

## Agent Identity: PhoenixMind (L6 Advisor)
- **Agent ID:** PM-001
- **Role:** Cybersecurity Forensic Analyzer & Explainer.
- **Input:**
  - `monitor.DriftScore` (Z-Score, EventType, Severity, Frequency)
  - `tcsScore` (Confidence index)
- **Output:**
  - Suggest Command (ISOLATE_PID, THROTTLE_NETWORK, REVOKE_UID, LOG_ONLY)
  - Confidence Score (0.0 to 1.0)
  - Reasoning (Text explaining anomaly)
- **Permissions:** Read-only access to L4 Trace DB and telemetry event feeds. No write permissions to Ledger, Warden, or raw filesystem.
- **Memory Scope:** Stateless inference per batch request (5-second aggregation window).
- **Allowed Actions:** Push JSON advice strings to log streams for manual operator inspection.
- **Escalation Path:** In case of LLM drift or anomalous command generation, the AI loop drops the advice and defaults to static Warden FSM containment rules.

## AI Alignment Mandate
AI is strictly advisory. Under **Axiom 3**, AI outputs must pass human review or warden threshold verification. AI can never autonomously escalate system state to recovery or containment.
"""

# 9. TEST_PLAN.md
core_docs["validation/TEST_PLAN.md"] = """# Test Plan & Verification Matrix

PhoenixOS requires rigorous verification to ensure determinism and zero state drift under adversarial stress.

## 1. Test Tiers

### Tier 1: Deterministic Replay Verification
- **Command:** `go test ./guard`
- **Method:** Replay identical sequence events through the mock adapter and check that the resulting Ledger genesis hashes and sequenceProof hashes are bit-for-bit identical.

### Tier 2: Concurrency & Race Testing
- **Command:** `go test -race ./warden ./ledger/src`
- **Method:** Stress the state controllers using parallel goroutines executing concurrent reads/writes (FSM actuations, ledger entry additions, budget resets) to ensure no data races or panics occur.

### Tier 3: TCS Math Stability Tests
- **Command:** `go test ./tcs`
- **Method:** Validate out-of-order sequence insertion, missing sequence IDs, and negative SeqID filters to check that TCS confidence metrics remain within $[0.0, 1.0]$.

## 2. Invariant Gates
- **Zero Race Conditions:** Build and verify with `-race` flag must return 0 issues.
- **FSM Cooldown:** Confirm that Warden rejects back-to-back state escalations inside the 10-tick cooldown window unless severity class >= ClassLocalIsolate.
"""

# 10. BUILD_PIPELINE.md
core_docs["deployment/BUILD_PIPELINE.md"] = """# Build Pipeline & Versioning

PhoenixOS build automation guarantees the integrity of code from compilation to standalone appliance rollout.

## Pipeline Steps

```
[Lint & Static Analysis] -> [Unit & Race Tests] -> [Go Compile] -> [Appliance Packaging]
```

1. **Lint & Static Analysis:** Run `golangci-lint` to verify code hygiene and find potential unhandled map locks.
2. **Race-Detector Tests:** Execute `go test -race ./...` to prevent concurrency regressions.
3. **Compilation:** Build main binary targeting static compilation (`CGO_ENABLED=1` for SQLite support).
4. **Appliance Packaging:** Package the binary as an initrd/busybox bootable appliance (Stage 3).

## Version Policy
PhoenixOS adheres to semantic versioning (Major.Minor.Patch). Documentation modifications mapping to state schema changes must increment the Minor version.
"""

# 11. MASTER_STATUS.md
core_docs["governance/MASTER_STATUS.md"] = """# Master Status: PhoenixOS

This status report tracks the current active state of the PhoenixOS runtime modules.

## Current State
- **Active Stage:** Stage 2 (Real Telemetry Runtime).
- **Core Status:** Concurrency hardening completed and verified using Go race detection.

## Completed Modules
- **Ledger V2:** Verifiable parent-hash chain, logical clock allocation checks, and RWMutex safety.
- **Warden FSM:** Rate-limited state transitions, 30-tick dwell hysteresis, 10-tick cooldowns, and manual operator recovery budget reset.
- **TCS:** Telemetry Confidence Score calculations, out-of-order stabilization, and negative sequence event filtering.
- **Main Loop:** Deterministic execution reorder window and logical clock standardizations.

## Active Subsystems
- eBPF probe adapters for syscall monitoring (Stage 2).
- LinuxKit appliance boot assembly (Stage 3).

## Active Risks
- **Overhead Risk:** Real-time eBPF packet parsing overhead under heavy networking load. Mitigation: strict event filters inside eBPF kernel space.
"""

# 12. DECISION_LOG.md
core_docs["governance/DECISION_LOG.md"] = """# Architectural Decision Log

This log lists critical architectural choices made during the PhoenixOS project.

## Decision 1: Concurrency Mutex vs. Channel-Only Synchronization
- **Status:** APPROVED.
- **Alternatives Considered:** Processing all SOC API requests sequentially via a single coordination channel.
- **Tradeoff Analysis:** While channel loops are clean, Warden FSM triggers need immediate microsecond-latency evaluation to protect platform integrity. Adding explicit Mutex locks on Warden and RWMutex locks on Ledger minimizes latency and guarantees map thread safety.

## Decision 2: TCS Sliding Window Dynamic Range Scanning
- **Status:** APPROVED.
- **Alternatives Considered:** Relying on slice boundaries `events[len-1] - events[0]` for sequence delta calculations.
- **Tradeoff Analysis:** Slice endpoints assume chronological events match sequence progression. Because network telemetry can arrive out of order, or system overflow events can inject negative SeqIDs, dynamically scanning active window events prevents underflows and caps loss rate estimates accurately.

## Decision 3: Advisory-Only LLM Loop
- **Status:** APPROVED.
- **Alternatives Considered:** Allowing the AI Orchestrator to directly trigger Warden state escalations if confidence > 90%.
- **Tradeoff Analysis:** Direct AI-actuated FSM control violates **Axiom 3** and introduces non-deterministic model failures. Keeping LLM outputs strictly advisory ensures audit logs are informative without compromising system predictability.
"""

# Write core docs
for path_suffix, content in core_docs.items():
    full_path = os.path.join(docs_root, path_suffix)
    with open(full_path, "w") as fh:
        fh.write(content)
    print(f"Populated core document: {full_path}")

print("Successfully generated Structured Documentation Architecture!")
