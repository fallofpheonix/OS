---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Events — Component Map

> Last verified: 2026-06-04

The `foundation/events` subproject provides concrete implementations of the event schemas, envelopes, and serialization formats specified in the contracts layer.

## Component Breakdown

```
foundation/events/
├── go.mod                     # Module configuration
├── schema.go                  # Core structs (Event, ArtifactManifest, Checkpoint)
└── docs/                      # Subproject documentation
```

### Component Details

1. **`Event`**:
   - Represents the atomic unit of state transition and audit trailing in Phoenix OS.
   - Encapsulates event ID, parent event ID, authority signatures, logical times, and JSON raw payload data.
   - Implements the `eventsv1.EventEnvelope` interface.

2. **`ArtifactManifest`**:
   - Provides immutable structure for binary artifacts, verifying hashes, versions, dependencies, and retaining signature properties.

3. **`Checkpoint`**:
   - Stores consolidated state hashes and replay sequence offsets to speed up ledger loading and consensus initialization.
