# Replay Identity Report

## Evidence Summary
- **Total Events Replayed**: 200,000
- **Duration**: 720.51ms
- **Graph Size**: 43,539 nodes
- **Graph Hash**: d992fd79d65e79cbdae9ddfa31f10242d16d5df6f6190bd40931a8df753cf6d8
- **Precision**: 100.00%
- **Recall**: 100.00%
- **Average Reaction Time**: 50ms
- **Optimal Moves**: 200,000 / 200,000

## Determinism Status
- **Userspace Replay**: PARTIAL VERIFIED
  - Evidence: 100% hash matching on process graph reconstruction across re-runs.
- **Global Determinism**: UNVERIFIED
  - Note: Kernel-level jitter and distributed consensus effects were not part of this replay scope.

## Validation Results
- Replay Identity: **VERIFIED**
- Truth Mutation: **UNVERIFIED**
- Ledger Tamper: **UNVERIFIED**
- Hash Fork: **UNVERIFIED**
- Replay Split: **UNVERIFIED**
