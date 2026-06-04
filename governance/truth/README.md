# PhoenixTruth — Evidence Validation Layer

> Evidence evaluation, truth scoring, confidence assessment, and contradiction detection for the PhoenixOS ecosystem.

## Overview

PhoenixTruth provides the evidence validation backbone for PhoenixOS. It evaluates evidence records, computes truth scores, detects contradictions, and maintains the truth model integrity.

**All evidence must be cryptographically signed and source-ranked.**

## Repository Structure

```
PhoenixTruth/
├── engine/
│   ├── evaluator.go        # Evidence assessment and truth scoring
│   ├── contradiction.go    # Contradiction detection
│   └── engine_test.go      # Engine tests
├── go.mod
└── go.sum
```

## Core Principles

1. **Evidence Weighting**: Evidence is weighted by source authority
2. **Source Ranking**: Trusted sources have higher weight
3. **Confidence Scoring**: Deterministic confidence assessment
4. **Contradiction Tracking**: All contradictions are recorded

## Build & Test

```bash
# Run all tests
go test ./...

# Run with race detector
go test -race ./...
```

## Dependencies

- **Depends on**: PhoenixCore (contracts, protobuf)
- **Depended by**: PhoenixMind (truth feature), PhoenixValidation (truth tests)

## Invariants

- Evidence must be signed
- Contradictions must be tracked
- Confidence scores must be deterministic
- Truth assessments must be reproducible

## License

PhoenixTruth is part of the PhoenixOS ecosystem.
