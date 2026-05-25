# PX-014 Replay Identity Lab

**Label**: core-validation, F0-exit
**Status**: IN_PROGRESS

## Problem
Empirical proof of byte-for-byte identity across 1000 consecutive replay runs is required for F0 exit.

## Tasks
- [x] Create `tests/replay/` suite
- [x] Implement `TestReplayIdentity1000`
- [x] Implement `TestReplayFork`
- [x] Implement `TestReplayHash`
- [x] Implement `TestReplayRecovery`
- [x] Implement `TestReplayStress`
- [x] Implement `TestCrossRunDeterminism`
- [ ] Execute 1000-run verification
- [ ] Generate `REPLAY_IDENTITY_REPORT.md`

## Verification
`go test -count=1000 ./tests/replay/`
