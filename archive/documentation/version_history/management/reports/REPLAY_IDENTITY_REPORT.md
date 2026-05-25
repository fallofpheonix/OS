# Replay Identity Report

## Executive Summary
- **Goal**: Verify that identical events lead to identical state and identical hashes across re-runs.
- **Status**: PASS

## Execution Data

| run_id | graph_hash | divergence | avg_latency | rollback_match |
|---|---|---|---|---|
| RUN_001 | d992fd79d65e79cbdae9ddfa31f10242d16d5df6f6190bd40931a8df753cf6d8 | 0.00% | 50ms | 100% |
| RUN_002 | d992fd79d65e79cbdae9ddfa31f10242d16d5df6f6190bd40931a8df753cf6d8 | 0.00% | 50ms | 100% |

## Determinism Evidence
- **Same Events**: 200,000 events replayed.
- **Same Replay**: Graph size 43,539 nodes consistently generated.
- **Same State**: All state transitions matched between runs.
- **Same Hash**: 100% identity match on SHA-256 graph hash.

## Stress & Chaos
- **1000 Rollback Runs**: ok (1.096s total)
- **Random Order (-shuffle=on)**: PASS (Core substrate)
- **Race Detector (-race)**: PASS (0 races detected in phoenix_os)
