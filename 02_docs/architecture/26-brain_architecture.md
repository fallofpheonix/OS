# SentinelOS LLM Brain — Architecture Summary

This document describes the LLM Brain concept, responsibilities, modules, and integration points with SentinelOS. It is a concise reference for architects and implementers.

Purpose
- Provide a cognitive coordination layer that understands telemetry, stitches evidence, reasons over attack graphs, plans defenses, and executes (safely) through the control runtime.

Stack placement
Telemetry → Evidence → Graph → Math → Physics → Game → Cyber → Cloud/Distributed → **LLM Brain** → Control → Kernel/Hybrid

Brain capabilities
- Observation: normalize and enrich telemetry into structured world-state snapshots.
- Memory: maintain session and long-term stores (vector DB, graph DB, replay DB, evidence DB).
- Reasoning: LLM + symbolic rules + graph traversal for explainable outputs.
- Planning: produce stepwise response plans and risk-adjusted actions.
- Execution: safe enactment APIs with human-in-the-loop controls and rollback.
- Learning: closed-loop improvement from replay and SOC feedback.

Multi-agent design
- Chief Agent coordinates policies and approvals.
- Domain agents (SOC, DFIR, Cloud, Game, Physics, Graph, Recovery, Research) implement specialised reasoning and actions.

Model tiers
- Fast: low-latency classification and routing
- Medium: plan and analysis
- Deep: research and simulation

Safety & governance
- Policy DB enforces allowed actions and risk budgets.
- All actions have `shadow`, `canary`, and `enforce` modes.
- Evidence-first: all decisions are recorded, signed, and reproducible via replay.

Implementation milestones
1. Observation adapters + mini LLM classification prototype (shadow-only).
2. RAG + graph connector to support enriched incident analysis.
3. SOC Agent + human approval UI. Integrate with `agents/surface/orchestrator` API.
4. DFIR Agent: replay export, forensics, and evidence stitching.
5. Execution wiring to `07_security/control` with canary rollout.
6. Federated learning & model governance.

Repository suggestion
- `20_brain/` (see README for layout)
- Keep documentation under `02_docs/architecture/26-brain_architecture.md` for cross-references.

Security notes
- LLM outputs must be treated as recommendations; do not allow blind automatic enforcement in Phase 0–3.
- Model explainability and decision logs are required for audit and compliance.

References
- See `02_docs/architecture/13-roadmap.md` for phase alignment and gating rules.
- See `02_docs/architecture/23-sentinel_game_features.md` and `02_docs/architecture/24-cloud-distributed_features.md` for micro-feature catalogs.

Questions for the team
- Which model runtimes are acceptable for the fast/medium/deep tiers (local vs cloud)?
- Preferred vector DB and graph DB stacks for low-latency RAG?
- Human approval workflows and SOC operator UX requirements.
