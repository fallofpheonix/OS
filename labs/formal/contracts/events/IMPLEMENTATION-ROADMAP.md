---\nStatus: Planned\nImplementation: 5%\nConfidence: Conceptual\n---\n# 10 Concrete Engineering TODOs
# Observable Causality Implementation Plan
# DATED: 13 May 2026
# AUTHORITY: Binding constraint—stress without causality = unobservable destruction

---

## Phase 1: Foundational Event Infrastructure (TODOs 1-4)
### Goal: Build the nervous system substrate
### Duration: ~1 week
### Success: Event causality survives a full runtime restart

---

## TODO 1: Build Canonical Event Emitter Library

**Location**: `modules/core/event-core/`

**Responsibility**: Single, correct way to emit events that satisfies schema invariants.

**Success Criteria**:
- [ ] Event dataclass maps to every field in runtime-event.schema.yaml
- [ ] emit_event(Event) validates against schema before returning
- [ ] UUID generation is cryptographically sound and deduplicable
- [ ] trace_id is thread-local and auto-propagates to children
- [ ] Deduplication guard: identical event_id returns early
- [ ] All validation errors are machine-parseable

**Expected Output**:
```
modules/core/event-core/
├── __init__.py
├── event.py              # Event dataclass
├── emitter.py            # Emitter class
├── schema_validator.py    # YAML validation
├── trace_context.py       # Thread-local trace management
└── tests/
```

---

## TODO 2: Implement Supervisor Lifecycle Emitters

**Location**: `control-plane/runtime/supervisor.py`

**Events to Emit**:
- startup (before services start)
- ready (initialized, accepting work)
- restart (restart sequence initiated)
- degraded (system degradation detected)
- shutdown (graceful termination)
- fatal (unrecoverable failure)

**Success Criteria**:
- [ ] Supervisor emits 6 lifecycle events
- [ ] runtime_id is consistent across all supervisor events
- [ ] runtime_id persists in checkpoint file
- [ ] restart event references previous runtime_id
- [ ] restart_count increments correctly
- [ ] Events persisted to runtime-state/supervisor/events/

**Why This Matters**:
- Establishes runtime identity
- Enables restart causality tracking
- Provides foundation for TODO 3-4

---

## TODO 3: Implement Runtime-State Ingestion Layer

**Location**: `control-plane/runtime-state/ingestion/`

**Responsibility**: Durably persist events in append-only format.

**Output Structure**:
```
runtime-state/
├── supervisor/events/YYYY/MM/DD/*.jsonl
├── queue/events/YYYY/MM/DD/*.jsonl
├── sink/events/YYYY/MM/DD/*.jsonl
└── .state_version
```

**Success Criteria**:
- [ ] Events persisted to runtime-state/{service}/events/
- [ ] Append-only: no event modification/deletion
- [ ] Deduplication by event_id
- [ ] Timestamp monotonicity validation
- [ ] Index files with checksum
- [ ] Rotation at 1GB or 24h

**Why This Matters**:
- Events become durable (survive restarts)
- Foundation for forensics and recovery
- Enables state derivation from historical data

---

## TODO 4: Build Trace Propagation Rules

**Location**: `modules/core/event-core/trace_context.py`

**Critical Rules**:
1. One incident = One trace tree
2. Restart preserves trace
3. Scenario isolation creates new trace
4. Child events reference parent correctly

**Success Criteria**:
- [ ] trace_id auto-threaded parent→child
- [ ] parent_event always valid reference
- [ ] Trace DAG is acyclic
- [ ] New traces at explicit boundaries only
- [ ] Child events in threads don't interfere
- [ ] Trace context inherited across async boundaries

**Why This Matters**:
- Enables complete incident reconstruction
- Supports multi-threaded runtime
- Forms foundation for forensics

---

### Phase 1 Checkpoint: When complete
- Supervisor lifecycle is observable
- Event causality is enforceable
- Events are durable
- Full incident traces reconstructible

**Expected Duration**: 1 week

---

## Phase 2: Component-Specific Emission (TODOs 5-6)
### Goal: Make queue and sink behavior observable
### Duration: ~3-4 days
### Success: All P0 emitters emit continuously

---

## TODO 5: Implement Queue Saturation Emitters

**Location**: `modules/core/queue/worker.py`, `backpressure.py`

**Events**:
- queue.depth.warning (70% threshold)
- queue.depth.critical (90% threshold)
- queue.dropped (backpressure)
- backpressure.applied / released
- queue.recovered

**Success Criteria**:
- [ ] Events at correct thresholds
- [ ] Dropped items tracked
- [ ] queue_depth accurate at emission
- [ ] Causally linked
- [ ] < 1% performance overhead

---

## TODO 6: Implement Sink Isolation Events

**Location**: `modules/core/sink/processor.py`

**Events**:
- sink.timeout (transient failure)
- sink.unreachable (persistent network error)
- sink.quarantined (isolation decision)
- sink.recovered (back to healthy)

**Success Criteria**:
- [ ] Timeouts with latency_ms
- [ ] Quarantine at threshold
- [ ] Recovery events captured
- [ ] All causally linked

---

