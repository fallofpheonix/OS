# PhoenixTruth — Evidence Validation Layer

## Agent Skills
### Issue Tracker
GitHub issue tracker. See `docs/agents/issue-tracker.md`.

### Triage Labels
Default triage label vocabulary. See `docs/agents/triage-labels.md`.

### Domain Docs
Multi-context layout. See `docs/agents/domain.md`.

## Build & Test
```bash
go test ./...
go test -race ./...
```

## Architecture
PhoenixTruth evaluates evidence records, computes truth scores, detects contradictions, and maintains truth model integrity.

## Key Components
- **engine/evaluator.go** — Evidence assessment and truth scoring
- **engine/contradiction.go** — Contradiction detection

## Invariants
- Evidence must be signed
- Contradictions must be tracked
- Confidence scores must be deterministic
- Truth assessments must be reproducible
