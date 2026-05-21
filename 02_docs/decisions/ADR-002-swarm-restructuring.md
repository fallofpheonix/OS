# ADR-002: Swarm Logic Restructuring

## Status

Proposed.

## Context

The `phoenix_os/agents/internal/swarm` package contained mixed concerns, causing package collisions and potential circular dependencies between the Warden (actuation) and the Arbiter (reasoning).

## Decision

Restructure `phoenix_os/agents/internal/swarm` into sub-packages to encapsulate agent logic and governance.

- Move entry point/agent logic to `phoenix_os/agents/internal/swarm/agent`.
- Move governance logic to `phoenix_os/agents/internal/swarm/governance`.

## Consequences

- Resolves package collisions.
- Eliminates circular dependencies between Warden and Arbiter.
- Improves code maintainability and modularity.
- Aligns with the Phoenix Matrix separation of concerns.
