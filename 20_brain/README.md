# SentinelOS Brain — Design and Repository Layout

Purpose
- Define the LLM "Brain" that coordinates SentinelOS: observation, memory, reasoning, planning, execution, and learning.
- Provide a practical repo layout for implementation and incremental prototyping.

Overview
The Brain sits above the telemetry, evidence, graph, math, physics, game, cyber, cloud and distribution layers and acts as a multi-agent coordination and decision-making layer. It is intentionally gated: decisions that affect kernel or production policy require policy constraints, simulation-first checks and human approval flows.

High-level flow
Users / Nodes → Telemetry → Evidence → Graph → Math → Physics → Game → Cyber → Cloud/Distributed → Brain → Control → Execution (Kernel/Hybrid)

Brain lifecycle
- Observe → Understand → Reason → Predict → Plan → Execute → Learn

Core modules (20_brain/)
- `observation/` — adapters, ingest transforms, schema normalization
- `memory/` — session cache, replay store, vector store connectors, graph DB connectors
- `rag/` — retrieval-augmented components, indexing, chunking
- `reasoning/` — LLM wrappers, symbolic rules, graph reasoning APIs
- `planning/` — plan synthesizers, step generators, policy composer
- `execution/` — safe action runner, approvals, shadow execution, rollbacks
- `learning/` — offline training, federated update orchestrator, drift detectors
- `agents/` — SOC, DFIR, Cloud, Graph, Physics, Game, Recovery, Research agents
- `models/` — curated model artifacts (small/medium/large) and deploy configs
- `knowledge/` — policies, replay, evidence, research artifacts

Memory architecture
- Realtime cache: low-latency in-memory store for current session state
- Session memory: short-lived context for ongoing incidents
- Replay memory: deterministic replay artifacts and checkpoints
- Evidence memory: signed, immutable evidence records
- Graph memory: graph DB snapshot and indexes
- Long-term knowledge: vector DB + policy DB + research archive

Model tiers and roles
- Fast (1B–8B): routing, classification, short reasoning, SOC realtime
- Medium (8B–30B): incident analysis, plan generation, graph reasoning
- Deep (30B+): research, simulation, global policy synthesis

Agents and responsibilities
- Chief Agent: policy, risk budget, approvals
- SOC Agent: alert triage, summarization, MITRE mapping
- DFIR Agent: evidence stitching, timeline, export
- Cloud Agent: scheduler advice, placement, migration plans
- Game Agent: defense allocation, honeypot placement
- Physics Agent: entropy/propagation analysis
- Graph Agent: attack chain reasoning and query execution
- Recovery Agent: rollback planning and choreography
- Research Agent: experiment orchestration, model training

Safety and gating
- All effectful actions must check: policy constraints, safety levels, simulation pass, and optionally require human approval.
- Modes: `shadow` (dry-run), `canary` (limited small group), `enforce` (full action)
- Rollback: all actions create a reversible checkpoint and decision log.

Integration points
- Telemetry adapters: `phoenix_os/monitor` → `20_brain/observation` connectors
- Evidence linkage: `phoenix_os/ledger` → `20_brain/memory` evidence ingestion
- Graph queries: `phoenix_os/trace` → `20_brain/reasoning` traversal API
- Execution: `20_brain/execution` → `07_security/control` & `phoenix_os/warden` → kernel helpers

Initial implementation plan (phased)
1. Prototype observation adapters, realtime cache, and a tiny LLM wrapper (fast model) for classification and alert summarization.
2. Add RAG index + graph connector for incident context enrichment.
3. Implement SOC Agent with shadow-mode actions (no enforcement) and a human approval flow.
4. Implement DFIR Agent with replay-execution and forensic export.
5. Hardening: policy DB, model registry, drift detectors, federated update.
6. Gate and scale with canary and feature flags before enabling active enforcement.

Repository layout (suggested)
- `20_brain/`
  - `observation/`
  - `memory/`
  - `rag/`
  - `reasoning/`
  - `planning/`
  - `execution/`
  - `learning/`
  - `agents/`
  - `models/`
  - `knowledge/`

Next steps
- Create a minimal prototype: ingest a telemetry sample, enrich with RAG, generate an incident summary, and propose a plan in `shadow` mode.
- Optionally, create a small CI job to run the prototype using a tiny open LLM (or local LLM runtime) for reproducible testing.

Notes
- Keep the brain modular and replaceable: LLM weights, vector DB, and policy engine must be pluggable.
- Document every decision in `20_brain/DECISION_LOG.md` and keep an audit trail for changes.
