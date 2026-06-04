# Subproject Manifest: Phoenix Contracts

## Purpose
Owns the public interfaces, compatibility rules, and versioning specifications for the PhoenixOS system, providing a clean separation of concern between implementation and usage.

## Owners
* Primary: Architecture & Coordination Team
* Secondary: Core Runtime Devs

## Dependencies
* None (Strict zero-dependency rule to avoid circular leakage).

## Consumers
* `PhoenixCore`
* `PhoenixValidation`
* `PhoenixGuard`
* `Phoenix.UI`
* `contract-tests`

## Build Command
```bash
go build ./core/Phoenix.Nucleus/PhoenixCore/contracts/...
```

## Test Command
```bash
go test -v ./contract-tests/...
```

## Success Criteria
* 100% of contract-tests pass successfully.
* Clean standalone build with zero implementation imports.

## Contract Surface
* `github.com/fallofpheonix/PhoenixCore/contracts/events/v1` (`Event`, `EventEnvelope`, `Serializer`)
* `github.com/fallofpheonix/PhoenixCore/contracts/replay/v1` (`Snapshot`, `Reconstructor`, `ReplayEngine`)
* `github.com/fallofpheonix/PhoenixCore/contracts/security/v1` (`Containment`, `Escalation`, `Actuator`)
