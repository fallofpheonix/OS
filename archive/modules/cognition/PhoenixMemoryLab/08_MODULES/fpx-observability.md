# Module: fpx-observability

## Source
Extracted from [[05_PROJECTS/COMPLETED/UDIE/Project]]

## Purpose
Event sourcing, CQRS projections, and real-time state materialization for building audit trails, operational dashboards, and derived views from an authoritative event log.

## Interface
```python
from fpx_observability import EventStore, Projection, Materializer

store = EventStore(backend="postgres")
store.append(event)  # Immutable append to event log

projection = Projection(
    name="risk-cells",
    source=store,
    reducer=risk_reducer_fn
)

materializer = Materializer(projections=[projection])
materializer.rebuild()  # Full replay
materializer.apply(new_event)  # Incremental update
```

## Depends On
- PostgreSQL (event store backend)
- Redis (optional hot-path caching)

## Used By
- Banking App (transaction audit trail, account balance projections)
- Network Security Scanner (threat event timeline, alert materialization)

## Extraction Status
NOT_STARTED

## Location
`~/engineering/infrastructure/shared-libraries/fpx-observability/`

## Key Files
| File | Role |
|------|------|
| `event_store.py` | Immutable event log with append-only semantics |
| `projection.py` | Derived state computation from event stream |
| `materializer.py` | Incremental + full-replay materialization engine |
| `workers.py` | Background workers for async projection updates |
| `spatial.py` | H3 spatial indexing utilities (from UDIE) |

## UDIE Architecture Being Extracted
```
Ingestion Substrate → Event Log (authoritative)
                          ↓
                    Projections & Workers
                          ↓
                    Materialized Views (queryable)
                          ↓
                    Operational Interface
```

## Quality Gates
- [ ] Tests passing
- [ ] Event store is backend-agnostic (not PostGIS-coupled)
- [ ] Projection replay is deterministic
- [ ] README with event sourcing tutorial
- [ ] Version pinned

> [!WARNING]
> This is the most complex extraction. UDIE's event sourcing is tightly coupled to spatial indexing (H3 + PostGIS). Extract the generic event sourcing first, keep spatial as an optional add-on.

#module #extracted-from/UDIE #priority/P2
