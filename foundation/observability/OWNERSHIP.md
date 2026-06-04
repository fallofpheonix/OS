# PhoenixTrace — Repository Ownership

## Owner
**Primary:** Phoenix.Nucleus Team
**Contact:** [team@phoenixos.dev]

## Purpose
Causal DAG lineage, trace tracking, evidence linking.

## Dependencies
- **Internal:** PhoenixCore (contracts, event bus)
- **External:** Go standard library

## Consumed By
- PhoenixMind (lineage for AI reasoning)
- PhoenixTruth (evidence linking)
- PhoenixValidation (replay verification)
- PhoenixOS (top-level orchestration)

## Invariants
- Time ordering mandatory
- Causal ordering mandatory
- Graph reconstruction must be deterministic
- Evidence linkage must be explicit

## Criticality: P0
This is the forensic backbone. If PhoenixTrace fails, the system cannot reconstruct state.

## Status: ACTIVE