### Phase 2 Checkpoint: When complete
- Queue saturation observable
- Sink failures attributed
- Backpressure measurable
- 1-hour runtime shows expected patterns

**Expected Duration**: 3-4 days

---

## Phase 3: First Stress Scenario (TODO 7)
### Goal: Prove causality survives destruction
### Duration: ~2-3 days
### Success: All events show under injection

---

## TODO 7: Create First Executable Stress Scenario

**Location**: `runtime-lab/scenarios/flood-basic.yaml`, `executor.py`

**Scenario**: Flood queue with 1M events, bounded capacity

**Expectations**:
- queue.depth.warning (multiple)
- queue.depth.critical (at least 1)
- queue.dropped (>10,000)
- sink.timeout (>1000)
- health-engine.degraded (< 5s after critical)
- backpressure.released (recovery)

**Success Criteria**:
- [ ] Scenario executes without hanging
- [ ] All expected events emitted
- [ ] Complete causality chain
- [ ] System detects degradation < 5s
- [ ] Recovery < 60s
- [ ] No contradictory transitions

**Output**:
```yaml
runtime-lab/results/flood-basic-2026-05-13.yaml
- scenario execution status
- events collected
- causality validation
- findings & next actions
```

---

### Phase 3 Checkpoint: When complete
- Stress scenario passes
- Causality proved under destruction
- System behavior measurable
- Incident traceable

**Expected Duration**: 2-3 days

---

## Phase 4: State Derivation & Automation (TODOs 8-10)
### Goal: System becomes self-describing
### Duration: ~1 week
### Success: Registry, state, incidents auto-generated

---

## TODO 8: Implement Runtime-State Derivation Engine

**Location**: `control-plane/runtime-state/derivers/`

**Derived State Types**:
- health-state (from lifecycle events)
- incident-state (from error events)
- queue-metrics (aggregated from queue events)
- availability (computed from restarts)

**Success Criteria**:
- [ ] Health state computed correctly
- [ ] Incidents auto-created on errors
- [ ] Queue metrics aggregated
- [ ] Availability % calculated
- [ ] Events never mutated
- [ ] Updated every 10 seconds

---

## TODO 9: Build Registry Compiler Prototype

**Location**: `control-plane/sync-engine/compiler/`

**Generates**: `control-plane/repo-registry.yaml` (auto-updated every 60s)

**Includes**:
- Service discovery
- Emitter identification
- Dependency extraction
- Active incidents

**Success Criteria**:
- [ ] All services discovered
- [ ] All emitters identified
- [ ] Dependencies extracted
- [ ] Status reflects reality
- [ ] Auto-synced continuously

---

## TODO 10: Build Runtime Incident Recorder

**Location**: `runtime-state/incidents/`

**Auto-records**:
- Incident creation on error events
- Degradation path
- Recovery path
- Forensic-quality YAML

**Incident Fields**:
- trace_id (= incident_id)
- timeline (start, detection, recovery, end)
- affected_components
- root_cause
- metrics (drops, timeouts, etc.)
- resolution status

**Success Criteria**:
- [ ] Incidents auto-created on errors
- [ ] Paths captured completely
- [ ] Auto-closed on recovery
- [ ] Severity classified P1/P2/P3
- [ ] YAML is forensically complete

---

### Phase 4 Checkpoint: When complete
- System is self-describing
- Registry auto-synced
- Incidents auto-recorded
- State automatically derived
- No architectural rework needed

**Expected Duration**: ~1 week

---

## CRITICAL EXECUTION RULE

**DO NOT parallelize across phases.**

Correct grouping:
- Phase 1 (TODOs 1-4): ~1 week
- Phase 2 (TODOs 5-6): ~3-4 days (after Phase 1)
- Phase 3 (TODO 7): ~2-3 days (after Phase 2)
- Phase 4 (TODOs 8-10): ~1 week (after Phase 3)

**Reason**: Each phase unblocks the next. Starting Phase 2 before Phase 1 is complete means:
- No emission library → emitters fail
- No event persistence → no data for derivers
- No causality → stress tests are noise

---

## Success Criteria (All 10 Complete)

- [ ] Prove causality survives runtime degradation
- [ ] Event stream is authoritative source of truth
- [ ] System state is derived (not stored separately)
- [ ] Complete forensics from event records
- [ ] Stress scenarios observable and measurable
- [ ] No rework needed for next phase

---

## Most Important Current Priority

> The most important thing now is:
> 
> **Proving event causality survives runtime degradation**
>
> NOT: dashboards, AI agents, distributed runtimes, control-plane expansion.
>
> Everything depends on whether the event substrate remains coherent under failure.

---

## Reference Architecture

After all 10 TODOs:

```
Event Emission (TODO 1-6)
    ↓
Event Ingestion (TODO 3)
    ↓
runtime-state/{service}/events/ (POPULATED)
    ↓
State Derivation (TODO 8)
    ↓
runtime-state/derived/ (health, incidents, metrics, availability)
    ↓
Registry Compiler (TODO 9)
    ↓
control-plane/repo-registry.yaml (AUTO-GENERATED)

+ Incident Recorder (TODO 10)
    ↓
runtime-state/incidents/ (FORENSIC RECORDS)
```

This forms a complete observability substrate.
