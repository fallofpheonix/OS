# Project: PhoenixOS Stage 1 - State & Intelligence

**Status:** In Progress
**Priority:** High
**Tags:** #engineering/os #ai/orchestrator

## Description
Pivot the execution logic of PhoenixOS to put the AI Orchestrator as the central main component coordinating all modular features (Warden, Monitor, Guard, Ledger, Trace, TCS).

## Tasks
- [x] Merge/rebase stable branches and define the generic `Feature` interface.
- [x] Refactor [main.go](file:///Users/fallofpheonix/os/phoenix_os/main.go) to register all subsystems as features and delegate ticks to `OrchestrateTick`.
- [x] Verify determinism hashes.
- [ ] Wire up Hermes Agent to read from the Obsidian vault and suggest policy modifications based on local logs.

## Notes & Design Decisions
- In the new design, the Warden FSM remains the deterministic boundary, but the orchestrator directs the chronological tick steps.
- We preserve the SHA-256 Merkle-DAG ledger outputs to guarantee replay equivalence.
