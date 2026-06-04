# Subproject Manifest: Phoenix Validation

## Purpose
Exposes independent validation pipelines (event replay, clock invariants, fuzz testing) to guarantee historical system correctness and prevent execution drifts.

## Owners
* Primary: Validation & Verification Team
* Secondary: Core Runtime Devs

## Dependencies
* `github.com/fallofpheonix/PhoenixCore/contracts`
* `github.com/fallofpheonix/PhoenixCore/event`

## Consumers
* Continuous audit verifiers and integration suites.

## Build Command
```bash
go build ./core/Phoenix.Nucleus/PhoenixValidation/...
```

## Test Command
```bash
go test -v ./core/Phoenix.Nucleus/PhoenixValidation/...
```

## Success Criteria
* No legacy event or custom internal type imports.
* Clean adapter-based interactions matching contracts.
* 100% test coverage for fuzz, invariants, and proofs suites.

## Contract Surface
* `github.com/fallofpheonix/PhoenixValidation/replay` (`Engine`, `NewEngine`, `CalculateStateHash`, `DivergenceDetector`, `AuthorityVerifier`)
