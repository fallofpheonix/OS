# PhoenixMind Intelligence Layer

## Scope
This document defines the execution logic for the PhoenixOS central orchestrator. It covers the telemetry processing pipeline, policy evaluation, and the collaborative reasoning cycle.

## Primary Responsibility
Orchestrate the 5-layer cognition stack to translate system entropy into deterministic, auditable actuations.

## The 5-Stage Execution Pipeline
Every logical tick in the `AIOrchestrator` follows this sequence:

### 1. Ingress Stage
Synchronizes the formal cognitive state. Ingests raw telemetry into the `Working` memory tier and updates the Causal Knowledge Graph.

### 2. Sensory Stage
Calculates thermodynamic drift (SDI Z-Score). Updates sensor reputation using Exponential Moving Average (EMA) recovery. Penalizes sensors providing high-drift, uncorrelated data.

### 3. Temporal Stage
Evaluates time-series confidence (TCS). Identifies confidence degradation across sliding windows and logs violations to the Ledger.

### 4. Strategic Stage (Policy)
Invokes the `Arbiter` to translate numeric scores into a `TargetState` and `ActuationClass`. This determines if the system is authorized to move from `SAFE` to `COMPROMISED`.

### 5. Cognitive Pipeline (Guarded Autonomy)
Triggers the async reasoning cycle. 
- **Reflection:** Performs a pre-action safety audit. If reality drift exceeds the threshold, the Oracle is blocked.
- **Oracle:** Queries the G0DM0D3 Nexus for a strategic directive.
- **User Gate:** Requests explicit human permission for critical changes (Axiom 3).
- **Actuation:** Executes the directive via the `Warden` and generates a `ReasonPath` for auditing.

## Verification
MTTE (Mean Time To Explain) is tracked at each stage. Determinism leaks trigger a sovereign audit.
