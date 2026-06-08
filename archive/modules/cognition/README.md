# Phoenix.Cognition — Autonomous Intelligence Layer

## Primary Responsibility
Phoenix.Cognition orchestrates the high-level cognitive cycles (L5.5–L7) of the PhoenixOS ecosystem. It is responsible for translating raw system telemetry into deterministic causal lineages, maintainable belief states, and auditable strategic actuations.

## System Architecture
The repository implements a 4-package cognitive substrate unified by a central AI Orchestrator:
1. **Knowledge (Causal/Semantic):** Manages the Causal DAG and Belief Engine.
2. **Memory (Episodic/Persistence):** Implements tiered storage with SQLite-Vec backing.
3. **Reasoning (Inference/Logic):** Provides LLM-agnostic inference bridging and the Explanation Layer.
4. **Reflection (Audit/Safety):** Monitors reality drift and enforces epistemic quarantine.

## Tech Stack
- **Language:** Go 1.26
- **Persistence:** SQLite 3 (with sqlite-vec support)
- **Messaging:** internal/bus (L4 nucleus dependency)
- **Protocol:** gRPC / REST (Bridge targets)

## System Boundaries (AI Context)
- **Northbound:** Consumes gRPC directives from the Distributed Nexus (L7).
- **Southbound:** Issues ActuationCertificates to the Warden FSM (L5) in Phoenix.Nucleus.
- **East/West:** Synchronizes state across peer nodes via the Consensus Ledger.

## Data Flow
Telemetry Event -> Monitor (Drift) -> Trace (Lineage) -> Arbiter (Policy) -> Oracle (Reasoning) -> Warden (Actuation).

## Setup Prerequisites
- Go 1.26+
- SQLite 3 development headers
- `go.work` correctly pointing to `Phoenix.Nucleus` and `Phoenix.Terminus`.

## External Dependencies
- `github.com/mattn/go-sqlite3`: Database driver.
- `github.com/fallofpheonix/PhoenixCore`: Bus and primitive types.
- `github.com/fallofpheonix/PhoenixGuard`: Warden and Actuation logic.
