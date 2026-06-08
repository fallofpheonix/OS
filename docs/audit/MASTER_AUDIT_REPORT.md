---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# PhoenixOS Security Review

## Threat Model & Trust Boundaries
PhoenixOS operates under a zero-trust model where all user-space processes are untrusted. The security enforcement engine (`Warden`) runs in a privileged daemon space, utilizing kernel-level eBPF helper maps and LSM hooks to enforce isolation.

```
       [ USER SPACE (Untrusted) ]
                   │
                   ▼ (Syscalls)
┌──────────────────────────────────────┐
│       1. KERNEL SENSING (eBPF)       │
└──────────────────┬───────────────────┘
                   │  (Ring Buffer)
                   ▼  [TRUST BOUNDARY]
┌──────────────────────────────────────┐
│         2. DAEMON SPACE (Go)         │
│  - Core Bus, Monitor, Arbiter        │
│  - AI Orchestrator                   │
│  - Warden FSM (Enforcement)          │
└──────────────────┬───────────────────┘
                   │  (Syscalls / Maps)
                   ▼  [TRUST BOUNDARY]
┌──────────────────────────────────────┐
│      3. KERNEL ENFORCEMENT (LSM)     │
└──────────────────────────────────────┘
```

---

## Formal Invariants & Proof Gates
Before the `Warden` executes any FSM state transition or triggers physical actuators, it passes the `AuthorityEscalationRequest` through three formal invariants:

1. **Evidence Weight Invariant (`EvidenceWeightInvariant`)**:
   Enforces the *Conservation of Authority Law*. Transitioning to a higher alert state requires a minimum evidence weight (Drift Score).
   - SAFE → WATCH: Threshold 0.5
   - WATCH → SUSPICIOUS: Threshold 0.7
   - SUSPICIOUS → CRITICAL: Threshold 0.9
   - CRITICAL → COMPROMISED: Threshold 1.0

2. **Certificate Invariant (`CertificateInvariant`)**:
   Verifies that the evidence weight and event parameters are backed by a cryptographically signed certificate matching the append-only `Ledger`. This prevents spoofed telemetry from causing denial-of-service escalations.

3. **Contextual Invariant (`ContextualInvariant`)**:
   Validates the causal Directed Acyclic Graph (DAG) path provided in the request (`GraphProof`).
   - Ensures that the path of execution is verified.
   - Namespace Drift Detection: Validates that the process namespace inode (`TargetNsproxy`) matches the lineage (`ExpectedNsproxy`), detecting container escapes or namespace breakouts instantly.

---

## Security Best Practices
- **No Hardcoded Credentials**: Checked and validated; all keys and configuration paths are resolved dynamically or passed via environment variables.
- **Input Validation**: Telemetry payloads parsed through strict JSON validation before processing.
- **Race Condition Mitigations**: Shared state protected by thread-safe primitives (sync.RWMutex, atomic operations).
- **Emergency Halt**: If any invariant check fails during transition verification, the Warden triggers an emergency shutdown, locking the FSM state permanently until process restart.
# Content-First Orphan Classification Report

**Date:** 2026-06-01
**Methodology:** CLASSIFY-002 (content inspected, dependency graph generated, runtime/test/roadmap usage checked, historical purpose documented)
**Status:** REVIEWED — decisions recorded in ADR-002.

---

## Review Outcome (ADR-002)

| Candidate | Original Recommendation | Review Decision | Authority |
|-----------|------------------------|-----------------|-----------|
| PhoenixSimulation | ARCHIVE | **Approved** — archive/research/PhoenixSimulation/ | Phase A |
| PhoenixGuard/physics | DELETE | **Approved** | Phase A |
| PhoenixOS/agents | External/separate repo | **Approved** — External/ClaudeCode/ or External/Agents/ | Phase A |
| PhoenixStimulation | ARCHIVE (with split) | **Approved with conditions** — keep stress/, scenarios/; archive remainder | Phase B |
| PhoenixOS/core (pheonixmind-core) | ARCHIVE | **Escalated** — audit required; archive/legacy/ if migration/reasoning found | Phase C |
| PhoenixOS/core (phoenix_os) | ARCHIVE | **Escalated** — audit required; archive/legacy/ if migration/reasoning found | Phase C |
| PhoenixOS/memory | ARCHIVE | **Not contested** — aligns with Phase C audit scope | Phase C |
| PhoenixGuard/integrated_model | DELETE | **Reversed** — diverged duplicate is not a duplicate; archive/legacy/ pending historical analysis | Phase C |

**New Rule:** CLASSIFY-003 — If a component was previously classified as a duplicate and later found to have diverged behavior, automatic deletion authority is revoked. Historical and capability analysis required before removal.

