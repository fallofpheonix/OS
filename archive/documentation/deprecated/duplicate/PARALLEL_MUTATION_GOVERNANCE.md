# PARALLEL_MUTATION_GOVERNANCE.md

## Overview
Astraeus enables parallel task execution, but repository mutations cannot happen simultaneously without risking file corruption, git state lockouts, or broken AST structures. This document defines the governance rules for parallel mutations.

## 1. Mutation Locks (Repository Level)
- **Global Write Lock**: Any operation that modifies the shared repository root (e.g., `TransactionRunner.apply` or `RollbackEngine.rollback_run`) must acquire an exclusive Global Write Lock.
- **Read-Write Segregation**: Tasks performing analysis (Read) can execute concurrently. Tasks applying mutations (Write) are serialized.

## 2. Path Ownership & Topology-Aware Locking
- **Path-level Locks (Future Enhancement)**: Instead of locking the entire repository, transactions will eventually request locks for specific file paths or directory subtrees.
- **Topology Awareness**: If Task A mutates `api/routes.py` and Task B mutates `api/models.py`, they can proceed in parallel *if and only if* the semantic topology proves they do not have overlapping structural dependencies that would break compilation.
- **Current Fallback**: Until granular path locking is verified, Astraeus will enforce a **Strict Serialization Policy** for all mutations. All transactions queue up and apply one at a time.

## 3. Branch-Aware Isolation
- When executing complex refactors, workers may operate on isolated Git branches (`feature/run_id_task_id`).
- Merging these branches back into the main run branch requires standard git conflict resolution, handled deterministically by the coordinator.

## 4. Rollback Conflict Prevention
- A rollback is the highest priority write operation.
- **Preemption**: If a rollback is initiated, all pending mutations in the queue are paused.
- **Pre-flight Checks**: As established in the Mutation Safety Hardening phase, rollbacks verify hash consistency to ensure they don't overwrite parallel, out-of-band mutations.

## 5. Implementation
1. Create a `LockManager` (using `fcntl` on a dedicated `.astraeus.lock` file in the repo root).
2. Wrap `TransactionRunner.apply` and `RollbackEngine` calls in this lock.
3. Ensure lock acquisition timeouts and failure recovery.
