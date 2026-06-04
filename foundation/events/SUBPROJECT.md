# Subproject Manifest: Phoenix Events

## Purpose
Governs the core event schemas, checkpoint structures, and metadata format declarations for the event log replication system.

## Owners
* Primary: Storage & Event Bus Team
* Secondary: Core Runtime Devs

## Dependencies
* None

## Consumers
* `PhoenixCore` runtime, game, and recovery
* `PhoenixValidation`
* `PhoenixGuard`

## Build Command
```bash
go build ./core/Phoenix.Nucleus/PhoenixCore/event/...
```

## Test Command
```bash
go test -v ./core/Phoenix.Nucleus/PhoenixCore/event/...
```

## Success Criteria
* Clean independent build.
* Seamless serialization compatibility with v1 events contracts.

## Contract Surface
* `github.com/fallofpheonix/PhoenixCore/event` (`Event`, `ArtifactManifest`, `Checkpoint`)
