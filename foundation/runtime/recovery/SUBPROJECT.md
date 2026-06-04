# Subproject Manifest: Phoenix Recovery

## Purpose
Implements the recovery protocol (P5.1/P5.2), permitting safe node state restoration and resurrection from historical checkpoint and ledger streams.

## Owners
* Primary: Core Runtime Devs
* Secondary: Ledger & Audit Team

## Dependencies
* `github.com/fallofpheonix/PhoenixCore/contracts`
* `github.com/fallofpheonix/PhoenixCore/event`
* `github.com/fallofpheonix/PhoenixCore/constitution`

## Consumers
* `PhoenixOS` Daemon (`phoenixd`)
* `PhoenixValidation` proofs suite

## Build Command
```bash
go build ./core/Phoenix.Nucleus/PhoenixCore/recovery/...
```

## Test Command
```bash
go test -v ./core/Phoenix.Nucleus/PhoenixValidation/proofs -run TestRecoveryProof
```

## Success Criteria
* State restoration reproduces identical state hashes matching checkpoints.
* Independent compilation and test verification.

## Contract Surface
* `github.com/fallofpheonix/PhoenixCore/recovery` (`Engine`, `NewEngine`, `Recover`)
