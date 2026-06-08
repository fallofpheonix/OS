# Knowledge Subsystem

## Primary Responsibility
The Knowledge subsystem maintains the deterministic world model for PhoenixOS. It is responsible for causal mapping (Graph) and high-level semantic conclusions (Beliefs) derived from the immutable Ledger.

## System Architecture
1. **Causal Graph:** A directed acyclic graph (DAG) representing the "Wait-For" and "Caused-By" relationships between system processes and network events.
2. **Belief Engine:** A versioned state machine for high-confidence conclusions. Beliefs are grounded in facts and are subject to invalidation if reality drift is detected.

## Tech Stack
- Go (Standard Library)
- encoding/json (Serialization)

## AI-Specific Context
- **System Boundaries:** Receives `ledger.Event` streams from the Nucleus. Provides state to `PhoenixMind` for prompt grounding.
- **Data Flow:** Ledger Entry -> Graph Node/Edge -> Belief Commitment -> Reasoning Context.