See [docs/adr/ADR-002-Classification-Decisions.md](docs/adr/ADR-002-Classification-Decisions.md) for full review.

---

## 1. PhoenixSimulation

| Field | Answer |
|-------|--------|
| **Purpose** | Self-healing CI/build pipeline — panic aggregation log tailer, forbidden-path guard, mock patch synthesis, test runner stubs |
| **Primary Capability** | Monitoring build failures and proposing code patches via a simulated evolution loop |
| **Production Usage** | No — no main package, no entry point, no runtime invocation |
| **Test Usage** | No — zero references from any test file in the workspace |
| **Roadmap Usage** | No — no OWNERSHIP, CLAUDE, or README; single TODO references an abandoned LLM integration idea |
| **Imports** | Standard library only (`sync`, `fmt`, `time`, `bufio`, `os`, `strings`) |
| **Imported By** | Zero |
| **Duplicate** | No — content is unique |
| **External** | No — all code appears Phoenix-authored |
| **Content Classification** | Prototype/experimental CI tooling |
| **Recommended Domain** | **Archive** (experimental prototype, zero consumers, no owner) |
| **Confidence** | 100% |
| **CLASSIFY-002 Blockers** | None — all 6 steps complete |

**Notes:** 10 of 15 Go files are stubs (4 lines, no code). The 5 real files implement a log tailer that reads panic logs, aggregates failures, and proposes patches — but nothing calls it, nothing owns it, and it references an LLM that was never integrated.

---

## 2. PhoenixStimulation

| Field | Answer |
|-------|--------|
| **Purpose** | Stress testing and scenario injection for replay verification |
| **Primary Capability** | `stress.NewReplaySurge(bus)` — generates synthetic event surges for replay stress tests; `scenarios.NewInjector(bus)` — injects simulated scenarios (e.g., SystemPulse) into the event bus |
| **Production Usage** | No — both importers are `cmd/` test binaries (replay_stress, pulse) |
| **Test Usage** | **Yes** — imported by PhoenixRedteam `cmd/replay_stress/main.go` and `cmd/pulse/main.go` |
| **Roadmap Usage** | Yes — documented in REPOSITORY_OWNERSHIP.md as repository #13 under Crucible |
| **Imports** | PhoenixCore/bus, plus internal PhoenixStimulation packages |
| **Imported By** | PhoenixRedteam (2 source files), PhoenixValidation (file path reference) |
| **Duplicate** | Contains many duplicate sub-modules (kernel, telemetry, security, ledger) that mirror Nucleus modules — these should be separated |
| **External** | No core logic is external; sub-modules duplicate Nucleus content |
| **Content Classification** | Testing + simulation infrastructure |
| **Recommended Domain** | **Crucible** — KEEP the `stress/` and `scenarios/` packages (test support). ARCHIVE duplicate sub-modules (kernel, telemetry, security, ledger) |
| **Confidence** | 90% — need to separate the test-support packages from the duplicate sub-modules |
| **CLASSIFY-002 Blockers** | Need to split: `stress/` and `scenarios/` stay in Crucible; `phoenix_os/` sub-modules go to archive |

**Notes:** The module path is `github.com/fallofpheonix/phoenix-os` — same as PhoenixOS, which is a naming collision. The module is NOT dead code — 2 source files in PhoenixRedteam import it. But ~80% of its internal structure is duplicate Nucleus modules that should be removed.

---

## 3. PhoenixOS/core/pheonixmind-core

