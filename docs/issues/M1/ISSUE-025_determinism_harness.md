# ISSUE-025: Determinism Harness

## Problem
Without proof of determinism, the entire replay-centric architecture is speculative. We cannot verify state invariant drift across runs.

## Current State
None.

## Missing
A test harness that asserts: Same Input + Same Config → Identical Output Hash.

## Required Work
- [ ] Implement deterministic replay trigger.
- [ ] Implement runtime hashing of all FSM transitions.
- [ ] Implement post-replay comparison of state hashes.

## Dependencies
Depends on: #ISSUE-001, #ISSUE-004

## Acceptance Criteria
[ ] Identical runs produce identical deterministic output hash.
[ ] Failure to match hash terminates replay immediately.

## Risk
High (P0 - Architecture Blocker)

## Labels
critical, determinism, research, P0
