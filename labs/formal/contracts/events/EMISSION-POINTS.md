---\nStatus: Planned\nImplementation: 5%\nConfidence: Conceptual\n---\n# Emission Point Mapping
# Phase 2a: Identify where runtime causality must originate
# DATED: 13 May 2026
# DEPENDENCY: runtime-event.schema.yaml (must be stable first)

---

## Purpose

This document identifies which runtime components MUST emit events, and which event types they are responsible for. This forms the **causal backbone** of the ecosystem.

Without complete emission point coverage:
- Causality chains will have gaps
- Incident forensics will be incomplete
- Runtime-state derivation will miss critical state transitions
- Stress scenarios cannot validate all failure modes

---

## Emission Point Classification

### Critical Path (MUST emit)
Events that must be emitted for operational causality:
- **Supervisor**: lifecycle transitions (startup, restart, shutdown)
- **Queue**: saturation, backpressure, drain transitions
- **Sink**: timeouts, failures, recovery
- **Health Engine**: degradation detection, containment decisions

### Observability Path (SHOULD emit)
Events that enable deep observability but not strictly critical:
- **Orchestration**: registry mutations, policy violations
- **Contracts**: drift detection, incompatibility alerts
- **Runtime-Lab**: scenario start/stop, injection points

### Optional Path (CAN emit)
Events for future capabilities:
- **Governance**: audit trail, mutation attribution
- **Modules**: capability startup/shutdown

---

## Emission Matrix

| Component | Event Type | Current Status | Emitter Location | Priority |
|-----------|-----------|-----------------|------------------|----------|
| supervisor | startup | ❌ Missing | supervisor/bootstrap.py | P0 |
| supervisor | restart | ❌ Missing | supervisor/restart.py | P0 |
| supervisor | shutdown | ❌ Missing | supervisor/lifecycle.py | P0 |
| supervisor | supervisor-takeover | ❌ Missing | supervisor/failover.py | P1 |
| queue | queue-saturated | ❌ Missing | queue/worker.py | P0 |
| queue | queue-drained | ❌ Missing | queue/worker.py | P0 |
| queue | backpressure-applied | ❌ Missing | queue/backpressure.py | P0 |
| queue | backpressure-released | ❌ Missing | queue/backpressure.py | P0 |
| sink | timeout | ❌ Missing | sink/processor.py | P0 |
| sink | failure | ❌ Missing | sink/processor.py | P0 |
| sink | recovery | ❌ Missing | sink/processor.py | P1 |
| health-engine | health-check | ❌ Missing | health-engine/checker.py | P1 |
| health-engine | degradation | ❌ Missing | health-engine/degradation.py | P0 |
| health-engine | quarantine | ❌ Missing | health-engine/containment.py | P0 |
| control-plane | registry-update | ❌ Missing | control-plane/registry.py | P1 |
| control-plane | policy-violation | ❌ Missing | control-plane/governance.py | P1 |
| contracts | contract-drift | ❌ Missing | contracts/validator.py | P1 |
| runtime-lab | scenario-started | ❌ Missing | runtime-lab/executor.py | P1 |
| runtime-lab | scenario-completed | ❌ Missing | runtime-lab/executor.py | P1 |
| runtime-lab | injection-point-reached | ❌ Missing | runtime-lab/chaos.py | P1 |

---

## Emission Context Requirements

Each emitter MUST provide:

```python
# Required context at emission time
emission_context = {
    "runtime_id": str,          # From supervisor checkpoint
    "service": str,             # From CURRENT_SERVICE env var or config
    "supervisor": str,          # From SUPERVISOR_ID or supervisor heartbeat
    "component": str,           # Service-specific identifier
    "event_type": str,          # From enum in schema
    "timestamp": datetime.utcnow(),
    "trace_id": uuid,           # From trace context or generate if origin
    "parent_event": uuid,       # From current span context or None
}
```

---

## Priority Sequence

### Phase 2a.1: Supervisor Lifecycle (P0)
**Goal**: Establish runtime identity and restart causality

Files to modify:
- `control-plane/runtime/supervisor.py`: Add startup/restart/shutdown emitters
- `control-plane/contracts/supervisor-context.py`: Capture supervisor state for emission

What this enables:
- runtime_id continuity tracking
- restart loop detection
- supervision chain reconstruction

### Phase 2a.2: Queue Saturation & Backpressure (P0)
**Goal**: Make queue state transitions observable

Files to modify:
- `modules/core/queue/worker.py`: Emit saturation/drain events
- `modules/core/queue/backpressure.py`: Emit pressure state transitions

What this enables:
- SLO violation detection
- Cascading failure root-cause
- Resource contention analysis

### Phase 2a.3: Sink Timeout/Failure Handling (P0)
**Goal**: Capture transaction boundaries where failure occurs

Files to modify:
- `modules/core/sink/processor.py`: Emit timeout and failure events
- Add metadata: work_item_id, timeout_threshold_ms, resource_used

What this enables:
- End-to-end trace reconstruction
- Failure classification
- SLA tracking

### Phase 2a.4: Health Engine Degradation Detection (P0)
**Goal**: Emit degradation and isolation decisions

Files to modify:
- `control-plane/health-engine/degradation.py`: Emit degradation events
- `control-plane/health-engine/containment.py`: Emit quarantine decisions

What this enables:
- Predictive failure detection
- Containment policy validation
- Cascade prevention verification

### Phase 2b: Optional Emission Points (P1)
Only after P0 is complete:
- Orchestration registry mutations
- Contract drift detection
- Scenario lifecycle

---

## Validation Checklist

For each emission point, before considering it "complete":

- [ ] Event schema validation passes at emission
- [ ] trace_id is correctly threaded from parent
- [ ] parent_event references valid previous event
- [ ] All required fields are populated
- [ ] Metadata does not exceed 64KB
- [ ] Event is appended to runtime-state/{service}/events (idempotent)
- [ ] Deduplication works correctly for retries
- [ ] Timestamp monotonicity is maintained within service

---

## Testing Strategy

For each emission point:

1. **Unit test**: Event is correctly structured, matches schema
2. **Integration test**: Event is persisted to runtime-state
3. **Causal test**: Parent-child relationships are preserved
4. **Scenario test**: Event appears under synthetic stress scenario

---

## Immediate Next Action

1. Select **supervisor lifecycle** emitters as first implementation target
2. Create minimal event persistence layer to `runtime-state/{service}/events`
3. Validate supervisor bootstrap → startup event → ready transition
4. This becomes the proof that the schema + emission pattern works
5. Then cascade pattern to queue, sink, health-engine

This establishes operational causality from the root.