| Field | Answer |
|-------|--------|
| **Purpose** | Define interface contracts for PhoenixMind subsystems (IAgent, ILedger, IMemory, IObservability, IRuntime, ISandbox, ISecurity, IValidator, IModelRouter) |
| **Primary Capability** | Provide pure Go interface definitions (8 interfaces, 1 method each) |
| **Production Usage** | No — zero references from any workspace module |
| **Test Usage** | No |
| **Roadmap Usage** | No — no OWNERSHIP, no README, no CLAUDE. Only documentation reference is an issue note saying "Interface contracts defined" |
| **Imports** | None (stdlib only) |
| **Imported By** | Zero |
| **Duplicate** | Partially — real PhoenixMind already has its own internal contracts |
| **External** | No |
| **Content Classification** | Abandoned interface definitions — obsolete architectural remnant |
| **Recommended Domain** | **Archive** (no consumers, no owner, superseded by PhoenixMind's own contracts) |
| **Confidence** | 95% |
| **CLASSIFY-002 Blockers** | None |

**Notes:** 8 interfaces, each with a single method. No implementations. No consumers. These are from an earlier PhoenixMind architecture that was replaced. The current PhoenixMind builds and tests without them.

---

## 4. PhoenixOS/core/phoenix_os

| Field | Answer |
|-------|--------|
| **Purpose** | Truth/evidence tracking, trace/lineage, causal graph, timeline, checkpoint management |
| **Primary Capability** | Partial implementation of evidence registry (sync-protected map) and trace data structures (in-memory graph, timeline, checkpoints) with TruthState constants |
| **Production Usage** | No — zero references from any workspace module |
| **Test Usage** | Self-contained test in phoenixmind-trace (tests its own trace types, no external dependency) |
| **Roadmap Usage** | No — no OWNERSHIP, no README, no CLAUDE |
| **Imports** | `sync`, `time`, plus external `github.com/fallofpheonix/phoenixmind-validator/truth/evidence` (non-workspace module) |
| **Imported By** | Zero |
| **Duplicate** | Similar functionality exists in `Phoenix.Nucleus/PhoenixTruth` and `Phoenix.Nucleus/PhoenixTrace` — the real workspace modules |
| **External** | Has external dependency to a module outside the workspace |
| **Content Classification** | Abandoned partial implementation — 12 of 17 files are broken stubs (no package declaration) |
| **Recommended Domain** | **Archive** (partial implementation, 70% broken files, external dependency to non-workspace module) |
| **Confidence** | 95% |
| **CLASSIFY-002 Blockers** | None |

**Notes:** The evidence types package (evidence_types.go, evidence_registry.go) has working code, but the trace package imports from `github.com/fallofpheonix/phoenixmind-validator/truth/evidence` which is not in the workspace. The real PhoenixTruth and PhoenixTrace modules in Nucleus already provide this capability.

---

## 5. PhoenixOS/core/trace/phoenixmind-trace

(Merged with #4 above — same Go module as phoenix_os/phoenix_os)

---

## 6. PhoenixOS/memory

| Field | Answer |
|-------|--------|
| **Purpose** | Memory subsystem scaffold (planning) |
| **Primary Capability** | None — zero Go files exist |
| **Production Usage** | No |
| **Test Usage** | No |
| **Roadmap Usage** | No — no OWNERSHIP, no README, no CLAUDE |
| **Imports** | N/A (no code) |
| **Imported By** | N/A |
| **Duplicate** | N/A |
| **External** | N/A |
| **Content Classification** | Empty scaffold — directory with only go.mod and DIRECTORY_NOTES.md |
| **Recommended Domain** | **Archive** (empty scaffold, no content) |
| **Confidence** | 100% |
| **CLASSIFY-002 Blockers** | None |

**Notes:** Single go.mod with module path `github.com/fallofpheonix/phoenix-os/phoenixmind-memory`, Go 1.25.0, no dependencies. Zero Go files. No content to preserve.

---

## 7. PhoenixOS/agents

| Field | Answer |
|-------|--------|
| **Purpose** | Agent persona definitions, command workflows, hooks, and skills for Claude Code ECC plugin system |
| **Primary Capability** | Provide 62 agent personas, 77 command definitions, 57+ skill definitions, and a lifecycle hook system — all in Markdown/JSON |
| **Production Usage** | No — not Go code, not compiled, not executed by PhoenixOS |
| **Test Usage** | No |
| **Roadmap Usage** | Yes — this is an operational agent plugin library (Claude Code ECC plugin) |
| **Imports** | N/A (no Go code) |
| **Imported By** | N/A (no Go code) |
| **Duplicate** | No |
| **External** | Yes — this is a Claude Code plugin library, not PhoenixOS runtime code |
| **Content Classification** | Configuration/knowledge library for Claude Code agent ecosystem |
| **Recommended Domain** | **External** or separate repository — this is NOT PhoenixOS runtime code. It's a Claude Code plugin library that happens to be stored here. |
| **Confidence** | 90% — needs decision on whether to keep in-repo or split |
| **CLASSIFY-002 Blockers** | This is a policy decision, not a technical one. The content is valid (agent definitions work for Claude Code), but it's not PhoenixOS architecture. |

**Notes:** Contains zero Go files. This is an extensive library of agent definitions, commands, hooks, and skills for the Claude Code ECC plugin system. It's operational — the hooks.json references Node.js scripts that run during agent lifecycle events. This is infrastructure for AI-assisted development, not PhoenixOS runtime.

---

## 8. PhoenixGuard/integrated_model

| Field | Answer |
|-------|--------|
| **Purpose** | Integrated security sentinel model demo |
| **Primary Capability** | Simulated sentinel model that integrates disorder entropy, trace graphs, and telemetry |
| **Production Usage** | No — demo/example code in `src/main.go` |
| **Test Usage** | No |
| **Roadmap Usage** | No — deprecated copy |
| **Imports** | `PhoenixCore/security/physics`, `PhoenixCore/telemetry/entropy_engine`, `PhoenixTrace/engine/process_graphs` |
| **Imported By** | Zero |
| **Duplicate** | **Yes, with divergence** — go.mod present (Guard) vs absent (Core). Import paths differ: `entropy_engine` (Guard) vs `entropy_engine_go` (Core). Function call differs: `AddNode(2 args)` vs `AddNode(3 args)`. No git history available. |
| **External** | No |
| **Content Classification** | Copy of PhoenixCore/security/integrated_model with minor divergence during copy |
| **Recommended Domain** | **Delete** (diverged duplicate, no consumers, no unique value) |
| **Confidence** | 90% — content slightly diverged from original but has no consumers and no roadmap ownership |
| **CLASSIFY-002 Blockers** | Would benefit from confirming the Core version is the canonical one |

---

## 9. PhoenixGuard/physics

| Field | Answer |
|-------|--------|
| **Purpose** | Thermodynamic modeling (SDI calculation via Shannon Entropy, system state computation) |
| **Primary Capability** | `disorder.CalculateSDI()` — Security Disorder Index from event distributions; `physics.ComputeState()` — thermodynamic system state |
| **Production Usage** | No — zero references from PhoenixGuard source files |
| **Test Usage** | No |
| **Roadmap Usage** | No — deprecated copy |
| **Imports** | `math`, `PhoenixCore/security/physics/disorder` |
| **Imported By** | Zero |
| **Duplicate** | **Yes, with minimal divergence** — go.mod present (Guard) vs absent (Core). Go files are byte-for-byte identical. The Core version is the canonical one (imported by PhoenixMind). |
| **External** | No |
| **Content Classification** | Duplicate of PhoenixCore/security/physics |
| **Recommended Domain** | **Delete** (duplicate, zero consumers, canonical version in PhoenixCore) |
| **Confidence** | 95% |
| **CLASSIFY-002 Blockers** | None — confirmed duplicate, Core version is the active one |

---

## Consolidated Action Table

| Candidate | Action | Rationale | Status |
|-----------|--------|-----------|--------|
| PhoenixSimulation | **ARCHIVE** | Abandoned prototype, zero consumers, no owner | ✅ Approved (Phase A) |
| PhoenixStimulation (stress/, scenarios/) | **KEEP in Crucible** | Active dependency of PhoenixRedteam test binaries | ✅ Approved (Phase B) |
| PhoenixStimulation (obsolete, prototypes, unused, abandoned) | **ARCHIVE** | Duplicate Nucleus modules, not imported | ✅ Approved (Phase B) |
| PhoenixOS/core/pheonixmind-core | **AUDIT then ARCHIVE** | Escalated — may contain historical reasoning | 🔶 Escalated (Phase C) |
| PhoenixOS/core/phoenix_os | **AUDIT then ARCHIVE** | Partial impl, may contain ADRs-as-code or unfinished migration | 🔶 Escalated (Phase C) |
| PhoenixOS/memory | **ARCHIVE** | Empty scaffold, zero content | 🔶 Escalated (Phase C) |
| PhoenixOS/agents | **External/ClaudeCode/** | Claude Code plugin library, not PhoenixOS runtime | ✅ Approved (Phase A) |
| PhoenixGuard/integrated_model | **archive/legacy/** | Diverged duplicate — needs historical analysis; not safe to delete | ❌ Reversed (Phase C) |
| PhoenixGuard/physics | **DELETE** | Byte-for-byte duplicate of PhoenixCore/security/physics | ✅ Approved (Phase A) |

## Previously Classified (Re-verified)

| Candidate | Classification | Status |
|-----------|---------------|--------|
| Phoenix.Cognition/go.mod | **ACTIVE — KEEP** | Re-verified: PhoenixMind imports from it |
| PhoenixRedteam | **ACTIVE — KEEP** | Used for validation, security testing, CI |
| Phoenix.UI/Service | **ACTIVE — KEEP** | Fixed and building |

## Contradictions Resolved

| Previous Claim | Corrected Claim |
|---------------|-----------------|
| "PhoenixStimulation → DEAD_CODE" | **CORRECTED:** stress/scenarios packages are active (imported by Redteam). Duplicate sub-modules are dead code. |
| "PhoenixRedteam → UNUSED" | **RETRACTED:** Redteam is test infrastructure. It has no runtime production path but is active for testing/validation. |
| "PhoenixGuard copies are byte-for-byte identical" | **CORRECTED:** integrated_model is NOT byte-for-btye (go.mod differs, import paths differ, function signatures differ). physics IS byte-for-byte. |
| "PhoenixOS/core → ABANDONED" | **CORRECTED:** 3 sub-modules, all abandoned, but the `evidence` and `trace` packages within `phoenix_os` have partial implementations that could inform future work. |

## CLASSIFY-002 Verification

| Candidate | Content Inspected | Dep Graph | Runtime Usage | Test Usage | Roadmap Ownership | Historical Purpose | Action Blocked? |
|-----------|:-:|:-:|:-:|:-:|:-:|:-:|:-:|
| PhoenixSimulation | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | No |
| PhoenixStimulation | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | Yes (needs split) |
| PhoenixOS core/pheonixmind-core | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | No |
| PhoenixOS core/phoenix_os | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | No |
| PhoenixOS memory | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | No |
| PhoenixOS agents | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | No |
| PhoenixGuard integrated_model | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | No |
| PhoenixGuard physics | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | No |
# PhoenixOS External Repository Integration Audit

This document serves as the master architectural blueprint for integrating the 70+ cloned open-source repositories into the PhoenixOS ecosystem. It details the precise location, technical methodology, strategic reasoning, and sequencing for every external component.

## Strategic Sequencing (The "When")

To prevent architectural collapse, integration must follow a strict dependency graph:

*   **Phase 1: Foundation & Orchestration (Weeks 1-4):** Establish the core network (gRPC/Zilla), the underlying eBPF telemetry (Pixie/Tetragon), and the master routing DAGs (LangGraph/AutoGen).
*   **Phase 2: Code Intelligence & Memory (Weeks 5-8):** Wire up the Vector DBs (Qdrant), long-term memory (Letta/Mem0), and the foundational AST engines (GitNexus/OpenSandbox) to give the OS a durable brain.
*   **Phase 3: Autonomous Engineering & Cyber Ops (Weeks 9-12):** Deploy the offensive/defensive agents (PentestGPT, Aegis, SWE-AF) into the sandboxed environments created in Phase 2.
*   **Phase 4: Optimization & Embodied Simulation (Weeks 13+):** Implement hardware optimizations (vLLM, DeepSpeed) and connect the cognitive layer to physical/simulated environments (IsaacLab, Godot).

---

## Layer 1: OS Core (eBPF · Kernel · Telemetry)
**Location:** `Phoenix.Cognition/External/Layer_1_OS_Core/`
**Goal:** Provide zero-overhead, deep-kernel observability and security without requiring application-level code changes.

| Repository | Why (Strategic Purpose) | How (Integration Method) | When |
| :--- | :--- | :--- | :--- |
| **`pixie`** | Provides auto-instrumented eBPF observability for the entire K8s cluster. | Deploy via Pixie Operator. Ingest traces directly into the PhoenixOS central logging bus. | Phase 1 |
| **`tetragon`** | Cilium's security observability. Crucial for causal process tracing and in-kernel enforcement. | Run as a DaemonSet. Configure custom YAML policies to block malicious syscalls from AI agents. | Phase 1 |
| **`falco`** | Cloud-native runtime security. Detects abnormal behavior in real-time. | Integrate Falco rules Engine via eBPF probes. Route alerts to `falcoclaw` for remediation. | Phase 2 |
| **`tracee`** | Deep forensic logging of security-relevant events (file access, net connections). | Deploy alongside Tetragon for forensic evidence collection during anomalous agent behavior. | Phase 2 |
| **`kubearmor`** | Enforces access control and behavioral policies using LSMs at the kernel level. | Integrate with PhoenixOS's identity layer to restrict workload capabilities (e.g., blocking `execve`). | Phase 2 |
| **`odigos`** | Distributed tracing with eBPF. Auto-instruments all services across 20+ languages. | Use as the primary OpenTelemetry collector to map the dependency graph of microservices. | Phase 1 |
| **`coroot`** | eBPF-based APM for generating service maps and dependency graphs. | Feed its service maps into `GitNexus` to correlate runtime dependencies with static code structure. | Phase 2 |
| **`beyla`** | eBPF-based service mesh tracing (HTTP/gRPC) without code changes. | Use for immediate, low-effort observability on new agent deployments. | Phase 1 |
| **`parca`** | Continuous memory/CPU profiling at scale using eBPF. | Run constantly to identify resource leaks in long-running AI orchestration loops. | Phase 2 |
| **`anteon`** | eBPF observability tuned for K8s performance and cost optimization. | Use its load engine to stress-test the multi-agent orchestration layers. | Phase 4 |
| **`agentsight`** | Specialized eBPF observability explicitly designed for monitoring LLM agents. | Core dependency for tracking API calls and system usage of isolated AI workers. | Phase 1 |

---

## Layer 2: Cyber Ops (Red Team · Blue Team · Forensics)
**Location:** `Phoenix.Cognition/External/Layer_2_Cyber_Ops/`
**Goal:** Secure the agentic environment, enforce guardrails, and provide autonomous security testing capabilities.

| Repository | Why (Strategic Purpose) | How (Integration Method) | When |
| :--- | :--- | :--- | :--- |
| **`ERA`** | MicroVM-based sandboxing. Essential for running untrusted, AI-generated code securely. | Replace standard Docker containers with `krunvm` microVMs for the execution engine. | Phase 2 |
| **`aegis`** | Dedicated EDR for AI agents. Monitors processes, files, and network behavior locally. | Install as a daemon on the host/worker nodes. Feed its trust scores into the orchestrator. | Phase 3 |
| **`agent-governance-toolkit`** | Microsoft's framework for deterministic, sub-millisecond policy enforcement against AI risks. | Wire into the `Proxy/Gateway` layer to intercept and validate all agent API requests. | Phase 2 |
| **`PentestGPT`** | Automated penetration testing agent. | Deploy as a specialized sub-agent within the `PhoenixRedteam` module for continuous security validation. | Phase 3 |
| **`PentestGPT-MCP`** | Enables PentestGPT to directly trigger tools via the Model Context Protocol. | Configure MCP servers on target test environments to allow safe, controlled exploitation. | Phase 3 |
| **`AutoGPT`** | Highly autonomous agent, useful for reconnaissance and fuzzing loops. | Run in a heavily sandboxed `ERA` environment for unbounded exploration tasks. | Phase 3 |
| **`falcoclaw`** | Automated remediation tool that acts on alerts generated by Falco. | Set up webhook integrations from Falco to FalcoClaw to automatically kill rogue agents. | Phase 3 |
| **`cai`** | Cybersecurity AI for red teaming and vulnerability discovery. | Integrate its scanning outputs into the PhoenixOS threat intelligence dashboard. | Phase 3 |
| **`yawning-titan`** | Graph-based autonomous cyber ops simulation. | Use for training defensive AI models against simulated network attacks. | Phase 4 |
| **`fasterpc`** | Fast, bidirectional JSON-RPC over WebSockets. | Use as the primary communication protocol between distinct, distributed AI agents. | Phase 1 |
| **`nats`** | Lightweight pub/sub messaging. | Implement as the central event bus for the entire PhoenixOS event-driven architecture. | Phase 1 |
| **`agentscope`** | Alibaba's meta-agent architecture framework. | Evaluate alongside `crewai` for specific, highly distributed multi-agent workloads. | Phase 2 |
| **`agentdb`** | Single-file cognitive container (learning state, vectors, audit trail). | Use for lightweight, portable agent memory in isolated environments. | Phase 2 |
| **`sqlite-vec`** | Pure C vector search SQLite extension. | Use for fast, local, dependency-free vector storage for edge agents. | Phase 2 |
| **`awesome-cybersecurity-agentic-ai`** | Curated list of agentic security tools. | Reference index; not executed. | N/A |

---

## Layer 3: Code Intelligence (AST · Review · Sandbox)
**Location:** `Phoenix.Cognition/External/Layer_3_Code_Intelligence/`
**Goal:** Empower the OS with deep understanding of code structure, autonomous editing capabilities, and secure execution environments.

| Repository | Why (Strategic Purpose) | How (Integration Method) | When |
| :--- | :--- | :--- | :--- |
| **`GitNexus`** | Zero-server code intelligence engine. Creates a knowledge graph of the codebase via AST. | Run as a background daemon exposing MCP endpoints to give agents (like Claude) structural awareness. | Phase 2 |
| **`OpenHands`** | Full AI coding agent with AST-aware editing and sandboxed execution. | Integrate as the primary "Software Engineer" persona within the orchestrator. | Phase 3 |
| **`swe-agent`** | specialized agent designed to autonomously navigate and patch repositories. | Deploy for automated bug fixing pipelines triggered by CI/CD failures. | Phase 3 |
| **`SWE-AF`** | Autonomous software engineering fleet (PM, Architect, Coder, Tester). | Use for complex, multi-issue epic orchestration that requires cross-repo coordination. | Phase 3 |
| **`aider`** | Hyper-focused CLI pair programming agent. Excels at surgical AST modifications. | Integrate into the developer's local terminal for fast, iterative code generation. | Phase 2 |
| **`opensandbox`** | Secure, multi-language sandbox runtime for AI agents. | Use as the default execution environment for `OpenHands` and `swe-agent` tasks. | Phase 2 |
| **`pr-agent`** | Automated PR review, security scanning, and complexity scoring. | Hook into the local or remote Git repository to gate all AI-generated code. | Phase 3 |
| **`armada`** | Kubernetes-native distributed job scheduler for high-throughput batch workloads. | Use to schedule and manage thousands of parallel AI inference or compilation jobs. | Phase 4 |
| **`skypilot`** | GPU-aware distributed training/inference orchestrator. | Deploy to manage cost-effective routing of LLM queries across different cloud providers. | Phase 4 |
| **`vllm`** | High-throughput LLM inference engine (PagedAttention). | Set up as the primary local inference server for serving Llama/Mistral models to agents. | Phase 1 |
| **`text-generation-inference`** | Optimized LLM serving (continuous batching, tensor parallelism). | Evaluate against `vllm` for specific model architectures requiring specific optimizations. | Phase 1 |
| **`swe-rl`** | Trains LLM agents via RL to inject and repair software bugs. | Use to continuously generate synthetic training data to improve the coding agents' accuracy. | Phase 4 |
| **`Self-Healing-SRE-Agent`** | Multi-agent orchestrator for detecting and fixing production issues. | Wire into `Tetragon`/`Pixie` alerts to automatically respond to infrastructure failures. | Phase 3 |

---

## Layer 4: AI/ML Engineering (Training · Inference)
**Location:** `Phoenix.Cognition/External/Layer_4_AI_ML_Engineering/`
**Goal:** Manage the lifecycle, optimization, fine-tuning, and memory of the Large Language Models.

| Repository | Why (Strategic Purpose) | How (Integration Method) | When |
| :--- | :--- | :--- | :--- |
| **`ChatDev`** | Multi-agent simulation (virtual software company). | Use as a sandbox for prototyping complex interaction protocols before deploying to production agents. | Phase 3 |
| **`ray`** | Distributed computing framework. | The foundational layer for scaling RLHF and distributed training tasks. | Phase 4 |
| **`DeepSpeed`** | Microsoft's model optimization library (ZeRO). | Apply to fine-tuning pipelines to allow training of large models on consumer GPUs. | Phase 4 |
| **`unsloth`** | Fast LoRA fine-tuning for Llama/Mistral models. | Use as the primary engine for continuous fine-tuning of `PhoenixMind` on the Failure Library. | Phase 4 |
| **`OpenRLHF`** | Production-ready RLHF framework built on Ray. | Use to align custom agent models with specific PhoenixOS operational guidelines. | Phase 4 |
| **`bitsandbytes`** | 4-bit/8-bit quantization libraries. | Essential dependency for `unsloth` and local inference to reduce VRAM requirements. | Phase 1 |
| **`ollama`** | Local LLM serving and management. | Manage the local model registry and serve lightweight models for edge agents. | Phase 1 |
| **`llama.cpp`** | GGUF quantization and CPU/GPU inference. | Fallback inference engine for environments without dedicated high-end GPUs. | Phase 1 |
| **`llama_index`** | RAG pipelines and vector store integration. | Build the contextual retrieval system connecting agents to project documentation (`GEMINI.md`). | Phase 2 |
| **`mem0`** | Persistent agent memory across sessions. | Implement as the short-to-medium term episodic memory store for active agents. | Phase 2 |
| **`train-llm-from-scratch`** | Educational foundation for LLM architecture and pre-training. | Ingest as core knowledge (via RAG) to enable PhoenixOS agents to understand and debug their own underlying transformer architectures and training data pipelines. | Phase 4 |

---

## Layer 5: Game/Simulation (RL · Physics)
**Location:** `Phoenix.Cognition/External/Layer_5_Game_Simulation/`
**Goal:** Provide embodied environments for Reinforcement Learning and complex state-space simulation.

| Repository | Why (Strategic Purpose) | How (Integration Method) | When |
| :--- | :--- | :--- | :--- |
| **`IsaacLab`** | NVIDIA's embodied AI RL framework. | Use for high-fidelity physics simulation when training robotic or spatial awareness agents. | Phase 4 |
| **`mujoco`** | Advanced physics engine. | Alternative to IsaacLab for specific continuous control reinforcement learning tasks. | Phase 4 |
| **`ml-agents`** | Unity game engine ML integration. | Use to generate complex synthetic data environments for visual/spatial agents. | Phase 4 |
| **`RLinf`** | Production-grade RL framework for embodied AI. | The bridge between the underlying simulation engines and the higher-level RL algorithms. | Phase 4 |
| **`Gymnasium`** | Standard RL environment API. | Define all simulation interfaces using this standard for cross-compatibility. | Phase 4 |
| **`stable-baselines3`** | Reliable implementations of RL algorithms (PPO, SAC). | The core algorithmic engine used to train agents within `Gymnasium` environments. | Phase 4 |
| **`FinRL`** | RL for financial simulation. | Adapt its continuous action-reward structures for optimizing cloud resource spending. | Phase 4 |
| **`godot`** | Open-source game engine. | The target "Engine Bridge" for visualizing agent logic or creating bespoke simulation UI. | Phase 4 |

---

## Layer 6: Self-Evolution (Memory · Orchestration)
**Location:** `Phoenix.Cognition/External/Layer_6_Self_Evolution/`
**Goal:** Manage state, multi-actor coordination, and the continuous improvement of the system.

| Repository | Why (Strategic Purpose) | How (Integration Method) | When |
| :--- | :--- | :--- | :--- |
| **`langgraph`** | Stateful graph-based agent workflows. | The Master Brain. Define the core cyclic loops (Plan -> Execute -> Reflect) for all top-level processes. | Phase 1 |
| **`autogen`** | Event-driven multi-agent framework. | Use for conversational team orchestration (e.g., setting up a debate between a security agent and a coder). | Phase 1 |
| **`crewai`** | Role-playing multi-agent orchestration. | Use for highly structured, sequential pipelines where agents pass tasks down an assembly line. | Phase 1 |
| **`letta`** | Long-term memory management across sessions. | The "State Scribe". Implement to manage infinite-context illusions and consolidate historical interactions. | Phase 2 |
| **`qdrant`** | High-scale, Rust-based vector database. | The primary, centralized semantic search engine for the entire OS knowledge graph. | Phase 2 |
| **`weaviate`** | Vector database + knowledge graph capabilities. | Evaluate against Qdrant if hybrid search (GraphQL + Vectors) becomes a critical requirement. | Phase 2 |
| **`chroma`** | Lightweight embedding store. | Use for ephemeral, agent-specific vector storage during isolated tasks. | Phase 2 |
| **`auto-code-rover`** | Autonomous codebase navigation and patch generation. | Integrate its navigation heuristics into `SWE-AF` for better spatial awareness in large repos. | Phase 3 |
| **`SWE-bench`** | Benchmark for self-patching agents. | Use as the primary evaluation metric to score the performance of the local coding agents. | Phase 3 |

---

## Layer 7: Network & Distributed Systems
**Location:** `Phoenix.Cognition/External/Layer_7_Network_Distributed/`
**Goal:** Ensure fast, reliable, and standardized communication across the entire distributed architecture.

| Repository | Why (Strategic Purpose) | How (Integration Method) | When |
| :--- | :--- | :--- | :--- |
| **`grpc`** | High-performance RPC framework. | The backbone protocol for all inter-layer communication (e.g., Code Intelligence calling Vector DB). | Phase 1 |
| **`zilla`** | Stateless, multi-protocol proxy. | Act as the universal Protocol Translation Layer (TCP/MQTT ↔ gRPC/REST) for edge agent ingress. | Phase 1 |
| **`envoy`** | Cloud-native L7 proxy and service mesh. | Deploy as the ingress gateway and internal service mesh to manage rate limiting and circuit breaking. | Phase 1 |
| **`versionize`** | Protocol/API translation on-the-fly. | Automate versioning and changelog generation across the various micro-repositories in PhoenixOS. | Phase 2 |
| **`flink`** | Distributed stream processing. | Handle real-time data pipelines (e.g., aggregating millions of eBPF telemetry events per second). | Phase 4 |

---

## Master Directories (Reference)
**Location:** `Phoenix.Cognition/External/Master_Directories/`
**Goal:** Maintain curated lists of ecosystem developments. *These are not executed, but tracked for future integration.*

*   `awesome-agent-orchestrators`
*   `awesome-agents`
*   `awesome-ai-agents-2026`
*   `awesome-ai-security`
*   `re-list` (Reverse engineering tools)
# WARDEN.EXE: Priority Audit Resolutions

### Q501. Canonical time model?
Hybrid. Ledger Sequence Number provides the logical ordering. Simulation Ticks provide the deterministic clock for engine state updates. Wall clock is forbidden inside the engines.

### Q511. Deterministic replay model?
Bit-for-bit reconstruction of state. RNG is seeded by the previous block hash. Fixed-point arithmetic ensures consistency. Divergence is detected via hash mismatch of the final projection state.

### Q521. Projection authority model?
The Ledger is the Authority. Projections are caches. Any projection can be discarded and reconstructed from the Ledger.

### Q551. Evidence lifecycle model?
Evidence -> Court Adjudication -> Verdict -> Fact Projection. Evidence is an event payload. A Fact is a derived projection that updates when a Verdict event supersedes previous ones.

### Q571. PhoenixMind role?
Constitutional Authority (L3 Service). It is a projections processor that observes the Ledger and emits directive events. It is bounded by the same Invariants as any other entity and can be audited or replaced if its projections consistently cause entropy spikes.
